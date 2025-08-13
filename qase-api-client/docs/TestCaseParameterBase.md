# TestCaseParameterBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedId** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Items** | [**[]ParameterSingle**](ParameterSingle.md) |  | 

## Methods

### NewTestCaseParameterBase

`func NewTestCaseParameterBase(type_ string, items []ParameterSingle, ) *TestCaseParameterBase`

NewTestCaseParameterBase instantiates a new TestCaseParameterBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCaseParameterBaseWithDefaults

`func NewTestCaseParameterBaseWithDefaults() *TestCaseParameterBase`

NewTestCaseParameterBaseWithDefaults instantiates a new TestCaseParameterBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedId

`func (o *TestCaseParameterBase) GetSharedId() string`

GetSharedId returns the SharedId field if non-nil, zero value otherwise.

### GetSharedIdOk

`func (o *TestCaseParameterBase) GetSharedIdOk() (*string, bool)`

GetSharedIdOk returns a tuple with the SharedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedId

`func (o *TestCaseParameterBase) SetSharedId(v string)`

SetSharedId sets SharedId field to given value.

### HasSharedId

`func (o *TestCaseParameterBase) HasSharedId() bool`

HasSharedId returns a boolean if a field has been set.

### SetSharedIdNil

`func (o *TestCaseParameterBase) SetSharedIdNil(b bool)`

 SetSharedIdNil sets the value for SharedId to be an explicit nil

### UnsetSharedId
`func (o *TestCaseParameterBase) UnsetSharedId()`

UnsetSharedId ensures that no value is present for SharedId, not even an explicit nil
### GetType

`func (o *TestCaseParameterBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestCaseParameterBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestCaseParameterBase) SetType(v string)`

SetType sets Type field to given value.


### GetItems

`func (o *TestCaseParameterBase) GetItems() []ParameterSingle`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TestCaseParameterBase) GetItemsOk() (*[]ParameterSingle, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TestCaseParameterBase) SetItems(v []ParameterSingle)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


