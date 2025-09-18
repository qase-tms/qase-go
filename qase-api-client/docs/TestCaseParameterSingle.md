# TestCaseParameterSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedId** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Item** | [**ParameterSingle**](ParameterSingle.md) |  | 

## Methods

### NewTestCaseParameterSingle

`func NewTestCaseParameterSingle(type_ string, item ParameterSingle, ) *TestCaseParameterSingle`

NewTestCaseParameterSingle instantiates a new TestCaseParameterSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCaseParameterSingleWithDefaults

`func NewTestCaseParameterSingleWithDefaults() *TestCaseParameterSingle`

NewTestCaseParameterSingleWithDefaults instantiates a new TestCaseParameterSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedId

`func (o *TestCaseParameterSingle) GetSharedId() string`

GetSharedId returns the SharedId field if non-nil, zero value otherwise.

### GetSharedIdOk

`func (o *TestCaseParameterSingle) GetSharedIdOk() (*string, bool)`

GetSharedIdOk returns a tuple with the SharedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedId

`func (o *TestCaseParameterSingle) SetSharedId(v string)`

SetSharedId sets SharedId field to given value.

### HasSharedId

`func (o *TestCaseParameterSingle) HasSharedId() bool`

HasSharedId returns a boolean if a field has been set.

### SetSharedIdNil

`func (o *TestCaseParameterSingle) SetSharedIdNil(b bool)`

 SetSharedIdNil sets the value for SharedId to be an explicit nil

### UnsetSharedId
`func (o *TestCaseParameterSingle) UnsetSharedId()`

UnsetSharedId ensures that no value is present for SharedId, not even an explicit nil
### GetType

`func (o *TestCaseParameterSingle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestCaseParameterSingle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestCaseParameterSingle) SetType(v string)`

SetType sets Type field to given value.


### GetItem

`func (o *TestCaseParameterSingle) GetItem() ParameterSingle`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TestCaseParameterSingle) GetItemOk() (*ParameterSingle, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TestCaseParameterSingle) SetItem(v ParameterSingle)`

SetItem sets Item field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


