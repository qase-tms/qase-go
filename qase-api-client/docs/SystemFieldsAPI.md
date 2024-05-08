# \SystemFieldsAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemFields**](SystemFieldsAPI.md#GetSystemFields) | **Get** /system_field | Get all System Fields



## GetSystemFields

> SystemFieldListResponse GetSystemFields(ctx).Execute()

Get all System Fields



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemFieldsAPI.GetSystemFields(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemFieldsAPI.GetSystemFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemFields`: SystemFieldListResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemFieldsAPI.GetSystemFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemFieldsRequest struct via the builder pattern


### Return type

[**SystemFieldListResponse**](SystemFieldListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

