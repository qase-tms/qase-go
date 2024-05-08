# PlanQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CasesCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPlanQuery

`func NewPlanQuery() *PlanQuery`

NewPlanQuery instantiates a new PlanQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanQueryWithDefaults

`func NewPlanQueryWithDefaults() *PlanQuery`

NewPlanQueryWithDefaults instantiates a new PlanQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanQuery) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanQuery) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanQuery) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PlanQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *PlanQuery) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlanQuery) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlanQuery) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlanQuery) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PlanQuery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanQuery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanQuery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanQuery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlanQuery) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlanQuery) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCasesCount

`func (o *PlanQuery) GetCasesCount() int32`

GetCasesCount returns the CasesCount field if non-nil, zero value otherwise.

### GetCasesCountOk

`func (o *PlanQuery) GetCasesCountOk() (*int32, bool)`

GetCasesCountOk returns a tuple with the CasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesCount

`func (o *PlanQuery) SetCasesCount(v int32)`

SetCasesCount sets CasesCount field to given value.

### HasCasesCount

`func (o *PlanQuery) HasCasesCount() bool`

HasCasesCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlanQuery) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanQuery) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanQuery) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlanQuery) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlanQuery) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlanQuery) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlanQuery) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlanQuery) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


