# TestCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Preconditions** | Pointer to **NullableString** |  | [optional] 
**Postconditions** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Layer** | Pointer to **int32** |  | [optional] 
**IsFlaky** | Pointer to **int32** |  | [optional] 
**Behavior** | Pointer to **int32** |  | [optional] 
**Automation** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**MilestoneId** | Pointer to **NullableInt64** |  | [optional] 
**SuiteId** | Pointer to **NullableInt64** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**StepsType** | Pointer to **NullableString** |  | [optional] 
**Steps** | Pointer to [**[]TestStep**](TestStep.md) |  | [optional] 
**Params** | Pointer to [**TestCaseParams**](TestCaseParams.md) |  | [optional] 
**Tags** | Pointer to [**[]TagValue**](TagValue.md) |  | [optional] 
**MemberId** | Pointer to **int64** | Deprecated, use &#x60;author_id&#x60; instead. | [optional] 
**AuthorId** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Deleted** | Pointer to **NullableString** |  | [optional] 
**Created** | Pointer to **string** | Deprecated, use the &#x60;created_at&#x60; property instead. | [optional] 
**Updated** | Pointer to **string** | Deprecated, use the &#x60;updated_at&#x60; property instead. | [optional] 
**ExternalIssues** | Pointer to [**[]ExternalIssue**](ExternalIssue.md) |  | [optional] 

## Methods

### NewTestCase

`func NewTestCase() *TestCase`

NewTestCase instantiates a new TestCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCaseWithDefaults

`func NewTestCaseWithDefaults() *TestCase`

NewTestCaseWithDefaults instantiates a new TestCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestCase) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestCase) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestCase) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TestCase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *TestCase) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TestCase) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TestCase) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *TestCase) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTitle

`func (o *TestCase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestCase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestCase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TestCase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TestCase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestCase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestCase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestCase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestCase) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestCase) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPreconditions

`func (o *TestCase) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *TestCase) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *TestCase) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *TestCase) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### SetPreconditionsNil

`func (o *TestCase) SetPreconditionsNil(b bool)`

 SetPreconditionsNil sets the value for Preconditions to be an explicit nil

### UnsetPreconditions
`func (o *TestCase) UnsetPreconditions()`

UnsetPreconditions ensures that no value is present for Preconditions, not even an explicit nil
### GetPostconditions

`func (o *TestCase) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *TestCase) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *TestCase) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *TestCase) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### SetPostconditionsNil

`func (o *TestCase) SetPostconditionsNil(b bool)`

 SetPostconditionsNil sets the value for Postconditions to be an explicit nil

### UnsetPostconditions
`func (o *TestCase) UnsetPostconditions()`

UnsetPostconditions ensures that no value is present for Postconditions, not even an explicit nil
### GetSeverity

`func (o *TestCase) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TestCase) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TestCase) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *TestCase) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetPriority

`func (o *TestCase) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestCase) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestCase) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TestCase) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetType

`func (o *TestCase) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestCase) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestCase) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *TestCase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayer

`func (o *TestCase) GetLayer() int32`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *TestCase) GetLayerOk() (*int32, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *TestCase) SetLayer(v int32)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *TestCase) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetIsFlaky

`func (o *TestCase) GetIsFlaky() int32`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *TestCase) GetIsFlakyOk() (*int32, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *TestCase) SetIsFlaky(v int32)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *TestCase) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### GetBehavior

`func (o *TestCase) GetBehavior() int32`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *TestCase) GetBehaviorOk() (*int32, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *TestCase) SetBehavior(v int32)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *TestCase) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetAutomation

`func (o *TestCase) GetAutomation() int32`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *TestCase) GetAutomationOk() (*int32, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *TestCase) SetAutomation(v int32)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *TestCase) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetStatus

`func (o *TestCase) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestCase) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestCase) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestCase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMilestoneId

`func (o *TestCase) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *TestCase) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *TestCase) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *TestCase) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *TestCase) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *TestCase) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetSuiteId

