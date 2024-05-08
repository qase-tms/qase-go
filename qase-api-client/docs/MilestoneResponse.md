# MilestoneResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**Milestone**](Milestone.md) |  | [optional] 

## Methods

### NewMilestoneResponse

`func NewMilestoneResponse() *MilestoneResponse`

NewMilestoneResponse instantiates a new MilestoneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilestoneResponseWithDefaults

`func NewMilestoneResponseWithDefaults() *MilestoneResponse`

NewMilestoneResponseWithDefaults instantiates a new MilestoneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MilestoneResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MilestoneResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MilestoneResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MilestoneResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *MilestoneResponse) GetResult() Milestone`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MilestoneResponse) GetResultOk() (*Milestone, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MilestoneResponse) SetResult(v Milestone)`

SetResult sets Result field to given value.

### HasResult

`func (o *MilestoneResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


