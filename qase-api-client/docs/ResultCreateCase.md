# ResultCreateCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**SuiteTitle** | Pointer to **NullableString** | Nested suites should be separated with &#x60;TAB&#x60; symbol. | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Preconditions** | Pointer to **NullableString** |  | [optional] 
**Postconditions** | Pointer to **NullableString** |  | [optional] 
**Layer** | Pointer to **string** | Slug of the layer. You can get it in the System Field settings. | [optional] 
**Severity** | Pointer to **string** | Slug of the severity. You can get it in the System Field settings. | [optional] 
**Priority** | Pointer to **string** | Slug of the priority. You can get it in the System Field settings. | [optional] 

## Methods

### NewResultCreateCase

`func NewResultCreateCase() *ResultCreateCase`

NewResultCreateCase instantiates a new ResultCreateCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCreateCaseWithDefaults

`func NewResultCreateCaseWithDefaults() *ResultCreateCase`

NewResultCreateCaseWithDefaults instantiates a new ResultCreateCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ResultCreateCase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResultCreateCase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResultCreateCase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResultCreateCase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSuiteTitle

`func (o *ResultCreateCase) GetSuiteTitle() string`

GetSuiteTitle returns the SuiteTitle field if non-nil, zero value otherwise.

### GetSuiteTitleOk

`func (o *ResultCreateCase) GetSuiteTitleOk() (*string, bool)`

GetSuiteTitleOk returns a tuple with the SuiteTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteTitle

`func (o *ResultCreateCase) SetSuiteTitle(v string)`

SetSuiteTitle sets SuiteTitle field to given value.

### HasSuiteTitle

`func (o *ResultCreateCase) HasSuiteTitle() bool`

HasSuiteTitle returns a boolean if a field has been set.

### SetSuiteTitleNil

`func (o *ResultCreateCase) SetSuiteTitleNil(b bool)`

 SetSuiteTitleNil sets the value for SuiteTitle to be an explicit nil

### UnsetSuiteTitle
`func (o *ResultCreateCase) UnsetSuiteTitle()`

UnsetSuiteTitle ensures that no value is present for SuiteTitle, not even an explicit nil
### GetDescription

`func (o *ResultCreateCase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResultCreateCase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResultCreateCase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResultCreateCase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResultCreateCase) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResultCreateCase) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPreconditions

`func (o *ResultCreateCase) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *ResultCreateCase) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *ResultCreateCase) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *ResultCreateCase) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### SetPreconditionsNil

`func (o *ResultCreateCase) SetPreconditionsNil(b bool)`

 SetPreconditionsNil sets the value for Preconditions to be an explicit nil

### UnsetPreconditions
`func (o *ResultCreateCase) UnsetPreconditions()`

UnsetPreconditions ensures that no value is present for Preconditions, not even an explicit nil
### GetPostconditions

`func (o *ResultCreateCase) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *ResultCreateCase) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *ResultCreateCase) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *ResultCreateCase) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### SetPostconditionsNil

`func (o *ResultCreateCase) SetPostconditionsNil(b bool)`

 SetPostconditionsNil sets the value for Postconditions to be an explicit nil

### UnsetPostconditions
`func (o *ResultCreateCase) UnsetPostconditions()`

UnsetPostconditions ensures that no value is present for Postconditions, not even an explicit nil
### GetLayer

`func (o *ResultCreateCase) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *ResultCreateCase) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *ResultCreateCase) SetLayer(v string)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *ResultCreateCase) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetSeverity

`func (o *ResultCreateCase) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ResultCreateCase) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ResultCreateCase) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ResultCreateCase) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetPriority

`func (o *ResultCreateCase) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ResultCreateCase) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ResultCreateCase) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ResultCreateCase) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


