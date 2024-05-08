# CustomFieldCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Value** | Pointer to [**[]CustomFieldCreateValueInner**](CustomFieldCreateValueInner.md) | Required if type one of: 3 - selectbox; 5 - radio; 6 - multiselect;  | [optional] 
**Entity** | **int32** | Possible values: 0 - case; 1 - run; 2 - defect;  | 
**Type** | **int32** | Possible values: 0 - number; 1 - string; 2 - text; 3 - selectbox; 4 - checkbox; 5 - radio; 6 - multiselect; 7 - url; 8 - user; 9 - datetime;  | 
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**IsFilterable** | Pointer to **bool** |  | [optional] 
**IsVisible** | Pointer to **bool** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**IsEnabledForAllProjects** | Pointer to **bool** |  | [optional] 
**ProjectsCodes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCustomFieldCreate

`func NewCustomFieldCreate(title string, entity int32, type_ int32, ) *CustomFieldCreate`

NewCustomFieldCreate instantiates a new CustomFieldCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldCreateWithDefaults

`func NewCustomFieldCreateWithDefaults() *CustomFieldCreate`

NewCustomFieldCreateWithDefaults instantiates a new CustomFieldCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CustomFieldCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomFieldCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomFieldCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetValue

`func (o *CustomFieldCreate) GetValue() []CustomFieldCreateValueInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldCreate) GetValueOk() (*[]CustomFieldCreateValueInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldCreate) SetValue(v []CustomFieldCreateValueInner)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldCreate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomFieldCreate) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomFieldCreate) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetEntity

`func (o *CustomFieldCreate) GetEntity() int32`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CustomFieldCreate) GetEntityOk() (*int32, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CustomFieldCreate) SetEntity(v int32)`

SetEntity sets Entity field to given value.


### GetType

`func (o *CustomFieldCreate) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldCreate) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldCreate) SetType(v int32)`

SetType sets Type field to given value.


### GetPlaceholder

`func (o *CustomFieldCreate) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *CustomFieldCreate) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *CustomFieldCreate) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *CustomFieldCreate) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *CustomFieldCreate) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *CustomFieldCreate) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetDefaultValue

`func (o *CustomFieldCreate) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CustomFieldCreate) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CustomFieldCreate) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CustomFieldCreate) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CustomFieldCreate) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CustomFieldCreate) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetIsFilterable

`func (o *CustomFieldCreate) GetIsFilterable() bool`

GetIsFilterable returns the IsFilterable field if non-nil, zero value otherwise.

### GetIsFilterableOk

`func (o *CustomFieldCreate) GetIsFilterableOk() (*bool, bool)`

GetIsFilterableOk returns a tuple with the IsFilterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFilterable

`func (o *CustomFieldCreate) SetIsFilterable(v bool)`

SetIsFilterable sets IsFilterable field to given value.

### HasIsFilterable

`func (o *CustomFieldCreate) HasIsFilterable() bool`

HasIsFilterable returns a boolean if a field has been set.

### GetIsVisible

`func (o *CustomFieldCreate) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *CustomFieldCreate) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *CustomFieldCreate) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *CustomFieldCreate) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetIsRequired

`func (o *CustomFieldCreate) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CustomFieldCreate) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CustomFieldCreate) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *CustomFieldCreate) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsEnabledForAllProjects

`func (o *CustomFieldCreate) GetIsEnabledForAllProjects() bool`

GetIsEnabledForAllProjects returns the IsEnabledForAllProjects field if non-nil, zero value otherwise.

### GetIsEnabledForAllProjectsOk

`func (o *CustomFieldCreate) GetIsEnabledForAllProjectsOk() (*bool, bool)`

GetIsEnabledForAllProjectsOk returns a tuple with the IsEnabledForAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForAllProjects

`func (o *CustomFieldCreate) SetIsEnabledForAllProjects(v bool)`

SetIsEnabledForAllProjects sets IsEnabledForAllProjects field to given value.

### HasIsEnabledForAllProjects

`func (o *CustomFieldCreate) HasIsEnabledForAllProjects() bool`

HasIsEnabledForAllProjects returns a boolean if a field has been set.

### GetProjectsCodes

`func (o *CustomFieldCreate) GetProjectsCodes() []string`

GetProjectsCodes returns the ProjectsCodes field if non-nil, zero value otherwise.

### GetProjectsCodesOk

`func (o *CustomFieldCreate) GetProjectsCodesOk() (*[]string, bool)`

GetProjectsCodesOk returns a tuple with the ProjectsCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCodes

`func (o *CustomFieldCreate) SetProjectsCodes(v []string)`

SetProjectsCodes sets ProjectsCodes field to given value.

### HasProjectsCodes

`func (o *CustomFieldCreate) HasProjectsCodes() bool`

HasProjectsCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


