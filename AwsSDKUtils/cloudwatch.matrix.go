package awssdkutils

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

type Agent struct {
	client     *cloudwatch.Client
	Metricdata []types.MetricDatum
}

func NewAgent(cfg aws.Config) *Agent {
	client := cloudwatch.NewFromConfig(cfg)
	agent := &Agent{
		client:     client,
		Metricdata: []types.MetricDatum{},
	}
	return agent
}

func (agent *Agent) AddMetric(metric types.MetricDatum) {
	agent.Metricdata = append(agent.Metricdata, metric)
}

func (agent *Agent) ClearMetrix() {
	agent.Metricdata = []types.MetricDatum{}
}

func (agent *Agent) PutMetric(namespace string) { // <- this namespace mean the cloudwatch metrix namespace, it will record dimension and Metrix data under this namespace
	_, err := agent.client.PutMetricData(context.TODO(), &cloudwatch.PutMetricDataInput{
		Namespace:  aws.String(namespace),
		MetricData: agent.Metricdata,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func GenerateDimension(dimensionName string, dimensionValue string) []types.Dimension {
	return []types.Dimension{ // this is test for send one dimension per times, if you have other requirement, you can custom to add more dimesions to array
		types.Dimension{
			Name:  aws.String(dimensionName),
			Value: aws.String(dimensionValue),
		},
	}
}

func GenerateMetric(metricName string, unit types.StandardUnit, value float64, dimension []types.Dimension) *types.MetricDatum {
	return &types.MetricDatum{
		MetricName: aws.String(metricName),
		Unit:       unit,
		Value:      aws.Float64(value),
		Dimensions: dimension,
	}
}
