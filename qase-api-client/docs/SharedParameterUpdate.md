# SharedParameterUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**ProjectCodes** | Pointer to **[]string** | List of project codes to associate with this shared parameter | [optional] 
**IsEnabledForAllProjects** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to [**SharedParameterParameter**](SharedParameterParameter.md) |  | [optional] 

## Methods

### NewSharedParameterUpdate

`func NewSharedParameterUpdate() *SharedParameterUpdate`

NewSharedParameterUpdate instantiates a new SharedParameterUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedParameterUpdateWithDefaults

`func NewSharedParameterUpdateWithDefaults() *SharedParameterUpdate`

NewSharedParameterUpdateWithDefaults instantiates a new SharedParameterUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SharedParameterUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SharedParameterUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SharedParameterUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SharedParameterUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProjectCodes

`func (o *SharedParameterUpdate) GetProjectCodes() []string`

GetProjectCodes returns the ProjectCodes field if non-nil, zero value otherwise.

### GetProjectCodesOk

`func (o *SharedParameterUpdate) GetProjectCodesOk() (*[]string, bool)`

GetProjectCodesOk returns a tuple with the ProjectCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCodes

`func (o *SharedParameterUpdate) SetProjectCodes(v []string)`

SetProjectCodes sets ProjectCodes field to given value.

### HasProjectCodes

`func (o *SharedParameterUpdate) HasProjectCodes() bool`

HasProjectCodes returns a boolean if a field has been set.

### GetIsEnabledForAllProjects

`func (o *SharedParameterUpdate) GetIsEnabledForAllProjects() bool`

GetIsEnabledForAllProjects returns the IsEnabledForAllProjects field if non-nil, zero value otherwise.

### GetIsEnabledForAllProjectsOk

`func (o *SharedParameterUpdate) GetIsEnabledForAllProjectsOk() (*bool, bool)`

GetIsEnabledForAllProjectsOk returns a tuple with the IsEnabledForAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForAllProjects

`func (o *SharedParameterUpdate) SetIsEnabledForAllProjects(v bool)`

SetIsEnabledForAllProjects sets IsEnabledForAllProjects field to given value.

### HasIsEnabledForAllProjects

`func (o *SharedParameterUpdate) HasIsEnabledForAllProjects() bool`

HasIsEnabledForAllProjects returns a boolean if a field has been set.

### GetParameters

`func (o *SharedParameterUpdate) GetParameters() SharedParameterParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SharedParameterUpdate) GetParametersOk() (*SharedParameterParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SharedParameterUpdate) SetParameters(v SharedParameterParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *SharedParameterUpdate) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


