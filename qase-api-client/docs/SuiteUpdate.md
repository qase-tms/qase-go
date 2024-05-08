# SuiteUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Test suite title. | [optional] 
**Description** | Pointer to **string** | Test suite description. | [optional] 
**Preconditions** | Pointer to **string** | Test suite preconditions | [optional] 
**ParentId** | Pointer to **NullableInt64** | Parent suite ID | [optional] 

## Methods

### NewSuiteUpdate

`func NewSuiteUpdate() *SuiteUpdate`

NewSuiteUpdate instantiates a new SuiteUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuiteUpdateWithDefaults

`func NewSuiteUpdateWithDefaults() *SuiteUpdate`

NewSuiteUpdateWithDefaults instantiates a new SuiteUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SuiteUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SuiteUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SuiteUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SuiteUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *SuiteUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SuiteUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SuiteUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SuiteUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPreconditions

`func (o *SuiteUpdate) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *SuiteUpdate) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *SuiteUpdate) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *SuiteUpdate) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### GetParentId

`func (o *SuiteUpdate) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SuiteUpdate) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SuiteUpdate) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SuiteUpdate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *SuiteUpdate) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *SuiteUpdate) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


