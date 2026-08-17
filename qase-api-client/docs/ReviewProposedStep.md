# ReviewProposedStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Step action text. Used for classic steps. For gherkin steps, use the \&quot;value\&quot; property instead. | [optional] 
**ExpectedResult** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | Gherkin scenario text. Used when steps_type is \&quot;gherkin\&quot;. | [optional] 
**Shared** | Pointer to **string** | Hash of the referenced shared step. | [optional] 
**Attachments** | Pointer to **[]string** | A list of Attachment hashes. | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** | Nested steps use the same structure. | [optional] 

## Methods

### NewReviewProposedStep

`func NewReviewProposedStep() *ReviewProposedStep`

NewReviewProposedStep instantiates a new ReviewProposedStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewProposedStepWithDefaults

`func NewReviewProposedStepWithDefaults() *ReviewProposedStep`

NewReviewProposedStepWithDefaults instantiates a new ReviewProposedStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ReviewProposedStep) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ReviewProposedStep) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ReviewProposedStep) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ReviewProposedStep) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExpectedResult

`func (o *ReviewProposedStep) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *ReviewProposedStep) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *ReviewProposedStep) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *ReviewProposedStep) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### GetData

`func (o *ReviewProposedStep) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReviewProposedStep) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReviewProposedStep) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ReviewProposedStep) HasData() bool`

HasData returns a boolean if a field has been set.

### GetValue

`func (o *ReviewProposedStep) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReviewProposedStep) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReviewProposedStep) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ReviewProposedStep) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetShared

`func (o *ReviewProposedStep) GetShared() string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *ReviewProposedStep) GetSharedOk() (*string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *ReviewProposedStep) SetShared(v string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *ReviewProposedStep) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetAttachments

`func (o *ReviewProposedStep) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ReviewProposedStep) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ReviewProposedStep) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ReviewProposedStep) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *ReviewProposedStep) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ReviewProposedStep) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ReviewProposedStep) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ReviewProposedStep) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


