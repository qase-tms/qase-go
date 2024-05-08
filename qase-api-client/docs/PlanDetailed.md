# PlanDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CasesCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Created** | Pointer to **string** | Deprecated, use the &#x60;created_at&#x60; property instead. | [optional] 
**Updated** | Pointer to **string** | Deprecated, use the &#x60;updated_at&#x60; property instead. | [optional] 
**AverageTime** | Pointer to **float32** |  | [optional] 
**Cases** | Pointer to [**[]PlanDetailedAllOfCases**](PlanDetailedAllOfCases.md) |  | [optional] 

## Methods

### NewPlanDetailed

`func NewPlanDetailed() *PlanDetailed`

NewPlanDetailed instantiates a new PlanDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanDetailedWithDefaults

`func NewPlanDetailedWithDefaults() *PlanDetailed`

NewPlanDetailedWithDefaults instantiates a new PlanDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanDetailed) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanDetailed) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanDetailed) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PlanDetailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *PlanDetailed) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlanDetailed) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlanDetailed) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlanDetailed) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PlanDetailed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanDetailed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanDetailed) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanDetailed) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlanDetailed) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlanDetailed) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCasesCount

`func (o *PlanDetailed) GetCasesCount() int32`

GetCasesCount returns the CasesCount field if non-nil, zero value otherwise.

### GetCasesCountOk

`func (o *PlanDetailed) GetCasesCountOk() (*int32, bool)`

GetCasesCountOk returns a tuple with the CasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesCount

`func (o *PlanDetailed) SetCasesCount(v int32)`

SetCasesCount sets CasesCount field to given value.

### HasCasesCount

`func (o *PlanDetailed) HasCasesCount() bool`

HasCasesCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlanDetailed) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanDetailed) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanDetailed) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlanDetailed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlanDetailed) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlanDetailed) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlanDetailed) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlanDetailed) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreated

`func (o *PlanDetailed) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PlanDetailed) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PlanDetailed) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PlanDetailed) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *PlanDetailed) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PlanDetailed) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PlanDetailed) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PlanDetailed) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAverageTime

`func (o *PlanDetailed) GetAverageTime() float32`

GetAverageTime returns the AverageTime field if non-nil, zero value otherwise.

### GetAverageTimeOk

`func (o *PlanDetailed) GetAverageTimeOk() (*float32, bool)`

GetAverageTimeOk returns a tuple with the AverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTime

`func (o *PlanDetailed) SetAverageTime(v float32)`

SetAverageTime sets AverageTime field to given value.

### HasAverageTime

`func (o *PlanDetailed) HasAverageTime() bool`

HasAverageTime returns a boolean if a field has been set.

### GetCases

`func (o *PlanDetailed) GetCases() []PlanDetailedAllOfCases`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *PlanDetailed) GetCasesOk() (*[]PlanDetailedAllOfCases, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *PlanDetailed) SetCases(v []PlanDetailedAllOfCases)`

SetCases sets Cases field to given value.

### HasCases

`func (o *PlanDetailed) HasCases() bool`

HasCases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


