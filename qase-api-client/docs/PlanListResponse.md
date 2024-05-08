# PlanListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**PlanListResponseAllOfResult**](PlanListResponseAllOfResult.md) |  | [optional] 

## Methods

### NewPlanListResponse

`func NewPlanListResponse() *PlanListResponse`

NewPlanListResponse instantiates a new PlanListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanListResponseWithDefaults

`func NewPlanListResponseWithDefaults() *PlanListResponse`

NewPlanListResponseWithDefaults instantiates a new PlanListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PlanListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlanListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlanListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlanListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *PlanListResponse) GetResult() PlanListResponseAllOfResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PlanListResponse) GetResultOk() (*PlanListResponseAllOfResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PlanListResponse) SetResult(v PlanListResponseAllOfResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *PlanListResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


