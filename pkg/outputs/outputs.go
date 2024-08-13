package outputs

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcsbucket"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/automationapi/autoapistackoutput"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

const BucketIdOutputName = "bucket-id"

func PulumiOutputsToStackOutputsConverter(stackOutput auto.OutputMap,
	input *gcsbucket.GcsBucketStackInput) *gcsbucket.GcsBucketStackOutputs {
	return &gcsbucket.GcsBucketStackOutputs{
		BucketId: autoapistackoutput.GetVal(stackOutput, BucketIdOutputName),
	}
}
