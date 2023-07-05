// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

//go:build go1.16 && integration
// +build go1.16,integration

package apigateway_test

import (
	"context"
	"testing"
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/awserr"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/awstesting/integration"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/apigateway"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_GetDomainNames(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := apigateway.New(sess)
	params := &apigateway.GetDomainNamesInput{}
	_, err := svc.GetDomainNamesWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_CreateUsagePlanKey(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := apigateway.New(sess)
	params := &apigateway.CreateUsagePlanKeyInput{
		KeyId:       aws.String("bar"),
		KeyType:     aws.String("fixx"),
		UsagePlanId: aws.String("foo"),
	}
	_, err := svc.CreateUsagePlanKeyWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if len(aerr.Code()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if len(aerr.Message()) == 0 {
		t.Errorf("expect non-empty error message")
	}
	if v := aerr.Code(); v == request.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
