# \ResultsAPI

All URIs are relative to *https://api.qase.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResultV2**](ResultsAPI.md#CreateResultV2) | **Post** /{project_code}/run/{run_id}/result | Create test run result
[**CreateResultsV2**](ResultsAPI.md#CreateResultsV2) | **Post** /{project_code}/run/{run_id}/results | Bulk create test run result



## CreateResultV2

> CreateResultV2(ctx, projectCode, runId).ResultCreate(resultCreate).Execute()

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
	projectCode := "projectCode_example" // string | 
	runId := int64(789) // int64 | 
	resultCreate := *openapiclient.NewResultCreate("Title_example", *openapiclient.NewResultExecution("Status_example")) // ResultCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResultsAPI.CreateResultV2(context.Background(), projectCode, runId).ResultCreate(resultCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultsAPI.CreateResultV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectCode** | **string** |  | 
**runId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResultV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resultCreate** | [**ResultCreate**](ResultCreate.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResultsV2

> CreateResultsV2(ctx, projectCode, runId).CreateResultsRequestV2(createResultsRequestV2).Execute()

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
	projectCode := "projectCode_example" // string | 
	runId := int64(789) // int64 | 
	createResultsRequestV2 := *openapiclient.NewCreateResultsRequestV2() // CreateResultsRequestV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResultsAPI.CreateResultsV2(context.Background(), projectCode, runId).CreateResultsRequestV2(createResultsRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultsAPI.CreateResultsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectCode** | **string** |  | 
**runId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResultsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createResultsRequestV2** | [**CreateResultsRequestV2**](CreateResultsRequestV2.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

