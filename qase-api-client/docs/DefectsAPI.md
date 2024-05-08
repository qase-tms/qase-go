# \DefectsAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefect**](DefectsAPI.md#CreateDefect) | **Post** /defect/{code} | Create a new defect
[**DeleteDefect**](DefectsAPI.md#DeleteDefect) | **Delete** /defect/{code}/{id} | Delete defect
[**GetDefect**](DefectsAPI.md#GetDefect) | **Get** /defect/{code}/{id} | Get a specific defect
[**GetDefects**](DefectsAPI.md#GetDefects) | **Get** /defect/{code} | Get all defects
[**ResolveDefect**](DefectsAPI.md#ResolveDefect) | **Patch** /defect/{code}/resolve/{id} | Resolve a specific defect
[**UpdateDefect**](DefectsAPI.md#UpdateDefect) | **Patch** /defect/{code}/{id} | Update defect
[**UpdateDefectStatus**](DefectsAPI.md#UpdateDefectStatus) | **Patch** /defect/{code}/status/{id} | Update a specific defect status



## CreateDefect

> IdResponse CreateDefect(ctx, code).DefectCreate(defectCreate).Execute()

Create a new defect



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
	defectCreate := *openapiclient.NewDefectCreate("Title_example", "ActualResult_example", int32(123)) // DefectCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefectsAPI.CreateDefect(context.Background(), code).DefectCreate(defectCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefectsAPI.CreateDefect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDefect`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `DefectsAPI.CreateDefect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDefectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defectCreate** | [**DefectCreate**](DefectCreate.md) |  | 

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


## DeleteDefect

> IdResponse DeleteDefect(ctx, code, id).Execute()

Delete defect



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
	resp, r, err := apiClient.DefectsAPI.DeleteDefect(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefectsAPI.DeleteDefect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDefect`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `DefectsAPI.DeleteDefect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefectRequest struct via the builder pattern


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


## GetDefect

> DefectResponse GetDefect(ctx, code, id).Execute()

Get a specific defect



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
	resp, r, err := apiClient.DefectsAPI.GetDefect(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefectsAPI.GetDefect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefect`: DefectResponse
	fmt.Fprintf(os.Stdout, "Response from `DefectsAPI.GetDefect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DefectResponse**](DefectResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefects

> DefectListResponse GetDefects(ctx, code).Status(status).Limit(limit).Offset(offset).Execute()

Get all defects



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
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 | A number of entities in result set. (optional) (default to 10)
	offset := int32(56) // int32 | How many entities should be skipped. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefectsAPI.GetDefects(context.Background(), code).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefectsAPI.GetDefects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefects`: DefectListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefectsAPI.GetDefects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]

### Return type

[**DefectListResponse**](DefectListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveDefect

> IdResponse ResolveDefect(ctx, code, id).Execute()

Resolve a specific defect



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
	resp, r, err := apiClient.DefectsAPI.ResolveDefect(context.Background(), code, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefectsAPI.ResolveDefect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveDefect`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `DefectsAPI.ResolveDefect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveDefectRequest struct via the builder pattern


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


## UpdateDefect

> IdResponse UpdateDefect(ctx, code, id).DefectUpdate(defectUpdate).Execute()

Update defect



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
	defectUpdate := *openapiclient.NewDefectUpdate() // DefectUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefectsAPI.UpdateDefect(context.Background(), code, id).DefectUpdate(defectUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefectsAPI.UpdateDefect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDefect`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `DefectsAPI.UpdateDefect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **defectUpdate** | [**DefectUpdate**](DefectUpdate.md) |  | 

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


## UpdateDefectStatus

> BaseResponse UpdateDefectStatus(ctx, code, id).DefectStatus(defectStatus).Execute()

Update a specific defect status



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
	defectStatus := *openapiclient.NewDefectStatus("Status_example") // DefectStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefectsAPI.UpdateDefectStatus(context.Background(), code, id).DefectStatus(defectStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefectsAPI.UpdateDefectStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDefectStatus`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `DefectsAPI.UpdateDefectStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefectStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **defectStatus** | [**DefectStatus**](DefectStatus.md) |  | 

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

