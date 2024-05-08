# EnvironmentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Slug** | **string** |  | 
**Host** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentCreate

`func NewEnvironmentCreate(title string, slug string, ) *EnvironmentCreate`

NewEnvironmentCreate instantiates a new EnvironmentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCreateWithDefaults

`func NewEnvironmentCreateWithDefaults() *EnvironmentCreate`

NewEnvironmentCreateWithDefaults instantiates a new EnvironmentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *EnvironmentCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EnvironmentCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EnvironmentCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *EnvironmentCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSlug

`func (o *EnvironmentCreate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EnvironmentCreate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EnvironmentCreate) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetHost

`func (o *EnvironmentCreate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EnvironmentCreate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EnvironmentCreate) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EnvironmentCreate) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


