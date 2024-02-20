package gcp

import (
	"context"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"

	code2cloudv1deploybktstackgcpmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/storagebucket/stack/gcp/model"
)

func Outputs(ctx context.Context, input *code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackInput) (*code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs, error) {
	return &code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs{}, nil
}

func Get(stackOutput map[string]interface{}, input *code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackInput) *code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply || stackOutput == nil {
		return &code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs{}
	}
	return &code2cloudv1deploybktstackgcpmodel.StorageBucketGcpStackOutputs{}
}
