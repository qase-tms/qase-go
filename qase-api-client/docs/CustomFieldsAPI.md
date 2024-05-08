# \CustomFieldsAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomField**](CustomFieldsAPI.md#CreateCustomField) | **Post** /custom_field | Create new Custom Field
[**DeleteCustomField**](CustomFieldsAPI.md#DeleteCustomField) | **Delete** /custom_field/{id} | Delete Custom Field by id
[**GetCustomField**](CustomFieldsAPI.md#GetCustomField) | **Get** /custom_field/{id} | Get Custom Field by id
[**GetCustomFields**](CustomFieldsAPI.md#GetCustomFields) | **Get** /custom_field | Get all Custom Fields
[**UpdateCustomField**](CustomFieldsAPI.md#UpdateCustomField) | **Patch** /custom_field/{id} | Update Custom Field by id



## CreateCustomField

> IdResponse CreateCustomField(ctx).CustomFieldCreate(customFieldCreate).Execute()

Create new Custom Field



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
	customFieldCreate := *openapiclient.NewCustomFieldCreate("Title_example", int32(123), int32(123)) // CustomFieldCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.CreateCustomField(context.Background()).CustomFieldCreate(customFieldCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.CreateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomField`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.CreateCustomField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customFieldCreate** | [**CustomFieldCreate**](CustomFieldCreate.md) |  | 

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


## DeleteCustomField

> BaseResponse DeleteCustomField(ctx, id).Execute()

Delete Custom Field by id



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
	id := int32(56) // int32 | Identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.DeleteCustomField(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.DeleteCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomField`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.DeleteCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFieldRequest struct via the builder pattern


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


## GetCustomField

> CustomFieldResponse GetCustomField(ctx, id).Execute()

Get Custom Field by id



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
	id := int32(56) // int32 | Identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.GetCustomField(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.GetCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomField`: CustomFieldResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.GetCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFieldResponse**](CustomFieldResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomFields

> CustomFieldListResponse GetCustomFields(ctx).Entity(entity).Type_(type_).Limit(limit).Offset(offset).Execute()

Get all Custom Fields



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
	entity := "entity_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	limit := int32(56) // int32 | A number of entities in result set. (optional) (default to 10)
	offset := int32(56) // int32 | How many entities should be skipped. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.GetCustomFields(context.Background()).Entity(entity).Type_(type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.GetCustomFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomFields`: CustomFieldListResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.GetCustomFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string** |  | 
 **type_** | **string** |  | 
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]

### Return type

[**CustomFieldListResponse**](CustomFieldListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomField

> BaseResponse UpdateCustomField(ctx, id).CustomFieldUpdate(customFieldUpdate).Execute()

Update Custom Field by id



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
	id := int32(56) // int32 | Identifier.
	customFieldUpdate := *openapiclient.NewCustomFieldUpdate("Title_example") // CustomFieldUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFieldsAPI.UpdateCustomField(context.Background(), id).CustomFieldUpdate(customFieldUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFieldsAPI.UpdateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomField`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFieldsAPI.UpdateCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFieldUpdate** | [**CustomFieldUpdate**](CustomFieldUpdate.md) |  | 

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

