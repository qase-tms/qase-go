# ResultUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**TimeMs** | Pointer to **NullableInt64** |  | [optional] 
**Defect** | Pointer to **NullableBool** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Stacktrace** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Steps** | Pointer to [**[]TestStepResultCreate**](TestStepResultCreate.md) |  | [optional] 

## Methods

### NewResultUpdate

`func NewResultUpdate() *ResultUpdate`

NewResultUpdate instantiates a new ResultUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultUpdateWithDefaults

`func NewResultUpdateWithDefaults() *ResultUpdate`

NewResultUpdateWithDefaults instantiates a new ResultUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResultUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResultUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeMs

`func (o *ResultUpdate) GetTimeMs() int64`

GetTimeMs returns the TimeMs field if non-nil, zero value otherwise.

### GetTimeMsOk

`func (o *ResultUpdate) GetTimeMsOk() (*int64, bool)`

GetTimeMsOk returns a tuple with the TimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMs

`func (o *ResultUpdate) SetTimeMs(v int64)`

SetTimeMs sets TimeMs field to given value.

### HasTimeMs

`func (o *ResultUpdate) HasTimeMs() bool`

HasTimeMs returns a boolean if a field has been set.

### SetTimeMsNil

`func (o *ResultUpdate) SetTimeMsNil(b bool)`

 SetTimeMsNil sets the value for TimeMs to be an explicit nil

### UnsetTimeMs
`func (o *ResultUpdate) UnsetTimeMs()`

UnsetTimeMs ensures that no value is present for TimeMs, not even an explicit nil
### GetDefect

`func (o *ResultUpdate) GetDefect() bool`

GetDefect returns the Defect field if non-nil, zero value otherwise.

### GetDefectOk

`func (o *ResultUpdate) GetDefectOk() (*bool, bool)`

GetDefectOk returns a tuple with the Defect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefect

`func (o *ResultUpdate) SetDefect(v bool)`

SetDefect sets Defect field to given value.

### HasDefect

`func (o *ResultUpdate) HasDefect() bool`

HasDefect returns a boolean if a field has been set.

### SetDefectNil

`func (o *ResultUpdate) SetDefectNil(b bool)`

 SetDefectNil sets the value for Defect to be an explicit nil

### UnsetDefect
`func (o *ResultUpdate) UnsetDefect()`

UnsetDefect ensures that no value is present for Defect, not even an explicit nil
### GetAttachments

`func (o *ResultUpdate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ResultUpdate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ResultUpdate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ResultUpdate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *ResultUpdate) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *ResultUpdate) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetStacktrace

`func (o *ResultUpdate) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *ResultUpdate) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *ResultUpdate) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *ResultUpdate) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### SetStacktraceNil

`func (o *ResultUpdate) SetStacktraceNil(b bool)`

 SetStacktraceNil sets the value for Stacktrace to be an explicit nil

### UnsetStacktrace
`func (o *ResultUpdate) UnsetStacktrace()`

UnsetStacktrace ensures that no value is present for Stacktrace, not even an explicit nil
### GetComment

`func (o *ResultUpdate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ResultUpdate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ResultUpdate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ResultUpdate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ResultUpdate) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ResultUpdate) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetSteps

`func (o *ResultUpdate) GetSteps() []TestStepResultCreate`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ResultUpdate) GetStepsOk() (*[]TestStepResultCreate, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ResultUpdate) SetSteps(v []TestStepResultCreate)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ResultUpdate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *ResultUpdate) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *ResultUpdate) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


