# TestStepCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ExpectedResult** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Attachments** | Pointer to **[]string** | A list of Attachment hashes. | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** | Nested steps may be passed here. Use same structure for them. | [optional] 

## Methods

### NewTestStepCreate

`func NewTestStepCreate() *TestStepCreate`

NewTestStepCreate instantiates a new TestStepCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStepCreateWithDefaults

`func NewTestStepCreateWithDefaults() *TestStepCreate`

NewTestStepCreateWithDefaults instantiates a new TestStepCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TestStepCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TestStepCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TestStepCreate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TestStepCreate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExpectedResult

`func (o *TestStepCreate) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *TestStepCreate) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *TestStepCreate) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *TestStepCreate) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### GetData

`func (o *TestStepCreate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TestStepCreate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TestStepCreate) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TestStepCreate) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPosition

`func (o *TestStepCreate) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TestStepCreate) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TestStepCreate) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *TestStepCreate) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetAttachments

`func (o *TestStepCreate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestStepCreate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestStepCreate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestStepCreate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *TestStepCreate) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TestStepCreate) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TestStepCreate) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TestStepCreate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


