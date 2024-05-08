# RunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**Run**](Run.md) |  | [optional] 

## Methods

### NewRunResponse

`func NewRunResponse() *RunResponse`

NewRunResponse instantiates a new RunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunResponseWithDefaults

`func NewRunResponseWithDefaults() *RunResponse`

NewRunResponseWithDefaults instantiates a new RunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RunResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RunResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *RunResponse) GetResult() Run`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RunResponse) GetResultOk() (*Run, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RunResponse) SetResult(v Run)`

SetResult sets Result field to given value.

### HasResult

`func (o *RunResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


