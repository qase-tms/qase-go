# SharedParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**ProjectCodes** | **[]string** |  | 
**IsEnabledForAllProjects** | **bool** |  | 
**Parameters** | [**SharedParameterParameter**](SharedParameterParameter.md) |  | 

## Methods

### NewSharedParameter

`func NewSharedParameter(id string, title string, type_ string, projectCodes []string, isEnabledForAllProjects bool, parameters SharedParameterParameter, ) *SharedParameter`

NewSharedParameter instantiates a new SharedParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedParameterWithDefaults

`func NewSharedParameterWithDefaults() *SharedParameter`

NewSharedParameterWithDefaults instantiates a new SharedParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedParameter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedParameter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedParameter) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *SharedParameter) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SharedParameter) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SharedParameter) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *SharedParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SharedParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SharedParameter) SetType(v string)`

SetType sets Type field to given value.


### GetProjectCodes

`func (o *SharedParameter) GetProjectCodes() []string`

GetProjectCodes returns the ProjectCodes field if non-nil, zero value otherwise.

### GetProjectCodesOk

`func (o *SharedParameter) GetProjectCodesOk() (*[]string, bool)`

GetProjectCodesOk returns a tuple with the ProjectCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCodes

`func (o *SharedParameter) SetProjectCodes(v []string)`

SetProjectCodes sets ProjectCodes field to given value.


### GetIsEnabledForAllProjects

`func (o *SharedParameter) GetIsEnabledForAllProjects() bool`

GetIsEnabledForAllProjects returns the IsEnabledForAllProjects field if non-nil, zero value otherwise.

### GetIsEnabledForAllProjectsOk

`func (o *SharedParameter) GetIsEnabledForAllProjectsOk() (*bool, bool)`

GetIsEnabledForAllProjectsOk returns a tuple with the IsEnabledForAllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForAllProjects

`func (o *SharedParameter) SetIsEnabledForAllProjects(v bool)`

SetIsEnabledForAllProjects sets IsEnabledForAllProjects field to given value.


### GetParameters

`func (o *SharedParameter) GetParameters() SharedParameterParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SharedParameter) GetParametersOk() (*SharedParameterParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SharedParameter) SetParameters(v SharedParameterParameter)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


