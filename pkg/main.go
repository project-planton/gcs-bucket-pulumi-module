package pkg

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/gcs-bucket-pulumi-module/pkg/outputs"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcsbucket/model"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	Input     *model.GcsBucketStackInput
	GcpLabels map[string]string
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	gcpProvider, err := pulumigoogleprovider.Get(ctx, s.Input.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	gcsBucket := s.Input.ApiResource

	createdBucket, err := storage.NewBucket(ctx, gcsBucket.Metadata.Name, &storage.BucketArgs{
		ForceDestroy:             pulumi.Bool(true),
		Labels:                   pulumi.ToStringMap(s.GcpLabels),
		Location:                 pulumi.String(gcsBucket.Spec.GcpRegion),
		Name:                     pulumi.String(gcsBucket.Metadata.Name),
		Project:                  pulumi.String(gcsBucket.Spec.GcpProjectId),
		UniformBucketLevelAccess: pulumi.Bool(!gcsBucket.Spec.IsPublic),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create bucket resource")
	}

	ctx.Export(outputs.BucketIdOutputName, createdBucket.Project)

	if !gcsBucket.Spec.IsPublic {
		return nil
	}

	_, err = storage.NewBucketAccessControl(ctx, fmt.Sprintf("%s-public", gcsBucket.Metadata.Name), &storage.BucketAccessControlArgs{
		Bucket: createdBucket.Name,
		Role:   pulumi.String("READER"),
		Entity: pulumi.String("allUsers"),
	}, pulumi.Parent(createdBucket))
	if err != nil {
		return errors.Wrap(err, "failed to create public access control rule")
	}
	return nil
}
