# ResultCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | Pointer to **int64** |  | [optional] 
**Case** | Pointer to [**ResultCreateCase**](ResultCreateCase.md) |  | [optional] 
**Status** | **string** | Can have the following values &#x60;passed&#x60;, &#x60;failed&#x60;, &#x60;blocked&#x60;, &#x60;skipped&#x60;, &#x60;invalid&#x60; + custom statuses | 
**StartTime** | Pointer to **NullableInt32** |  | [optional] 
**Time** | Pointer to **NullableInt64** |  | [optional] 
**TimeMs** | Pointer to **NullableInt64** |  | [optional] 
**Defect** | Pointer to **NullableBool** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Stacktrace** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Param** | Pointer to **map[string]string** | A map of parameters (name &#x3D;&gt; value) | [optional] 
**ParamGroups** | Pointer to **[][]string** | List parameter groups by name only. Add their values in the &#39;param&#39; field | [optional] 
**Steps** | Pointer to [**[]TestStepResultCreate**](TestStepResultCreate.md) |  | [optional] 
**AuthorId** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewResultCreate

`func NewResultCreate(status string, ) *ResultCreate`

NewResultCreate instantiates a new ResultCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCreateWithDefaults

`func NewResultCreateWithDefaults() *ResultCreate`

NewResultCreateWithDefaults instantiates a new ResultCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *ResultCreate) GetCaseId() int64`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ResultCreate) GetCaseIdOk() (*int64, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ResultCreate) SetCaseId(v int64)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *ResultCreate) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetCase

`func (o *ResultCreate) GetCase() ResultCreateCase`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *ResultCreate) GetCaseOk() (*ResultCreateCase, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *ResultCreate) SetCase(v ResultCreateCase)`

SetCase sets Case field to given value.

### HasCase

`func (o *ResultCreate) HasCase() bool`

HasCase returns a boolean if a field has been set.

### GetStatus

`func (o *ResultCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultCreate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartTime

`func (o *ResultCreate) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ResultCreate) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ResultCreate) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ResultCreate) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ResultCreate) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ResultCreate) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetTime

`func (o *ResultCreate) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResultCreate) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResultCreate) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResultCreate) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *ResultCreate) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *ResultCreate) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetTimeMs

`func (o *ResultCreate) GetTimeMs() int64`

GetTimeMs returns the TimeMs field if non-nil, zero value otherwise.

### GetTimeMsOk

`func (o *ResultCreate) GetTimeMsOk() (*int64, bool)`

GetTimeMsOk returns a tuple with the TimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMs

`func (o *ResultCreate) SetTimeMs(v int64)`

SetTimeMs sets TimeMs field to given value.

### HasTimeMs

`func (o *ResultCreate) HasTimeMs() bool`

HasTimeMs returns a boolean if a field has been set.

### SetTimeMsNil

`func (o *ResultCreate) SetTimeMsNil(b bool)`

 SetTimeMsNil sets the value for TimeMs to be an explicit nil

### UnsetTimeMs
`func (o *ResultCreate) UnsetTimeMs()`

UnsetTimeMs ensures that no value is present for TimeMs, not even an explicit nil
### GetDefect

`func (o *ResultCreate) GetDefect() bool`

GetDefect returns the Defect field if non-nil, zero value otherwise.

### GetDefectOk

`func (o *ResultCreate) GetDefectOk() (*bool, bool)`

GetDefectOk returns a tuple with the Defect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefect

`func (o *ResultCreate) SetDefect(v bool)`

SetDefect sets Defect field to given value.

### HasDefect

`func (o *ResultCreate) HasDefect() bool`

HasDefect returns a boolean if a field has been set.

### SetDefectNil

`func (o *ResultCreate) SetDefectNil(b bool)`

 SetDefectNil sets the value for Defect to be an explicit nil

### UnsetDefect
`func (o *ResultCreate) UnsetDefect()`

UnsetDefect ensures that no value is present for Defect, not even an explicit nil
### GetAttachments

`func (o *ResultCreate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ResultCreate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ResultCreate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ResultCreate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *ResultCreate) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *ResultCreate) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetStacktrace

`func (o *ResultCreate) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *ResultCreate) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *ResultCreate) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *ResultCreate) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### SetStacktraceNil

`func (o *ResultCreate) SetStacktraceNil(b bool)`

 SetStacktraceNil sets the value for Stacktrace to be an explicit nil

### UnsetStacktrace
`func (o *ResultCreate) UnsetStacktrace()`

UnsetStacktrace ensures that no value is present for Stacktrace, not even an explicit nil
### GetComment

`func (o *ResultCreate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ResultCreate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ResultCreate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ResultCreate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ResultCreate) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ResultCreate) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetParam

`func (o *ResultCreate) GetParam() map[string]string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ResultCreate) GetParamOk() (*map[string]string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ResultCreate) SetParam(v map[string]string)`

SetParam sets Param field to given value.

### HasParam

`func (o *ResultCreate) HasParam() bool`

HasParam returns a boolean if a field has been set.

### SetParamNil

`func (o *ResultCreate) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *ResultCreate) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil
### GetParamGroups

`func (o *ResultCreate) GetParamGroups() [][]string`

GetParamGroups returns the ParamGroups field if non-nil, zero value otherwise.

### GetParamGroupsOk

`func (o *ResultCreate) GetParamGroupsOk() (*[][]string, bool)`

GetParamGroupsOk returns a tuple with the ParamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamGroups

`func (o *ResultCreate) SetParamGroups(v [][]string)`

SetParamGroups sets ParamGroups field to given value.

### HasParamGroups

`func (o *ResultCreate) HasParamGroups() bool`

HasParamGroups returns a boolean if a field has been set.

### SetParamGroupsNil

`func (o *ResultCreate) SetParamGroupsNil(b bool)`

 SetParamGroupsNil sets the value for ParamGroups to be an explicit nil

### UnsetParamGroups
`func (o *ResultCreate) UnsetParamGroups()`

UnsetParamGroups ensures that no value is present for ParamGroups, not even an explicit nil
### GetSteps

`func (o *ResultCreate) GetSteps() []TestStepResultCreate`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ResultCreate) GetStepsOk() (*[]TestStepResultCreate, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ResultCreate) SetSteps(v []TestStepResultCreate)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ResultCreate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *ResultCreate) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *ResultCreate) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetAuthorId

`func (o *ResultCreate) GetAuthorId() int64`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *ResultCreate) GetAuthorIdOk() (*int64, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *ResultCreate) SetAuthorId(v int64)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *ResultCreate) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### SetAuthorIdNil

`func (o *ResultCreate) SetAuthorIdNil(b bool)`

 SetAuthorIdNil sets the value for AuthorId to be an explicit nil

### UnsetAuthorId
`func (o *ResultCreate) UnsetAuthorId()`

UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


