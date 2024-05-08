# TestStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**SharedStepHash** | Pointer to **NullableString** |  | [optional] 
**SharedStepNestedHash** | Pointer to **NullableString** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**ExpectedResult** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** | Nested steps will be here. The same structure is used for them. | [optional] 

## Methods

### NewTestStep

`func NewTestStep() *TestStep`

NewTestStep instantiates a new TestStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStepWithDefaults

`func NewTestStepWithDefaults() *TestStep`

NewTestStepWithDefaults instantiates a new TestStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *TestStep) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *TestStep) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *TestStep) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *TestStep) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetSharedStepHash

`func (o *TestStep) GetSharedStepHash() string`

GetSharedStepHash returns the SharedStepHash field if non-nil, zero value otherwise.

### GetSharedStepHashOk

`func (o *TestStep) GetSharedStepHashOk() (*string, bool)`

GetSharedStepHashOk returns a tuple with the SharedStepHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepHash

`func (o *TestStep) SetSharedStepHash(v string)`

SetSharedStepHash sets SharedStepHash field to given value.

### HasSharedStepHash

`func (o *TestStep) HasSharedStepHash() bool`

HasSharedStepHash returns a boolean if a field has been set.

### SetSharedStepHashNil

`func (o *TestStep) SetSharedStepHashNil(b bool)`

 SetSharedStepHashNil sets the value for SharedStepHash to be an explicit nil

### UnsetSharedStepHash
`func (o *TestStep) UnsetSharedStepHash()`

UnsetSharedStepHash ensures that no value is present for SharedStepHash, not even an explicit nil
### GetSharedStepNestedHash

`func (o *TestStep) GetSharedStepNestedHash() string`

GetSharedStepNestedHash returns the SharedStepNestedHash field if non-nil, zero value otherwise.

### GetSharedStepNestedHashOk

`func (o *TestStep) GetSharedStepNestedHashOk() (*string, bool)`

GetSharedStepNestedHashOk returns a tuple with the SharedStepNestedHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepNestedHash

`func (o *TestStep) SetSharedStepNestedHash(v string)`

SetSharedStepNestedHash sets SharedStepNestedHash field to given value.

### HasSharedStepNestedHash

`func (o *TestStep) HasSharedStepNestedHash() bool`

HasSharedStepNestedHash returns a boolean if a field has been set.

### SetSharedStepNestedHashNil

`func (o *TestStep) SetSharedStepNestedHashNil(b bool)`

 SetSharedStepNestedHashNil sets the value for SharedStepNestedHash to be an explicit nil

### UnsetSharedStepNestedHash
`func (o *TestStep) UnsetSharedStepNestedHash()`

UnsetSharedStepNestedHash ensures that no value is present for SharedStepNestedHash, not even an explicit nil
### GetPosition

`func (o *TestStep) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TestStep) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TestStep) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *TestStep) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetAction

`func (o *TestStep) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TestStep) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TestStep) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TestStep) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExpectedResult

`func (o *TestStep) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *TestStep) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *TestStep) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *TestStep) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### SetExpectedResultNil

`func (o *TestStep) SetExpectedResultNil(b bool)`

 SetExpectedResultNil sets the value for ExpectedResult to be an explicit nil

### UnsetExpectedResult
`func (o *TestStep) UnsetExpectedResult()`

UnsetExpectedResult ensures that no value is present for ExpectedResult, not even an explicit nil
### GetData

`func (o *TestStep) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TestStep) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TestStep) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TestStep) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *TestStep) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TestStep) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetAttachments

`func (o *TestStep) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestStep) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestStep) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestStep) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *TestStep) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TestStep) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TestStep) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TestStep) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


