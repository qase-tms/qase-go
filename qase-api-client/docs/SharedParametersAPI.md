# \SharedParametersAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSharedParameter**](SharedParametersAPI.md#CreateSharedParameter) | **Post** /shared_parameter | Create a new shared parameter
[**DeleteSharedParameter**](SharedParametersAPI.md#DeleteSharedParameter) | **Delete** /shared_parameter/{id} | Delete shared parameter
[**GetSharedParameter**](SharedParametersAPI.md#GetSharedParameter) | **Get** /shared_parameter/{id} | Get a specific shared parameter
[**GetSharedParameters**](SharedParametersAPI.md#GetSharedParameters) | **Get** /shared_parameter | Get all shared parameters
[**UpdateSharedParameter**](SharedParametersAPI.md#UpdateSharedParameter) | **Patch** /shared_parameter/{id} | Update shared parameter



## CreateSharedParameter

> UuidResponse CreateSharedParameter(ctx).SharedParameterCreate(sharedParameterCreate).Execute()

Create a new shared parameter

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
	sharedParameterCreate := *openapiclient.NewSharedParameterCreate("Title_example", "Type_example", false, openapiclient.SharedParameterParameter{ArrayOfParameterSingle: new([]ParameterSingle)}) // SharedParameterCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedParametersAPI.CreateSharedParameter(context.Background()).SharedParameterCreate(sharedParameterCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedParametersAPI.CreateSharedParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSharedParameter`: UuidResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedParametersAPI.CreateSharedParameter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSharedParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sharedParameterCreate** | [**SharedParameterCreate**](SharedParameterCreate.md) |  | 

### Return type

[**UuidResponse**](UuidResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSharedParameter

> UuidResponse1 DeleteSharedParameter(ctx, id).Execute()

Delete shared parameter



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedParametersAPI.DeleteSharedParameter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedParametersAPI.DeleteSharedParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSharedParameter`: UuidResponse1
	fmt.Fprintf(os.Stdout, "Response from `SharedParametersAPI.DeleteSharedParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharedParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UuidResponse1**](UuidResponse1.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedParameter

> SharedParameterResponse GetSharedParameter(ctx, id).Execute()

Get a specific shared parameter

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedParametersAPI.GetSharedParameter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedParametersAPI.GetSharedParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedParameter`: SharedParameterResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedParametersAPI.GetSharedParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SharedParameterResponse**](SharedParameterResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedParameters

> SharedParameterListResponse GetSharedParameters(ctx).Limit(limit).Offset(offset).FiltersSearch(filtersSearch).FiltersType(filtersType).FiltersProjectCodes(filtersProjectCodes).Execute()

Get all shared parameters

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
	limit := int32(56) // int32 | A number of entities in result set. (optional) (default to 10)
	offset := int32(56) // int32 | How many entities should be skipped. (optional) (default to 0)
	filtersSearch := "filtersSearch_example" // string |  (optional)
	filtersType := "filtersType_example" // string |  (optional)
	filtersProjectCodes := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedParametersAPI.GetSharedParameters(context.Background()).Limit(limit).Offset(offset).FiltersSearch(filtersSearch).FiltersType(filtersType).FiltersProjectCodes(filtersProjectCodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedParametersAPI.GetSharedParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedParameters`: SharedParameterListResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedParametersAPI.GetSharedParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]
 **filtersSearch** | **string** |  | 
 **filtersType** | **string** |  | 
 **filtersProjectCodes** | **[]string** |  | 

### Return type

[**SharedParameterListResponse**](SharedParameterListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSharedParameter

> UuidResponse1 UpdateSharedParameter(ctx, id).SharedParameterUpdate(sharedParameterUpdate).Execute()

Update shared parameter

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifier.
	sharedParameterUpdate := *openapiclient.NewSharedParameterUpdate() // SharedParameterUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedParametersAPI.UpdateSharedParameter(context.Background(), id).SharedParameterUpdate(sharedParameterUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedParametersAPI.UpdateSharedParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSharedParameter`: UuidResponse1
	fmt.Fprintf(os.Stdout, "Response from `SharedParametersAPI.UpdateSharedParameter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSharedParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharedParameterUpdate** | [**SharedParameterUpdate**](SharedParameterUpdate.md) |  | 

### Return type

[**UuidResponse1**](UuidResponse1.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

