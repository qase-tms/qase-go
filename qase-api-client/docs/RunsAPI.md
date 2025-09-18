# \RunsAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteRun**](RunsAPI.md#CompleteRun) | **Post** /run/{code}/{id}/complete | Complete a specific run
[**CreateRun**](RunsAPI.md#CreateRun) | **Post** /run/{code} | Create a new run
[**DeleteRun**](RunsAPI.md#DeleteRun) | **Delete** /run/{code}/{id} | Delete run
[**GetRun**](RunsAPI.md#GetRun) | **Get** /run/{code}/{id} | Get a specific run
[**GetRuns**](RunsAPI.md#GetRuns) | **Get** /run/{code} | Get all runs
[**RunUpdateExternalIssue**](RunsAPI.md#RunUpdateExternalIssue) | **Post** /run/{code}/external-issue | Update external issues for runs
[**UpdateRun**](RunsAPI.md#UpdateRun) | **Patch** /run/{code}/{id} | Update a specific run
[**UpdateRunPublicity**](RunsAPI.md#UpdateRunPublicity) | **Patch** /run/{code}/{id}/public | Update publicity of a specific run



## CompleteRun

> BaseResponse CompleteRun(ctx, code, id).Execute()

Complete a specific run



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
	resp, r, err := apiClient.RunsAPI.CompleteRun(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.CompleteRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteRun`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.CompleteRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRun

> IdResponse CreateRun(ctx, code).RunCreate(runCreate).Execute()

Create a new run



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
	runCreate := *openapiclient.NewRunCreate("Title_example") // RunCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPI.CreateRun(context.Background(), code).RunCreate(runCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.CreateRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRun`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.CreateRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runCreate** | [**RunCreate**](RunCreate.md) |  | 

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


## DeleteRun

> IdResponse DeleteRun(ctx, code, id).Execute()

Delete run



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
	resp, r, err := apiClient.RunsAPI.DeleteRun(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.DeleteRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRun`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.DeleteRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRunRequest struct via the builder pattern


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


## GetRun

> RunResponse GetRun(ctx, code, id).Include(include).Execute()

Get a specific run



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
	include := "include_example" // string | Include a list of related entities IDs into response. Should be separated by comma. Possible values: cases, defects, external_issue  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPI.GetRun(context.Background(), code, id).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.GetRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRun`: RunResponse
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.GetRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **string** | Include a list of related entities IDs into response. Should be separated by comma. Possible values: cases, defects, external_issue  | 

### Return type

[**RunResponse**](RunResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuns

> RunListResponse GetRuns(ctx, code).Search(search).Status(status).Milestone(milestone).Environment(environment).FromStartTime(fromStartTime).ToStartTime(toStartTime).Limit(limit).Offset(offset).Include(include).Execute()

Get all runs



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
	search := "search_example" // string |  (optional)
	status := "status_example" // string | A list of status values separated by comma. Possible values: in_progress, passed, failed, aborted, active (deprecated), complete (deprecated), abort (deprecated).  (optional)
	milestone := int32(56) // int32 |  (optional)
	environment := int32(56) // int32 |  (optional)
	fromStartTime := int64(789) // int64 |  (optional)
	toStartTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | A number of entities in result set. (optional) (default to 10)
	offset := int32(56) // int32 | How many entities should be skipped. (optional) (default to 0)
	include := "include_example" // string | Include a list of related entities IDs into response. Should be separated by comma. Possible values: cases, defects, external_issue  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPI.GetRuns(context.Background(), code).Search(search).Status(status).Milestone(milestone).Environment(environment).FromStartTime(fromStartTime).ToStartTime(toStartTime).Limit(limit).Offset(offset).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.GetRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuns`: RunListResponse
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.GetRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **status** | **string** | A list of status values separated by comma. Possible values: in_progress, passed, failed, aborted, active (deprecated), complete (deprecated), abort (deprecated).  | 
 **milestone** | **int32** |  | 
 **environment** | **int32** |  | 
 **fromStartTime** | **int64** |  | 
 **toStartTime** | **int64** |  | 
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]
 **include** | **string** | Include a list of related entities IDs into response. Should be separated by comma. Possible values: cases, defects, external_issue  | 

### Return type

[**RunListResponse**](RunListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunUpdateExternalIssue

> RunUpdateExternalIssue(ctx, code).RunExternalIssues(runExternalIssues).Execute()

Update external issues for runs



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
	runExternalIssues := *openapiclient.NewRunExternalIssues("Type_example", []openapiclient.RunExternalIssuesLinksInner{*openapiclient.NewRunExternalIssuesLinksInner(int64(123))}) // RunExternalIssues | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunsAPI.RunUpdateExternalIssue(context.Background(), code).RunExternalIssues(runExternalIssues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.RunUpdateExternalIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunUpdateExternalIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runExternalIssues** | [**RunExternalIssues**](RunExternalIssues.md) |  | 

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


## UpdateRun

> BaseResponse UpdateRun(ctx, code, id).Runupdate(runupdate).Execute()

Update a specific run



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
	runupdate := *openapiclient.NewRunupdate() // Runupdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPI.UpdateRun(context.Background(), code, id).Runupdate(runupdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.UpdateRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRun`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.UpdateRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **runupdate** | [**Runupdate**](Runupdate.md) |  | 

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


## UpdateRunPublicity

> RunPublicResponse UpdateRunPublicity(ctx, code, id).RunPublic(runPublic).Execute()

Update publicity of a specific run



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
	runPublic := *openapiclient.NewRunPublic(false) // RunPublic | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPI.UpdateRunPublicity(context.Background(), code, id).RunPublic(runPublic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPI.UpdateRunPublicity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRunPublicity`: RunPublicResponse
	fmt.Fprintf(os.Stdout, "Response from `RunsAPI.UpdateRunPublicity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRunPublicityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **runPublic** | [**RunPublic**](RunPublic.md) |  | 

### Return type

[**RunPublicResponse**](RunPublicResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

