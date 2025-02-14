package services

import (
	"context"

	"github.com/IAmRiteshKoushik/mercury/cmd"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var AwsConfig aws.Config

func SetupSession() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cmd.EnvVars.AwsRegionName))
	if err != nil {
		return err
	}
	AwsConfig = cfg
	return nil
}
