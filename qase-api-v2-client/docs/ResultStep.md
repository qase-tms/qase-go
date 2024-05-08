# ResultStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResultStepData**](ResultStepData.md) |  | [optional] 
**Execution** | Pointer to [**ResultStepExecution**](ResultStepExecution.md) |  | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** | Nested steps will be here. The same structure is used for them. | [optional] 

## Methods

### NewResultStep

`func NewResultStep() *ResultStep`

NewResultStep instantiates a new ResultStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultStepWithDefaults

`func NewResultStepWithDefaults() *ResultStep`

NewResultStepWithDefaults instantiates a new ResultStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResultStep) GetData() ResultStepData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultStep) GetDataOk() (*ResultStepData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultStep) SetData(v ResultStepData)`

SetData sets Data field to given value.

### HasData

`func (o *ResultStep) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExecution

`func (o *ResultStep) GetExecution() ResultStepExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *ResultStep) GetExecutionOk() (*ResultStepExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *ResultStep) SetExecution(v ResultStepExecution)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *ResultStep) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### GetSteps

`func (o *ResultStep) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ResultStep) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ResultStep) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ResultStep) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


