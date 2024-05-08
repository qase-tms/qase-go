# \ConfigurationsAPI

All URIs are relative to *https://api.qase.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](ConfigurationsAPI.md#CreateConfiguration) | **Post** /configuration/{code} | Create a new configuration in a particular group.
[**CreateConfigurationGroup**](ConfigurationsAPI.md#CreateConfigurationGroup) | **Post** /configuration/{code}/group | Create a new configuration group.
[**GetConfigurations**](ConfigurationsAPI.md#GetConfigurations) | **Get** /configuration/{code} | Get all configuration groups with configurations.



## CreateConfiguration

> IdResponse CreateConfiguration(ctx, code).ConfigurationCreate(configurationCreate).Execute()

Create a new configuration in a particular group.



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
	configurationCreate := *openapiclient.NewConfigurationCreate("Title_example", int32(123)) // ConfigurationCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.CreateConfiguration(context.Background(), code).ConfigurationCreate(configurationCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.CreateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConfiguration`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.CreateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationCreate** | [**ConfigurationCreate**](ConfigurationCreate.md) |  | 

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


## CreateConfigurationGroup

> IdResponse CreateConfigurationGroup(ctx, code).ConfigurationGroupCreate(configurationGroupCreate).Execute()

Create a new configuration group.



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
	configurationGroupCreate := *openapiclient.NewConfigurationGroupCreate("Title_example") // ConfigurationGroupCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.CreateConfigurationGroup(context.Background(), code).ConfigurationGroupCreate(configurationGroupCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.CreateConfigurationGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConfigurationGroup`: IdResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.CreateConfigurationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationGroupCreate** | [**ConfigurationGroupCreate**](ConfigurationGroupCreate.md) |  | 

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


## GetConfigurations

> ConfigurationListResponse GetConfigurations(ctx, code).Execute()

Get all configuration groups with configurations.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.GetConfigurations(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.GetConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurations`: ConfigurationListResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.GetConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code of project, where to search entities. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationListResponse**](ConfigurationListResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

