# CustomFieldUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Value** | Pointer to [**[]CustomFieldCreateValueInner**](CustomFieldCreateValueInner.md) |  | [optional] 
**ReplaceValues** | Pointer to **map[string]string** | Dictionary of old values and their replacemants | [optional] 
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**IsFilterable** | Pointer to **bool** |  | [optional] 
**IsVisible** | Pointer to **bool** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**IsEnabledForAllProjects** | Pointer to **bool** |  | [optional] 
**ProjectsCodes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCustomFieldUpdate

`func NewCustomFieldUpdate(title string, ) *CustomFieldUpdate`

NewCustomFieldUpdate instantiates a new CustomFieldUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldUpdateWithDefaults

`func NewCustomFieldUpdateWithDefaults() *CustomFieldUpdate`

NewCustomFieldUpdateWithDefaults instantiates a new CustomFieldUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CustomFieldUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomFieldUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomFieldUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetValue

`func (o *CustomFieldUpdate) GetValue() []CustomFieldCreateValueInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldUpdate) GetValueOk() (*[]CustomFieldCreateValueInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldUpdate) SetValue(v []CustomFieldCreateValueInner)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldUpdate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomFieldUpdate) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomFieldUpdate) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetReplaceValues

`func (o *CustomFieldUpdate) GetReplaceValues() map[string]string`

GetReplaceValues returns the ReplaceValues field if non-nil, zero value otherwise.

### GetReplaceValuesOk

`func (o *CustomFieldUpdate) GetReplaceValuesOk() (*map[string]string, bool)`

GetReplaceValuesOk returns a tuple with the ReplaceValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceValues

`func (o *CustomFieldUpdate) SetReplaceValues(v map[string]string)`

SetReplaceValues sets ReplaceValues field to given value.

### HasReplaceValues

`func (o *CustomFieldUpdate) HasReplaceValues() bool`

HasReplaceValues returns a boolean if a field has been set.

### SetReplaceValuesNil

`func (o *CustomFieldUpdate) SetReplaceValuesNil(b bool)`

 SetReplaceValuesNil sets the value for ReplaceValues to be an explicit nil

### UnsetReplaceValues
`func (o *CustomFieldUpdate) UnsetReplaceValues()`

UnsetReplaceValues ensures that no value is present for ReplaceValues, not even an explicit nil
### GetPlaceholder

`func (o *CustomFieldUpdate) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *CustomFieldUpdate) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *CustomFieldUpdate) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *CustomFieldUpdate) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *CustomFieldUpdate) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *CustomFieldUpdate) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetDefaultValue

`func (o *CustomFieldUpdate) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CustomFieldUpdate) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CustomFieldUpdate) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CustomFieldUpdate) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CustomFieldUpdate) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CustomFieldUpdate) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetIsFilterable

`func (o *CustomFieldUpdate) GetIsFilterable() bool`

GetIsFilterable returns the IsFilterable field if non-nil, zero value otherwise.

### GetIsFilterableOk

`func (o *CustomFieldUpdate) GetIsFilterableOk() (*bool, bool)`

GetIsFilterableOk returns a tuple with the IsFilterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFilterable

`func (o *CustomFieldUpdate) SetIsFilterable(v bool)`

SetIsFilterable sets IsFilterable field to given value.

### HasIsFilterable

`func (o *CustomFieldUpdate) HasIsFilterable() bool`

HasIsFilterable returns a boolean if a field has been set.

### GetIsVisible

`func (o *CustomFieldUpdate) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *CustomFieldUpdate) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *CustomFieldUpdate) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *CustomFieldUpdate) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetIsRequired

`func (o *CustomFieldUpdate) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CustomFieldUpdate) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CustomFieldUpdate) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *CustomFieldUpdate) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsEnabledForAllProjects

`func (o *CustomFieldUpdate) GetIsEnabledForAllProjects() bool`

GetIsEnabledForAllProjects returns the IsEnabledForAllProjects field if non-nil, zero value otherwise.

### GetIsEnabledForAllProjectsOk

`func (o *CustomFieldUpdate) GetIsEnabledForAllProjectsOk() (*bool, bool)`

GetIsEnabledForAllProjectsOk returns a tuple with the IsEnabledForAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForAllProjects

`func (o *CustomFieldUpdate) SetIsEnabledForAllProjects(v bool)`

SetIsEnabledForAllProjects sets IsEnabledForAllProjects field to given value.

### HasIsEnabledForAllProjects

`func (o *CustomFieldUpdate) HasIsEnabledForAllProjects() bool`

HasIsEnabledForAllProjects returns a boolean if a field has been set.

### GetProjectsCodes

`func (o *CustomFieldUpdate) GetProjectsCodes() []string`

GetProjectsCodes returns the ProjectsCodes field if non-nil, zero value otherwise.

### GetProjectsCodesOk

`func (o *CustomFieldUpdate) GetProjectsCodesOk() (*[]string, bool)`

GetProjectsCodesOk returns a tuple with the ProjectsCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCodes

`func (o *CustomFieldUpdate) SetProjectsCodes(v []string)`

SetProjectsCodes sets ProjectsCodes field to given value.

### HasProjectsCodes

`func (o *CustomFieldUpdate) HasProjectsCodes() bool`

HasProjectsCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


