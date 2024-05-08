# ResultStepData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ExpectedResult** | Pointer to **string** |  | [optional] 
**InputData** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResultStepData

`func NewResultStepData(action string, ) *ResultStepData`

NewResultStepData instantiates a new ResultStepData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultStepDataWithDefaults

`func NewResultStepDataWithDefaults() *ResultStepData`

NewResultStepDataWithDefaults instantiates a new ResultStepData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ResultStepData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ResultStepData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ResultStepData) SetAction(v string)`

SetAction sets Action field to given value.


### GetExpectedResult

`func (o *ResultStepData) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *ResultStepData) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *ResultStepData) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *ResultStepData) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### GetInputData

`func (o *ResultStepData) GetInputData() string`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *ResultStepData) GetInputDataOk() (*string, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *ResultStepData) SetInputData(v string)`

SetInputData sets InputData field to given value.

### HasInputData

`func (o *ResultStepData) HasInputData() bool`

HasInputData returns a boolean if a field has been set.

### GetAttachments

`func (o *ResultStepData) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ResultStepData) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ResultStepData) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ResultStepData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


