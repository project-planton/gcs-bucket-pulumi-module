package gcp

import (
	"context"
	sbv1stack "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/storagebucket/stack/gcp"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/rpc/enums"
)

func Outputs(ctx context.Context, input *sbv1stack.StorageBucketGcpStackInput) (*sbv1stack.StorageBucketGcpStackOutputs, error) {
	return &sbv1stack.StorageBucketGcpStackOutputs{}, nil
}

func Get(stackOutput map[string]interface{}, input *sbv1stack.StorageBucketGcpStackInput) *sbv1stack.StorageBucketGcpStackOutputs {
	if input.StackJob.OperationType != enums.StackOperationType_apply || stackOutput == nil {
		return &sbv1stack.StorageBucketGcpStackOutputs{}
	}
	return &sbv1stack.StorageBucketGcpStackOutputs{}
}
