# \MilestonesAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMilestone**](MilestonesAPI.md#CreateMilestone) | **Post** /milestone/{code} | Create a new milestone
[**DeleteMilestone**](MilestonesAPI.md#DeleteMilestone) | **Delete** /milestone/{code}/{id} | Delete milestone
[**GetMilestone**](MilestonesAPI.md#GetMilestone) | **Get** /milestone/{code}/{id} | Get a specific milestone
[**GetMilestones**](MilestonesAPI.md#GetMilestones) | **Get** /milestone/{code} | Get all milestones
[**UpdateMilestone**](MilestonesAPI.md#UpdateMilestone) | **Patch** /milestone/{code}/{id} | Update milestone



## CreateMilestone

> IdResponse CreateMilestone(ctx, code).MilestoneCreate(milestoneCreate).Execute()

Create a new milestone



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
	milestoneCreate := *openapiclient.NewMilestoneCreate("Title_example") // MilestoneCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilestonesAPI.CreateMilestone(context.Background(), code).MilestoneCreate(milestoneCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.CreateMilestone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMilestone`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.CreateMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **milestoneCreate** | [**MilestoneCreate**](MilestoneCreate.md) |  | 

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


## DeleteMilestone

> IdResponse DeleteMilestone(ctx, code, id).Execute()

Delete milestone



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
	resp, r, err := apiClient.MilestonesAPI.DeleteMilestone(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.DeleteMilestone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMilestone`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.DeleteMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMilestoneRequest struct via the builder pattern


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


## GetMilestone

> MilestoneResponse GetMilestone(ctx, code, id).Execute()

Get a specific milestone



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
	resp, r, err := apiClient.MilestonesAPI.GetMilestone(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.GetMilestone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMilestone`: MilestoneResponse
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.GetMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MilestoneResponse**](MilestoneResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMilestones

> MilestoneListResponse GetMilestones(ctx, code).Search(search).Limit(limit).Offset(offset).Execute()

Get all milestones



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
	resp, r, err := apiClient.MilestonesAPI.GetMilestones(context.Background(), code).Search(search).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.GetMilestones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMilestones`: MilestoneListResponse
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.GetMilestones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMilestonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Provide a string that will be used to search by name. | 
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]

### Return type

[**MilestoneListResponse**](MilestoneListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMilestone

> IdResponse UpdateMilestone(ctx, code, id).MilestoneUpdate(milestoneUpdate).Execute()

Update milestone



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
	milestoneUpdate := *openapiclient.NewMilestoneUpdate() // MilestoneUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilestonesAPI.UpdateMilestone(context.Background(), code, id).MilestoneUpdate(milestoneUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.UpdateMilestone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMilestone`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.UpdateMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **milestoneUpdate** | [**MilestoneUpdate**](MilestoneUpdate.md) |  | 

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

