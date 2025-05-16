# TestStepResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | 1 - passed, 2 - failed, 3 - blocked, 5 - skipped, 7 - in_progress | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** | Nested steps results will be here. The same structure is used for them for them. | [optional] 

## Methods

### NewTestStepResult

`func NewTestStepResult() *TestStepResult`

NewTestStepResult instantiates a new TestStepResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStepResultWithDefaults

`func NewTestStepResultWithDefaults() *TestStepResult`

NewTestStepResultWithDefaults instantiates a new TestStepResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TestStepResult) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestStepResult) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestStepResult) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestStepResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPosition

`func (o *TestStepResult) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TestStepResult) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TestStepResult) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *TestStepResult) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetAttachments

`func (o *TestStepResult) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestStepResult) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestStepResult) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestStepResult) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *TestStepResult) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TestStepResult) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TestStepResult) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TestStepResult) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


