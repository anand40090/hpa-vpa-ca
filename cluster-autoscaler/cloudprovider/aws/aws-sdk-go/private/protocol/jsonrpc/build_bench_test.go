//go:build bench
// +build bench

package jsonrpc_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/awstesting"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol/json/jsonutil"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol/jsonrpc"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/dynamodb"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func BenchmarkJSONRPCBuild_Simple_dynamodbPutItem(b *testing.B) {
	svc := awstesting.NewClient()

	params := getDynamodbPutItemParams(b)

	for i := 0; i < b.N; i++ {
		r := svc.NewRequest(&request.Operation{Name: "Operation"}, params, nil)
		jsonrpc.Build(r)
		if r.Error != nil {
			b.Fatal("Unexpected error", r.Error)
		}
	}
}

func BenchmarkJSONUtilBuild_Simple_dynamodbPutItem(b *testing.B) {
	svc := awstesting.NewClient()

	params := getDynamodbPutItemParams(b)

	for i := 0; i < b.N; i++ {
		r := svc.NewRequest(&request.Operation{Name: "Operation"}, params, nil)
		_, err := jsonutil.BuildJSON(r.Params)
		if err != nil {
			b.Fatal("Unexpected error", err)
		}
	}
}

func BenchmarkEncodingJSONMarshal_Simple_dynamodbPutItem(b *testing.B) {
	params := getDynamodbPutItemParams(b)

	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		encoder := json.NewEncoder(buf)
		if err := encoder.Encode(params); err != nil {
			b.Fatal("Unexpected error", err)
		}
	}
}

func getDynamodbPutItemParams(b *testing.B) *dynamodb.PutItemInput {
	av, err := dynamodbattribute.ConvertToMap(struct {
		Key  string
		Data string
	}{Key: "MyKey", Data: "MyData"})
	if err != nil {
		b.Fatal("benchPutItem, expect no ConvertToMap errors", err)
	}
	return &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("tablename"),
	}
}
