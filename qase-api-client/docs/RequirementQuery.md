# RequirementQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**RequirementId** | **int64** |  | 
**ParentId** | Pointer to **NullableInt64** |  | [optional] 
**MemberId** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewRequirementQuery

`func NewRequirementQuery(requirementId int64, ) *RequirementQuery`

NewRequirementQuery instantiates a new RequirementQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementQueryWithDefaults

`func NewRequirementQueryWithDefaults() *RequirementQuery`

NewRequirementQueryWithDefaults instantiates a new RequirementQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequirementQuery) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequirementQuery) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequirementQuery) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RequirementQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequirementId

`func (o *RequirementQuery) GetRequirementId() int64`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *RequirementQuery) GetRequirementIdOk() (*int64, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *RequirementQuery) SetRequirementId(v int64)`

SetRequirementId sets RequirementId field to given value.


### GetParentId

`func (o *RequirementQuery) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *RequirementQuery) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *RequirementQuery) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *RequirementQuery) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *RequirementQuery) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *RequirementQuery) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetMemberId

`func (o *RequirementQuery) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *RequirementQuery) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *RequirementQuery) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *RequirementQuery) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetTitle

`func (o *RequirementQuery) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequirementQuery) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequirementQuery) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RequirementQuery) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *RequirementQuery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequirementQuery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequirementQuery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequirementQuery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RequirementQuery) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RequirementQuery) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *RequirementQuery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequirementQuery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequirementQuery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequirementQuery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *RequirementQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequirementQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequirementQuery) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RequirementQuery) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RequirementQuery) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequirementQuery) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequirementQuery) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequirementQuery) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RequirementQuery) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RequirementQuery) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RequirementQuery) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RequirementQuery) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *RequirementQuery) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *RequirementQuery) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


