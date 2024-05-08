# Plan

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

## Methods

### NewPlan

`func NewPlan() *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Plan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Plan) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Plan) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Plan) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Plan) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Plan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Plan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Plan) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Plan) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCasesCount

`func (o *Plan) GetCasesCount() int32`

GetCasesCount returns the CasesCount field if non-nil, zero value otherwise.

### GetCasesCountOk

`func (o *Plan) GetCasesCountOk() (*int32, bool)`

GetCasesCountOk returns a tuple with the CasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesCount

`func (o *Plan) SetCasesCount(v int32)`

SetCasesCount sets CasesCount field to given value.

### HasCasesCount

`func (o *Plan) HasCasesCount() bool`

HasCasesCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Plan) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Plan) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Plan) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Plan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Plan) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Plan) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Plan) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Plan) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreated

`func (o *Plan) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Plan) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Plan) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Plan) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Plan) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Plan) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Plan) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Plan) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


