# ResultExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **NullableFloat64** |  | [optional] 
**EndTime** | Pointer to **NullableFloat64** |  | [optional] 
**Status** | **string** | Can have the following values passed, failed, blocked, skipped, invalid + custom statuses | 
**Duration** | Pointer to **NullableInt64** |  | [optional] 
**Stacktrace** | Pointer to **NullableString** |  | [optional] 
**Thread** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResultExecution

`func NewResultExecution(status string, ) *ResultExecution`

NewResultExecution instantiates a new ResultExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultExecutionWithDefaults

`func NewResultExecutionWithDefaults() *ResultExecution`

NewResultExecutionWithDefaults instantiates a new ResultExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *ResultExecution) GetStartTime() float64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ResultExecution) GetStartTimeOk() (*float64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ResultExecution) SetStartTime(v float64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ResultExecution) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ResultExecution) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ResultExecution) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *ResultExecution) GetEndTime() float64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ResultExecution) GetEndTimeOk() (*float64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ResultExecution) SetEndTime(v float64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ResultExecution) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *ResultExecution) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *ResultExecution) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStatus

`func (o *ResultExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultExecution) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDuration

`func (o *ResultExecution) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ResultExecution) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ResultExecution) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ResultExecution) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ResultExecution) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ResultExecution) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetStacktrace

`func (o *ResultExecution) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *ResultExecution) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *ResultExecution) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *ResultExecution) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### SetStacktraceNil

`func (o *ResultExecution) SetStacktraceNil(b bool)`

 SetStacktraceNil sets the value for Stacktrace to be an explicit nil

### UnsetStacktrace
`func (o *ResultExecution) UnsetStacktrace()`

UnsetStacktrace ensures that no value is present for Stacktrace, not even an explicit nil
### GetThread

`func (o *ResultExecution) GetThread() string`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *ResultExecution) GetThreadOk() (*string, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *ResultExecution) SetThread(v string)`

SetThread sets Thread field to given value.

### HasThread

`func (o *ResultExecution) HasThread() bool`

HasThread returns a boolean if a field has been set.

### SetThreadNil

`func (o *ResultExecution) SetThreadNil(b bool)`

 SetThreadNil sets the value for Thread to be an explicit nil

### UnsetThread
`func (o *ResultExecution) UnsetThread()`

UnsetThread ensures that no value is present for Thread, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


