# SearchResponseAllOfResultEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**RunId** | **int64** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusText** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**Stats** | Pointer to [**RunStats**](RunStats.md) |  | [optional] 
**TimeSpent** | Pointer to **int64** | Time in ms. | [optional] 
**Environment** | Pointer to [**NullableRunEnvironment**](RunEnvironment.md) |  | [optional] 
**Milestone** | Pointer to [**NullableRunMilestone**](RunMilestone.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 
**Tags** | Pointer to [**[]TagValue**](TagValue.md) |  | [optional] 
**Cases** | Pointer to **[]int64** |  | [optional] 
**PlanId** | **int64** |  | 
**Hash** | Pointer to **string** |  | [optional] 
**ResultHash** | **string** |  | 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Stacktrace** | Pointer to **NullableString** |  | [optional] 
**CaseId** | Pointer to **int64** |  | [optional] 
**Steps** | Pointer to [**[]TestStep**](TestStep.md) |  | [optional] 
**IsApiResult** | Pointer to **bool** |  | [optional] 
**TimeSpentMs** | Pointer to **int64** |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**RequirementId** | **int64** |  | 
**ParentId** | Pointer to **NullableInt64** |  | [optional] 
**MemberId** | Pointer to **int64** | Deprecated, use &#x60;author_id&#x60; instead. | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**TestCaseId** | **int64** |  | 
**Position** | Pointer to **int32** |  | [optional] 
**Preconditions** | Pointer to **NullableString** |  | [optional] 
**Postconditions** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Layer** | Pointer to **int32** |  | [optional] 
**IsFlaky** | Pointer to **int32** |  | [optional] 
**Behavior** | Pointer to **int32** |  | [optional] 
**Automation** | Pointer to **int32** |  | [optional] 
**MilestoneId** | Pointer to **NullableInt64** |  | [optional] 
**SuiteId** | Pointer to **NullableInt64** |  | [optional] 
**StepsType** | Pointer to **NullableString** |  | [optional] 
**Params** | Pointer to [**TestCaseParams**](TestCaseParams.md) |  | [optional] 
**AuthorId** | Pointer to **int64** |  | [optional] 
**DefectId** | **int64** |  | 
**ActualResult** | Pointer to **string** |  | [optional] 
**Resolved** | Pointer to **NullableTime** |  | [optional] 
**ExternalData** | Pointer to **string** |  | [optional] 
**CasesCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchResponseAllOfResultEntities

`func NewSearchResponseAllOfResultEntities(runId int64, planId int64, resultHash string, requirementId int64, testCaseId int64, defectId int64, ) *SearchResponseAllOfResultEntities`

NewSearchResponseAllOfResultEntities instantiates a new SearchResponseAllOfResultEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseAllOfResultEntitiesWithDefaults

`func NewSearchResponseAllOfResultEntitiesWithDefaults() *SearchResponseAllOfResultEntities`

NewSearchResponseAllOfResultEntitiesWithDefaults instantiates a new SearchResponseAllOfResultEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchResponseAllOfResultEntities) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchResponseAllOfResultEntities) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchResponseAllOfResultEntities) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SearchResponseAllOfResultEntities) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRunId

`func (o *SearchResponseAllOfResultEntities) GetRunId() int64`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *SearchResponseAllOfResultEntities) GetRunIdOk() (*int64, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *SearchResponseAllOfResultEntities) SetRunId(v int64)`

SetRunId sets RunId field to given value.


### GetTitle

`func (o *SearchResponseAllOfResultEntities) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SearchResponseAllOfResultEntities) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SearchResponseAllOfResultEntities) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SearchResponseAllOfResultEntities) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *SearchResponseAllOfResultEntities) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchResponseAllOfResultEntities) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchResponseAllOfResultEntities) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchResponseAllOfResultEntities) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SearchResponseAllOfResultEntities) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SearchResponseAllOfResultEntities) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *SearchResponseAllOfResultEntities) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchResponseAllOfResultEntities) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchResponseAllOfResultEntities) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchResponseAllOfResultEntities) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusText

