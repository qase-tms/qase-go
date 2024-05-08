# \ResultsAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResult**](ResultsAPI.md#CreateResult) | **Post** /result/{code}/{id} | Create test run result
[**CreateResultBulk**](ResultsAPI.md#CreateResultBulk) | **Post** /result/{code}/{id}/bulk | Bulk create test run result
[**DeleteResult**](ResultsAPI.md#DeleteResult) | **Delete** /result/{code}/{id}/{hash} | Delete test run result
[**GetResult**](ResultsAPI.md#GetResult) | **Get** /result/{code}/{hash} | Get test run result by code
[**GetResults**](ResultsAPI.md#GetResults) | **Get** /result/{code} | Get all test run results
[**UpdateResult**](ResultsAPI.md#UpdateResult) | **Patch** /result/{code}/{id}/{hash} | Update test run result



## CreateResult

> ResultCreateResponse CreateResult(ctx, code, id).ResultCreate(resultCreate).Execute()

Create test run result



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
	resultCreate := *openapiclient.NewResultCreate("Status_example") // ResultCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultsAPI.CreateResult(context.Background(), code, id).ResultCreate(resultCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultsAPI.CreateResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResult`: ResultCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ResultsAPI.CreateResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resultCreate** | [**ResultCreate**](ResultCreate.md) |  | 

### Return type

[**ResultCreateResponse**](ResultCreateResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResultBulk

> BaseResponse CreateResultBulk(ctx, code, id).ResultcreateBulk(resultcreateBulk).Execute()

Bulk create test run result



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
	resultcreateBulk := *openapiclient.NewResultcreateBulk([]openapiclient.ResultCreate{*openapiclient.NewResultCreate("Status_example")}) // ResultcreateBulk | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultsAPI.CreateResultBulk(context.Background(), code, id).ResultcreateBulk(resultcreateBulk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultsAPI.CreateResultBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResultBulk`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ResultsAPI.CreateResultBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResultBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resultcreateBulk** | [**ResultcreateBulk**](ResultcreateBulk.md) |  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResult

> HashResponse DeleteResult(ctx, code, id, hash).Execute()

Delete test run result



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
	hash := "hash_example" // string | Hash.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultsAPI.DeleteResult(context.Background(), code, id, hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultsAPI.DeleteResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteResult`: HashResponse
	fmt.Fprintf(os.Stdout, "Response from `ResultsAPI.DeleteResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 
**hash** | **string** | Hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResultRequest struct via the builder pattern


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


## GetResult

> ResultResponse GetResult(ctx, code, hash).Execute()

Get test run result by code



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
	resp, r, err := apiClient.ResultsAPI.GetResult(context.Background(), code, hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultsAPI.GetResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResult`: ResultResponse
	fmt.Fprintf(os.Stdout, "Response from `ResultsAPI.GetResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**hash** | **string** | Hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResultResponse**](ResultResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResults

> ResultListResponse GetResults(ctx, code).Status(status).Run(run).CaseId(caseId).Member(member).Api(api).FromEndTime(fromEndTime).ToEndTime(toEndTime).Limit(limit).Offset(offset).Execute()

Get all test run results



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
	status := "status_example" // string | A single test run result status. Possible values: in_progress, passed, failed, blocked, skipped, invalid.  (optional)
	run := "run_example" // string | A list of run IDs separated by comma. (optional)
	caseId := "caseId_example" // string | A list of case IDs separated by comma. (optional)
	member := "member_example" // string | A list of member IDs separated by comma. (optional)
	api := true // bool |  (optional)
	fromEndTime := "fromEndTime_example" // string | Will return all results created after provided datetime. Allowed format: `Y-m-d H:i:s`.  (optional)
	toEndTime := "toEndTime_example" // string | Will return all results created before provided datetime. Allowed format: `Y-m-d H:i:s`.  (optional)
	limit := int32(56) // int32 | A number of entities in result set. (optional) (default to 10)
	offset := int32(56) // int32 | How many entities should be skipped. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultsAPI.GetResults(context.Background(), code).Status(status).Run(run).CaseId(caseId).Member(member).Api(api).FromEndTime(fromEndTime).ToEndTime(toEndTime).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultsAPI.GetResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResults`: ResultListResponse
	fmt.Fprintf(os.Stdout, "Response from `ResultsAPI.GetResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | A single test run result status. Possible values: in_progress, passed, failed, blocked, skipped, invalid.  | 
 **run** | **string** | A list of run IDs separated by comma. | 
 **caseId** | **string** | A list of case IDs separated by comma. | 
 **member** | **string** | A list of member IDs separated by comma. | 
 **api** | **bool** |  | 
 **fromEndTime** | **string** | Will return all results created after provided datetime. Allowed format: &#x60;Y-m-d H:i:s&#x60;.  | 
 **toEndTime** | **string** | Will return all results created before provided datetime. Allowed format: &#x60;Y-m-d H:i:s&#x60;.  | 
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]

### Return type

[**ResultListResponse**](ResultListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResult

> HashResponse UpdateResult(ctx, code, id, hash).ResultUpdate(resultUpdate).Execute()

Update test run result



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
	hash := "hash_example" // string | Hash.
	resultUpdate := *openapiclient.NewResultUpdate() // ResultUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultsAPI.UpdateResult(context.Background(), code, id, hash).ResultUpdate(resultUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultsAPI.UpdateResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResult`: HashResponse
	fmt.Fprintf(os.Stdout, "Response from `ResultsAPI.UpdateResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 
**hash** | **string** | Hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resultUpdate** | [**ResultUpdate**](ResultUpdate.md) |  | 

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

