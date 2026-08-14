# ReviewProposedCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Preconditions** | Pointer to **NullableString** |  | [optional] 
**Postconditions** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **NullableInt32** |  | [optional] 
**Priority** | Pointer to **NullableInt32** |  | [optional] 
**Behavior** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **NullableInt32** |  | [optional] 
**Layer** | Pointer to **NullableInt32** |  | [optional] 
**IsFlaky** | Pointer to **NullableInt32** |  | [optional] 
**IsMuted** | Pointer to **bool** |  | [optional] 
**SuiteId** | Pointer to **NullableInt64** |  | [optional] 
**MilestoneId** | Pointer to **NullableInt64** |  | [optional] 
**IsManual** | Pointer to **bool** | &#x60;true&#x60; if the case is manual, &#x60;false&#x60; if it is automated. | [optional] 
**IsToBeAutomated** | Pointer to **bool** | &#x60;true&#x60; if a manual case is planned to be automated. | [optional] 
**Status** | Pointer to **NullableInt32** |  | [optional] 
**StepsType** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to **[]string** | Attachment hashes. | [optional] 
**Steps** | Pointer to [**[]ReviewProposedStep**](ReviewProposedStep.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Parameters** | Pointer to [**[]TestCaseParameter**](TestCaseParameter.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewReviewProposedCase

`func NewReviewProposedCase() *ReviewProposedCase`

NewReviewProposedCase instantiates a new ReviewProposedCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewProposedCaseWithDefaults

`func NewReviewProposedCaseWithDefaults() *ReviewProposedCase`

NewReviewProposedCaseWithDefaults instantiates a new ReviewProposedCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ReviewProposedCase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReviewProposedCase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReviewProposedCase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReviewProposedCase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ReviewProposedCase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReviewProposedCase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReviewProposedCase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReviewProposedCase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ReviewProposedCase) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ReviewProposedCase) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPreconditions

`func (o *ReviewProposedCase) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *ReviewProposedCase) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *ReviewProposedCase) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *ReviewProposedCase) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### SetPreconditionsNil

`func (o *ReviewProposedCase) SetPreconditionsNil(b bool)`

 SetPreconditionsNil sets the value for Preconditions to be an explicit nil

### UnsetPreconditions
`func (o *ReviewProposedCase) UnsetPreconditions()`

UnsetPreconditions ensures that no value is present for Preconditions, not even an explicit nil
### GetPostconditions

`func (o *ReviewProposedCase) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *ReviewProposedCase) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *ReviewProposedCase) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *ReviewProposedCase) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### SetPostconditionsNil

`func (o *ReviewProposedCase) SetPostconditionsNil(b bool)`

 SetPostconditionsNil sets the value for Postconditions to be an explicit nil

### UnsetPostconditions
`func (o *ReviewProposedCase) UnsetPostconditions()`

UnsetPostconditions ensures that no value is present for Postconditions, not even an explicit nil
### GetSeverity

`func (o *ReviewProposedCase) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ReviewProposedCase) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ReviewProposedCase) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ReviewProposedCase) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *ReviewProposedCase) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *ReviewProposedCase) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetPriority

`func (o *ReviewProposedCase) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ReviewProposedCase) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ReviewProposedCase) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ReviewProposedCase) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ReviewProposedCase) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ReviewProposedCase) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetBehavior

`func (o *ReviewProposedCase) GetBehavior() int32`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *ReviewProposedCase) GetBehaviorOk() (*int32, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *ReviewProposedCase) SetBehavior(v int32)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *ReviewProposedCase) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### SetBehaviorNil

`func (o *ReviewProposedCase) SetBehaviorNil(b bool)`

 SetBehaviorNil sets the value for Behavior to be an explicit nil

### UnsetBehavior
`func (o *ReviewProposedCase) UnsetBehavior()`

UnsetBehavior ensures that no value is present for Behavior, not even an explicit nil
### GetType