`func (o *SearchResponseAllOfResultEntities) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *SearchResponseAllOfResultEntities) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *SearchResponseAllOfResultEntities) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *SearchResponseAllOfResultEntities) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetStartTime

`func (o *SearchResponseAllOfResultEntities) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SearchResponseAllOfResultEntities) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SearchResponseAllOfResultEntities) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SearchResponseAllOfResultEntities) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *SearchResponseAllOfResultEntities) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *SearchResponseAllOfResultEntities) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *SearchResponseAllOfResultEntities) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SearchResponseAllOfResultEntities) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SearchResponseAllOfResultEntities) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SearchResponseAllOfResultEntities) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *SearchResponseAllOfResultEntities) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *SearchResponseAllOfResultEntities) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetPublic

`func (o *SearchResponseAllOfResultEntities) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *SearchResponseAllOfResultEntities) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *SearchResponseAllOfResultEntities) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *SearchResponseAllOfResultEntities) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetStats

`func (o *SearchResponseAllOfResultEntities) GetStats() RunStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SearchResponseAllOfResultEntities) GetStatsOk() (*RunStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SearchResponseAllOfResultEntities) SetStats(v RunStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SearchResponseAllOfResultEntities) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetTimeSpent

`func (o *SearchResponseAllOfResultEntities) GetTimeSpent() int64`

GetTimeSpent returns the TimeSpent field if non-nil, zero value otherwise.

### GetTimeSpentOk

`func (o *SearchResponseAllOfResultEntities) GetTimeSpentOk() (*int64, bool)`

GetTimeSpentOk returns a tuple with the TimeSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpent

`func (o *SearchResponseAllOfResultEntities) SetTimeSpent(v int64)`

SetTimeSpent sets TimeSpent field to given value.

### HasTimeSpent

`func (o *SearchResponseAllOfResultEntities) HasTimeSpent() bool`

HasTimeSpent returns a boolean if a field has been set.

### GetEnvironment

`func (o *SearchResponseAllOfResultEntities) GetEnvironment() RunEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SearchResponseAllOfResultEntities) GetEnvironmentOk() (*RunEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SearchResponseAllOfResultEntities) SetEnvironment(v RunEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SearchResponseAllOfResultEntities) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *SearchResponseAllOfResultEntities) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *SearchResponseAllOfResultEntities) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetMilestone

`func (o *SearchResponseAllOfResultEntities) GetMilestone() RunMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *SearchResponseAllOfResultEntities) GetMilestoneOk() (*RunMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *SearchResponseAllOfResultEntities) SetMilestone(v RunMilestone)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *SearchResponseAllOfResultEntities) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### SetMilestoneNil

`func (o *SearchResponseAllOfResultEntities) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *SearchResponseAllOfResultEntities) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetCustomFields

`func (o *SearchResponseAllOfResultEntities) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SearchResponseAllOfResultEntities) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SearchResponseAllOfResultEntities) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SearchResponseAllOfResultEntities) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *SearchResponseAllOfResultEntities) GetTags() []TagValue`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SearchResponseAllOfResultEntities) GetTagsOk() (*[]TagValue, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SearchResponseAllOfResultEntities) SetTags(v []TagValue)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SearchResponseAllOfResultEntities) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCases

`func (o *SearchResponseAllOfResultEntities) GetCases() []int64`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *SearchResponseAllOfResultEntities) GetCasesOk() (*[]int64, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *SearchResponseAllOfResultEntities) SetCases(v []int64)`

SetCases sets Cases field to given value.

### HasCases

`func (o *SearchResponseAllOfResultEntities) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetPlanId

`func (o *SearchResponseAllOfResultEntities) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SearchResponseAllOfResultEntities) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SearchResponseAllOfResultEntities) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.


### GetHash

`func (o *SearchResponseAllOfResultEntities) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SearchResponseAllOfResultEntities) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SearchResponseAllOfResultEntities) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SearchResponseAllOfResultEntities) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetResultHash

`func (o *SearchResponseAllOfResultEntities) GetResultHash() string`

GetResultHash returns the ResultHash field if non-nil, zero value otherwise.

### GetResultHashOk

`func (o *SearchResponseAllOfResultEntities) GetResultHashOk() (*string, bool)`

GetResultHashOk returns a tuple with the ResultHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultHash

`func (o *SearchResponseAllOfResultEntities) SetResultHash(v string)`

SetResultHash sets ResultHash field to given value.


### GetComment

`func (o *SearchResponseAllOfResultEntities) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SearchResponseAllOfResultEntities) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SearchResponseAllOfResultEntities) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SearchResponseAllOfResultEntities) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *SearchResponseAllOfResultEntities) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *SearchResponseAllOfResultEntities) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetStacktrace

