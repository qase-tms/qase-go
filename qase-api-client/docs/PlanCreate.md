# PlanCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Cases** | **[]int64** |  | 

## Methods

### NewPlanCreate

`func NewPlanCreate(title string, cases []int64, ) *PlanCreate`

NewPlanCreate instantiates a new PlanCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanCreateWithDefaults

`func NewPlanCreateWithDefaults() *PlanCreate`

NewPlanCreateWithDefaults instantiates a new PlanCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PlanCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlanCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlanCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *PlanCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlanCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlanCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCases

`func (o *PlanCreate) GetCases() []int64`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *PlanCreate) GetCasesOk() (*[]int64, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *PlanCreate) SetCases(v []int64)`

SetCases sets Cases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


