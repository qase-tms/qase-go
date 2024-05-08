# Suite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Preconditions** | Pointer to **NullableString** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**CasesCount** | Pointer to **int32** |  | [optional] 
**ParentId** | Pointer to **NullableInt64** |  | [optional] 
**Created** | Pointer to **string** | Deprecated, use the &#x60;created_at&#x60; property instead. | [optional] 
**Updated** | Pointer to **NullableString** | Deprecated, use the &#x60;updated_at&#x60; property instead. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSuite

`func NewSuite() *Suite`

NewSuite instantiates a new Suite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuiteWithDefaults

`func NewSuiteWithDefaults() *Suite`

NewSuiteWithDefaults instantiates a new Suite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Suite) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Suite) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Suite) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Suite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Suite) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Suite) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Suite) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Suite) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Suite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Suite) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Suite) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Suite) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Suite) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Suite) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPreconditions

`func (o *Suite) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *Suite) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *Suite) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *Suite) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### SetPreconditionsNil

`func (o *Suite) SetPreconditionsNil(b bool)`

 SetPreconditionsNil sets the value for Preconditions to be an explicit nil

### UnsetPreconditions
`func (o *Suite) UnsetPreconditions()`

UnsetPreconditions ensures that no value is present for Preconditions, not even an explicit nil
### GetPosition

`func (o *Suite) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Suite) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Suite) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Suite) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetCasesCount

`func (o *Suite) GetCasesCount() int32`

GetCasesCount returns the CasesCount field if non-nil, zero value otherwise.

### GetCasesCountOk

`func (o *Suite) GetCasesCountOk() (*int32, bool)`

GetCasesCountOk returns a tuple with the CasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesCount

`func (o *Suite) SetCasesCount(v int32)`

SetCasesCount sets CasesCount field to given value.

### HasCasesCount

`func (o *Suite) HasCasesCount() bool`

HasCasesCount returns a boolean if a field has been set.

### GetParentId

`func (o *Suite) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Suite) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Suite) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Suite) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *Suite) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *Suite) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetCreated

`func (o *Suite) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Suite) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Suite) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Suite) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Suite) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Suite) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Suite) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Suite) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *Suite) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *Suite) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetCreatedAt

`func (o *Suite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Suite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Suite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Suite) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Suite) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Suite) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Suite) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Suite) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


