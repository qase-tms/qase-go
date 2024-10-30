# ResultCreateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Preconditions** | Pointer to **string** |  | [optional] 
**Postconditions** | Pointer to **string** |  | [optional] 
**Layer** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Behavior** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Muted** | Pointer to **string** |  | [optional] 
**IsFlaky** | Pointer to **string** |  | [optional] 

## Methods

### NewResultCreateFields

`func NewResultCreateFields() *ResultCreateFields`

NewResultCreateFields instantiates a new ResultCreateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCreateFieldsWithDefaults

`func NewResultCreateFieldsWithDefaults() *ResultCreateFields`

NewResultCreateFieldsWithDefaults instantiates a new ResultCreateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ResultCreateFields) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ResultCreateFields) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ResultCreateFields) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ResultCreateFields) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDescription

`func (o *ResultCreateFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResultCreateFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResultCreateFields) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResultCreateFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPreconditions

`func (o *ResultCreateFields) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *ResultCreateFields) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *ResultCreateFields) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *ResultCreateFields) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### GetPostconditions

`func (o *ResultCreateFields) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *ResultCreateFields) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *ResultCreateFields) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *ResultCreateFields) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### GetLayer

`func (o *ResultCreateFields) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *ResultCreateFields) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *ResultCreateFields) SetLayer(v string)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *ResultCreateFields) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetSeverity

`func (o *ResultCreateFields) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ResultCreateFields) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ResultCreateFields) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ResultCreateFields) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetPriority

`func (o *ResultCreateFields) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ResultCreateFields) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ResultCreateFields) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ResultCreateFields) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetBehavior

`func (o *ResultCreateFields) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *ResultCreateFields) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *ResultCreateFields) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *ResultCreateFields) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetType

`func (o *ResultCreateFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResultCreateFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResultCreateFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResultCreateFields) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMuted

`func (o *ResultCreateFields) GetMuted() string`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ResultCreateFields) GetMutedOk() (*string, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ResultCreateFields) SetMuted(v string)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *ResultCreateFields) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetIsFlaky

`func (o *ResultCreateFields) GetIsFlaky() string`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *ResultCreateFields) GetIsFlakyOk() (*string, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *ResultCreateFields) SetIsFlaky(v string)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *ResultCreateFields) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


