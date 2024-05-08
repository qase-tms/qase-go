# \SuitesAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSuite**](SuitesAPI.md#CreateSuite) | **Post** /suite/{code} | Create a new test suite
[**DeleteSuite**](SuitesAPI.md#DeleteSuite) | **Delete** /suite/{code}/{id} | Delete test suite
[**GetSuite**](SuitesAPI.md#GetSuite) | **Get** /suite/{code}/{id} | Get a specific test suite
[**GetSuites**](SuitesAPI.md#GetSuites) | **Get** /suite/{code} | Get all test suites
[**UpdateSuite**](SuitesAPI.md#UpdateSuite) | **Patch** /suite/{code}/{id} | Update test suite



## CreateSuite

> IdResponse CreateSuite(ctx, code).SuiteCreate(suiteCreate).Execute()

Create a new test suite



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
	suiteCreate := *openapiclient.NewSuiteCreate("Title_example") // SuiteCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuitesAPI.CreateSuite(context.Background(), code).SuiteCreate(suiteCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuitesAPI.CreateSuite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSuite`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `SuitesAPI.CreateSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSuiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suiteCreate** | [**SuiteCreate**](SuiteCreate.md) |  | 

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


## DeleteSuite

> IdResponse DeleteSuite(ctx, code, id).SuiteDelete(suiteDelete).Execute()

Delete test suite



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
	suiteDelete := *openapiclient.NewSuiteDelete() // SuiteDelete |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuitesAPI.DeleteSuite(context.Background(), code, id).SuiteDelete(suiteDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuitesAPI.DeleteSuite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSuite`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `SuitesAPI.DeleteSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSuiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suiteDelete** | [**SuiteDelete**](SuiteDelete.md) |  | 

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


## GetSuite

> SuiteResponse GetSuite(ctx, code, id).Execute()

Get a specific test suite



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
	resp, r, err := apiClient.SuitesAPI.GetSuite(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuitesAPI.GetSuite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuite`: SuiteResponse
	fmt.Fprintf(os.Stdout, "Response from `SuitesAPI.GetSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SuiteResponse**](SuiteResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuites

> SuiteListResponse GetSuites(ctx, code).Search(search).Limit(limit).Offset(offset).Execute()

Get all test suites



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
	resp, r, err := apiClient.SuitesAPI.GetSuites(context.Background(), code).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuitesAPI.GetSuites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuites`: SuiteListResponse
	fmt.Fprintf(os.Stdout, "Response from `SuitesAPI.GetSuites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Provide a string that will be used to search by name. | 
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]

### Return type

[**SuiteListResponse**](SuiteListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSuite

> IdResponse UpdateSuite(ctx, code, id).SuiteUpdate(suiteUpdate).Execute()

Update test suite



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
	suiteUpdate := *openapiclient.NewSuiteUpdate() // SuiteUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuitesAPI.UpdateSuite(context.Background(), code, id).SuiteUpdate(suiteUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuitesAPI.UpdateSuite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSuite`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `SuitesAPI.UpdateSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suiteUpdate** | [**SuiteUpdate**](SuiteUpdate.md) |  | 

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

