package pkg

import (
	"github.com/plantoncloud/project-planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcsbucket"
	"github.com/plantoncloud/project-planton/apis/zzgo/cloud/planton/apis/commons/apiresource/enums/apiresourcekind"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/gcplabelkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"strconv"
)

type Locals struct {
	GcpLabels map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *gcsbucket.GcsBucketStackInput) *Locals {
	locals := &Locals{}
	gcsBucket := stackInput.Target
	locals.GcpLabels = map[string]string{
		gcplabelkeys.Resource:     strconv.FormatBool(true),
		gcplabelkeys.Organization: gcsBucket.Spec.EnvironmentInfo.OrgId,
		gcplabelkeys.ResourceKind: apiresourcekind.ApiResourceKind_gcs_bucket.String(),
		gcplabelkeys.ResourceId:   gcsBucket.Metadata.Id,
	}
	return locals
}
