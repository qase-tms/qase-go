# CustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**IsVisible** | Pointer to **bool** |  | [optional] 
**IsFilterable** | Pointer to **bool** |  | [optional] 
**IsEnabledForAllProjects** | Pointer to **bool** |  | [optional] 
**Created** | Pointer to **string** | Deprecated, use the &#x60;created_at&#x60; property instead. | [optional] 
**Updated** | Pointer to **NullableString** | Deprecated, use the &#x60;updated_at&#x60; property instead. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ProjectsCodes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCustomField

`func NewCustomField() *CustomField`

NewCustomField instantiates a new CustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldWithDefaults

`func NewCustomFieldWithDefaults() *CustomField`

NewCustomFieldWithDefaults instantiates a new CustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomField) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomField) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomField) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CustomField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *CustomField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomField) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomField) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEntity

`func (o *CustomField) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CustomField) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CustomField) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CustomField) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetType

`func (o *CustomField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPlaceholder

`func (o *CustomField) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *CustomField) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *CustomField) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *CustomField) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *CustomField) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *CustomField) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetDefaultValue

`func (o *CustomField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CustomField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CustomField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CustomField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *CustomField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *CustomField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetValue

`func (o *CustomField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomField) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomField) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsRequired

`func (o *CustomField) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CustomField) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CustomField) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *CustomField) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsVisible

`func (o *CustomField) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *CustomField) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *CustomField) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *CustomField) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetIsFilterable

`func (o *CustomField) GetIsFilterable() bool`

GetIsFilterable returns the IsFilterable field if non-nil, zero value otherwise.

### GetIsFilterableOk

`func (o *CustomField) GetIsFilterableOk() (*bool, bool)`

GetIsFilterableOk returns a tuple with the IsFilterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFilterable

`func (o *CustomField) SetIsFilterable(v bool)`

SetIsFilterable sets IsFilterable field to given value.

### HasIsFilterable

`func (o *CustomField) HasIsFilterable() bool`

HasIsFilterable returns a boolean if a field has been set.

### GetIsEnabledForAllProjects

`func (o *CustomField) GetIsEnabledForAllProjects() bool`

GetIsEnabledForAllProjects returns the IsEnabledForAllProjects field if non-nil, zero value otherwise.

### GetIsEnabledForAllProjectsOk

`func (o *CustomField) GetIsEnabledForAllProjectsOk() (*bool, bool)`

GetIsEnabledForAllProjectsOk returns a tuple with the IsEnabledForAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForAllProjects

`func (o *CustomField) SetIsEnabledForAllProjects(v bool)`

SetIsEnabledForAllProjects sets IsEnabledForAllProjects field to given value.

### HasIsEnabledForAllProjects

`func (o *CustomField) HasIsEnabledForAllProjects() bool`

HasIsEnabledForAllProjects returns a boolean if a field has been set.

### GetCreated

`func (o *CustomField) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomField) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomField) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CustomField) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *CustomField) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CustomField) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CustomField) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CustomField) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *CustomField) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *CustomField) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetCreatedAt

`func (o *CustomField) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomField) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomField) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomField) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomField) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomField) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomField) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomField) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProjectsCodes

`func (o *CustomField) GetProjectsCodes() []string`

GetProjectsCodes returns the ProjectsCodes field if non-nil, zero value otherwise.

### GetProjectsCodesOk

`func (o *CustomField) GetProjectsCodesOk() (*[]string, bool)`

GetProjectsCodesOk returns a tuple with the ProjectsCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCodes

`func (o *CustomField) SetProjectsCodes(v []string)`

SetProjectsCodes sets ProjectsCodes field to given value.

### HasProjectsCodes

`func (o *CustomField) HasProjectsCodes() bool`

HasProjectsCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


