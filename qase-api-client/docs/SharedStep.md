# SharedStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**ExpectedResult** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]SharedStepContent**](SharedStepContent.md) |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Cases** | Pointer to **[]int64** |  | [optional] 
**CasesCount** | Pointer to **int32** |  | [optional] 
**Created** | Pointer to **string** | Deprecated, use the &#x60;created_at&#x60; property instead. | [optional] 
**Updated** | Pointer to **string** | Deprecated, use the &#x60;updated_at&#x60; property instead. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSharedStep

`func NewSharedStep() *SharedStep`

NewSharedStep instantiates a new SharedStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepWithDefaults

`func NewSharedStepWithDefaults() *SharedStep`

NewSharedStepWithDefaults instantiates a new SharedStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *SharedStep) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SharedStep) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SharedStep) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SharedStep) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetTitle

`func (o *SharedStep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SharedStep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SharedStep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SharedStep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAction

`func (o *SharedStep) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SharedStep) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SharedStep) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SharedStep) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExpectedResult

`func (o *SharedStep) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *SharedStep) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *SharedStep) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *SharedStep) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### GetSteps

`func (o *SharedStep) GetSteps() []SharedStepContent`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SharedStep) GetStepsOk() (*[]SharedStepContent, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SharedStep) SetSteps(v []SharedStepContent)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *SharedStep) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetData

`func (o *SharedStep) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SharedStep) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SharedStep) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SharedStep) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCases

`func (o *SharedStep) GetCases() []int64`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *SharedStep) GetCasesOk() (*[]int64, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *SharedStep) SetCases(v []int64)`

SetCases sets Cases field to given value.

### HasCases

`func (o *SharedStep) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetCasesCount

`func (o *SharedStep) GetCasesCount() int32`

GetCasesCount returns the CasesCount field if non-nil, zero value otherwise.

### GetCasesCountOk

`func (o *SharedStep) GetCasesCountOk() (*int32, bool)`

GetCasesCountOk returns a tuple with the CasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesCount

`func (o *SharedStep) SetCasesCount(v int32)`

SetCasesCount sets CasesCount field to given value.

### HasCasesCount

`func (o *SharedStep) HasCasesCount() bool`

HasCasesCount returns a boolean if a field has been set.

### GetCreated

`func (o *SharedStep) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SharedStep) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SharedStep) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SharedStep) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *SharedStep) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SharedStep) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SharedStep) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SharedStep) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SharedStep) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SharedStep) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SharedStep) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SharedStep) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SharedStep) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SharedStep) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SharedStep) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SharedStep) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


