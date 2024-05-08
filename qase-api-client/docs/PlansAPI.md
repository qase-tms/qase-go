# \PlansAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlan**](PlansAPI.md#CreatePlan) | **Post** /plan/{code} | Create a new plan
[**DeletePlan**](PlansAPI.md#DeletePlan) | **Delete** /plan/{code}/{id} | Delete plan
[**GetPlan**](PlansAPI.md#GetPlan) | **Get** /plan/{code}/{id} | Get a specific plan
[**GetPlans**](PlansAPI.md#GetPlans) | **Get** /plan/{code} | Get all plans
[**UpdatePlan**](PlansAPI.md#UpdatePlan) | **Patch** /plan/{code}/{id} | Update plan



## CreatePlan

> IdResponse CreatePlan(ctx, code).PlanCreate(planCreate).Execute()

Create a new plan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	code := "code_example" // string | Code of project, where to search entities.
	planCreate := *openapiclient.NewPlanCreate("Title_example", []int64{int64(123)}) // PlanCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.CreatePlan(context.Background(), code).PlanCreate(planCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CreatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlan`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CreatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **planCreate** | [**PlanCreate**](PlanCreate.md) |  | 

### Return type

[**IdResponse**](IdResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlan

> IdResponse DeletePlan(ctx, code, id).Execute()

Delete plan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	code := "code_example" // string | Code of project, where to search entities.
	id := int32(56) // int32 | Identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.DeletePlan(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.DeletePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePlan`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.DeletePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdResponse**](IdResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlan

> PlanResponse GetPlan(ctx, code, id).Execute()

Get a specific plan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	code := "code_example" // string | Code of project, where to search entities.
	id := int32(56) // int32 | Identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.GetPlan(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlan`: PlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlanResponse**](PlanResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlans

> PlanListResponse GetPlans(ctx, code).Limit(limit).Offset(offset).Execute()

Get all plans



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	code := "code_example" // string | Code of project, where to search entities.
	limit := int32(56) // int32 | A number of entities in result set. (optional) (default to 10)
	offset := int32(56) // int32 | How many entities should be skipped. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.GetPlans(context.Background(), code).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlans`: PlanListResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]

### Return type

[**PlanListResponse**](PlanListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlan

> IdResponse UpdatePlan(ctx, code, id).PlanUpdate(planUpdate).Execute()

Update plan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	code := "code_example" // string | Code of project, where to search entities.
	id := int32(56) // int32 | Identifier.
	planUpdate := *openapiclient.NewPlanUpdate() // PlanUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.UpdatePlan(context.Background(), code, id).PlanUpdate(planUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.UpdatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlan`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.UpdatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **planUpdate** | [**PlanUpdate**](PlanUpdate.md) |  | 

### Return type

[**IdResponse**](IdResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

