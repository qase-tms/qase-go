# Runupdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnvironmentId** | Pointer to **NullableInt64** |  | [optional] 
**EnvironmentSlug** | Pointer to **NullableString** |  | [optional] 
**MilestoneId** | Pointer to **NullableInt64** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Configurations** | Pointer to **[]int64** |  | [optional] 
**CustomField** | Pointer to **map[string]string** | A map of custom fields values (id &#x3D;&gt; value) | [optional] 

## Methods

### NewRunupdate

`func NewRunupdate() *Runupdate`

NewRunupdate instantiates a new Runupdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunupdateWithDefaults

`func NewRunupdateWithDefaults() *Runupdate`

NewRunupdateWithDefaults instantiates a new Runupdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Runupdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Runupdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Runupdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Runupdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Runupdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Runupdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Runupdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Runupdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Runupdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Runupdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnvironmentId

`func (o *Runupdate) GetEnvironmentId() int64`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Runupdate) GetEnvironmentIdOk() (*int64, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Runupdate) SetEnvironmentId(v int64)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *Runupdate) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *Runupdate) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *Runupdate) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetEnvironmentSlug

`func (o *Runupdate) GetEnvironmentSlug() string`

GetEnvironmentSlug returns the EnvironmentSlug field if non-nil, zero value otherwise.

### GetEnvironmentSlugOk

`func (o *Runupdate) GetEnvironmentSlugOk() (*string, bool)`

GetEnvironmentSlugOk returns a tuple with the EnvironmentSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentSlug

`func (o *Runupdate) SetEnvironmentSlug(v string)`

SetEnvironmentSlug sets EnvironmentSlug field to given value.

### HasEnvironmentSlug

`func (o *Runupdate) HasEnvironmentSlug() bool`

HasEnvironmentSlug returns a boolean if a field has been set.

### SetEnvironmentSlugNil

`func (o *Runupdate) SetEnvironmentSlugNil(b bool)`

 SetEnvironmentSlugNil sets the value for EnvironmentSlug to be an explicit nil

### UnsetEnvironmentSlug
`func (o *Runupdate) UnsetEnvironmentSlug()`

UnsetEnvironmentSlug ensures that no value is present for EnvironmentSlug, not even an explicit nil
### GetMilestoneId

`func (o *Runupdate) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *Runupdate) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *Runupdate) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *Runupdate) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *Runupdate) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *Runupdate) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetTags

`func (o *Runupdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Runupdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Runupdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Runupdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Runupdate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Runupdate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetConfigurations

`func (o *Runupdate) GetConfigurations() []int64`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *Runupdate) GetConfigurationsOk() (*[]int64, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *Runupdate) SetConfigurations(v []int64)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *Runupdate) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### SetConfigurationsNil

`func (o *Runupdate) SetConfigurationsNil(b bool)`

 SetConfigurationsNil sets the value for Configurations to be an explicit nil

### UnsetConfigurations
`func (o *Runupdate) UnsetConfigurations()`

UnsetConfigurations ensures that no value is present for Configurations, not even an explicit nil
### GetCustomField

`func (o *Runupdate) GetCustomField() map[string]string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *Runupdate) GetCustomFieldOk() (*map[string]string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *Runupdate) SetCustomField(v map[string]string)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *Runupdate) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.

### SetCustomFieldNil

`func (o *Runupdate) SetCustomFieldNil(b bool)`

 SetCustomFieldNil sets the value for CustomField to be an explicit nil

### UnsetCustomField
`func (o *Runupdate) UnsetCustomField()`

UnsetCustomField ensures that no value is present for CustomField, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


