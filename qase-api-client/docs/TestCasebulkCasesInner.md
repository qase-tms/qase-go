# TestCasebulkCasesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Preconditions** | Pointer to **string** |  | [optional] 
**Postconditions** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Severity** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Behavior** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Layer** | Pointer to **int32** |  | [optional] 
**IsFlaky** | Pointer to **int32** |  | [optional] 
**AuthorId** | Pointer to **int32** |  | [optional] 
**SuiteId** | Pointer to **int64** |  | [optional] 
**MilestoneId** | Pointer to **int64** |  | [optional] 
**Automation** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Attachments** | Pointer to **[]string** | A list of Attachment hashes. | [optional] 
**Steps** | Pointer to [**[]TestStepCreate**](TestStepCreate.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Params** | Pointer to **map[string][]string** |  | [optional] 
**CustomField** | Pointer to **map[string]string** | A map of custom fields values (id &#x3D;&gt; value) | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTestCasebulkCasesInner

`func NewTestCasebulkCasesInner(title string, ) *TestCasebulkCasesInner`

NewTestCasebulkCasesInner instantiates a new TestCasebulkCasesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCasebulkCasesInnerWithDefaults

`func NewTestCasebulkCasesInnerWithDefaults() *TestCasebulkCasesInner`

NewTestCasebulkCasesInnerWithDefaults instantiates a new TestCasebulkCasesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TestCasebulkCasesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestCasebulkCasesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestCasebulkCasesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestCasebulkCasesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPreconditions

`func (o *TestCasebulkCasesInner) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *TestCasebulkCasesInner) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *TestCasebulkCasesInner) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *TestCasebulkCasesInner) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### GetPostconditions

`func (o *TestCasebulkCasesInner) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *TestCasebulkCasesInner) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *TestCasebulkCasesInner) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *TestCasebulkCasesInner) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### GetTitle

`func (o *TestCasebulkCasesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestCasebulkCasesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestCasebulkCasesInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSeverity

`func (o *TestCasebulkCasesInner) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TestCasebulkCasesInner) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TestCasebulkCasesInner) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *TestCasebulkCasesInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetPriority

`func (o *TestCasebulkCasesInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestCasebulkCasesInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestCasebulkCasesInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TestCasebulkCasesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetBehavior

`func (o *TestCasebulkCasesInner) GetBehavior() int32`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *TestCasebulkCasesInner) GetBehaviorOk() (*int32, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *TestCasebulkCasesInner) SetBehavior(v int32)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *TestCasebulkCasesInner) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetType

`func (o *TestCasebulkCasesInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestCasebulkCasesInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestCasebulkCasesInner) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *TestCasebulkCasesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayer

`func (o *TestCasebulkCasesInner) GetLayer() int32`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *TestCasebulkCasesInner) GetLayerOk() (*int32, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *TestCasebulkCasesInner) SetLayer(v int32)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *TestCasebulkCasesInner) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetIsFlaky

`func (o *TestCasebulkCasesInner) GetIsFlaky() int32`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *TestCasebulkCasesInner) GetIsFlakyOk() (*int32, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *TestCasebulkCasesInner) SetIsFlaky(v int32)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *TestCasebulkCasesInner) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### GetAuthorId

`func (o *TestCasebulkCasesInner) GetAuthorId() int32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *TestCasebulkCasesInner) GetAuthorIdOk() (*int32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *TestCasebulkCasesInner) SetAuthorId(v int32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *TestCasebulkCasesInner) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetSuiteId

`func (o *TestCasebulkCasesInner) GetSuiteId() int64`

GetSuiteId returns the SuiteId field if non-nil, zero value otherwise.

### GetSuiteIdOk

`func (o *TestCasebulkCasesInner) GetSuiteIdOk() (*int64, bool)`

GetSuiteIdOk returns a tuple with the SuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteId

`func (o *TestCasebulkCasesInner) SetSuiteId(v int64)`

SetSuiteId sets SuiteId field to given value.

### HasSuiteId

`func (o *TestCasebulkCasesInner) HasSuiteId() bool`

HasSuiteId returns a boolean if a field has been set.

### GetMilestoneId

`func (o *TestCasebulkCasesInner) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *TestCasebulkCasesInner) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *TestCasebulkCasesInner) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *TestCasebulkCasesInner) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### GetAutomation

`func (o *TestCasebulkCasesInner) GetAutomation() int32`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *TestCasebulkCasesInner) GetAutomationOk() (*int32, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *TestCasebulkCasesInner) SetAutomation(v int32)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *TestCasebulkCasesInner) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetStatus

`func (o *TestCasebulkCasesInner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestCasebulkCasesInner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestCasebulkCasesInner) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestCasebulkCasesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAttachments

`func (o *TestCasebulkCasesInner) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestCasebulkCasesInner) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestCasebulkCasesInner) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestCasebulkCasesInner) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *TestCasebulkCasesInner) GetSteps() []TestStepCreate`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TestCasebulkCasesInner) GetStepsOk() (*[]TestStepCreate, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TestCasebulkCasesInner) SetSteps(v []TestStepCreate)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TestCasebulkCasesInner) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *TestCasebulkCasesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestCasebulkCasesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestCasebulkCasesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestCasebulkCasesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParams

`func (o *TestCasebulkCasesInner) GetParams() map[string][]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TestCasebulkCasesInner) GetParamsOk() (*map[string][]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TestCasebulkCasesInner) SetParams(v map[string][]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *TestCasebulkCasesInner) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *TestCasebulkCasesInner) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *TestCasebulkCasesInner) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetCustomField

`func (o *TestCasebulkCasesInner) GetCustomField() map[string]string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *TestCasebulkCasesInner) GetCustomFieldOk() (*map[string]string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *TestCasebulkCasesInner) SetCustomField(v map[string]string)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *TestCasebulkCasesInner) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TestCasebulkCasesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestCasebulkCasesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestCasebulkCasesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TestCasebulkCasesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TestCasebulkCasesInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TestCasebulkCasesInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TestCasebulkCasesInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TestCasebulkCasesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetId

`func (o *TestCasebulkCasesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestCasebulkCasesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestCasebulkCasesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TestCasebulkCasesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TestCasebulkCasesInner) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TestCasebulkCasesInner) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


