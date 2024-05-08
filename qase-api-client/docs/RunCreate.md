# RunCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IncludeAllCases** | Pointer to **bool** |  | [optional] 
**Cases** | Pointer to **[]int64** |  | [optional] 
**IsAutotest** | Pointer to **bool** |  | [optional] 
**EnvironmentId** | Pointer to **int64** |  | [optional] 
**MilestoneId** | Pointer to **int64** |  | [optional] 
**PlanId** | Pointer to **int64** |  | [optional] 
**AuthorId** | Pointer to **int64** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Configurations** | Pointer to **[]int64** |  | [optional] 
**CustomField** | Pointer to **map[string]string** | A map of custom fields values (id &#x3D;&gt; value) | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 

## Methods

### NewRunCreate

`func NewRunCreate(title string, ) *RunCreate`

NewRunCreate instantiates a new RunCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunCreateWithDefaults

`func NewRunCreateWithDefaults() *RunCreate`

NewRunCreateWithDefaults instantiates a new RunCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RunCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RunCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RunCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *RunCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RunCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RunCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RunCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludeAllCases

`func (o *RunCreate) GetIncludeAllCases() bool`

GetIncludeAllCases returns the IncludeAllCases field if non-nil, zero value otherwise.

### GetIncludeAllCasesOk

`func (o *RunCreate) GetIncludeAllCasesOk() (*bool, bool)`

GetIncludeAllCasesOk returns a tuple with the IncludeAllCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllCases

`func (o *RunCreate) SetIncludeAllCases(v bool)`

SetIncludeAllCases sets IncludeAllCases field to given value.

### HasIncludeAllCases

`func (o *RunCreate) HasIncludeAllCases() bool`

HasIncludeAllCases returns a boolean if a field has been set.

### GetCases

`func (o *RunCreate) GetCases() []int64`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *RunCreate) GetCasesOk() (*[]int64, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *RunCreate) SetCases(v []int64)`

SetCases sets Cases field to given value.

### HasCases

`func (o *RunCreate) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetIsAutotest

`func (o *RunCreate) GetIsAutotest() bool`

GetIsAutotest returns the IsAutotest field if non-nil, zero value otherwise.

### GetIsAutotestOk

`func (o *RunCreate) GetIsAutotestOk() (*bool, bool)`

GetIsAutotestOk returns a tuple with the IsAutotest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutotest

`func (o *RunCreate) SetIsAutotest(v bool)`

SetIsAutotest sets IsAutotest field to given value.

### HasIsAutotest

`func (o *RunCreate) HasIsAutotest() bool`

HasIsAutotest returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *RunCreate) GetEnvironmentId() int64`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *RunCreate) GetEnvironmentIdOk() (*int64, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *RunCreate) SetEnvironmentId(v int64)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *RunCreate) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetMilestoneId

`func (o *RunCreate) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *RunCreate) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *RunCreate) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *RunCreate) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### GetPlanId

`func (o *RunCreate) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RunCreate) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RunCreate) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RunCreate) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetAuthorId

`func (o *RunCreate) GetAuthorId() int64`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *RunCreate) GetAuthorIdOk() (*int64, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *RunCreate) SetAuthorId(v int64)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *RunCreate) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetTags

`func (o *RunCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RunCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RunCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RunCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfigurations

`func (o *RunCreate) GetConfigurations() []int64`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *RunCreate) GetConfigurationsOk() (*[]int64, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *RunCreate) SetConfigurations(v []int64)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *RunCreate) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetCustomField

`func (o *RunCreate) GetCustomField() map[string]string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *RunCreate) GetCustomFieldOk() (*map[string]string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *RunCreate) SetCustomField(v map[string]string)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *RunCreate) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.

### GetStartTime

`func (o *RunCreate) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RunCreate) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RunCreate) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RunCreate) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *RunCreate) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RunCreate) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RunCreate) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RunCreate) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


