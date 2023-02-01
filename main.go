package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"

	awssdkutils "cloudwatch-agent/AwsSDKUtils"
)

var (
	region = "us-west-2"
)

func main() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
	)

	if err != nil {
		panic("Error to create load config : " + err.Error())
	}

	// use cfg to create cloudwatch agent
	cloudwatchAgent := awssdkutils.NewAgent(cfg)

	// create dimension
	dimensionName := fmt.Sprintf("test-Dimension")
	dimensionValue := fmt.Sprintf("test-index-1")
	dimension := awssdkutils.GenerateDimension(dimensionName, dimensionValue)

	// create Metric
	MetricName := "Test Value Record"
	Unit := types.StandardUnitBytes
	Value := float64(25535)
	Metric := awssdkutils.GenerateMetric(MetricName, Unit, Value, dimension)

	cloudwatchAgent.AddMetric(*Metric)
	cloudwatchAgent.PutMetric("Test")
}
