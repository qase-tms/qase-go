# PlanUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Cases** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewPlanUpdate

`func NewPlanUpdate() *PlanUpdate`

NewPlanUpdate instantiates a new PlanUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanUpdateWithDefaults

`func NewPlanUpdateWithDefaults() *PlanUpdate`

NewPlanUpdateWithDefaults instantiates a new PlanUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PlanUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlanUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlanUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlanUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PlanUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlanUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlanUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCases

`func (o *PlanUpdate) GetCases() []int64`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *PlanUpdate) GetCasesOk() (*[]int64, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *PlanUpdate) SetCases(v []int64)`

SetCases sets Cases field to given value.

### HasCases

`func (o *PlanUpdate) HasCases() bool`

HasCases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


