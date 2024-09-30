package pkg

import (
	gcsbucketv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/apis/provider/gcp/gcsbucket/v1"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpLabels map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcsbucketv1.GcsBucketStackInput) *Locals {
	locals := &Locals{}
	gcsBucket := stackInput.Target
	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.Organization: gcsBucket.Spec.EnvironmentInfo.OrgId,
		gcplabelkeys.ResourceKind: "gcs_bucket",
		gcplabelkeys.ResourceId:   gcsBucket.Metadata.Id,
	}
	return locals
}
