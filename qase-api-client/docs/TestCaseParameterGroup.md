# TestCaseParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedId** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Items** | **map[string]interface{}** |  | 

## Methods

### NewTestCaseParameterGroup

`func NewTestCaseParameterGroup(type_ string, items map[string]interface{}, ) *TestCaseParameterGroup`

NewTestCaseParameterGroup instantiates a new TestCaseParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCaseParameterGroupWithDefaults

`func NewTestCaseParameterGroupWithDefaults() *TestCaseParameterGroup`

NewTestCaseParameterGroupWithDefaults instantiates a new TestCaseParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedId

`func (o *TestCaseParameterGroup) GetSharedId() string`

GetSharedId returns the SharedId field if non-nil, zero value otherwise.

### GetSharedIdOk

`func (o *TestCaseParameterGroup) GetSharedIdOk() (*string, bool)`

GetSharedIdOk returns a tuple with the SharedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedId

`func (o *TestCaseParameterGroup) SetSharedId(v string)`

SetSharedId sets SharedId field to given value.

### HasSharedId

`func (o *TestCaseParameterGroup) HasSharedId() bool`

HasSharedId returns a boolean if a field has been set.

### SetSharedIdNil

`func (o *TestCaseParameterGroup) SetSharedIdNil(b bool)`

 SetSharedIdNil sets the value for SharedId to be an explicit nil

### UnsetSharedId
`func (o *TestCaseParameterGroup) UnsetSharedId()`

UnsetSharedId ensures that no value is present for SharedId, not even an explicit nil
### GetType

`func (o *TestCaseParameterGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestCaseParameterGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestCaseParameterGroup) SetType(v string)`

SetType sets Type field to given value.


### GetItems

`func (o *TestCaseParameterGroup) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TestCaseParameterGroup) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TestCaseParameterGroup) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


