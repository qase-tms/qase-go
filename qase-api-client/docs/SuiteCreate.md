# SuiteCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Test suite title. | 
**Description** | Pointer to **string** | Test suite description. | [optional] 
**Preconditions** | Pointer to **string** | Test suite preconditions | [optional] 
**ParentId** | Pointer to **NullableInt64** | Parent suite ID | [optional] 

## Methods

### NewSuiteCreate

`func NewSuiteCreate(title string, ) *SuiteCreate`

NewSuiteCreate instantiates a new SuiteCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuiteCreateWithDefaults

`func NewSuiteCreateWithDefaults() *SuiteCreate`

NewSuiteCreateWithDefaults instantiates a new SuiteCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SuiteCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SuiteCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SuiteCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SuiteCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SuiteCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SuiteCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SuiteCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPreconditions

`func (o *SuiteCreate) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *SuiteCreate) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *SuiteCreate) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *SuiteCreate) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### GetParentId

`func (o *SuiteCreate) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SuiteCreate) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SuiteCreate) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SuiteCreate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *SuiteCreate) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *SuiteCreate) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


