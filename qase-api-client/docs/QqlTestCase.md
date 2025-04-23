# QqlTestCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**TestCaseId** | **int64** |  | 
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
**UpdatedBy** | Pointer to **int64** | Author ID of the last update. | [optional] 

## Methods

### NewQqlTestCase

`func NewQqlTestCase(testCaseId int64, ) *QqlTestCase`

NewQqlTestCase instantiates a new QqlTestCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQqlTestCaseWithDefaults

`func NewQqlTestCaseWithDefaults() *QqlTestCase`

NewQqlTestCaseWithDefaults instantiates a new QqlTestCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QqlTestCase) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QqlTestCase) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QqlTestCase) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QqlTestCase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTestCaseId

`func (o *QqlTestCase) GetTestCaseId() int64`

GetTestCaseId returns the TestCaseId field if non-nil, zero value otherwise.

### GetTestCaseIdOk

`func (o *QqlTestCase) GetTestCaseIdOk() (*int64, bool)`

GetTestCaseIdOk returns a tuple with the TestCaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCaseId

`func (o *QqlTestCase) SetTestCaseId(v int64)`

SetTestCaseId sets TestCaseId field to given value.


### GetPosition

`func (o *QqlTestCase) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *QqlTestCase) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *QqlTestCase) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *QqlTestCase) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTitle

`func (o *QqlTestCase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QqlTestCase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QqlTestCase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QqlTestCase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *QqlTestCase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QqlTestCase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QqlTestCase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QqlTestCase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *QqlTestCase) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QqlTestCase) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPreconditions

`func (o *QqlTestCase) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *QqlTestCase) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *QqlTestCase) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *QqlTestCase) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### SetPreconditionsNil

`func (o *QqlTestCase) SetPreconditionsNil(b bool)`

 SetPreconditionsNil sets the value for Preconditions to be an explicit nil

### UnsetPreconditions
`func (o *QqlTestCase) UnsetPreconditions()`

UnsetPreconditions ensures that no value is present for Preconditions, not even an explicit nil
### GetPostconditions

`func (o *QqlTestCase) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *QqlTestCase) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *QqlTestCase) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *QqlTestCase) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### SetPostconditionsNil

`func (o *QqlTestCase) SetPostconditionsNil(b bool)`

 SetPostconditionsNil sets the value for Postconditions to be an explicit nil

### UnsetPostconditions
`func (o *QqlTestCase) UnsetPostconditions()`

UnsetPostconditions ensures that no value is present for Postconditions, not even an explicit nil
### GetSeverity

`func (o *QqlTestCase) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *QqlTestCase) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *QqlTestCase) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *QqlTestCase) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetPriority

`func (o *QqlTestCase) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *QqlTestCase) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *QqlTestCase) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *QqlTestCase) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetType

`func (o *QqlTestCase) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QqlTestCase) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QqlTestCase) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *QqlTestCase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayer

`func (o *QqlTestCase) GetLayer() int32`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *QqlTestCase) GetLayerOk() (*int32, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *QqlTestCase) SetLayer(v int32)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *QqlTestCase) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetIsFlaky

`func (o *QqlTestCase) GetIsFlaky() int32`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *QqlTestCase) GetIsFlakyOk() (*int32, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *QqlTestCase) SetIsFlaky(v int32)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *QqlTestCase) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### GetBehavior

`func (o *QqlTestCase) GetBehavior() int32`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *QqlTestCase) GetBehaviorOk() (*int32, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *QqlTestCase) SetBehavior(v int32)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *QqlTestCase) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetAutomation

`func (o *QqlTestCase) GetAutomation() int32`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *QqlTestCase) GetAutomationOk() (*int32, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *QqlTestCase) SetAutomation(v int32)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *QqlTestCase) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetStatus

`func (o *QqlTestCase) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QqlTestCase) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QqlTestCase) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QqlTestCase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMilestoneId

`func (o *QqlTestCase) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *QqlTestCase) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *QqlTestCase) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *QqlTestCase) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *QqlTestCase) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *QqlTestCase) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetSuiteId

`func (o *QqlTestCase) GetSuiteId() int64`

GetSuiteId returns the SuiteId field if non-nil, zero value otherwise.

### GetSuiteIdOk

`func (o *QqlTestCase) GetSuiteIdOk() (*int64, bool)`

GetSuiteIdOk returns a tuple with the SuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteId

`func (o *QqlTestCase) SetSuiteId(v int64)`

SetSuiteId sets SuiteId field to given value.

### HasSuiteId

`func (o *QqlTestCase) HasSuiteId() bool`

HasSuiteId returns a boolean if a field has been set.

### SetSuiteIdNil

`func (o *QqlTestCase) SetSuiteIdNil(b bool)`

 SetSuiteIdNil sets the value for SuiteId to be an explicit nil

### UnsetSuiteId
`func (o *QqlTestCase) UnsetSuiteId()`

UnsetSuiteId ensures that no value is present for SuiteId, not even an explicit nil
### GetCustomFields

`func (o *QqlTestCase) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QqlTestCase) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QqlTestCase) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QqlTestCase) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAttachments

`func (o *QqlTestCase) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *QqlTestCase) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *QqlTestCase) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *QqlTestCase) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetStepsType

`func (o *QqlTestCase) GetStepsType() string`

GetStepsType returns the StepsType field if non-nil, zero value otherwise.

### GetStepsTypeOk

`func (o *QqlTestCase) GetStepsTypeOk() (*string, bool)`

GetStepsTypeOk returns a tuple with the StepsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsType

`func (o *QqlTestCase) SetStepsType(v string)`

SetStepsType sets StepsType field to given value.

### HasStepsType

`func (o *QqlTestCase) HasStepsType() bool`

HasStepsType returns a boolean if a field has been set.

### SetStepsTypeNil

`func (o *QqlTestCase) SetStepsTypeNil(b bool)`

 SetStepsTypeNil sets the value for StepsType to be an explicit nil

### UnsetStepsType
`func (o *QqlTestCase) UnsetStepsType()`

UnsetStepsType ensures that no value is present for StepsType, not even an explicit nil
### GetSteps

`func (o *QqlTestCase) GetSteps() []TestStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *QqlTestCase) GetStepsOk() (*[]TestStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *QqlTestCase) SetSteps(v []TestStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *QqlTestCase) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetParams

`func (o *QqlTestCase) GetParams() TestCaseParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *QqlTestCase) GetParamsOk() (*TestCaseParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *QqlTestCase) SetParams(v TestCaseParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *QqlTestCase) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetTags

`func (o *QqlTestCase) GetTags() []TagValue`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QqlTestCase) GetTagsOk() (*[]TagValue, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QqlTestCase) SetTags(v []TagValue)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QqlTestCase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMemberId

`func (o *QqlTestCase) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *QqlTestCase) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *QqlTestCase) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *QqlTestCase) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetAuthorId

`func (o *QqlTestCase) GetAuthorId() int64`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *QqlTestCase) GetAuthorIdOk() (*int64, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *QqlTestCase) SetAuthorId(v int64)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *QqlTestCase) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QqlTestCase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QqlTestCase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QqlTestCase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QqlTestCase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *QqlTestCase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QqlTestCase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QqlTestCase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *QqlTestCase) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *QqlTestCase) GetUpdatedBy() int64`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *QqlTestCase) GetUpdatedByOk() (*int64, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *QqlTestCase) SetUpdatedBy(v int64)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *QqlTestCase) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


