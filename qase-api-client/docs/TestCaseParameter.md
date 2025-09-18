# TestCaseParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedId** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Item** | [**ParameterSingle**](ParameterSingle.md) |  | 
**Items** | [**[]ParameterSingle**](ParameterSingle.md) |  | 

## Methods

### NewTestCaseParameter

`func NewTestCaseParameter(type_ string, item ParameterSingle, items []ParameterSingle, ) *TestCaseParameter`

NewTestCaseParameter instantiates a new TestCaseParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCaseParameterWithDefaults

`func NewTestCaseParameterWithDefaults() *TestCaseParameter`

NewTestCaseParameterWithDefaults instantiates a new TestCaseParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedId

`func (o *TestCaseParameter) GetSharedId() string`

GetSharedId returns the SharedId field if non-nil, zero value otherwise.

### GetSharedIdOk

`func (o *TestCaseParameter) GetSharedIdOk() (*string, bool)`

GetSharedIdOk returns a tuple with the SharedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedId

`func (o *TestCaseParameter) SetSharedId(v string)`

SetSharedId sets SharedId field to given value.

### HasSharedId

`func (o *TestCaseParameter) HasSharedId() bool`

HasSharedId returns a boolean if a field has been set.

### SetSharedIdNil

`func (o *TestCaseParameter) SetSharedIdNil(b bool)`

 SetSharedIdNil sets the value for SharedId to be an explicit nil

### UnsetSharedId
`func (o *TestCaseParameter) UnsetSharedId()`

UnsetSharedId ensures that no value is present for SharedId, not even an explicit nil
### GetType

`func (o *TestCaseParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestCaseParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestCaseParameter) SetType(v string)`

SetType sets Type field to given value.


### GetItem

`func (o *TestCaseParameter) GetItem() ParameterSingle`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TestCaseParameter) GetItemOk() (*ParameterSingle, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TestCaseParameter) SetItem(v ParameterSingle)`

SetItem sets Item field to given value.


### GetItems

`func (o *TestCaseParameter) GetItems() []ParameterSingle`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TestCaseParameter) GetItemsOk() (*[]ParameterSingle, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TestCaseParameter) SetItems(v []ParameterSingle)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


