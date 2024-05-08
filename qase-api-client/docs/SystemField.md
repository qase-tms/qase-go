# SystemField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**Entity** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **int64** |  | [optional] 
**Options** | Pointer to [**[]SystemFieldOption**](SystemFieldOption.md) |  | [optional] 

## Methods

### NewSystemField

`func NewSystemField() *SystemField`

NewSystemField instantiates a new SystemField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemFieldWithDefaults

`func NewSystemFieldWithDefaults() *SystemField`

NewSystemFieldWithDefaults instantiates a new SystemField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SystemField) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SystemField) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SystemField) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SystemField) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSlug

`func (o *SystemField) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SystemField) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SystemField) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *SystemField) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SystemField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SystemField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SystemField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SystemField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *SystemField) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *SystemField) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetIsRequired

`func (o *SystemField) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *SystemField) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *SystemField) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *SystemField) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetEntity

`func (o *SystemField) GetEntity() int64`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SystemField) GetEntityOk() (*int64, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SystemField) SetEntity(v int64)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SystemField) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetType

`func (o *SystemField) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemField) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemField) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *SystemField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *SystemField) GetOptions() []SystemFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SystemField) GetOptionsOk() (*[]SystemFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SystemField) SetOptions(v []SystemFieldOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SystemField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *SystemField) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *SystemField) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


