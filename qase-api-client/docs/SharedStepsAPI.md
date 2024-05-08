# \SharedStepsAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSharedStep**](SharedStepsAPI.md#CreateSharedStep) | **Post** /shared_step/{code} | Create a new shared step
[**DeleteSharedStep**](SharedStepsAPI.md#DeleteSharedStep) | **Delete** /shared_step/{code}/{hash} | Delete shared step
[**GetSharedStep**](SharedStepsAPI.md#GetSharedStep) | **Get** /shared_step/{code}/{hash} | Get a specific shared step
[**GetSharedSteps**](SharedStepsAPI.md#GetSharedSteps) | **Get** /shared_step/{code} | Get all shared steps
[**UpdateSharedStep**](SharedStepsAPI.md#UpdateSharedStep) | **Patch** /shared_step/{code}/{hash} | Update shared step



## CreateSharedStep

> HashResponse CreateSharedStep(ctx, code).SharedStepCreate(sharedStepCreate).Execute()

Create a new shared step



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
	sharedStepCreate := *openapiclient.NewSharedStepCreate("Title_example") // SharedStepCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedStepsAPI.CreateSharedStep(context.Background(), code).SharedStepCreate(sharedStepCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedStepsAPI.CreateSharedStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSharedStep`: HashResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedStepsAPI.CreateSharedStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSharedStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharedStepCreate** | [**SharedStepCreate**](SharedStepCreate.md) |  | 

### Return type

[**HashResponse**](HashResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSharedStep

> HashResponse DeleteSharedStep(ctx, code, hash).Execute()

Delete shared step



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
	hash := "hash_example" // string | Hash.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedStepsAPI.DeleteSharedStep(context.Background(), code, hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedStepsAPI.DeleteSharedStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSharedStep`: HashResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedStepsAPI.DeleteSharedStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**hash** | **string** | Hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharedStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HashResponse**](HashResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedStep

> SharedStepResponse GetSharedStep(ctx, code, hash).Execute()

Get a specific shared step



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
	hash := "hash_example" // string | Hash.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedStepsAPI.GetSharedStep(context.Background(), code, hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedStepsAPI.GetSharedStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedStep`: SharedStepResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedStepsAPI.GetSharedStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**hash** | **string** | Hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SharedStepResponse**](SharedStepResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedSteps

> SharedStepListResponse GetSharedSteps(ctx, code).Search(search).Limit(limit).Offset(offset).Execute()

Get all shared steps



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
	search := "search_example" // string | Provide a string that will be used to search by name. (optional)
	limit := int32(56) // int32 | A number of entities in result set. (optional) (default to 10)
	offset := int32(56) // int32 | How many entities should be skipped. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedStepsAPI.GetSharedSteps(context.Background(), code).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedStepsAPI.GetSharedSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedSteps`: SharedStepListResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedStepsAPI.GetSharedSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Provide a string that will be used to search by name. | 
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]

### Return type

[**SharedStepListResponse**](SharedStepListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSharedStep

> HashResponse UpdateSharedStep(ctx, code, hash).SharedStepUpdate(sharedStepUpdate).Execute()

Update shared step



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
	hash := "hash_example" // string | Hash.
	sharedStepUpdate := *openapiclient.NewSharedStepUpdate("Title_example") // SharedStepUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedStepsAPI.UpdateSharedStep(context.Background(), code, hash).SharedStepUpdate(sharedStepUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedStepsAPI.UpdateSharedStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSharedStep`: HashResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedStepsAPI.UpdateSharedStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**hash** | **string** | Hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSharedStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sharedStepUpdate** | [**SharedStepUpdate**](SharedStepUpdate.md) |  | 

### Return type

[**HashResponse**](HashResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

