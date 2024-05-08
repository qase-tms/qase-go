# ProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Project title. | 
**Code** | **string** | Project code. Unique for team. Digits and special characters are not allowed. | 
**Description** | Pointer to **string** | Project description. | [optional] 
**Access** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** | Team group hash. Required if access param is set to group. | [optional] 
**Settings** | Pointer to **map[string]interface{}** | Additional project settings. | [optional] 

## Methods

### NewProjectCreate

`func NewProjectCreate(title string, code string, ) *ProjectCreate`

NewProjectCreate instantiates a new ProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateWithDefaults

`func NewProjectCreateWithDefaults() *ProjectCreate`

NewProjectCreateWithDefaults instantiates a new ProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ProjectCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCode

`func (o *ProjectCreate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProjectCreate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProjectCreate) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *ProjectCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccess

`func (o *ProjectCreate) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ProjectCreate) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ProjectCreate) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ProjectCreate) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetGroup

`func (o *ProjectCreate) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ProjectCreate) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ProjectCreate) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ProjectCreate) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSettings

`func (o *ProjectCreate) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ProjectCreate) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ProjectCreate) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ProjectCreate) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


