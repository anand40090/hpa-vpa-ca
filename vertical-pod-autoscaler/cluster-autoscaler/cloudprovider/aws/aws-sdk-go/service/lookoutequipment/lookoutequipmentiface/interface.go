// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lookoutequipmentiface provides an interface to enable mocking the Amazon Lookout for Equipment service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lookoutequipmentiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/lookoutequipment"
)

// LookoutEquipmentAPI provides an interface to enable mocking the
// lookoutequipment.LookoutEquipment service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Lookout for Equipment.
//	func myFunc(svc lookoutequipmentiface.LookoutEquipmentAPI) bool {
//	    // Make svc.CreateDataset request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := lookoutequipment.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockLookoutEquipmentClient struct {
//	    lookoutequipmentiface.LookoutEquipmentAPI
//	}
//	func (m *mockLookoutEquipmentClient) CreateDataset(input *lookoutequipment.CreateDatasetInput) (*lookoutequipment.CreateDatasetOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockLookoutEquipmentClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type LookoutEquipmentAPI interface {
	CreateDataset(*lookoutequipment.CreateDatasetInput) (*lookoutequipment.CreateDatasetOutput, error)
	CreateDatasetWithContext(aws.Context, *lookoutequipment.CreateDatasetInput, ...request.Option) (*lookoutequipment.CreateDatasetOutput, error)
	CreateDatasetRequest(*lookoutequipment.CreateDatasetInput) (*request.Request, *lookoutequipment.CreateDatasetOutput)

	CreateInferenceScheduler(*lookoutequipment.CreateInferenceSchedulerInput) (*lookoutequipment.CreateInferenceSchedulerOutput, error)
	CreateInferenceSchedulerWithContext(aws.Context, *lookoutequipment.CreateInferenceSchedulerInput, ...request.Option) (*lookoutequipment.CreateInferenceSchedulerOutput, error)
	CreateInferenceSchedulerRequest(*lookoutequipment.CreateInferenceSchedulerInput) (*request.Request, *lookoutequipment.CreateInferenceSchedulerOutput)

	CreateModel(*lookoutequipment.CreateModelInput) (*lookoutequipment.CreateModelOutput, error)
	CreateModelWithContext(aws.Context, *lookoutequipment.CreateModelInput, ...request.Option) (*lookoutequipment.CreateModelOutput, error)
	CreateModelRequest(*lookoutequipment.CreateModelInput) (*request.Request, *lookoutequipment.CreateModelOutput)

	DeleteDataset(*lookoutequipment.DeleteDatasetInput) (*lookoutequipment.DeleteDatasetOutput, error)
	DeleteDatasetWithContext(aws.Context, *lookoutequipment.DeleteDatasetInput, ...request.Option) (*lookoutequipment.DeleteDatasetOutput, error)
	DeleteDatasetRequest(*lookoutequipment.DeleteDatasetInput) (*request.Request, *lookoutequipment.DeleteDatasetOutput)

	DeleteInferenceScheduler(*lookoutequipment.DeleteInferenceSchedulerInput) (*lookoutequipment.DeleteInferenceSchedulerOutput, error)
	DeleteInferenceSchedulerWithContext(aws.Context, *lookoutequipment.DeleteInferenceSchedulerInput, ...request.Option) (*lookoutequipment.DeleteInferenceSchedulerOutput, error)
	DeleteInferenceSchedulerRequest(*lookoutequipment.DeleteInferenceSchedulerInput) (*request.Request, *lookoutequipment.DeleteInferenceSchedulerOutput)

	DeleteModel(*lookoutequipment.DeleteModelInput) (*lookoutequipment.DeleteModelOutput, error)
	DeleteModelWithContext(aws.Context, *lookoutequipment.DeleteModelInput, ...request.Option) (*lookoutequipment.DeleteModelOutput, error)
	DeleteModelRequest(*lookoutequipment.DeleteModelInput) (*request.Request, *lookoutequipment.DeleteModelOutput)

	DescribeDataIngestionJob(*lookoutequipment.DescribeDataIngestionJobInput) (*lookoutequipment.DescribeDataIngestionJobOutput, error)
	DescribeDataIngestionJobWithContext(aws.Context, *lookoutequipment.DescribeDataIngestionJobInput, ...request.Option) (*lookoutequipment.DescribeDataIngestionJobOutput, error)
	DescribeDataIngestionJobRequest(*lookoutequipment.DescribeDataIngestionJobInput) (*request.Request, *lookoutequipment.DescribeDataIngestionJobOutput)

	DescribeDataset(*lookoutequipment.DescribeDatasetInput) (*lookoutequipment.DescribeDatasetOutput, error)
	DescribeDatasetWithContext(aws.Context, *lookoutequipment.DescribeDatasetInput, ...request.Option) (*lookoutequipment.DescribeDatasetOutput, error)
	DescribeDatasetRequest(*lookoutequipment.DescribeDatasetInput) (*request.Request, *lookoutequipment.DescribeDatasetOutput)

	DescribeInferenceScheduler(*lookoutequipment.DescribeInferenceSchedulerInput) (*lookoutequipment.DescribeInferenceSchedulerOutput, error)
	DescribeInferenceSchedulerWithContext(aws.Context, *lookoutequipment.DescribeInferenceSchedulerInput, ...request.Option) (*lookoutequipment.DescribeInferenceSchedulerOutput, error)
	DescribeInferenceSchedulerRequest(*lookoutequipment.DescribeInferenceSchedulerInput) (*request.Request, *lookoutequipment.DescribeInferenceSchedulerOutput)

	DescribeModel(*lookoutequipment.DescribeModelInput) (*lookoutequipment.DescribeModelOutput, error)
	DescribeModelWithContext(aws.Context, *lookoutequipment.DescribeModelInput, ...request.Option) (*lookoutequipment.DescribeModelOutput, error)
	DescribeModelRequest(*lookoutequipment.DescribeModelInput) (*request.Request, *lookoutequipment.DescribeModelOutput)

	ListDataIngestionJobs(*lookoutequipment.ListDataIngestionJobsInput) (*lookoutequipment.ListDataIngestionJobsOutput, error)
	ListDataIngestionJobsWithContext(aws.Context, *lookoutequipment.ListDataIngestionJobsInput, ...request.Option) (*lookoutequipment.ListDataIngestionJobsOutput, error)
	ListDataIngestionJobsRequest(*lookoutequipment.ListDataIngestionJobsInput) (*request.Request, *lookoutequipment.ListDataIngestionJobsOutput)

	ListDataIngestionJobsPages(*lookoutequipment.ListDataIngestionJobsInput, func(*lookoutequipment.ListDataIngestionJobsOutput, bool) bool) error
	ListDataIngestionJobsPagesWithContext(aws.Context, *lookoutequipment.ListDataIngestionJobsInput, func(*lookoutequipment.ListDataIngestionJobsOutput, bool) bool, ...request.Option) error

	ListDatasets(*lookoutequipment.ListDatasetsInput) (*lookoutequipment.ListDatasetsOutput, error)
	ListDatasetsWithContext(aws.Context, *lookoutequipment.ListDatasetsInput, ...request.Option) (*lookoutequipment.ListDatasetsOutput, error)
	ListDatasetsRequest(*lookoutequipment.ListDatasetsInput) (*request.Request, *lookoutequipment.ListDatasetsOutput)

	ListDatasetsPages(*lookoutequipment.ListDatasetsInput, func(*lookoutequipment.ListDatasetsOutput, bool) bool) error
	ListDatasetsPagesWithContext(aws.Context, *lookoutequipment.ListDatasetsInput, func(*lookoutequipment.ListDatasetsOutput, bool) bool, ...request.Option) error

	ListInferenceExecutions(*lookoutequipment.ListInferenceExecutionsInput) (*lookoutequipment.ListInferenceExecutionsOutput, error)
	ListInferenceExecutionsWithContext(aws.Context, *lookoutequipment.ListInferenceExecutionsInput, ...request.Option) (*lookoutequipment.ListInferenceExecutionsOutput, error)
	ListInferenceExecutionsRequest(*lookoutequipment.ListInferenceExecutionsInput) (*request.Request, *lookoutequipment.ListInferenceExecutionsOutput)

	ListInferenceExecutionsPages(*lookoutequipment.ListInferenceExecutionsInput, func(*lookoutequipment.ListInferenceExecutionsOutput, bool) bool) error
	ListInferenceExecutionsPagesWithContext(aws.Context, *lookoutequipment.ListInferenceExecutionsInput, func(*lookoutequipment.ListInferenceExecutionsOutput, bool) bool, ...request.Option) error

	ListInferenceSchedulers(*lookoutequipment.ListInferenceSchedulersInput) (*lookoutequipment.ListInferenceSchedulersOutput, error)
	ListInferenceSchedulersWithContext(aws.Context, *lookoutequipment.ListInferenceSchedulersInput, ...request.Option) (*lookoutequipment.ListInferenceSchedulersOutput, error)
	ListInferenceSchedulersRequest(*lookoutequipment.ListInferenceSchedulersInput) (*request.Request, *lookoutequipment.ListInferenceSchedulersOutput)

	ListInferenceSchedulersPages(*lookoutequipment.ListInferenceSchedulersInput, func(*lookoutequipment.ListInferenceSchedulersOutput, bool) bool) error
	ListInferenceSchedulersPagesWithContext(aws.Context, *lookoutequipment.ListInferenceSchedulersInput, func(*lookoutequipment.ListInferenceSchedulersOutput, bool) bool, ...request.Option) error

	ListModels(*lookoutequipment.ListModelsInput) (*lookoutequipment.ListModelsOutput, error)
	ListModelsWithContext(aws.Context, *lookoutequipment.ListModelsInput, ...request.Option) (*lookoutequipment.ListModelsOutput, error)
	ListModelsRequest(*lookoutequipment.ListModelsInput) (*request.Request, *lookoutequipment.ListModelsOutput)

	ListModelsPages(*lookoutequipment.ListModelsInput, func(*lookoutequipment.ListModelsOutput, bool) bool) error
	ListModelsPagesWithContext(aws.Context, *lookoutequipment.ListModelsInput, func(*lookoutequipment.ListModelsOutput, bool) bool, ...request.Option) error

	ListSensorStatistics(*lookoutequipment.ListSensorStatisticsInput) (*lookoutequipment.ListSensorStatisticsOutput, error)
	ListSensorStatisticsWithContext(aws.Context, *lookoutequipment.ListSensorStatisticsInput, ...request.Option) (*lookoutequipment.ListSensorStatisticsOutput, error)
	ListSensorStatisticsRequest(*lookoutequipment.ListSensorStatisticsInput) (*request.Request, *lookoutequipment.ListSensorStatisticsOutput)

	ListSensorStatisticsPages(*lookoutequipment.ListSensorStatisticsInput, func(*lookoutequipment.ListSensorStatisticsOutput, bool) bool) error
	ListSensorStatisticsPagesWithContext(aws.Context, *lookoutequipment.ListSensorStatisticsInput, func(*lookoutequipment.ListSensorStatisticsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*lookoutequipment.ListTagsForResourceInput) (*lookoutequipment.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *lookoutequipment.ListTagsForResourceInput, ...request.Option) (*lookoutequipment.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*lookoutequipment.ListTagsForResourceInput) (*request.Request, *lookoutequipment.ListTagsForResourceOutput)

	StartDataIngestionJob(*lookoutequipment.StartDataIngestionJobInput) (*lookoutequipment.StartDataIngestionJobOutput, error)
	StartDataIngestionJobWithContext(aws.Context, *lookoutequipment.StartDataIngestionJobInput, ...request.Option) (*lookoutequipment.StartDataIngestionJobOutput, error)
	StartDataIngestionJobRequest(*lookoutequipment.StartDataIngestionJobInput) (*request.Request, *lookoutequipment.StartDataIngestionJobOutput)

	StartInferenceScheduler(*lookoutequipment.StartInferenceSchedulerInput) (*lookoutequipment.StartInferenceSchedulerOutput, error)
	StartInferenceSchedulerWithContext(aws.Context, *lookoutequipment.StartInferenceSchedulerInput, ...request.Option) (*lookoutequipment.StartInferenceSchedulerOutput, error)
	StartInferenceSchedulerRequest(*lookoutequipment.StartInferenceSchedulerInput) (*request.Request, *lookoutequipment.StartInferenceSchedulerOutput)

	StopInferenceScheduler(*lookoutequipment.StopInferenceSchedulerInput) (*lookoutequipment.StopInferenceSchedulerOutput, error)
	StopInferenceSchedulerWithContext(aws.Context, *lookoutequipment.StopInferenceSchedulerInput, ...request.Option) (*lookoutequipment.StopInferenceSchedulerOutput, error)
	StopInferenceSchedulerRequest(*lookoutequipment.StopInferenceSchedulerInput) (*request.Request, *lookoutequipment.StopInferenceSchedulerOutput)

	TagResource(*lookoutequipment.TagResourceInput) (*lookoutequipment.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *lookoutequipment.TagResourceInput, ...request.Option) (*lookoutequipment.TagResourceOutput, error)
	TagResourceRequest(*lookoutequipment.TagResourceInput) (*request.Request, *lookoutequipment.TagResourceOutput)

	UntagResource(*lookoutequipment.UntagResourceInput) (*lookoutequipment.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *lookoutequipment.UntagResourceInput, ...request.Option) (*lookoutequipment.UntagResourceOutput, error)
	UntagResourceRequest(*lookoutequipment.UntagResourceInput) (*request.Request, *lookoutequipment.UntagResourceOutput)

	UpdateInferenceScheduler(*lookoutequipment.UpdateInferenceSchedulerInput) (*lookoutequipment.UpdateInferenceSchedulerOutput, error)
	UpdateInferenceSchedulerWithContext(aws.Context, *lookoutequipment.UpdateInferenceSchedulerInput, ...request.Option) (*lookoutequipment.UpdateInferenceSchedulerOutput, error)
	UpdateInferenceSchedulerRequest(*lookoutequipment.UpdateInferenceSchedulerInput) (*request.Request, *lookoutequipment.UpdateInferenceSchedulerOutput)
}

var _ LookoutEquipmentAPI = (*lookoutequipment.LookoutEquipment)(nil)
