# DefectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**Defect**](Defect.md) |  | [optional] 

## Methods

### NewDefectResponse

`func NewDefectResponse() *DefectResponse`

NewDefectResponse instantiates a new DefectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefectResponseWithDefaults

`func NewDefectResponseWithDefaults() *DefectResponse`

NewDefectResponseWithDefaults instantiates a new DefectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DefectResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DefectResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DefectResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DefectResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *DefectResponse) GetResult() Defect`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DefectResponse) GetResultOk() (*Defect, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DefectResponse) SetResult(v Defect)`

SetResult sets Result field to given value.

### HasResult

`func (o *DefectResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


