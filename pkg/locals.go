package pkg

import (
	gcsbucketv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcsbucket/v1"
	"github.com/project-planton/pulumi-module-golang-commons/pkg/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpLabels map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcsbucketv1.GcsBucketStackInput) *Locals {
	locals := &Locals{}

	//if the id is empty, use name as id
	if stackInput.Target.Metadata.Id == "" {
		stackInput.Target.Metadata.Id = stackInput.Target.Metadata.Name
	}

	gcsBucket := stackInput.Target
	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.ResourceId:   gcsBucket.Metadata.Id,
		gcplabelkeys.ResourceKind: "gcs_bucket",
	}

	if gcsBucket.Spec.EnvironmentInfo != nil && gcsBucket.Spec.EnvironmentInfo.OrgId != "" {
		locals.GcpLabels[gcplabelkeys.Organization] = gcsBucket.Spec.EnvironmentInfo.OrgId
	}

	return locals
}
