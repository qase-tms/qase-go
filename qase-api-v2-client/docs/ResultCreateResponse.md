# ResultCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**ResultCreateResponseAllOfResult**](ResultCreateResponseAllOfResult.md) |  | [optional] 

## Methods

### NewResultCreateResponse

`func NewResultCreateResponse() *ResultCreateResponse`

NewResultCreateResponse instantiates a new ResultCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCreateResponseWithDefaults

`func NewResultCreateResponseWithDefaults() *ResultCreateResponse`

NewResultCreateResponseWithDefaults instantiates a new ResultCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResultCreateResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultCreateResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultCreateResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResultCreateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ResultCreateResponse) GetResult() ResultCreateResponseAllOfResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResultCreateResponse) GetResultOk() (*ResultCreateResponseAllOfResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResultCreateResponse) SetResult(v ResultCreateResponseAllOfResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ResultCreateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


