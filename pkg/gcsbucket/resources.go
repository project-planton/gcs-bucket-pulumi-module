package gcsbucket

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcsbucket/model"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/english/enums/englishword"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/gcp/pulumigoogleprovider"
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

	addedBucket, err := storage.NewBucket(ctx, gcsBucket.Metadata.Name, &storage.BucketArgs{
		ForceDestroy:             pulumi.Bool(true),
		Labels:                   pulumi.ToStringMap(s.GcpLabels),
		Location:                 pulumi.String(gcsBucket.Spec.GcpRegion),
		Name:                     pulumi.String(gcsBucket.Metadata.Name),
		Project:                  pulumi.String(gcsBucket.Spec.GcpProjectId),
		UniformBucketLevelAccess: pulumi.Bool(!gcsBucket.Spec.IsPublic),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to bucket resource")
	}

	ctx.Export(GetGcsBucketIdOutputName(), addedBucket.Project)

	if !gcsBucket.Spec.IsPublic {
		return nil
	}

	_, err = storage.NewBucketAccessControl(ctx, fmt.Sprintf("%s-public", gcsBucket.Metadata.Name), &storage.BucketAccessControlArgs{
		Bucket: addedBucket.Name,
		Role:   pulumi.String("READER"),
		Entity: pulumi.String("allUsers"),
	}, pulumi.Parent(addedBucket))
	if err != nil {
		return errors.Wrap(err, "failed to add public access control rule")
	}
	return nil
}

func GetGcsBucketIdOutputName() string {
	return pulumigoogleprovider.PulumiOutputName(storage.Bucket{}, englishword.EnglishWord_id.String())
}
