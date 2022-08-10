package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type STSAssumeRoleAPI interface {
	AssumeRole(ctx context.Context,
		params *sts.AssumeRoleInput,
		optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

func takeRole(c context.Context, api STSAssumeRoleAPI, input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	return api.AssumeRole(c, input)
}

func assumeRole(roleARN *string, sessionName *string, serialNumber *string, token *string) (*sts.AssumeRoleOutput, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sts.NewFromConfig(cfg)

	input := &sts.AssumeRoleInput{
		RoleArn:         roleARN,
		RoleSessionName: sessionName,
		SerialNumber:    serialNumber,
		TokenCode:       token,
	}

	result, err := takeRole(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error assuming the role:")
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}