`func (o *SearchResponseAllOfResultEntities) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *SearchResponseAllOfResultEntities) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *SearchResponseAllOfResultEntities) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *SearchResponseAllOfResultEntities) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### SetStacktraceNil

`func (o *SearchResponseAllOfResultEntities) SetStacktraceNil(b bool)`

 SetStacktraceNil sets the value for Stacktrace to be an explicit nil

### UnsetStacktrace
`func (o *SearchResponseAllOfResultEntities) UnsetStacktrace()`

UnsetStacktrace ensures that no value is present for Stacktrace, not even an explicit nil
### GetCaseId

`func (o *SearchResponseAllOfResultEntities) GetCaseId() int64`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *SearchResponseAllOfResultEntities) GetCaseIdOk() (*int64, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *SearchResponseAllOfResultEntities) SetCaseId(v int64)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *SearchResponseAllOfResultEntities) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetSteps

`func (o *SearchResponseAllOfResultEntities) GetSteps() []TestStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SearchResponseAllOfResultEntities) GetStepsOk() (*[]TestStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SearchResponseAllOfResultEntities) SetSteps(v []TestStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *SearchResponseAllOfResultEntities) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetIsApiResult

`func (o *SearchResponseAllOfResultEntities) GetIsApiResult() bool`

GetIsApiResult returns the IsApiResult field if non-nil, zero value otherwise.

### GetIsApiResultOk

`func (o *SearchResponseAllOfResultEntities) GetIsApiResultOk() (*bool, bool)`

GetIsApiResultOk returns a tuple with the IsApiResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApiResult

`func (o *SearchResponseAllOfResultEntities) SetIsApiResult(v bool)`

SetIsApiResult sets IsApiResult field to given value.

### HasIsApiResult

`func (o *SearchResponseAllOfResultEntities) HasIsApiResult() bool`

HasIsApiResult returns a boolean if a field has been set.

### GetTimeSpentMs

`func (o *SearchResponseAllOfResultEntities) GetTimeSpentMs() int64`

GetTimeSpentMs returns the TimeSpentMs field if non-nil, zero value otherwise.

### GetTimeSpentMsOk

`func (o *SearchResponseAllOfResultEntities) GetTimeSpentMsOk() (*int64, bool)`

GetTimeSpentMsOk returns a tuple with the TimeSpentMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpentMs

`func (o *SearchResponseAllOfResultEntities) SetTimeSpentMs(v int64)`

SetTimeSpentMs sets TimeSpentMs field to given value.

### HasTimeSpentMs

`func (o *SearchResponseAllOfResultEntities) HasTimeSpentMs() bool`

HasTimeSpentMs returns a boolean if a field has been set.

### GetAttachments

`func (o *SearchResponseAllOfResultEntities) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SearchResponseAllOfResultEntities) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SearchResponseAllOfResultEntities) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SearchResponseAllOfResultEntities) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetRequirementId

`func (o *SearchResponseAllOfResultEntities) GetRequirementId() int64`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *SearchResponseAllOfResultEntities) GetRequirementIdOk() (*int64, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *SearchResponseAllOfResultEntities) SetRequirementId(v int64)`

SetRequirementId sets RequirementId field to given value.


### GetParentId

`func (o *SearchResponseAllOfResultEntities) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SearchResponseAllOfResultEntities) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SearchResponseAllOfResultEntities) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SearchResponseAllOfResultEntities) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *SearchResponseAllOfResultEntities) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *SearchResponseAllOfResultEntities) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetMemberId

`func (o *SearchResponseAllOfResultEntities) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *SearchResponseAllOfResultEntities) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *SearchResponseAllOfResultEntities) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *SearchResponseAllOfResultEntities) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetType

`func (o *SearchResponseAllOfResultEntities) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchResponseAllOfResultEntities) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchResponseAllOfResultEntities) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SearchResponseAllOfResultEntities) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SearchResponseAllOfResultEntities) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SearchResponseAllOfResultEntities) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SearchResponseAllOfResultEntities) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SearchResponseAllOfResultEntities) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SearchResponseAllOfResultEntities) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SearchResponseAllOfResultEntities) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SearchResponseAllOfResultEntities) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SearchResponseAllOfResultEntities) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTestCaseId

`func (o *SearchResponseAllOfResultEntities) GetTestCaseId() int64`

GetTestCaseId returns the TestCaseId field if non-nil, zero value otherwise.

### GetTestCaseIdOk

`func (o *SearchResponseAllOfResultEntities) GetTestCaseIdOk() (*int64, bool)`

GetTestCaseIdOk returns a tuple with the TestCaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCaseId

`func (o *SearchResponseAllOfResultEntities) SetTestCaseId(v int64)`

SetTestCaseId sets TestCaseId field to given value.


### GetPosition

`func (o *SearchResponseAllOfResultEntities) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SearchResponseAllOfResultEntities) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SearchResponseAllOfResultEntities) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *SearchResponseAllOfResultEntities) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPreconditions

