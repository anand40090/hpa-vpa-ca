// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package costexploreriface provides an interface to enable mocking the AWS Cost Explorer Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package costexploreriface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/costexplorer"
)

// CostExplorerAPI provides an interface to enable mocking the
// costexplorer.CostExplorer service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Cost Explorer Service.
//	func myFunc(svc costexploreriface.CostExplorerAPI) bool {
//	    // Make svc.CreateAnomalyMonitor request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := costexplorer.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCostExplorerClient struct {
//	    costexploreriface.CostExplorerAPI
//	}
//	func (m *mockCostExplorerClient) CreateAnomalyMonitor(input *costexplorer.CreateAnomalyMonitorInput) (*costexplorer.CreateAnomalyMonitorOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCostExplorerClient{}
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
type CostExplorerAPI interface {
	CreateAnomalyMonitor(*costexplorer.CreateAnomalyMonitorInput) (*costexplorer.CreateAnomalyMonitorOutput, error)
	CreateAnomalyMonitorWithContext(aws.Context, *costexplorer.CreateAnomalyMonitorInput, ...request.Option) (*costexplorer.CreateAnomalyMonitorOutput, error)
	CreateAnomalyMonitorRequest(*costexplorer.CreateAnomalyMonitorInput) (*request.Request, *costexplorer.CreateAnomalyMonitorOutput)

	CreateAnomalySubscription(*costexplorer.CreateAnomalySubscriptionInput) (*costexplorer.CreateAnomalySubscriptionOutput, error)
	CreateAnomalySubscriptionWithContext(aws.Context, *costexplorer.CreateAnomalySubscriptionInput, ...request.Option) (*costexplorer.CreateAnomalySubscriptionOutput, error)
	CreateAnomalySubscriptionRequest(*costexplorer.CreateAnomalySubscriptionInput) (*request.Request, *costexplorer.CreateAnomalySubscriptionOutput)

	CreateCostCategoryDefinition(*costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error)
	CreateCostCategoryDefinitionWithContext(aws.Context, *costexplorer.CreateCostCategoryDefinitionInput, ...request.Option) (*costexplorer.CreateCostCategoryDefinitionOutput, error)
	CreateCostCategoryDefinitionRequest(*costexplorer.CreateCostCategoryDefinitionInput) (*request.Request, *costexplorer.CreateCostCategoryDefinitionOutput)

	DeleteAnomalyMonitor(*costexplorer.DeleteAnomalyMonitorInput) (*costexplorer.DeleteAnomalyMonitorOutput, error)
	DeleteAnomalyMonitorWithContext(aws.Context, *costexplorer.DeleteAnomalyMonitorInput, ...request.Option) (*costexplorer.DeleteAnomalyMonitorOutput, error)
	DeleteAnomalyMonitorRequest(*costexplorer.DeleteAnomalyMonitorInput) (*request.Request, *costexplorer.DeleteAnomalyMonitorOutput)

	DeleteAnomalySubscription(*costexplorer.DeleteAnomalySubscriptionInput) (*costexplorer.DeleteAnomalySubscriptionOutput, error)
	DeleteAnomalySubscriptionWithContext(aws.Context, *costexplorer.DeleteAnomalySubscriptionInput, ...request.Option) (*costexplorer.DeleteAnomalySubscriptionOutput, error)
	DeleteAnomalySubscriptionRequest(*costexplorer.DeleteAnomalySubscriptionInput) (*request.Request, *costexplorer.DeleteAnomalySubscriptionOutput)

	DeleteCostCategoryDefinition(*costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error)
	DeleteCostCategoryDefinitionWithContext(aws.Context, *costexplorer.DeleteCostCategoryDefinitionInput, ...request.Option) (*costexplorer.DeleteCostCategoryDefinitionOutput, error)
	DeleteCostCategoryDefinitionRequest(*costexplorer.DeleteCostCategoryDefinitionInput) (*request.Request, *costexplorer.DeleteCostCategoryDefinitionOutput)

	DescribeCostCategoryDefinition(*costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error)
	DescribeCostCategoryDefinitionWithContext(aws.Context, *costexplorer.DescribeCostCategoryDefinitionInput, ...request.Option) (*costexplorer.DescribeCostCategoryDefinitionOutput, error)
	DescribeCostCategoryDefinitionRequest(*costexplorer.DescribeCostCategoryDefinitionInput) (*request.Request, *costexplorer.DescribeCostCategoryDefinitionOutput)

	GetAnomalies(*costexplorer.GetAnomaliesInput) (*costexplorer.GetAnomaliesOutput, error)
	GetAnomaliesWithContext(aws.Context, *costexplorer.GetAnomaliesInput, ...request.Option) (*costexplorer.GetAnomaliesOutput, error)
	GetAnomaliesRequest(*costexplorer.GetAnomaliesInput) (*request.Request, *costexplorer.GetAnomaliesOutput)

	GetAnomalyMonitors(*costexplorer.GetAnomalyMonitorsInput) (*costexplorer.GetAnomalyMonitorsOutput, error)
	GetAnomalyMonitorsWithContext(aws.Context, *costexplorer.GetAnomalyMonitorsInput, ...request.Option) (*costexplorer.GetAnomalyMonitorsOutput, error)
	GetAnomalyMonitorsRequest(*costexplorer.GetAnomalyMonitorsInput) (*request.Request, *costexplorer.GetAnomalyMonitorsOutput)

	GetAnomalySubscriptions(*costexplorer.GetAnomalySubscriptionsInput) (*costexplorer.GetAnomalySubscriptionsOutput, error)
	GetAnomalySubscriptionsWithContext(aws.Context, *costexplorer.GetAnomalySubscriptionsInput, ...request.Option) (*costexplorer.GetAnomalySubscriptionsOutput, error)
	GetAnomalySubscriptionsRequest(*costexplorer.GetAnomalySubscriptionsInput) (*request.Request, *costexplorer.GetAnomalySubscriptionsOutput)

	GetCostAndUsage(*costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error)
	GetCostAndUsageWithContext(aws.Context, *costexplorer.GetCostAndUsageInput, ...request.Option) (*costexplorer.GetCostAndUsageOutput, error)
	GetCostAndUsageRequest(*costexplorer.GetCostAndUsageInput) (*request.Request, *costexplorer.GetCostAndUsageOutput)

	GetCostAndUsageWithResources(*costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error)
	GetCostAndUsageWithResourcesWithContext(aws.Context, *costexplorer.GetCostAndUsageWithResourcesInput, ...request.Option) (*costexplorer.GetCostAndUsageWithResourcesOutput, error)
	GetCostAndUsageWithResourcesRequest(*costexplorer.GetCostAndUsageWithResourcesInput) (*request.Request, *costexplorer.GetCostAndUsageWithResourcesOutput)

	GetCostCategories(*costexplorer.GetCostCategoriesInput) (*costexplorer.GetCostCategoriesOutput, error)
	GetCostCategoriesWithContext(aws.Context, *costexplorer.GetCostCategoriesInput, ...request.Option) (*costexplorer.GetCostCategoriesOutput, error)
	GetCostCategoriesRequest(*costexplorer.GetCostCategoriesInput) (*request.Request, *costexplorer.GetCostCategoriesOutput)

	GetCostForecast(*costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error)
	GetCostForecastWithContext(aws.Context, *costexplorer.GetCostForecastInput, ...request.Option) (*costexplorer.GetCostForecastOutput, error)
	GetCostForecastRequest(*costexplorer.GetCostForecastInput) (*request.Request, *costexplorer.GetCostForecastOutput)

	GetDimensionValues(*costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error)
	GetDimensionValuesWithContext(aws.Context, *costexplorer.GetDimensionValuesInput, ...request.Option) (*costexplorer.GetDimensionValuesOutput, error)
	GetDimensionValuesRequest(*costexplorer.GetDimensionValuesInput) (*request.Request, *costexplorer.GetDimensionValuesOutput)

	GetReservationCoverage(*costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error)
	GetReservationCoverageWithContext(aws.Context, *costexplorer.GetReservationCoverageInput, ...request.Option) (*costexplorer.GetReservationCoverageOutput, error)
	GetReservationCoverageRequest(*costexplorer.GetReservationCoverageInput) (*request.Request, *costexplorer.GetReservationCoverageOutput)

	GetReservationPurchaseRecommendation(*costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error)
	GetReservationPurchaseRecommendationWithContext(aws.Context, *costexplorer.GetReservationPurchaseRecommendationInput, ...request.Option) (*costexplorer.GetReservationPurchaseRecommendationOutput, error)
	GetReservationPurchaseRecommendationRequest(*costexplorer.GetReservationPurchaseRecommendationInput) (*request.Request, *costexplorer.GetReservationPurchaseRecommendationOutput)

	GetReservationUtilization(*costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error)
	GetReservationUtilizationWithContext(aws.Context, *costexplorer.GetReservationUtilizationInput, ...request.Option) (*costexplorer.GetReservationUtilizationOutput, error)
	GetReservationUtilizationRequest(*costexplorer.GetReservationUtilizationInput) (*request.Request, *costexplorer.GetReservationUtilizationOutput)

	GetRightsizingRecommendation(*costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error)
	GetRightsizingRecommendationWithContext(aws.Context, *costexplorer.GetRightsizingRecommendationInput, ...request.Option) (*costexplorer.GetRightsizingRecommendationOutput, error)
	GetRightsizingRecommendationRequest(*costexplorer.GetRightsizingRecommendationInput) (*request.Request, *costexplorer.GetRightsizingRecommendationOutput)

	GetSavingsPlansCoverage(*costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error)
	GetSavingsPlansCoverageWithContext(aws.Context, *costexplorer.GetSavingsPlansCoverageInput, ...request.Option) (*costexplorer.GetSavingsPlansCoverageOutput, error)
	GetSavingsPlansCoverageRequest(*costexplorer.GetSavingsPlansCoverageInput) (*request.Request, *costexplorer.GetSavingsPlansCoverageOutput)

	GetSavingsPlansCoveragePages(*costexplorer.GetSavingsPlansCoverageInput, func(*costexplorer.GetSavingsPlansCoverageOutput, bool) bool) error
	GetSavingsPlansCoveragePagesWithContext(aws.Context, *costexplorer.GetSavingsPlansCoverageInput, func(*costexplorer.GetSavingsPlansCoverageOutput, bool) bool, ...request.Option) error

	GetSavingsPlansPurchaseRecommendation(*costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error)
	GetSavingsPlansPurchaseRecommendationWithContext(aws.Context, *costexplorer.GetSavingsPlansPurchaseRecommendationInput, ...request.Option) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error)
	GetSavingsPlansPurchaseRecommendationRequest(*costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*request.Request, *costexplorer.GetSavingsPlansPurchaseRecommendationOutput)

	GetSavingsPlansUtilization(*costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error)
	GetSavingsPlansUtilizationWithContext(aws.Context, *costexplorer.GetSavingsPlansUtilizationInput, ...request.Option) (*costexplorer.GetSavingsPlansUtilizationOutput, error)
	GetSavingsPlansUtilizationRequest(*costexplorer.GetSavingsPlansUtilizationInput) (*request.Request, *costexplorer.GetSavingsPlansUtilizationOutput)

	GetSavingsPlansUtilizationDetails(*costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error)
	GetSavingsPlansUtilizationDetailsWithContext(aws.Context, *costexplorer.GetSavingsPlansUtilizationDetailsInput, ...request.Option) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error)
	GetSavingsPlansUtilizationDetailsRequest(*costexplorer.GetSavingsPlansUtilizationDetailsInput) (*request.Request, *costexplorer.GetSavingsPlansUtilizationDetailsOutput)

	GetSavingsPlansUtilizationDetailsPages(*costexplorer.GetSavingsPlansUtilizationDetailsInput, func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) bool) error
	GetSavingsPlansUtilizationDetailsPagesWithContext(aws.Context, *costexplorer.GetSavingsPlansUtilizationDetailsInput, func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) bool, ...request.Option) error

	GetTags(*costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error)
	GetTagsWithContext(aws.Context, *costexplorer.GetTagsInput, ...request.Option) (*costexplorer.GetTagsOutput, error)
	GetTagsRequest(*costexplorer.GetTagsInput) (*request.Request, *costexplorer.GetTagsOutput)

	GetUsageForecast(*costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error)
	GetUsageForecastWithContext(aws.Context, *costexplorer.GetUsageForecastInput, ...request.Option) (*costexplorer.GetUsageForecastOutput, error)
	GetUsageForecastRequest(*costexplorer.GetUsageForecastInput) (*request.Request, *costexplorer.GetUsageForecastOutput)

	ListCostCategoryDefinitions(*costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error)
	ListCostCategoryDefinitionsWithContext(aws.Context, *costexplorer.ListCostCategoryDefinitionsInput, ...request.Option) (*costexplorer.ListCostCategoryDefinitionsOutput, error)
	ListCostCategoryDefinitionsRequest(*costexplorer.ListCostCategoryDefinitionsInput) (*request.Request, *costexplorer.ListCostCategoryDefinitionsOutput)

	ListCostCategoryDefinitionsPages(*costexplorer.ListCostCategoryDefinitionsInput, func(*costexplorer.ListCostCategoryDefinitionsOutput, bool) bool) error
	ListCostCategoryDefinitionsPagesWithContext(aws.Context, *costexplorer.ListCostCategoryDefinitionsInput, func(*costexplorer.ListCostCategoryDefinitionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*costexplorer.ListTagsForResourceInput) (*costexplorer.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *costexplorer.ListTagsForResourceInput, ...request.Option) (*costexplorer.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*costexplorer.ListTagsForResourceInput) (*request.Request, *costexplorer.ListTagsForResourceOutput)

	ProvideAnomalyFeedback(*costexplorer.ProvideAnomalyFeedbackInput) (*costexplorer.ProvideAnomalyFeedbackOutput, error)
	ProvideAnomalyFeedbackWithContext(aws.Context, *costexplorer.ProvideAnomalyFeedbackInput, ...request.Option) (*costexplorer.ProvideAnomalyFeedbackOutput, error)
	ProvideAnomalyFeedbackRequest(*costexplorer.ProvideAnomalyFeedbackInput) (*request.Request, *costexplorer.ProvideAnomalyFeedbackOutput)

	TagResource(*costexplorer.TagResourceInput) (*costexplorer.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *costexplorer.TagResourceInput, ...request.Option) (*costexplorer.TagResourceOutput, error)
	TagResourceRequest(*costexplorer.TagResourceInput) (*request.Request, *costexplorer.TagResourceOutput)

	UntagResource(*costexplorer.UntagResourceInput) (*costexplorer.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *costexplorer.UntagResourceInput, ...request.Option) (*costexplorer.UntagResourceOutput, error)
	UntagResourceRequest(*costexplorer.UntagResourceInput) (*request.Request, *costexplorer.UntagResourceOutput)

	UpdateAnomalyMonitor(*costexplorer.UpdateAnomalyMonitorInput) (*costexplorer.UpdateAnomalyMonitorOutput, error)
	UpdateAnomalyMonitorWithContext(aws.Context, *costexplorer.UpdateAnomalyMonitorInput, ...request.Option) (*costexplorer.UpdateAnomalyMonitorOutput, error)
	UpdateAnomalyMonitorRequest(*costexplorer.UpdateAnomalyMonitorInput) (*request.Request, *costexplorer.UpdateAnomalyMonitorOutput)

	UpdateAnomalySubscription(*costexplorer.UpdateAnomalySubscriptionInput) (*costexplorer.UpdateAnomalySubscriptionOutput, error)
	UpdateAnomalySubscriptionWithContext(aws.Context, *costexplorer.UpdateAnomalySubscriptionInput, ...request.Option) (*costexplorer.UpdateAnomalySubscriptionOutput, error)
	UpdateAnomalySubscriptionRequest(*costexplorer.UpdateAnomalySubscriptionInput) (*request.Request, *costexplorer.UpdateAnomalySubscriptionOutput)

	UpdateCostCategoryDefinition(*costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error)
	UpdateCostCategoryDefinitionWithContext(aws.Context, *costexplorer.UpdateCostCategoryDefinitionInput, ...request.Option) (*costexplorer.UpdateCostCategoryDefinitionOutput, error)
	UpdateCostCategoryDefinitionRequest(*costexplorer.UpdateCostCategoryDefinitionInput) (*request.Request, *costexplorer.UpdateCostCategoryDefinitionOutput)
}

var _ CostExplorerAPI = (*costexplorer.CostExplorer)(nil)
