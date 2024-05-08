# \CasesAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Bulk**](CasesAPI.md#Bulk) | **Post** /case/{code}/bulk | Create test cases in bulk
[**CaseAttachExternalIssue**](CasesAPI.md#CaseAttachExternalIssue) | **Post** /case/{code}/external-issue/attach | Attach the external issues to the test cases
[**CaseDetachExternalIssue**](CasesAPI.md#CaseDetachExternalIssue) | **Post** /case/{code}/external-issue/detach | Detach the external issues from the test cases
[**CreateCase**](CasesAPI.md#CreateCase) | **Post** /case/{code} | Create a new test case
[**DeleteCase**](CasesAPI.md#DeleteCase) | **Delete** /case/{code}/{id} | Delete test case
[**GetCase**](CasesAPI.md#GetCase) | **Get** /case/{code}/{id} | Get a specific test case
[**GetCases**](CasesAPI.md#GetCases) | **Get** /case/{code} | Get all test cases
[**UpdateCase**](CasesAPI.md#UpdateCase) | **Patch** /case/{code}/{id} | Update test case



## Bulk

> Bulk200Response Bulk(ctx, code).TestCasebulk(testCasebulk).Execute()

Create test cases in bulk



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
	testCasebulk := *openapiclient.NewTestCasebulk([]openapiclient.TestCasebulkCasesInner{*openapiclient.NewTestCasebulkCasesInner("Title_example")}) // TestCasebulk | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.Bulk(context.Background(), code).TestCasebulk(testCasebulk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.Bulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Bulk`: Bulk200Response
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.Bulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testCasebulk** | [**TestCasebulk**](TestCasebulk.md) |  | 

### Return type

[**Bulk200Response**](Bulk200Response.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CaseAttachExternalIssue

> BaseResponse CaseAttachExternalIssue(ctx, code).TestCaseexternalIssues(testCaseexternalIssues).Execute()

Attach the external issues to the test cases

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
	testCaseexternalIssues := *openapiclient.NewTestCaseexternalIssues("Type_example", []openapiclient.TestCaseExternalIssuesLinksInner{*openapiclient.NewTestCaseExternalIssuesLinksInner(int64(123), []string{"ExternalIssues_example"})}) // TestCaseexternalIssues | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.CaseAttachExternalIssue(context.Background(), code).TestCaseexternalIssues(testCaseexternalIssues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.CaseAttachExternalIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaseAttachExternalIssue`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.CaseAttachExternalIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCaseAttachExternalIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testCaseexternalIssues** | [**TestCaseexternalIssues**](TestCaseexternalIssues.md) |  | 

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


## CaseDetachExternalIssue

> BaseResponse CaseDetachExternalIssue(ctx, code).TestCaseexternalIssues(testCaseexternalIssues).Execute()

Detach the external issues from the test cases

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
	testCaseexternalIssues := *openapiclient.NewTestCaseexternalIssues("Type_example", []openapiclient.TestCaseExternalIssuesLinksInner{*openapiclient.NewTestCaseExternalIssuesLinksInner(int64(123), []string{"ExternalIssues_example"})}) // TestCaseexternalIssues | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.CaseDetachExternalIssue(context.Background(), code).TestCaseexternalIssues(testCaseexternalIssues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.CaseDetachExternalIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaseDetachExternalIssue`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.CaseDetachExternalIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCaseDetachExternalIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testCaseexternalIssues** | [**TestCaseexternalIssues**](TestCaseexternalIssues.md) |  | 

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


## CreateCase

> IdResponse CreateCase(ctx, code).TestCaseCreate(testCaseCreate).Execute()

Create a new test case



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
	testCaseCreate := *openapiclient.NewTestCaseCreate("Title_example") // TestCaseCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.CreateCase(context.Background(), code).TestCaseCreate(testCaseCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.CreateCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCase`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.CreateCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testCaseCreate** | [**TestCaseCreate**](TestCaseCreate.md) |  | 

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


## DeleteCase

> IdResponse DeleteCase(ctx, code, id).Execute()

Delete test case



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
	resp, r, err := apiClient.CasesAPI.DeleteCase(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.DeleteCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCase`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.DeleteCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaseRequest struct via the builder pattern


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


## GetCase

> TestCaseResponse GetCase(ctx, code, id).Execute()

Get a specific test case



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
	resp, r, err := apiClient.CasesAPI.GetCase(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.GetCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCase`: TestCaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.GetCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TestCaseResponse**](TestCaseResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCases

> TestCaseListResponse GetCases(ctx, code).Search(search).MilestoneId(milestoneId).SuiteId(suiteId).Severity(severity).Priority(priority).Type_(type_).Behavior(behavior).Automation(automation).Status(status).ExternalIssuesType(externalIssuesType).ExternalIssuesIds(externalIssuesIds).Include(include).Limit(limit).Offset(offset).Execute()

Get all test cases



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
	milestoneId := int32(56) // int32 | ID of milestone. (optional)
	suiteId := int32(56) // int32 | ID of test suite. (optional)
	severity := "severity_example" // string | A list of severity values separated by comma. Possible values: undefined, blocker, critical, major, normal, minor, trivial  (optional)
	priority := "priority_example" // string | A list of priority values separated by comma. Possible values: undefined, high, medium, low  (optional)
	type_ := "type__example" // string | A list of type values separated by comma. Possible values: other, functional smoke, regression, security, usability, performance, acceptance  (optional)
	behavior := "behavior_example" // string | A list of behavior values separated by comma. Possible values: undefined, positive negative, destructive  (optional)
	automation := "automation_example" // string | A list of values separated by comma. Possible values: is-not-automated, automated to-be-automated  (optional)
	status := "status_example" // string | A list of values separated by comma. Possible values: actual, draft deprecated  (optional)
	externalIssuesType := "externalIssuesType_example" // string | An integration type.  (optional)
	externalIssuesIds := []string{"Inner_example"} // []string | A list of issue IDs. (optional)
	include := "include_example" // string | A list of entities to include in response separated by comma. Possible values: external_issues.  (optional)
	limit := int32(56) // int32 | A number of entities in result set. (optional) (default to 10)
	offset := int32(56) // int32 | How many entities should be skipped. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.GetCases(context.Background(), code).Search(search).MilestoneId(milestoneId).SuiteId(suiteId).Severity(severity).Priority(priority).Type_(type_).Behavior(behavior).Automation(automation).Status(status).ExternalIssuesType(externalIssuesType).ExternalIssuesIds(externalIssuesIds).Include(include).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.GetCases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCases`: TestCaseListResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.GetCases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Provide a string that will be used to search by name. | 
 **milestoneId** | **int32** | ID of milestone. | 
 **suiteId** | **int32** | ID of test suite. | 
 **severity** | **string** | A list of severity values separated by comma. Possible values: undefined, blocker, critical, major, normal, minor, trivial  | 
 **priority** | **string** | A list of priority values separated by comma. Possible values: undefined, high, medium, low  | 
 **type_** | **string** | A list of type values separated by comma. Possible values: other, functional smoke, regression, security, usability, performance, acceptance  | 
 **behavior** | **string** | A list of behavior values separated by comma. Possible values: undefined, positive negative, destructive  | 
 **automation** | **string** | A list of values separated by comma. Possible values: is-not-automated, automated to-be-automated  | 
 **status** | **string** | A list of values separated by comma. Possible values: actual, draft deprecated  | 
 **externalIssuesType** | **string** | An integration type.  | 
 **externalIssuesIds** | **[]string** | A list of issue IDs. | 
 **include** | **string** | A list of entities to include in response separated by comma. Possible values: external_issues.  | 
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]

### Return type

[**TestCaseListResponse**](TestCaseListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCase

> IdResponse UpdateCase(ctx, code, id).TestCaseUpdate(testCaseUpdate).Execute()

Update test case



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
	testCaseUpdate := *openapiclient.NewTestCaseUpdate() // TestCaseUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CasesAPI.UpdateCase(context.Background(), code, id).TestCaseUpdate(testCaseUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CasesAPI.UpdateCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCase`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `CasesAPI.UpdateCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testCaseUpdate** | [**TestCaseUpdate**](TestCaseUpdate.md) |  | 

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