`func (o *SearchResponseAllOfResultEntities) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *SearchResponseAllOfResultEntities) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *SearchResponseAllOfResultEntities) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *SearchResponseAllOfResultEntities) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### SetPreconditionsNil

`func (o *SearchResponseAllOfResultEntities) SetPreconditionsNil(b bool)`

 SetPreconditionsNil sets the value for Preconditions to be an explicit nil

### UnsetPreconditions
`func (o *SearchResponseAllOfResultEntities) UnsetPreconditions()`

UnsetPreconditions ensures that no value is present for Preconditions, not even an explicit nil
### GetPostconditions

`func (o *SearchResponseAllOfResultEntities) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *SearchResponseAllOfResultEntities) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *SearchResponseAllOfResultEntities) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *SearchResponseAllOfResultEntities) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### SetPostconditionsNil

`func (o *SearchResponseAllOfResultEntities) SetPostconditionsNil(b bool)`

 SetPostconditionsNil sets the value for Postconditions to be an explicit nil

### UnsetPostconditions
`func (o *SearchResponseAllOfResultEntities) UnsetPostconditions()`

UnsetPostconditions ensures that no value is present for Postconditions, not even an explicit nil
### GetSeverity

`func (o *SearchResponseAllOfResultEntities) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SearchResponseAllOfResultEntities) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SearchResponseAllOfResultEntities) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SearchResponseAllOfResultEntities) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetPriority

`func (o *SearchResponseAllOfResultEntities) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SearchResponseAllOfResultEntities) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SearchResponseAllOfResultEntities) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SearchResponseAllOfResultEntities) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetLayer

`func (o *SearchResponseAllOfResultEntities) GetLayer() int32`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *SearchResponseAllOfResultEntities) GetLayerOk() (*int32, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *SearchResponseAllOfResultEntities) SetLayer(v int32)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *SearchResponseAllOfResultEntities) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetIsFlaky

`func (o *SearchResponseAllOfResultEntities) GetIsFlaky() int32`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *SearchResponseAllOfResultEntities) GetIsFlakyOk() (*int32, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *SearchResponseAllOfResultEntities) SetIsFlaky(v int32)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *SearchResponseAllOfResultEntities) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### GetBehavior

`func (o *SearchResponseAllOfResultEntities) GetBehavior() int32`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *SearchResponseAllOfResultEntities) GetBehaviorOk() (*int32, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *SearchResponseAllOfResultEntities) SetBehavior(v int32)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *SearchResponseAllOfResultEntities) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetAutomation

`func (o *SearchResponseAllOfResultEntities) GetAutomation() int32`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *SearchResponseAllOfResultEntities) GetAutomationOk() (*int32, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *SearchResponseAllOfResultEntities) SetAutomation(v int32)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *SearchResponseAllOfResultEntities) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetMilestoneId

