# SharedParameterCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Type** | **string** |  | 
**ProjectCodes** | Pointer to **[]string** | List of project codes to associate with this shared parameter | [optional] 
**IsEnabledForAllProjects** | **bool** |  | 
**Parameters** | [**SharedParameterParameter**](SharedParameterParameter.md) |  | 

## Methods

### NewSharedParameterCreate

`func NewSharedParameterCreate(title string, type_ string, isEnabledForAllProjects bool, parameters SharedParameterParameter, ) *SharedParameterCreate`

NewSharedParameterCreate instantiates a new SharedParameterCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedParameterCreateWithDefaults

`func NewSharedParameterCreateWithDefaults() *SharedParameterCreate`

NewSharedParameterCreateWithDefaults instantiates a new SharedParameterCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SharedParameterCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SharedParameterCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SharedParameterCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *SharedParameterCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SharedParameterCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SharedParameterCreate) SetType(v string)`

SetType sets Type field to given value.


### GetProjectCodes

`func (o *SharedParameterCreate) GetProjectCodes() []string`

GetProjectCodes returns the ProjectCodes field if non-nil, zero value otherwise.

### GetProjectCodesOk

`func (o *SharedParameterCreate) GetProjectCodesOk() (*[]string, bool)`

GetProjectCodesOk returns a tuple with the ProjectCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCodes

`func (o *SharedParameterCreate) SetProjectCodes(v []string)`

SetProjectCodes sets ProjectCodes field to given value.

### HasProjectCodes

`func (o *SharedParameterCreate) HasProjectCodes() bool`

HasProjectCodes returns a boolean if a field has been set.

### GetIsEnabledForAllProjects

`func (o *SharedParameterCreate) GetIsEnabledForAllProjects() bool`

GetIsEnabledForAllProjects returns the IsEnabledForAllProjects field if non-nil, zero value otherwise.

### GetIsEnabledForAllProjectsOk

`func (o *SharedParameterCreate) GetIsEnabledForAllProjectsOk() (*bool, bool)`

GetIsEnabledForAllProjectsOk returns a tuple with the IsEnabledForAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForAllProjects

`func (o *SharedParameterCreate) SetIsEnabledForAllProjects(v bool)`

SetIsEnabledForAllProjects sets IsEnabledForAllProjects field to given value.


### GetParameters

`func (o *SharedParameterCreate) GetParameters() SharedParameterParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SharedParameterCreate) GetParametersOk() (*SharedParameterParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SharedParameterCreate) SetParameters(v SharedParameterParameter)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


