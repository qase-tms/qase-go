# ResultCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | If passed, used as an idempotency key | [optional] 
**Title** | **string** |  | 
**Signature** | Pointer to **string** |  | [optional] 
**TestopsId** | Pointer to **NullableInt64** |  | [optional] 
**Execution** | [**ResultExecution**](ResultExecution.md) |  | 
**Fields** | Pointer to **map[string]string** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Steps** | Pointer to [**[]ResultStep**](ResultStep.md) |  | [optional] 
**StepsType** | Pointer to [**NullableResultStepsType**](ResultStepsType.md) |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**Relations** | Pointer to [**NullableResultRelations**](ResultRelations.md) |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableFloat64** |  | [optional] 

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

`func (o *ResultCreate) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ResultCreate) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ResultCreate) SetFields(v map[string]string)`

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

### GetAuthor

`func (o *ResultCreate) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ResultCreate) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ResultCreate) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ResultCreate) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

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
### GetMuted

`func (o *ResultCreate) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ResultCreate) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ResultCreate) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *ResultCreate) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

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
### GetCreatedAt

`func (o *ResultCreate) GetCreatedAt() float64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResultCreate) GetCreatedAtOk() (*float64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResultCreate) SetCreatedAt(v float64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResultCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ResultCreate) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ResultCreate) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


