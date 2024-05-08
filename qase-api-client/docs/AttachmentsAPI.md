# \AttachmentsAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAttachment**](AttachmentsAPI.md#DeleteAttachment) | **Delete** /attachment/{hash} | Remove attachment by Hash
[**GetAttachment**](AttachmentsAPI.md#GetAttachment) | **Get** /attachment/{hash} | Get attachment by Hash
[**GetAttachments**](AttachmentsAPI.md#GetAttachments) | **Get** /attachment | Get all attachments
[**UploadAttachment**](AttachmentsAPI.md#UploadAttachment) | **Post** /attachment/{code} | Upload attachment



## DeleteAttachment

> HashResponse DeleteAttachment(ctx, hash).Execute()

Remove attachment by Hash



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
	hash := "hash_example" // string | Hash.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.DeleteAttachment(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.DeleteAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAttachment`: HashResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.DeleteAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentRequest struct via the builder pattern


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


## GetAttachment

> AttachmentResponse GetAttachment(ctx, hash).Execute()

Get attachment by Hash



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
	hash := "hash_example" // string | Hash.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.GetAttachment(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.GetAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachment`: AttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttachmentResponse**](AttachmentResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachments

> AttachmentListResponse GetAttachments(ctx).Limit(limit).Offset(offset).Execute()

Get all attachments



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.GetAttachments(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.GetAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachments`: AttachmentListResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.GetAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | A number of entities in result set. | [default to 10]
 **offset** | **int32** | How many entities should be skipped. | [default to 0]

### Return type

[**AttachmentListResponse**](AttachmentListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadAttachment

> AttachmentUploadsResponse UploadAttachment(ctx, code).File(file).Execute()

Upload attachment



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
	file := []*os.File{"TODO"} // []*os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.UploadAttachment(context.Background(), code).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.UploadAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadAttachment`: AttachmentUploadsResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.UploadAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **[]*os.File** |  | 

### Return type

[**AttachmentUploadsResponse**](AttachmentUploadsResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

