# TestCaseParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedId** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Items** | **map[string]interface{}** |  | 

## Methods

### NewTestCaseParameter

`func NewTestCaseParameter(type_ string, items map[string]interface{}, ) *TestCaseParameter`

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


### GetItems

`func (o *TestCaseParameter) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TestCaseParameter) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TestCaseParameter) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