`func (o *ReviewProposedCase) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReviewProposedCase) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReviewProposedCase) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ReviewProposedCase) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ReviewProposedCase) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ReviewProposedCase) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetLayer

`func (o *ReviewProposedCase) GetLayer() int32`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *ReviewProposedCase) GetLayerOk() (*int32, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *ReviewProposedCase) SetLayer(v int32)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *ReviewProposedCase) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### SetLayerNil

`func (o *ReviewProposedCase) SetLayerNil(b bool)`

 SetLayerNil sets the value for Layer to be an explicit nil

### UnsetLayer
`func (o *ReviewProposedCase) UnsetLayer()`

UnsetLayer ensures that no value is present for Layer, not even an explicit nil
### GetIsFlaky

`func (o *ReviewProposedCase) GetIsFlaky() int32`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *ReviewProposedCase) GetIsFlakyOk() (*int32, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *ReviewProposedCase) SetIsFlaky(v int32)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *ReviewProposedCase) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *ReviewProposedCase) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *ReviewProposedCase) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetIsMuted

`func (o *ReviewProposedCase) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *ReviewProposedCase) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *ReviewProposedCase) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *ReviewProposedCase) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetSuiteId

`func (o *ReviewProposedCase) GetSuiteId() int64`

GetSuiteId returns the SuiteId field if non-nil, zero value otherwise.

### GetSuiteIdOk

`func (o *ReviewProposedCase) GetSuiteIdOk() (*int64, bool)`

GetSuiteIdOk returns a tuple with the SuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteId

`func (o *ReviewProposedCase) SetSuiteId(v int64)`

SetSuiteId sets SuiteId field to given value.

### HasSuiteId

`func (o *ReviewProposedCase) HasSuiteId() bool`

HasSuiteId returns a boolean if a field has been set.

### SetSuiteIdNil

`func (o *ReviewProposedCase) SetSuiteIdNil(b bool)`

 SetSuiteIdNil sets the value for SuiteId to be an explicit nil

### UnsetSuiteId
`func (o *ReviewProposedCase) UnsetSuiteId()`

UnsetSuiteId ensures that no value is present for SuiteId, not even an explicit nil
### GetMilestoneId

`func (o *ReviewProposedCase) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *ReviewProposedCase) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *ReviewProposedCase) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *ReviewProposedCase) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *ReviewProposedCase) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *ReviewProposedCase) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetIsManual

`func (o *ReviewProposedCase) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *ReviewProposedCase) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *ReviewProposedCase) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *ReviewProposedCase) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### GetIsToBeAutomated

`func (o *ReviewProposedCase) GetIsToBeAutomated() bool`

GetIsToBeAutomated returns the IsToBeAutomated field if non-nil, zero value otherwise.

### GetIsToBeAutomatedOk

`func (o *ReviewProposedCase) GetIsToBeAutomatedOk() (*bool, bool)`

GetIsToBeAutomatedOk returns a tuple with the IsToBeAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsToBeAutomated

`func (o *ReviewProposedCase) SetIsToBeAutomated(v bool)`

SetIsToBeAutomated sets IsToBeAutomated field to given value.

### HasIsToBeAutomated

`func (o *ReviewProposedCase) HasIsToBeAutomated() bool`

HasIsToBeAutomated returns a boolean if a field has been set.

### GetStatus

`func (o *ReviewProposedCase) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReviewProposedCase) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReviewProposedCase) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReviewProposedCase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ReviewProposedCase) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReviewProposedCase) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStepsType

`func (o *ReviewProposedCase) GetStepsType() string`

GetStepsType returns the StepsType field if non-nil, zero value otherwise.

### GetStepsTypeOk

`func (o *ReviewProposedCase) GetStepsTypeOk() (*string, bool)`

GetStepsTypeOk returns a tuple with the StepsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsType

`func (o *ReviewProposedCase) SetStepsType(v string)`

SetStepsType sets StepsType field to given value.

### HasStepsType

`func (o *ReviewProposedCase) HasStepsType() bool`

HasStepsType returns a boolean if a field has been set.

### GetAttachments

`func (o *ReviewProposedCase) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ReviewProposedCase) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ReviewProposedCase) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ReviewProposedCase) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *ReviewProposedCase) GetSteps() []ReviewProposedStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ReviewProposedCase) GetStepsOk() (*[]ReviewProposedStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ReviewProposedCase) SetSteps(v []ReviewProposedStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ReviewProposedCase) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *ReviewProposedCase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReviewProposedCase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReviewProposedCase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReviewProposedCase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ReviewProposedCase) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ReviewProposedCase) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetParameters

`func (o *ReviewProposedCase) GetParameters() []TestCaseParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ReviewProposedCase) GetParametersOk() (*[]TestCaseParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ReviewProposedCase) SetParameters(v []TestCaseParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ReviewProposedCase) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetCustomFields

`func (o *ReviewProposedCase) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ReviewProposedCase) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ReviewProposedCase) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ReviewProposedCase) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


