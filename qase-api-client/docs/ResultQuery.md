# ResultQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**ResultHash** | **string** |  | 
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

### NewResultQuery

`func NewResultQuery(resultHash string, ) *ResultQuery`

NewResultQuery instantiates a new ResultQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultQueryWithDefaults

`func NewResultQueryWithDefaults() *ResultQuery`

NewResultQueryWithDefaults instantiates a new ResultQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *ResultQuery) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ResultQuery) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ResultQuery) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ResultQuery) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetResultHash

`func (o *ResultQuery) GetResultHash() string`

GetResultHash returns the ResultHash field if non-nil, zero value otherwise.

### GetResultHashOk

`func (o *ResultQuery) GetResultHashOk() (*string, bool)`

GetResultHashOk returns a tuple with the ResultHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultHash

`func (o *ResultQuery) SetResultHash(v string)`

SetResultHash sets ResultHash field to given value.


### GetComment

`func (o *ResultQuery) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ResultQuery) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ResultQuery) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ResultQuery) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ResultQuery) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ResultQuery) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetStacktrace

`func (o *ResultQuery) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *ResultQuery) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *ResultQuery) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *ResultQuery) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### SetStacktraceNil

`func (o *ResultQuery) SetStacktraceNil(b bool)`

 SetStacktraceNil sets the value for Stacktrace to be an explicit nil

### UnsetStacktrace
`func (o *ResultQuery) UnsetStacktrace()`

UnsetStacktrace ensures that no value is present for Stacktrace, not even an explicit nil
### GetRunId

`func (o *ResultQuery) GetRunId() int64`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ResultQuery) GetRunIdOk() (*int64, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ResultQuery) SetRunId(v int64)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ResultQuery) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetCaseId

`func (o *ResultQuery) GetCaseId() int64`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ResultQuery) GetCaseIdOk() (*int64, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ResultQuery) SetCaseId(v int64)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *ResultQuery) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetSteps

`func (o *ResultQuery) GetSteps() []TestStepResult`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ResultQuery) GetStepsOk() (*[]TestStepResult, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ResultQuery) SetSteps(v []TestStepResult)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ResultQuery) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *ResultQuery) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *ResultQuery) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetStatus

`func (o *ResultQuery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultQuery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultQuery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResultQuery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsApiResult

`func (o *ResultQuery) GetIsApiResult() bool`

GetIsApiResult returns the IsApiResult field if non-nil, zero value otherwise.

### GetIsApiResultOk

`func (o *ResultQuery) GetIsApiResultOk() (*bool, bool)`

GetIsApiResultOk returns a tuple with the IsApiResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApiResult

`func (o *ResultQuery) SetIsApiResult(v bool)`

SetIsApiResult sets IsApiResult field to given value.

### HasIsApiResult

`func (o *ResultQuery) HasIsApiResult() bool`

HasIsApiResult returns a boolean if a field has been set.

### GetTimeSpentMs

`func (o *ResultQuery) GetTimeSpentMs() int64`

GetTimeSpentMs returns the TimeSpentMs field if non-nil, zero value otherwise.

### GetTimeSpentMsOk

`func (o *ResultQuery) GetTimeSpentMsOk() (*int64, bool)`

GetTimeSpentMsOk returns a tuple with the TimeSpentMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpentMs

`func (o *ResultQuery) SetTimeSpentMs(v int64)`

SetTimeSpentMs sets TimeSpentMs field to given value.

### HasTimeSpentMs

`func (o *ResultQuery) HasTimeSpentMs() bool`

HasTimeSpentMs returns a boolean if a field has been set.

### GetEndTime

`func (o *ResultQuery) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ResultQuery) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ResultQuery) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ResultQuery) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *ResultQuery) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *ResultQuery) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetAttachments

`func (o *ResultQuery) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ResultQuery) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ResultQuery) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ResultQuery) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


