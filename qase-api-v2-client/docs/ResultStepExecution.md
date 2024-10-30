# ResultStepExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **NullableFloat64** | Unix epoch time in seconds (whole part) and milliseconds (fractional part). | [optional] 
**EndTime** | Pointer to **NullableFloat64** | Unix epoch time in seconds (whole part) and milliseconds (fractional part). | [optional] 
**Status** | [**ResultStepStatus**](ResultStepStatus.md) |  | 
**Duration** | Pointer to **NullableInt64** | Duration of the test step execution in milliseconds. | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResultStepExecution

`func NewResultStepExecution(status ResultStepStatus, ) *ResultStepExecution`

NewResultStepExecution instantiates a new ResultStepExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultStepExecutionWithDefaults

`func NewResultStepExecutionWithDefaults() *ResultStepExecution`

NewResultStepExecutionWithDefaults instantiates a new ResultStepExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *ResultStepExecution) GetStartTime() float64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ResultStepExecution) GetStartTimeOk() (*float64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ResultStepExecution) SetStartTime(v float64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ResultStepExecution) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ResultStepExecution) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ResultStepExecution) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *ResultStepExecution) GetEndTime() float64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ResultStepExecution) GetEndTimeOk() (*float64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ResultStepExecution) SetEndTime(v float64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ResultStepExecution) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *ResultStepExecution) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *ResultStepExecution) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStatus

`func (o *ResultStepExecution) GetStatus() ResultStepStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultStepExecution) GetStatusOk() (*ResultStepStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultStepExecution) SetStatus(v ResultStepStatus)`

SetStatus sets Status field to given value.


### GetDuration

`func (o *ResultStepExecution) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ResultStepExecution) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ResultStepExecution) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ResultStepExecution) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ResultStepExecution) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ResultStepExecution) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetComment

`func (o *ResultStepExecution) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ResultStepExecution) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ResultStepExecution) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ResultStepExecution) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAttachments

`func (o *ResultStepExecution) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ResultStepExecution) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ResultStepExecution) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ResultStepExecution) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


