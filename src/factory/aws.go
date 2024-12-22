package factory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/yoelsusanto/the-workshop/src/awssdk"
)

func CreateAWSConfig(ctx context.Context) aws.Config {
	var (
		cfg aws.Config
		err error
	)
	switch CheckRuntimeEnvironment() {
	case EnvLocal:
		cfg, err = awssdk.NewLocalstackConfig(ctx)
	}

	if err != nil {
		panic("factory: failed to create AWS config")
	}

	return cfg
}
