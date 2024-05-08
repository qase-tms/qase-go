# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment() *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Environment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Environment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Environment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Environment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Environment) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Environment) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Environment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Environment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Environment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Environment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Environment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Environment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSlug

`func (o *Environment) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Environment) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Environment) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Environment) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetHost

`func (o *Environment) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Environment) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Environment) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Environment) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *Environment) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *Environment) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