`func (o *TestCase) GetSuiteId() int64`

GetSuiteId returns the SuiteId field if non-nil, zero value otherwise.

### GetSuiteIdOk

`func (o *TestCase) GetSuiteIdOk() (*int64, bool)`

GetSuiteIdOk returns a tuple with the SuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteId

`func (o *TestCase) SetSuiteId(v int64)`

SetSuiteId sets SuiteId field to given value.

### HasSuiteId

`func (o *TestCase) HasSuiteId() bool`

HasSuiteId returns a boolean if a field has been set.

### SetSuiteIdNil

`func (o *TestCase) SetSuiteIdNil(b bool)`

 SetSuiteIdNil sets the value for SuiteId to be an explicit nil

### UnsetSuiteId
`func (o *TestCase) UnsetSuiteId()`

UnsetSuiteId ensures that no value is present for SuiteId, not even an explicit nil
### GetCustomFields

`func (o *TestCase) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TestCase) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TestCase) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TestCase) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAttachments

`func (o *TestCase) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestCase) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestCase) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestCase) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetStepsType

`func (o *TestCase) GetStepsType() string`

GetStepsType returns the StepsType field if non-nil, zero value otherwise.

### GetStepsTypeOk

`func (o *TestCase) GetStepsTypeOk() (*string, bool)`

GetStepsTypeOk returns a tuple with the StepsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsType

`func (o *TestCase) SetStepsType(v string)`

SetStepsType sets StepsType field to given value.

### HasStepsType

`func (o *TestCase) HasStepsType() bool`

HasStepsType returns a boolean if a field has been set.

### SetStepsTypeNil

`func (o *TestCase) SetStepsTypeNil(b bool)`

 SetStepsTypeNil sets the value for StepsType to be an explicit nil

### UnsetStepsType
`func (o *TestCase) UnsetStepsType()`

UnsetStepsType ensures that no value is present for StepsType, not even an explicit nil
### GetSteps

`func (o *TestCase) GetSteps() []TestStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TestCase) GetStepsOk() (*[]TestStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TestCase) SetSteps(v []TestStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TestCase) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetParams

`func (o *TestCase) GetParams() TestCaseParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TestCase) GetParamsOk() (*TestCaseParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TestCase) SetParams(v TestCaseParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *TestCase) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetTags

`func (o *TestCase) GetTags() []TagValue`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestCase) GetTagsOk() (*[]TagValue, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestCase) SetTags(v []TagValue)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestCase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMemberId

`func (o *TestCase) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *TestCase) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *TestCase) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *TestCase) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetAuthorId

`func (o *TestCase) GetAuthorId() int64`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *TestCase) GetAuthorIdOk() (*int64, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *TestCase) SetAuthorId(v int64)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *TestCase) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TestCase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestCase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestCase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TestCase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TestCase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TestCase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TestCase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TestCase) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *TestCase) GetDeleted() string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *TestCase) GetDeletedOk() (*string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *TestCase) SetDeleted(v string)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *TestCase) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *TestCase) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *TestCase) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetCreated

`func (o *TestCase) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TestCase) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TestCase) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TestCase) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *TestCase) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TestCase) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TestCase) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *TestCase) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetExternalIssues

`func (o *TestCase) GetExternalIssues() []ExternalIssue`

GetExternalIssues returns the ExternalIssues field if non-nil, zero value otherwise.

### GetExternalIssuesOk

`func (o *TestCase) GetExternalIssuesOk() (*[]ExternalIssue, bool)`

GetExternalIssuesOk returns a tuple with the ExternalIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIssues

`func (o *TestCase) SetExternalIssues(v []ExternalIssue)`

SetExternalIssues sets ExternalIssues field to given value.

### HasExternalIssues

`func (o *TestCase) HasExternalIssues() bool`

HasExternalIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


