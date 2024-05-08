# TestStepResultCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to **int32** |  | [optional] 
**Status** | **string** |  | 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**ExpectedResult** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** | Nested steps results may be passed here. Use same structure for them. | [optional] 

## Methods

### NewTestStepResultCreate

`func NewTestStepResultCreate(status string, ) *TestStepResultCreate`

NewTestStepResultCreate instantiates a new TestStepResultCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStepResultCreateWithDefaults

`func NewTestStepResultCreateWithDefaults() *TestStepResultCreate`

NewTestStepResultCreateWithDefaults instantiates a new TestStepResultCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *TestStepResultCreate) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TestStepResultCreate) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TestStepResultCreate) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *TestStepResultCreate) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStatus

`func (o *TestStepResultCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestStepResultCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestStepResultCreate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetComment

`func (o *TestStepResultCreate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestStepResultCreate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestStepResultCreate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestStepResultCreate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestStepResultCreate) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestStepResultCreate) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetAttachments

`func (o *TestStepResultCreate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestStepResultCreate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestStepResultCreate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestStepResultCreate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestStepResultCreate) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestStepResultCreate) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetAction

`func (o *TestStepResultCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TestStepResultCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TestStepResultCreate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TestStepResultCreate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExpectedResult

`func (o *TestStepResultCreate) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *TestStepResultCreate) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *TestStepResultCreate) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *TestStepResultCreate) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### SetExpectedResultNil

`func (o *TestStepResultCreate) SetExpectedResultNil(b bool)`

 SetExpectedResultNil sets the value for ExpectedResult to be an explicit nil

### UnsetExpectedResult
`func (o *TestStepResultCreate) UnsetExpectedResult()`

UnsetExpectedResult ensures that no value is present for ExpectedResult, not even an explicit nil
### GetData

`func (o *TestStepResultCreate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TestStepResultCreate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TestStepResultCreate) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TestStepResultCreate) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *TestStepResultCreate) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TestStepResultCreate) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetSteps

`func (o *TestStepResultCreate) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TestStepResultCreate) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TestStepResultCreate) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TestStepResultCreate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