`func (o *SearchResponseAllOfResultEntities) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *SearchResponseAllOfResultEntities) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *SearchResponseAllOfResultEntities) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *SearchResponseAllOfResultEntities) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *SearchResponseAllOfResultEntities) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *SearchResponseAllOfResultEntities) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetSuiteId

`func (o *SearchResponseAllOfResultEntities) GetSuiteId() int64`

GetSuiteId returns the SuiteId field if non-nil, zero value otherwise.

### GetSuiteIdOk

`func (o *SearchResponseAllOfResultEntities) GetSuiteIdOk() (*int64, bool)`

GetSuiteIdOk returns a tuple with the SuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteId

`func (o *SearchResponseAllOfResultEntities) SetSuiteId(v int64)`

SetSuiteId sets SuiteId field to given value.

### HasSuiteId

`func (o *SearchResponseAllOfResultEntities) HasSuiteId() bool`

HasSuiteId returns a boolean if a field has been set.

### SetSuiteIdNil

`func (o *SearchResponseAllOfResultEntities) SetSuiteIdNil(b bool)`

 SetSuiteIdNil sets the value for SuiteId to be an explicit nil

### UnsetSuiteId
`func (o *SearchResponseAllOfResultEntities) UnsetSuiteId()`

UnsetSuiteId ensures that no value is present for SuiteId, not even an explicit nil
### GetStepsType

`func (o *SearchResponseAllOfResultEntities) GetStepsType() string`

GetStepsType returns the StepsType field if non-nil, zero value otherwise.

### GetStepsTypeOk

`func (o *SearchResponseAllOfResultEntities) GetStepsTypeOk() (*string, bool)`

GetStepsTypeOk returns a tuple with the StepsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsType

`func (o *SearchResponseAllOfResultEntities) SetStepsType(v string)`

SetStepsType sets StepsType field to given value.

### HasStepsType

`func (o *SearchResponseAllOfResultEntities) HasStepsType() bool`

HasStepsType returns a boolean if a field has been set.

### SetStepsTypeNil

`func (o *SearchResponseAllOfResultEntities) SetStepsTypeNil(b bool)`

 SetStepsTypeNil sets the value for StepsType to be an explicit nil

### UnsetStepsType
`func (o *SearchResponseAllOfResultEntities) UnsetStepsType()`

UnsetStepsType ensures that no value is present for StepsType, not even an explicit nil
### GetParams

`func (o *SearchResponseAllOfResultEntities) GetParams() TestCaseParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SearchResponseAllOfResultEntities) GetParamsOk() (*TestCaseParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SearchResponseAllOfResultEntities) SetParams(v TestCaseParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *SearchResponseAllOfResultEntities) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetAuthorId

`func (o *SearchResponseAllOfResultEntities) GetAuthorId() int64`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *SearchResponseAllOfResultEntities) GetAuthorIdOk() (*int64, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *SearchResponseAllOfResultEntities) SetAuthorId(v int64)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *SearchResponseAllOfResultEntities) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetDefectId

`func (o *SearchResponseAllOfResultEntities) GetDefectId() int64`

GetDefectId returns the DefectId field if non-nil, zero value otherwise.

### GetDefectIdOk

`func (o *SearchResponseAllOfResultEntities) GetDefectIdOk() (*int64, bool)`

GetDefectIdOk returns a tuple with the DefectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectId

`func (o *SearchResponseAllOfResultEntities) SetDefectId(v int64)`

SetDefectId sets DefectId field to given value.


### GetActualResult

`func (o *SearchResponseAllOfResultEntities) GetActualResult() string`

GetActualResult returns the ActualResult field if non-nil, zero value otherwise.

### GetActualResultOk

`func (o *SearchResponseAllOfResultEntities) GetActualResultOk() (*string, bool)`

GetActualResultOk returns a tuple with the ActualResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualResult

`func (o *SearchResponseAllOfResultEntities) SetActualResult(v string)`

SetActualResult sets ActualResult field to given value.

### HasActualResult

`func (o *SearchResponseAllOfResultEntities) HasActualResult() bool`

HasActualResult returns a boolean if a field has been set.

### GetResolved

`func (o *SearchResponseAllOfResultEntities) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *SearchResponseAllOfResultEntities) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *SearchResponseAllOfResultEntities) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *SearchResponseAllOfResultEntities) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *SearchResponseAllOfResultEntities) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *SearchResponseAllOfResultEntities) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetExternalData

`func (o *SearchResponseAllOfResultEntities) GetExternalData() string`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *SearchResponseAllOfResultEntities) GetExternalDataOk() (*string, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *SearchResponseAllOfResultEntities) SetExternalData(v string)`

SetExternalData sets ExternalData field to given value.

### HasExternalData

`func (o *SearchResponseAllOfResultEntities) HasExternalData() bool`

HasExternalData returns a boolean if a field has been set.

### GetCasesCount

`func (o *SearchResponseAllOfResultEntities) GetCasesCount() int32`

GetCasesCount returns the CasesCount field if non-nil, zero value otherwise.

### GetCasesCountOk

`func (o *SearchResponseAllOfResultEntities) GetCasesCountOk() (*int32, bool)`

GetCasesCountOk returns a tuple with the CasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesCount

`func (o *SearchResponseAllOfResultEntities) SetCasesCount(v int32)`

SetCasesCount sets CasesCount field to given value.

### HasCasesCount

`func (o *SearchResponseAllOfResultEntities) HasCasesCount() bool`

HasCasesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


