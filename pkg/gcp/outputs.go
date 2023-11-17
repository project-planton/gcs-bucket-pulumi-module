package gcp

import (
	"context"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/enums/operationtype"

	sbv1stack "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/storagebucket/stack/gcp"
)

func Outputs(ctx context.Context, input *sbv1stack.StorageBucketGcpStackInput) (*sbv1stack.StorageBucketGcpStackOutputs, error) {
	return &sbv1stack.StorageBucketGcpStackOutputs{}, nil
}

func Get(stackOutput map[string]interface{}, input *sbv1stack.StorageBucketGcpStackInput) *sbv1stack.StorageBucketGcpStackOutputs {
	if input.StackJob.OperationType != operationtype.StackJobOperationType_apply || stackOutput == nil {
		return &sbv1stack.StorageBucketGcpStackOutputs{}
	}
	return &sbv1stack.StorageBucketGcpStackOutputs{}
}
