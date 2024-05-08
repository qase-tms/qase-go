# SharedStepContentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**Action** | **string** |  | 
**ExpectedResult** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to **[]string** | A list of Attachment hashes. | [optional] 

## Methods

### NewSharedStepContentCreate

`func NewSharedStepContentCreate(action string, ) *SharedStepContentCreate`

NewSharedStepContentCreate instantiates a new SharedStepContentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepContentCreateWithDefaults

`func NewSharedStepContentCreateWithDefaults() *SharedStepContentCreate`

NewSharedStepContentCreateWithDefaults instantiates a new SharedStepContentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *SharedStepContentCreate) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SharedStepContentCreate) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SharedStepContentCreate) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SharedStepContentCreate) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetAction

`func (o *SharedStepContentCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SharedStepContentCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SharedStepContentCreate) SetAction(v string)`

SetAction sets Action field to given value.


### GetExpectedResult

`func (o *SharedStepContentCreate) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *SharedStepContentCreate) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *SharedStepContentCreate) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *SharedStepContentCreate) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### GetData

`func (o *SharedStepContentCreate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SharedStepContentCreate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SharedStepContentCreate) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SharedStepContentCreate) HasData() bool`

HasData returns a boolean if a field has been set.

### GetAttachments

`func (o *SharedStepContentCreate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SharedStepContentCreate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SharedStepContentCreate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SharedStepContentCreate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


