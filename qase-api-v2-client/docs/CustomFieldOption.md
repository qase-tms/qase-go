# CustomFieldOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**ReadOnly** | Pointer to **NullableBool** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**IsInternal** | Pointer to **NullableBool** |  | [optional] 
**Behaviour** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCustomFieldOption

`func NewCustomFieldOption() *CustomFieldOption`

NewCustomFieldOption instantiates a new CustomFieldOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldOptionWithDefaults

`func NewCustomFieldOptionWithDefaults() *CustomFieldOption`

NewCustomFieldOptionWithDefaults instantiates a new CustomFieldOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldOption) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldOption) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldOption) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFieldOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *CustomFieldOption) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomFieldOption) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomFieldOption) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomFieldOption) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CustomFieldOption) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CustomFieldOption) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSlug

`func (o *CustomFieldOption) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CustomFieldOption) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CustomFieldOption) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CustomFieldOption) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *CustomFieldOption) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *CustomFieldOption) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetColor

`func (o *CustomFieldOption) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CustomFieldOption) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CustomFieldOption) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CustomFieldOption) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CustomFieldOption) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CustomFieldOption) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIcon

`func (o *CustomFieldOption) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CustomFieldOption) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CustomFieldOption) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CustomFieldOption) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CustomFieldOption) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CustomFieldOption) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetIsDefault

`func (o *CustomFieldOption) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CustomFieldOption) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CustomFieldOption) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CustomFieldOption) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *CustomFieldOption) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *CustomFieldOption) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetReadOnly

`func (o *CustomFieldOption) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *CustomFieldOption) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *CustomFieldOption) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *CustomFieldOption) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *CustomFieldOption) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *CustomFieldOption) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetIsActive

`func (o *CustomFieldOption) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CustomFieldOption) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CustomFieldOption) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CustomFieldOption) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CustomFieldOption) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CustomFieldOption) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsInternal

`func (o *CustomFieldOption) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *CustomFieldOption) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *CustomFieldOption) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.

### HasIsInternal

`func (o *CustomFieldOption) HasIsInternal() bool`

HasIsInternal returns a boolean if a field has been set.

### SetIsInternalNil

`func (o *CustomFieldOption) SetIsInternalNil(b bool)`

 SetIsInternalNil sets the value for IsInternal to be an explicit nil

### UnsetIsInternal
`func (o *CustomFieldOption) UnsetIsInternal()`

UnsetIsInternal ensures that no value is present for IsInternal, not even an explicit nil
### GetBehaviour

`func (o *CustomFieldOption) GetBehaviour() int32`

GetBehaviour returns the Behaviour field if non-nil, zero value otherwise.

### GetBehaviourOk

`func (o *CustomFieldOption) GetBehaviourOk() (*int32, bool)`

GetBehaviourOk returns a tuple with the Behaviour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviour

`func (o *CustomFieldOption) SetBehaviour(v int32)`

SetBehaviour sets Behaviour field to given value.

### HasBehaviour

`func (o *CustomFieldOption) HasBehaviour() bool`

HasBehaviour returns a boolean if a field has been set.

### SetBehaviourNil

`func (o *CustomFieldOption) SetBehaviourNil(b bool)`

 SetBehaviourNil sets the value for Behaviour to be an explicit nil

### UnsetBehaviour
`func (o *CustomFieldOption) UnsetBehaviour()`

UnsetBehaviour ensures that no value is present for Behaviour, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


