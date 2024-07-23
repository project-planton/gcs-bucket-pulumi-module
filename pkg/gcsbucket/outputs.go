package gcsbucket

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcsbucket/model"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/automationapi/autoapistackoutput"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func OutputMapTransformer(stackOutput auto.OutputMap,
	input *model.GcsBucketStackInput) *model.GcsBucketStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply ||
		stackOutput == nil {
		return &model.GcsBucketStackOutputs{}
	}
	return &model.GcsBucketStackOutputs{
		BucketId: autoapistackoutput.GetVal(stackOutput, GetGcsBucketIdOutputName()),
	}
}
