# SharedStepCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Action** | Pointer to **string** | Deprecated, use the &#x60;steps&#x60; property instead. | [optional] 
**ExpectedResult** | Pointer to **string** | Deprecated, use the &#x60;steps&#x60; property instead. | [optional] 
**Data** | Pointer to **string** | Deprecated, use the &#x60;steps&#x60; property instead. | [optional] 
**Steps** | Pointer to [**[]SharedStepContentCreate**](SharedStepContentCreate.md) |  | [optional] 

## Methods

### NewSharedStepCreate

`func NewSharedStepCreate(title string, ) *SharedStepCreate`

NewSharedStepCreate instantiates a new SharedStepCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepCreateWithDefaults

`func NewSharedStepCreateWithDefaults() *SharedStepCreate`

NewSharedStepCreateWithDefaults instantiates a new SharedStepCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SharedStepCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SharedStepCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SharedStepCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAction

`func (o *SharedStepCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SharedStepCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SharedStepCreate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SharedStepCreate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExpectedResult

`func (o *SharedStepCreate) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *SharedStepCreate) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *SharedStepCreate) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *SharedStepCreate) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### GetData

`func (o *SharedStepCreate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SharedStepCreate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SharedStepCreate) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SharedStepCreate) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSteps

`func (o *SharedStepCreate) GetSteps() []SharedStepContentCreate`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SharedStepCreate) GetStepsOk() (*[]SharedStepContentCreate, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SharedStepCreate) SetSteps(v []SharedStepContentCreate)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *SharedStepCreate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


