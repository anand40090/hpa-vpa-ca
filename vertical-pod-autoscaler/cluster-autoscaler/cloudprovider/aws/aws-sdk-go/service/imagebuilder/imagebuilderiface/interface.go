// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package imagebuilderiface provides an interface to enable mocking the EC2 Image Builder service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package imagebuilderiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/imagebuilder"
)

// ImagebuilderAPI provides an interface to enable mocking the
// imagebuilder.Imagebuilder service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// EC2 Image Builder.
//	func myFunc(svc imagebuilderiface.ImagebuilderAPI) bool {
//	    // Make svc.CancelImageCreation request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := imagebuilder.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockImagebuilderClient struct {
//	    imagebuilderiface.ImagebuilderAPI
//	}
//	func (m *mockImagebuilderClient) CancelImageCreation(input *imagebuilder.CancelImageCreationInput) (*imagebuilder.CancelImageCreationOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockImagebuilderClient{}
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
type ImagebuilderAPI interface {
	CancelImageCreation(*imagebuilder.CancelImageCreationInput) (*imagebuilder.CancelImageCreationOutput, error)
	CancelImageCreationWithContext(aws.Context, *imagebuilder.CancelImageCreationInput, ...request.Option) (*imagebuilder.CancelImageCreationOutput, error)
	CancelImageCreationRequest(*imagebuilder.CancelImageCreationInput) (*request.Request, *imagebuilder.CancelImageCreationOutput)

	CreateComponent(*imagebuilder.CreateComponentInput) (*imagebuilder.CreateComponentOutput, error)
	CreateComponentWithContext(aws.Context, *imagebuilder.CreateComponentInput, ...request.Option) (*imagebuilder.CreateComponentOutput, error)
	CreateComponentRequest(*imagebuilder.CreateComponentInput) (*request.Request, *imagebuilder.CreateComponentOutput)

	CreateContainerRecipe(*imagebuilder.CreateContainerRecipeInput) (*imagebuilder.CreateContainerRecipeOutput, error)
	CreateContainerRecipeWithContext(aws.Context, *imagebuilder.CreateContainerRecipeInput, ...request.Option) (*imagebuilder.CreateContainerRecipeOutput, error)
	CreateContainerRecipeRequest(*imagebuilder.CreateContainerRecipeInput) (*request.Request, *imagebuilder.CreateContainerRecipeOutput)

	CreateDistributionConfiguration(*imagebuilder.CreateDistributionConfigurationInput) (*imagebuilder.CreateDistributionConfigurationOutput, error)
	CreateDistributionConfigurationWithContext(aws.Context, *imagebuilder.CreateDistributionConfigurationInput, ...request.Option) (*imagebuilder.CreateDistributionConfigurationOutput, error)
	CreateDistributionConfigurationRequest(*imagebuilder.CreateDistributionConfigurationInput) (*request.Request, *imagebuilder.CreateDistributionConfigurationOutput)

	CreateImage(*imagebuilder.CreateImageInput) (*imagebuilder.CreateImageOutput, error)
	CreateImageWithContext(aws.Context, *imagebuilder.CreateImageInput, ...request.Option) (*imagebuilder.CreateImageOutput, error)
	CreateImageRequest(*imagebuilder.CreateImageInput) (*request.Request, *imagebuilder.CreateImageOutput)

	CreateImagePipeline(*imagebuilder.CreateImagePipelineInput) (*imagebuilder.CreateImagePipelineOutput, error)
	CreateImagePipelineWithContext(aws.Context, *imagebuilder.CreateImagePipelineInput, ...request.Option) (*imagebuilder.CreateImagePipelineOutput, error)
	CreateImagePipelineRequest(*imagebuilder.CreateImagePipelineInput) (*request.Request, *imagebuilder.CreateImagePipelineOutput)

	CreateImageRecipe(*imagebuilder.CreateImageRecipeInput) (*imagebuilder.CreateImageRecipeOutput, error)
	CreateImageRecipeWithContext(aws.Context, *imagebuilder.CreateImageRecipeInput, ...request.Option) (*imagebuilder.CreateImageRecipeOutput, error)
	CreateImageRecipeRequest(*imagebuilder.CreateImageRecipeInput) (*request.Request, *imagebuilder.CreateImageRecipeOutput)

	CreateInfrastructureConfiguration(*imagebuilder.CreateInfrastructureConfigurationInput) (*imagebuilder.CreateInfrastructureConfigurationOutput, error)
	CreateInfrastructureConfigurationWithContext(aws.Context, *imagebuilder.CreateInfrastructureConfigurationInput, ...request.Option) (*imagebuilder.CreateInfrastructureConfigurationOutput, error)
	CreateInfrastructureConfigurationRequest(*imagebuilder.CreateInfrastructureConfigurationInput) (*request.Request, *imagebuilder.CreateInfrastructureConfigurationOutput)

	DeleteComponent(*imagebuilder.DeleteComponentInput) (*imagebuilder.DeleteComponentOutput, error)
	DeleteComponentWithContext(aws.Context, *imagebuilder.DeleteComponentInput, ...request.Option) (*imagebuilder.DeleteComponentOutput, error)
	DeleteComponentRequest(*imagebuilder.DeleteComponentInput) (*request.Request, *imagebuilder.DeleteComponentOutput)

	DeleteContainerRecipe(*imagebuilder.DeleteContainerRecipeInput) (*imagebuilder.DeleteContainerRecipeOutput, error)
	DeleteContainerRecipeWithContext(aws.Context, *imagebuilder.DeleteContainerRecipeInput, ...request.Option) (*imagebuilder.DeleteContainerRecipeOutput, error)
	DeleteContainerRecipeRequest(*imagebuilder.DeleteContainerRecipeInput) (*request.Request, *imagebuilder.DeleteContainerRecipeOutput)

	DeleteDistributionConfiguration(*imagebuilder.DeleteDistributionConfigurationInput) (*imagebuilder.DeleteDistributionConfigurationOutput, error)
	DeleteDistributionConfigurationWithContext(aws.Context, *imagebuilder.DeleteDistributionConfigurationInput, ...request.Option) (*imagebuilder.DeleteDistributionConfigurationOutput, error)
	DeleteDistributionConfigurationRequest(*imagebuilder.DeleteDistributionConfigurationInput) (*request.Request, *imagebuilder.DeleteDistributionConfigurationOutput)

	DeleteImage(*imagebuilder.DeleteImageInput) (*imagebuilder.DeleteImageOutput, error)
	DeleteImageWithContext(aws.Context, *imagebuilder.DeleteImageInput, ...request.Option) (*imagebuilder.DeleteImageOutput, error)
	DeleteImageRequest(*imagebuilder.DeleteImageInput) (*request.Request, *imagebuilder.DeleteImageOutput)

	DeleteImagePipeline(*imagebuilder.DeleteImagePipelineInput) (*imagebuilder.DeleteImagePipelineOutput, error)
	DeleteImagePipelineWithContext(aws.Context, *imagebuilder.DeleteImagePipelineInput, ...request.Option) (*imagebuilder.DeleteImagePipelineOutput, error)
	DeleteImagePipelineRequest(*imagebuilder.DeleteImagePipelineInput) (*request.Request, *imagebuilder.DeleteImagePipelineOutput)

	DeleteImageRecipe(*imagebuilder.DeleteImageRecipeInput) (*imagebuilder.DeleteImageRecipeOutput, error)
	DeleteImageRecipeWithContext(aws.Context, *imagebuilder.DeleteImageRecipeInput, ...request.Option) (*imagebuilder.DeleteImageRecipeOutput, error)
	DeleteImageRecipeRequest(*imagebuilder.DeleteImageRecipeInput) (*request.Request, *imagebuilder.DeleteImageRecipeOutput)

	DeleteInfrastructureConfiguration(*imagebuilder.DeleteInfrastructureConfigurationInput) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error)
	DeleteInfrastructureConfigurationWithContext(aws.Context, *imagebuilder.DeleteInfrastructureConfigurationInput, ...request.Option) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error)
	DeleteInfrastructureConfigurationRequest(*imagebuilder.DeleteInfrastructureConfigurationInput) (*request.Request, *imagebuilder.DeleteInfrastructureConfigurationOutput)

	GetComponent(*imagebuilder.GetComponentInput) (*imagebuilder.GetComponentOutput, error)
	GetComponentWithContext(aws.Context, *imagebuilder.GetComponentInput, ...request.Option) (*imagebuilder.GetComponentOutput, error)
	GetComponentRequest(*imagebuilder.GetComponentInput) (*request.Request, *imagebuilder.GetComponentOutput)

	GetComponentPolicy(*imagebuilder.GetComponentPolicyInput) (*imagebuilder.GetComponentPolicyOutput, error)
	GetComponentPolicyWithContext(aws.Context, *imagebuilder.GetComponentPolicyInput, ...request.Option) (*imagebuilder.GetComponentPolicyOutput, error)
	GetComponentPolicyRequest(*imagebuilder.GetComponentPolicyInput) (*request.Request, *imagebuilder.GetComponentPolicyOutput)

	GetContainerRecipe(*imagebuilder.GetContainerRecipeInput) (*imagebuilder.GetContainerRecipeOutput, error)
	GetContainerRecipeWithContext(aws.Context, *imagebuilder.GetContainerRecipeInput, ...request.Option) (*imagebuilder.GetContainerRecipeOutput, error)
	GetContainerRecipeRequest(*imagebuilder.GetContainerRecipeInput) (*request.Request, *imagebuilder.GetContainerRecipeOutput)

	GetContainerRecipePolicy(*imagebuilder.GetContainerRecipePolicyInput) (*imagebuilder.GetContainerRecipePolicyOutput, error)
	GetContainerRecipePolicyWithContext(aws.Context, *imagebuilder.GetContainerRecipePolicyInput, ...request.Option) (*imagebuilder.GetContainerRecipePolicyOutput, error)
	GetContainerRecipePolicyRequest(*imagebuilder.GetContainerRecipePolicyInput) (*request.Request, *imagebuilder.GetContainerRecipePolicyOutput)

	GetDistributionConfiguration(*imagebuilder.GetDistributionConfigurationInput) (*imagebuilder.GetDistributionConfigurationOutput, error)
	GetDistributionConfigurationWithContext(aws.Context, *imagebuilder.GetDistributionConfigurationInput, ...request.Option) (*imagebuilder.GetDistributionConfigurationOutput, error)
	GetDistributionConfigurationRequest(*imagebuilder.GetDistributionConfigurationInput) (*request.Request, *imagebuilder.GetDistributionConfigurationOutput)

	GetImage(*imagebuilder.GetImageInput) (*imagebuilder.GetImageOutput, error)
	GetImageWithContext(aws.Context, *imagebuilder.GetImageInput, ...request.Option) (*imagebuilder.GetImageOutput, error)
	GetImageRequest(*imagebuilder.GetImageInput) (*request.Request, *imagebuilder.GetImageOutput)

	GetImagePipeline(*imagebuilder.GetImagePipelineInput) (*imagebuilder.GetImagePipelineOutput, error)
	GetImagePipelineWithContext(aws.Context, *imagebuilder.GetImagePipelineInput, ...request.Option) (*imagebuilder.GetImagePipelineOutput, error)
	GetImagePipelineRequest(*imagebuilder.GetImagePipelineInput) (*request.Request, *imagebuilder.GetImagePipelineOutput)

	GetImagePolicy(*imagebuilder.GetImagePolicyInput) (*imagebuilder.GetImagePolicyOutput, error)
	GetImagePolicyWithContext(aws.Context, *imagebuilder.GetImagePolicyInput, ...request.Option) (*imagebuilder.GetImagePolicyOutput, error)
	GetImagePolicyRequest(*imagebuilder.GetImagePolicyInput) (*request.Request, *imagebuilder.GetImagePolicyOutput)

	GetImageRecipe(*imagebuilder.GetImageRecipeInput) (*imagebuilder.GetImageRecipeOutput, error)
	GetImageRecipeWithContext(aws.Context, *imagebuilder.GetImageRecipeInput, ...request.Option) (*imagebuilder.GetImageRecipeOutput, error)
	GetImageRecipeRequest(*imagebuilder.GetImageRecipeInput) (*request.Request, *imagebuilder.GetImageRecipeOutput)

	GetImageRecipePolicy(*imagebuilder.GetImageRecipePolicyInput) (*imagebuilder.GetImageRecipePolicyOutput, error)
	GetImageRecipePolicyWithContext(aws.Context, *imagebuilder.GetImageRecipePolicyInput, ...request.Option) (*imagebuilder.GetImageRecipePolicyOutput, error)
	GetImageRecipePolicyRequest(*imagebuilder.GetImageRecipePolicyInput) (*request.Request, *imagebuilder.GetImageRecipePolicyOutput)

	GetInfrastructureConfiguration(*imagebuilder.GetInfrastructureConfigurationInput) (*imagebuilder.GetInfrastructureConfigurationOutput, error)
	GetInfrastructureConfigurationWithContext(aws.Context, *imagebuilder.GetInfrastructureConfigurationInput, ...request.Option) (*imagebuilder.GetInfrastructureConfigurationOutput, error)
	GetInfrastructureConfigurationRequest(*imagebuilder.GetInfrastructureConfigurationInput) (*request.Request, *imagebuilder.GetInfrastructureConfigurationOutput)

	ImportComponent(*imagebuilder.ImportComponentInput) (*imagebuilder.ImportComponentOutput, error)
	ImportComponentWithContext(aws.Context, *imagebuilder.ImportComponentInput, ...request.Option) (*imagebuilder.ImportComponentOutput, error)
	ImportComponentRequest(*imagebuilder.ImportComponentInput) (*request.Request, *imagebuilder.ImportComponentOutput)

	ImportVmImage(*imagebuilder.ImportVmImageInput) (*imagebuilder.ImportVmImageOutput, error)
	ImportVmImageWithContext(aws.Context, *imagebuilder.ImportVmImageInput, ...request.Option) (*imagebuilder.ImportVmImageOutput, error)
	ImportVmImageRequest(*imagebuilder.ImportVmImageInput) (*request.Request, *imagebuilder.ImportVmImageOutput)

	ListComponentBuildVersions(*imagebuilder.ListComponentBuildVersionsInput) (*imagebuilder.ListComponentBuildVersionsOutput, error)
	ListComponentBuildVersionsWithContext(aws.Context, *imagebuilder.ListComponentBuildVersionsInput, ...request.Option) (*imagebuilder.ListComponentBuildVersionsOutput, error)
	ListComponentBuildVersionsRequest(*imagebuilder.ListComponentBuildVersionsInput) (*request.Request, *imagebuilder.ListComponentBuildVersionsOutput)

	ListComponentBuildVersionsPages(*imagebuilder.ListComponentBuildVersionsInput, func(*imagebuilder.ListComponentBuildVersionsOutput, bool) bool) error
	ListComponentBuildVersionsPagesWithContext(aws.Context, *imagebuilder.ListComponentBuildVersionsInput, func(*imagebuilder.ListComponentBuildVersionsOutput, bool) bool, ...request.Option) error

	ListComponents(*imagebuilder.ListComponentsInput) (*imagebuilder.ListComponentsOutput, error)
	ListComponentsWithContext(aws.Context, *imagebuilder.ListComponentsInput, ...request.Option) (*imagebuilder.ListComponentsOutput, error)
	ListComponentsRequest(*imagebuilder.ListComponentsInput) (*request.Request, *imagebuilder.ListComponentsOutput)

	ListComponentsPages(*imagebuilder.ListComponentsInput, func(*imagebuilder.ListComponentsOutput, bool) bool) error
	ListComponentsPagesWithContext(aws.Context, *imagebuilder.ListComponentsInput, func(*imagebuilder.ListComponentsOutput, bool) bool, ...request.Option) error

	ListContainerRecipes(*imagebuilder.ListContainerRecipesInput) (*imagebuilder.ListContainerRecipesOutput, error)
	ListContainerRecipesWithContext(aws.Context, *imagebuilder.ListContainerRecipesInput, ...request.Option) (*imagebuilder.ListContainerRecipesOutput, error)
	ListContainerRecipesRequest(*imagebuilder.ListContainerRecipesInput) (*request.Request, *imagebuilder.ListContainerRecipesOutput)

	ListContainerRecipesPages(*imagebuilder.ListContainerRecipesInput, func(*imagebuilder.ListContainerRecipesOutput, bool) bool) error
	ListContainerRecipesPagesWithContext(aws.Context, *imagebuilder.ListContainerRecipesInput, func(*imagebuilder.ListContainerRecipesOutput, bool) bool, ...request.Option) error

	ListDistributionConfigurations(*imagebuilder.ListDistributionConfigurationsInput) (*imagebuilder.ListDistributionConfigurationsOutput, error)
	ListDistributionConfigurationsWithContext(aws.Context, *imagebuilder.ListDistributionConfigurationsInput, ...request.Option) (*imagebuilder.ListDistributionConfigurationsOutput, error)
	ListDistributionConfigurationsRequest(*imagebuilder.ListDistributionConfigurationsInput) (*request.Request, *imagebuilder.ListDistributionConfigurationsOutput)

	ListDistributionConfigurationsPages(*imagebuilder.ListDistributionConfigurationsInput, func(*imagebuilder.ListDistributionConfigurationsOutput, bool) bool) error
	ListDistributionConfigurationsPagesWithContext(aws.Context, *imagebuilder.ListDistributionConfigurationsInput, func(*imagebuilder.ListDistributionConfigurationsOutput, bool) bool, ...request.Option) error

	ListImageBuildVersions(*imagebuilder.ListImageBuildVersionsInput) (*imagebuilder.ListImageBuildVersionsOutput, error)
	ListImageBuildVersionsWithContext(aws.Context, *imagebuilder.ListImageBuildVersionsInput, ...request.Option) (*imagebuilder.ListImageBuildVersionsOutput, error)
	ListImageBuildVersionsRequest(*imagebuilder.ListImageBuildVersionsInput) (*request.Request, *imagebuilder.ListImageBuildVersionsOutput)

	ListImageBuildVersionsPages(*imagebuilder.ListImageBuildVersionsInput, func(*imagebuilder.ListImageBuildVersionsOutput, bool) bool) error
	ListImageBuildVersionsPagesWithContext(aws.Context, *imagebuilder.ListImageBuildVersionsInput, func(*imagebuilder.ListImageBuildVersionsOutput, bool) bool, ...request.Option) error

	ListImagePackages(*imagebuilder.ListImagePackagesInput) (*imagebuilder.ListImagePackagesOutput, error)
	ListImagePackagesWithContext(aws.Context, *imagebuilder.ListImagePackagesInput, ...request.Option) (*imagebuilder.ListImagePackagesOutput, error)
	ListImagePackagesRequest(*imagebuilder.ListImagePackagesInput) (*request.Request, *imagebuilder.ListImagePackagesOutput)

	ListImagePackagesPages(*imagebuilder.ListImagePackagesInput, func(*imagebuilder.ListImagePackagesOutput, bool) bool) error
	ListImagePackagesPagesWithContext(aws.Context, *imagebuilder.ListImagePackagesInput, func(*imagebuilder.ListImagePackagesOutput, bool) bool, ...request.Option) error

	ListImagePipelineImages(*imagebuilder.ListImagePipelineImagesInput) (*imagebuilder.ListImagePipelineImagesOutput, error)
	ListImagePipelineImagesWithContext(aws.Context, *imagebuilder.ListImagePipelineImagesInput, ...request.Option) (*imagebuilder.ListImagePipelineImagesOutput, error)
	ListImagePipelineImagesRequest(*imagebuilder.ListImagePipelineImagesInput) (*request.Request, *imagebuilder.ListImagePipelineImagesOutput)

	ListImagePipelineImagesPages(*imagebuilder.ListImagePipelineImagesInput, func(*imagebuilder.ListImagePipelineImagesOutput, bool) bool) error
	ListImagePipelineImagesPagesWithContext(aws.Context, *imagebuilder.ListImagePipelineImagesInput, func(*imagebuilder.ListImagePipelineImagesOutput, bool) bool, ...request.Option) error

	ListImagePipelines(*imagebuilder.ListImagePipelinesInput) (*imagebuilder.ListImagePipelinesOutput, error)
	ListImagePipelinesWithContext(aws.Context, *imagebuilder.ListImagePipelinesInput, ...request.Option) (*imagebuilder.ListImagePipelinesOutput, error)
	ListImagePipelinesRequest(*imagebuilder.ListImagePipelinesInput) (*request.Request, *imagebuilder.ListImagePipelinesOutput)

	ListImagePipelinesPages(*imagebuilder.ListImagePipelinesInput, func(*imagebuilder.ListImagePipelinesOutput, bool) bool) error
	ListImagePipelinesPagesWithContext(aws.Context, *imagebuilder.ListImagePipelinesInput, func(*imagebuilder.ListImagePipelinesOutput, bool) bool, ...request.Option) error

	ListImageRecipes(*imagebuilder.ListImageRecipesInput) (*imagebuilder.ListImageRecipesOutput, error)
	ListImageRecipesWithContext(aws.Context, *imagebuilder.ListImageRecipesInput, ...request.Option) (*imagebuilder.ListImageRecipesOutput, error)
	ListImageRecipesRequest(*imagebuilder.ListImageRecipesInput) (*request.Request, *imagebuilder.ListImageRecipesOutput)

	ListImageRecipesPages(*imagebuilder.ListImageRecipesInput, func(*imagebuilder.ListImageRecipesOutput, bool) bool) error
	ListImageRecipesPagesWithContext(aws.Context, *imagebuilder.ListImageRecipesInput, func(*imagebuilder.ListImageRecipesOutput, bool) bool, ...request.Option) error

	ListImages(*imagebuilder.ListImagesInput) (*imagebuilder.ListImagesOutput, error)
	ListImagesWithContext(aws.Context, *imagebuilder.ListImagesInput, ...request.Option) (*imagebuilder.ListImagesOutput, error)
	ListImagesRequest(*imagebuilder.ListImagesInput) (*request.Request, *imagebuilder.ListImagesOutput)

	ListImagesPages(*imagebuilder.ListImagesInput, func(*imagebuilder.ListImagesOutput, bool) bool) error
	ListImagesPagesWithContext(aws.Context, *imagebuilder.ListImagesInput, func(*imagebuilder.ListImagesOutput, bool) bool, ...request.Option) error

	ListInfrastructureConfigurations(*imagebuilder.ListInfrastructureConfigurationsInput) (*imagebuilder.ListInfrastructureConfigurationsOutput, error)
	ListInfrastructureConfigurationsWithContext(aws.Context, *imagebuilder.ListInfrastructureConfigurationsInput, ...request.Option) (*imagebuilder.ListInfrastructureConfigurationsOutput, error)
	ListInfrastructureConfigurationsRequest(*imagebuilder.ListInfrastructureConfigurationsInput) (*request.Request, *imagebuilder.ListInfrastructureConfigurationsOutput)

	ListInfrastructureConfigurationsPages(*imagebuilder.ListInfrastructureConfigurationsInput, func(*imagebuilder.ListInfrastructureConfigurationsOutput, bool) bool) error
	ListInfrastructureConfigurationsPagesWithContext(aws.Context, *imagebuilder.ListInfrastructureConfigurationsInput, func(*imagebuilder.ListInfrastructureConfigurationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*imagebuilder.ListTagsForResourceInput) (*imagebuilder.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *imagebuilder.ListTagsForResourceInput, ...request.Option) (*imagebuilder.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*imagebuilder.ListTagsForResourceInput) (*request.Request, *imagebuilder.ListTagsForResourceOutput)

	PutComponentPolicy(*imagebuilder.PutComponentPolicyInput) (*imagebuilder.PutComponentPolicyOutput, error)
	PutComponentPolicyWithContext(aws.Context, *imagebuilder.PutComponentPolicyInput, ...request.Option) (*imagebuilder.PutComponentPolicyOutput, error)
	PutComponentPolicyRequest(*imagebuilder.PutComponentPolicyInput) (*request.Request, *imagebuilder.PutComponentPolicyOutput)

	PutContainerRecipePolicy(*imagebuilder.PutContainerRecipePolicyInput) (*imagebuilder.PutContainerRecipePolicyOutput, error)
	PutContainerRecipePolicyWithContext(aws.Context, *imagebuilder.PutContainerRecipePolicyInput, ...request.Option) (*imagebuilder.PutContainerRecipePolicyOutput, error)
	PutContainerRecipePolicyRequest(*imagebuilder.PutContainerRecipePolicyInput) (*request.Request, *imagebuilder.PutContainerRecipePolicyOutput)

	PutImagePolicy(*imagebuilder.PutImagePolicyInput) (*imagebuilder.PutImagePolicyOutput, error)
	PutImagePolicyWithContext(aws.Context, *imagebuilder.PutImagePolicyInput, ...request.Option) (*imagebuilder.PutImagePolicyOutput, error)
	PutImagePolicyRequest(*imagebuilder.PutImagePolicyInput) (*request.Request, *imagebuilder.PutImagePolicyOutput)

	PutImageRecipePolicy(*imagebuilder.PutImageRecipePolicyInput) (*imagebuilder.PutImageRecipePolicyOutput, error)
	PutImageRecipePolicyWithContext(aws.Context, *imagebuilder.PutImageRecipePolicyInput, ...request.Option) (*imagebuilder.PutImageRecipePolicyOutput, error)
	PutImageRecipePolicyRequest(*imagebuilder.PutImageRecipePolicyInput) (*request.Request, *imagebuilder.PutImageRecipePolicyOutput)

	StartImagePipelineExecution(*imagebuilder.StartImagePipelineExecutionInput) (*imagebuilder.StartImagePipelineExecutionOutput, error)
	StartImagePipelineExecutionWithContext(aws.Context, *imagebuilder.StartImagePipelineExecutionInput, ...request.Option) (*imagebuilder.StartImagePipelineExecutionOutput, error)
	StartImagePipelineExecutionRequest(*imagebuilder.StartImagePipelineExecutionInput) (*request.Request, *imagebuilder.StartImagePipelineExecutionOutput)

	TagResource(*imagebuilder.TagResourceInput) (*imagebuilder.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *imagebuilder.TagResourceInput, ...request.Option) (*imagebuilder.TagResourceOutput, error)
	TagResourceRequest(*imagebuilder.TagResourceInput) (*request.Request, *imagebuilder.TagResourceOutput)

	UntagResource(*imagebuilder.UntagResourceInput) (*imagebuilder.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *imagebuilder.UntagResourceInput, ...request.Option) (*imagebuilder.UntagResourceOutput, error)
	UntagResourceRequest(*imagebuilder.UntagResourceInput) (*request.Request, *imagebuilder.UntagResourceOutput)

	UpdateDistributionConfiguration(*imagebuilder.UpdateDistributionConfigurationInput) (*imagebuilder.UpdateDistributionConfigurationOutput, error)
	UpdateDistributionConfigurationWithContext(aws.Context, *imagebuilder.UpdateDistributionConfigurationInput, ...request.Option) (*imagebuilder.UpdateDistributionConfigurationOutput, error)
	UpdateDistributionConfigurationRequest(*imagebuilder.UpdateDistributionConfigurationInput) (*request.Request, *imagebuilder.UpdateDistributionConfigurationOutput)

	UpdateImagePipeline(*imagebuilder.UpdateImagePipelineInput) (*imagebuilder.UpdateImagePipelineOutput, error)
	UpdateImagePipelineWithContext(aws.Context, *imagebuilder.UpdateImagePipelineInput, ...request.Option) (*imagebuilder.UpdateImagePipelineOutput, error)
	UpdateImagePipelineRequest(*imagebuilder.UpdateImagePipelineInput) (*request.Request, *imagebuilder.UpdateImagePipelineOutput)

	UpdateInfrastructureConfiguration(*imagebuilder.UpdateInfrastructureConfigurationInput) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error)
	UpdateInfrastructureConfigurationWithContext(aws.Context, *imagebuilder.UpdateInfrastructureConfigurationInput, ...request.Option) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error)
	UpdateInfrastructureConfigurationRequest(*imagebuilder.UpdateInfrastructureConfigurationInput) (*request.Request, *imagebuilder.UpdateInfrastructureConfigurationOutput)
}

var _ ImagebuilderAPI = (*imagebuilder.Imagebuilder)(nil)
