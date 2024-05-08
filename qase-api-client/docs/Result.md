# Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Stacktrace** | Pointer to **NullableString** |  | [optional] 
**RunId** | Pointer to **int64** |  | [optional] 
**CaseId** | Pointer to **int64** |  | [optional] 
**Steps** | Pointer to [**[]TestStepResult**](TestStepResult.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**IsApiResult** | Pointer to **bool** |  | [optional] 
**TimeSpentMs** | Pointer to **int64** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 

## Methods

### NewResult

`func NewResult() *Result`

NewResult instantiates a new Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultWithDefaults

`func NewResultWithDefaults() *Result`

NewResultWithDefaults instantiates a new Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *Result) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Result) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Result) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Result) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetComment

`func (o *Result) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Result) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Result) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Result) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Result) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Result) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetStacktrace

`func (o *Result) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *Result) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *Result) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *Result) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### SetStacktraceNil

`func (o *Result) SetStacktraceNil(b bool)`

 SetStacktraceNil sets the value for Stacktrace to be an explicit nil

### UnsetStacktrace
`func (o *Result) UnsetStacktrace()`

UnsetStacktrace ensures that no value is present for Stacktrace, not even an explicit nil
### GetRunId

`func (o *Result) GetRunId() int64`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *Result) GetRunIdOk() (*int64, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *Result) SetRunId(v int64)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *Result) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetCaseId

`func (o *Result) GetCaseId() int64`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *Result) GetCaseIdOk() (*int64, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *Result) SetCaseId(v int64)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *Result) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetSteps

`func (o *Result) GetSteps() []TestStepResult`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Result) GetStepsOk() (*[]TestStepResult, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Result) SetSteps(v []TestStepResult)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Result) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *Result) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *Result) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetStatus

`func (o *Result) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Result) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Result) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Result) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsApiResult

`func (o *Result) GetIsApiResult() bool`

GetIsApiResult returns the IsApiResult field if non-nil, zero value otherwise.

### GetIsApiResultOk

`func (o *Result) GetIsApiResultOk() (*bool, bool)`

GetIsApiResultOk returns a tuple with the IsApiResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApiResult

`func (o *Result) SetIsApiResult(v bool)`

SetIsApiResult sets IsApiResult field to given value.

### HasIsApiResult

`func (o *Result) HasIsApiResult() bool`

HasIsApiResult returns a boolean if a field has been set.

### GetTimeSpentMs

`func (o *Result) GetTimeSpentMs() int64`

GetTimeSpentMs returns the TimeSpentMs field if non-nil, zero value otherwise.

### GetTimeSpentMsOk

`func (o *Result) GetTimeSpentMsOk() (*int64, bool)`

GetTimeSpentMsOk returns a tuple with the TimeSpentMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpentMs

`func (o *Result) SetTimeSpentMs(v int64)`

SetTimeSpentMs sets TimeSpentMs field to given value.

### HasTimeSpentMs

`func (o *Result) HasTimeSpentMs() bool`

HasTimeSpentMs returns a boolean if a field has been set.

### GetEndTime

`func (o *Result) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Result) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Result) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Result) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *Result) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *Result) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetAttachments

`func (o *Result) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Result) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Result) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Result) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


