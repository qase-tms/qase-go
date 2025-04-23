# ResultCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | If passed, used as an idempotency key | [optional] 
**Title** | **string** |  | 
**Signature** | Pointer to **string** |  | [optional] 
**TestopsId** | Pointer to **NullableInt64** | ID of the test case. Cannot be specified together with testopd_ids. | [optional] 
**TestopsIds** | Pointer to **[]int64** | IDs of the test cases. Cannot be specified together with testopd_id. | [optional] 
**Execution** | [**ResultExecution**](ResultExecution.md) |  | 
**Fields** | Pointer to [**ResultCreateFields**](ResultCreateFields.md) |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Steps** | Pointer to [**[]ResultStep**](ResultStep.md) |  | [optional] 
**StepsType** | Pointer to [**NullableResultStepsType**](ResultStepsType.md) |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**ParamGroups** | Pointer to **[][]string** | List parameter groups by name only. Add their values in the &#39;params&#39; field | [optional] 
**Relations** | Pointer to [**NullableResultRelations**](ResultRelations.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Defect** | Pointer to **bool** | If true and the result is failed, the defect associated with the result will be created | [optional] 

## Methods

### NewResultCreate

`func NewResultCreate(title string, execution ResultExecution, ) *ResultCreate`

NewResultCreate instantiates a new ResultCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCreateWithDefaults

`func NewResultCreateWithDefaults() *ResultCreate`

NewResultCreateWithDefaults instantiates a new ResultCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResultCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ResultCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResultCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResultCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSignature

`func (o *ResultCreate) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ResultCreate) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ResultCreate) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ResultCreate) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTestopsId

`func (o *ResultCreate) GetTestopsId() int64`

GetTestopsId returns the TestopsId field if non-nil, zero value otherwise.

### GetTestopsIdOk

`func (o *ResultCreate) GetTestopsIdOk() (*int64, bool)`

GetTestopsIdOk returns a tuple with the TestopsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestopsId

`func (o *ResultCreate) SetTestopsId(v int64)`

SetTestopsId sets TestopsId field to given value.

### HasTestopsId

`func (o *ResultCreate) HasTestopsId() bool`

HasTestopsId returns a boolean if a field has been set.

### SetTestopsIdNil

`func (o *ResultCreate) SetTestopsIdNil(b bool)`

 SetTestopsIdNil sets the value for TestopsId to be an explicit nil

### UnsetTestopsId
`func (o *ResultCreate) UnsetTestopsId()`

UnsetTestopsId ensures that no value is present for TestopsId, not even an explicit nil
### GetTestopsIds

`func (o *ResultCreate) GetTestopsIds() []int64`

GetTestopsIds returns the TestopsIds field if non-nil, zero value otherwise.

### GetTestopsIdsOk

`func (o *ResultCreate) GetTestopsIdsOk() (*[]int64, bool)`

GetTestopsIdsOk returns a tuple with the TestopsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestopsIds

`func (o *ResultCreate) SetTestopsIds(v []int64)`

SetTestopsIds sets TestopsIds field to given value.

### HasTestopsIds

`func (o *ResultCreate) HasTestopsIds() bool`

HasTestopsIds returns a boolean if a field has been set.

### SetTestopsIdsNil

`func (o *ResultCreate) SetTestopsIdsNil(b bool)`

 SetTestopsIdsNil sets the value for TestopsIds to be an explicit nil

### UnsetTestopsIds
`func (o *ResultCreate) UnsetTestopsIds()`

UnsetTestopsIds ensures that no value is present for TestopsIds, not even an explicit nil
### GetExecution

`func (o *ResultCreate) GetExecution() ResultExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *ResultCreate) GetExecutionOk() (*ResultExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *ResultCreate) SetExecution(v ResultExecution)`

SetExecution sets Execution field to given value.


### GetFields

`func (o *ResultCreate) GetFields() ResultCreateFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ResultCreate) GetFieldsOk() (*ResultCreateFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ResultCreate) SetFields(v ResultCreateFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ResultCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetAttachments

`func (o *ResultCreate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ResultCreate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ResultCreate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ResultCreate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *ResultCreate) GetSteps() []ResultStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ResultCreate) GetStepsOk() (*[]ResultStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ResultCreate) SetSteps(v []ResultStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ResultCreate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStepsType

`func (o *ResultCreate) GetStepsType() ResultStepsType`

GetStepsType returns the StepsType field if non-nil, zero value otherwise.

### GetStepsTypeOk

`func (o *ResultCreate) GetStepsTypeOk() (*ResultStepsType, bool)`

GetStepsTypeOk returns a tuple with the StepsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsType

`func (o *ResultCreate) SetStepsType(v ResultStepsType)`

SetStepsType sets StepsType field to given value.

### HasStepsType

`func (o *ResultCreate) HasStepsType() bool`

HasStepsType returns a boolean if a field has been set.

### SetStepsTypeNil

`func (o *ResultCreate) SetStepsTypeNil(b bool)`

 SetStepsTypeNil sets the value for StepsType to be an explicit nil

### UnsetStepsType
`func (o *ResultCreate) UnsetStepsType()`

UnsetStepsType ensures that no value is present for StepsType, not even an explicit nil
### GetParams

`func (o *ResultCreate) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ResultCreate) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ResultCreate) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *ResultCreate) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetParamGroups

`func (o *ResultCreate) GetParamGroups() [][]string`

GetParamGroups returns the ParamGroups field if non-nil, zero value otherwise.

### GetParamGroupsOk

`func (o *ResultCreate) GetParamGroupsOk() (*[][]string, bool)`

GetParamGroupsOk returns a tuple with the ParamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamGroups

`func (o *ResultCreate) SetParamGroups(v [][]string)`

SetParamGroups sets ParamGroups field to given value.

### HasParamGroups

`func (o *ResultCreate) HasParamGroups() bool`

HasParamGroups returns a boolean if a field has been set.

### SetParamGroupsNil

`func (o *ResultCreate) SetParamGroupsNil(b bool)`

 SetParamGroupsNil sets the value for ParamGroups to be an explicit nil

### UnsetParamGroups
`func (o *ResultCreate) UnsetParamGroups()`

UnsetParamGroups ensures that no value is present for ParamGroups, not even an explicit nil
### GetRelations

`func (o *ResultCreate) GetRelations() ResultRelations`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *ResultCreate) GetRelationsOk() (*ResultRelations, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *ResultCreate) SetRelations(v ResultRelations)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *ResultCreate) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### SetRelationsNil

`func (o *ResultCreate) SetRelationsNil(b bool)`

 SetRelationsNil sets the value for Relations to be an explicit nil

### UnsetRelations
`func (o *ResultCreate) UnsetRelations()`

UnsetRelations ensures that no value is present for Relations, not even an explicit nil
### GetMessage

`func (o *ResultCreate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResultCreate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResultCreate) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResultCreate) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ResultCreate) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ResultCreate) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetDefect

`func (o *ResultCreate) GetDefect() bool`

GetDefect returns the Defect field if non-nil, zero value otherwise.

### GetDefectOk

`func (o *ResultCreate) GetDefectOk() (*bool, bool)`

GetDefectOk returns a tuple with the Defect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefect

`func (o *ResultCreate) SetDefect(v bool)`

SetDefect sets Defect field to given value.

### HasDefect

`func (o *ResultCreate) HasDefect() bool`

HasDefect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


