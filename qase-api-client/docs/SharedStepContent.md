# SharedStepContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**ExpectedResult** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentHash**](AttachmentHash.md) |  | [optional] 

## Methods

### NewSharedStepContent

`func NewSharedStepContent() *SharedStepContent`

NewSharedStepContent instantiates a new SharedStepContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepContentWithDefaults

`func NewSharedStepContentWithDefaults() *SharedStepContent`

NewSharedStepContentWithDefaults instantiates a new SharedStepContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SharedStepContent) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SharedStepContent) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SharedStepContent) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SharedStepContent) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHash

`func (o *SharedStepContent) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SharedStepContent) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SharedStepContent) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SharedStepContent) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetAction

`func (o *SharedStepContent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SharedStepContent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SharedStepContent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SharedStepContent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExpectedResult

`func (o *SharedStepContent) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *SharedStepContent) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *SharedStepContent) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *SharedStepContent) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### GetAttachments

`func (o *SharedStepContent) GetAttachments() []AttachmentHash`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SharedStepContent) GetAttachmentsOk() (*[]AttachmentHash, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SharedStepContent) SetAttachments(v []AttachmentHash)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SharedStepContent) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


