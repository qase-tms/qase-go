# ConfigurationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**ConfigurationListResponseAllOfResult**](ConfigurationListResponseAllOfResult.md) |  | [optional] 

## Methods

### NewConfigurationListResponse

`func NewConfigurationListResponse() *ConfigurationListResponse`

NewConfigurationListResponse instantiates a new ConfigurationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationListResponseWithDefaults

`func NewConfigurationListResponseWithDefaults() *ConfigurationListResponse`

NewConfigurationListResponseWithDefaults instantiates a new ConfigurationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConfigurationListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigurationListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigurationListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigurationListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ConfigurationListResponse) GetResult() ConfigurationListResponseAllOfResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ConfigurationListResponse) GetResultOk() (*ConfigurationListResponseAllOfResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ConfigurationListResponse) SetResult(v ConfigurationListResponseAllOfResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ConfigurationListResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


