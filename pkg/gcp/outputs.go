package gcp

import (
	"context"
	code2cloudv1deploybktstackgcpmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/code2cloud/deploy/storagebucket/stack/gcp/model"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/v1/stack/job/enums/operationtype"
)

func Outputs(ctx context.Context, input *code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackInput) (*code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs, error) {
	return &code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs{}, nil
}

func Get(stackOutput map[string]interface{}, input *code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackInput) *code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs {
	if input.StackJob.Spec.OperationType != operationtype.StackJobOperationType_apply || stackOutput == nil {
		return &code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs{}
	}
	return &code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs{}
}
