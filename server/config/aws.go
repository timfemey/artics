package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func AWS() (aws.Config, error) {
	staticProvider := credentials.NewStaticCredentialsProvider(envAWSAccess(), envAWSecret(), "")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(staticProvider))
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}
