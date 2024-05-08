# TestCaseListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**TestCaseListResponseAllOfResult**](TestCaseListResponseAllOfResult.md) |  | [optional] 

## Methods

### NewTestCaseListResponse

`func NewTestCaseListResponse() *TestCaseListResponse`

NewTestCaseListResponse instantiates a new TestCaseListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCaseListResponseWithDefaults

`func NewTestCaseListResponseWithDefaults() *TestCaseListResponse`

NewTestCaseListResponseWithDefaults instantiates a new TestCaseListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TestCaseListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestCaseListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestCaseListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestCaseListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TestCaseListResponse) GetResult() TestCaseListResponseAllOfResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TestCaseListResponse) GetResultOk() (*TestCaseListResponseAllOfResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TestCaseListResponse) SetResult(v TestCaseListResponseAllOfResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *TestCaseListResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


