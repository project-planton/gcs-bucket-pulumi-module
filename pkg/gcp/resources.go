package gcp

import (
	"fmt"

	"github.com/pkg/errors"
	pulumigcpprovider "github.com/plantoncloud-inc/pulumi-stack-runner-go-sdk/pkg/automation/provider/google"
	puluminameoutputgcp "github.com/plantoncloud-inc/pulumi-stack-runner-go-sdk/pkg/name/provider/cloud/gcp/output"
	code2cloudv1deploybktstackgcpmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/storagebucket/stack/gcp/model"
	pulumigcp "github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	Input     *code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackInput
	GcpLabels map[string]string
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	gcpProvider, err := pulumigcpprovider.Get(ctx, s.Input.CredentialsInput.Google)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	if err := addBucket(ctx, s.Input.ResourceInput, s.GcpLabels, gcpProvider); err != nil {
		return errors.Wrap(err, "failed to add bucket")
	}
	return nil
}

func addBucket(ctx *pulumi.Context, input *code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackResourceInput,
	labels map[string]string, gcpProvider *pulumigcp.Provider) error {
	addedBucket, err := storage.NewBucket(ctx, input.StorageBucket.Metadata.Name, &storage.BucketArgs{
		ForceDestroy:             pulumi.Bool(true),
		Labels:                   pulumi.ToStringMap(labels),
		Location:                 pulumi.String(input.StorageBucket.Spec.Gcp.Region),
		Name:                     pulumi.String(input.StorageBucket.Metadata.Name),
		Project:                  pulumi.String(input.StorageBucket.Spec.Gcp.ProjectId),
		UniformBucketLevelAccess: pulumi.Bool(!input.StorageBucket.Spec.IsExternal),
	}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to bucket resource")
	}
	ctx.Export(getGcpProjectIdOutputName(), addedBucket.Project)
	if !input.StorageBucket.Spec.IsExternal {
		return nil
	}
	_, err = storage.NewBucketAccessControl(ctx, fmt.Sprintf("%s-public", input.StorageBucket.Metadata.Name), &storage.BucketAccessControlArgs{
		Bucket: addedBucket.Name,
		Role:   pulumi.String("READER"),
		Entity: pulumi.String("allUsers"),
	}, pulumi.Parent(addedBucket))
	if err != nil {
		return errors.Wrap(err, "failed to add public access control rule")
	}
	return nil
}

func getGcpProjectIdOutputName() string {
	return puluminameoutputgcp.Name(storage.Bucket{}, "gcp-project-id")
}
