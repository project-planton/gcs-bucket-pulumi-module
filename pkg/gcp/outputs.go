package gcp

import (
	sbv1stack "buf.build/gen/go/plantoncloud/planton-cloud-apis/protocolbuffers/go/cloud/planton/apis/v1/code2cloud/deploy/storagebucket/stack/gcp"
	"buf.build/gen/go/plantoncloud/planton-cloud-apis/protocolbuffers/go/cloud/planton/apis/v1/stack/rpc/enums"
	"context"
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
