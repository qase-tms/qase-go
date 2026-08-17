# ReviewStepData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Step action text. Classic steps only. | [optional] 
**Shared** | Pointer to **string** | Hash of an existing shared step to insert at this position. | [optional] 
**ExpectedResult** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | Gherkin scenario text. Used when steps_type is \&quot;gherkin\&quot;. Example: \&quot;Given a user exists\\nWhen they log in\\nThen they see the dashboard\&quot; | [optional] 
**Attachments** | Pointer to **[]string** | A list of Attachment hashes. | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** | Nested steps may be passed here. Use same structure for them. | [optional] 

## Methods

### NewReviewStepData

`func NewReviewStepData() *ReviewStepData`

NewReviewStepData instantiates a new ReviewStepData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewStepDataWithDefaults

`func NewReviewStepDataWithDefaults() *ReviewStepData`

NewReviewStepDataWithDefaults instantiates a new ReviewStepData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ReviewStepData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ReviewStepData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ReviewStepData) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ReviewStepData) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetShared

`func (o *ReviewStepData) GetShared() string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *ReviewStepData) GetSharedOk() (*string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *ReviewStepData) SetShared(v string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *ReviewStepData) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetExpectedResult

`func (o *ReviewStepData) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *ReviewStepData) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *ReviewStepData) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *ReviewStepData) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### GetData

`func (o *ReviewStepData) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReviewStepData) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReviewStepData) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ReviewStepData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetValue

`func (o *ReviewStepData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReviewStepData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReviewStepData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ReviewStepData) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAttachments

`func (o *ReviewStepData) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ReviewStepData) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ReviewStepData) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ReviewStepData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *ReviewStepData) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ReviewStepData) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ReviewStepData) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ReviewStepData) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


