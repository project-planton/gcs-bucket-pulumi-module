package main

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/gcs-bucket-pulumi-module/pkg"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcsbucket"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/encoding/protojson"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	StackInputPulumiConfigKey = "planton-cloud:stack-input"
	StackInputFilePathEnvVar  = "STACK_INPUT_FILE_PATH"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInputString, ok := ctx.GetConfig(StackInputPulumiConfigKey)
		var jsonBytes []byte
		var err error

		if !ok {
			stackInputFilePath := os.Getenv("STACK_INPUT_FILE_PATH")
			if stackInputFilePath == "" {
				return errors.Errorf("stack-input not found in pulumi config %s or in %s environment variable",
					StackInputPulumiConfigKey, StackInputFilePathEnvVar)
			}
			inputFileBytes, err := os.ReadFile(stackInputFilePath)
			if err != nil {
				return errors.Wrap(err, "failed to read input file")
			}
			jsonBytes, err = yaml.YAMLToJSON(inputFileBytes)
			if err != nil {
				return errors.Wrap(err, "failed to load yaml to json")
			}
		} else {
			jsonBytes, err = yaml.YAMLToJSON([]byte(stackInputString))
			if err != nil {
				return errors.Wrap(err, "failed to load yaml to json")
			}
		}

		stackInput := &gcsbucket.GcsBucketStackInput{}

		if err := protojson.Unmarshal(jsonBytes, stackInput); err != nil {
			return errors.Wrap(err, "failed to load json into proto message")
		}

		s := &pkg.ResourceStack{
			StackInput: stackInput,
		}

		return s.Resources(ctx)
	})
}
