package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/gcs-bucket-pulumi-module/pkg"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcsbucket"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInputString, ok := ctx.GetConfig("planton-cloud:stack-input")
		if !ok {
			return errors.New("stack-input not found in pulumi config")
		}
		stackInput := &gcsbucket.GcsBucketStackInput{}
		if err := protojson.Unmarshal([]byte(stackInputString), stackInput); err != nil {
			return fmt.Errorf("failed to unmarshal stack input: %v", err)
		}
		s := &pkg.ResourceStack{
			StackInput: stackInput,
		}
		return s.Resources(ctx)
	})
}
