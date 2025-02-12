# RunQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**RunId** | **int64** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
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
**PlanId** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewRunQuery

`func NewRunQuery(runId int64, ) *RunQuery`

NewRunQuery instantiates a new RunQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunQueryWithDefaults

`func NewRunQueryWithDefaults() *RunQuery`

NewRunQueryWithDefaults instantiates a new RunQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunQuery) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunQuery) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunQuery) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RunQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRunId

`func (o *RunQuery) GetRunId() int64`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *RunQuery) GetRunIdOk() (*int64, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *RunQuery) SetRunId(v int64)`

SetRunId sets RunId field to given value.


### GetTitle

`func (o *RunQuery) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RunQuery) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RunQuery) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RunQuery) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *RunQuery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RunQuery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RunQuery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RunQuery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RunQuery) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RunQuery) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *RunQuery) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunQuery) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunQuery) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RunQuery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusText

`func (o *RunQuery) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *RunQuery) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *RunQuery) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *RunQuery) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetStartTime

`func (o *RunQuery) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RunQuery) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RunQuery) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RunQuery) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *RunQuery) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *RunQuery) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *RunQuery) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RunQuery) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RunQuery) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RunQuery) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *RunQuery) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *RunQuery) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetPublic

`func (o *RunQuery) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *RunQuery) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *RunQuery) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *RunQuery) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetStats

`func (o *RunQuery) GetStats() RunStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *RunQuery) GetStatsOk() (*RunStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *RunQuery) SetStats(v RunStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *RunQuery) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetTimeSpent

`func (o *RunQuery) GetTimeSpent() int64`

GetTimeSpent returns the TimeSpent field if non-nil, zero value otherwise.

### GetTimeSpentOk

`func (o *RunQuery) GetTimeSpentOk() (*int64, bool)`

GetTimeSpentOk returns a tuple with the TimeSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSpent

`func (o *RunQuery) SetTimeSpent(v int64)`

SetTimeSpent sets TimeSpent field to given value.

### HasTimeSpent

`func (o *RunQuery) HasTimeSpent() bool`

HasTimeSpent returns a boolean if a field has been set.

### GetEnvironment

`func (o *RunQuery) GetEnvironment() RunEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RunQuery) GetEnvironmentOk() (*RunEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RunQuery) SetEnvironment(v RunEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RunQuery) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *RunQuery) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *RunQuery) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetMilestone

`func (o *RunQuery) GetMilestone() RunMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *RunQuery) GetMilestoneOk() (*RunMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *RunQuery) SetMilestone(v RunMilestone)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *RunQuery) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### SetMilestoneNil

`func (o *RunQuery) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *RunQuery) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetCustomFields

`func (o *RunQuery) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RunQuery) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RunQuery) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RunQuery) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *RunQuery) GetTags() []TagValue`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RunQuery) GetTagsOk() (*[]TagValue, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RunQuery) SetTags(v []TagValue)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RunQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCases

`func (o *RunQuery) GetCases() []int64`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *RunQuery) GetCasesOk() (*[]int64, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *RunQuery) SetCases(v []int64)`

SetCases sets Cases field to given value.

### HasCases

`func (o *RunQuery) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetPlanId

`func (o *RunQuery) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RunQuery) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RunQuery) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RunQuery) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *RunQuery) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *RunQuery) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


