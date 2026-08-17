# ReviewCaseData

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
**IsMuted** | Pointer to **bool** | Mute state of the proposed test case. | [optional] 
**SuiteId** | Pointer to **NullableInt64** |  | [optional] 
**MilestoneId** | Pointer to **NullableInt64** |  | [optional] 
**IsManual** | Pointer to **bool** | &#x60;true&#x60; if the case is manual, &#x60;false&#x60; if it is automated. | [optional] 
**IsToBeAutomated** | Pointer to **bool** | &#x60;true&#x60; if a manual case is planned to be automated. | [optional] 
**Status** | Pointer to **NullableInt32** |  | [optional] 
**StepsType** | Pointer to **string** | Format of the steps field. Omit to keep the current one, &#x60;classic&#x60; for a new-case draft; changing it requires sending &#x60;steps&#x60; in the same request. | [optional] 
**Attachments** | Pointer to **[]string** | A list of Attachment hashes. | [optional] 
**Steps** | Pointer to [**[]ReviewStepData**](ReviewStepData.md) | For gherkin steps send the scenario in &#x60;value&#x60;. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Parameters** | Pointer to [**[]TestCaseParameterCreate**](TestCaseParameterCreate.md) |  | [optional] 
**CustomField** | Pointer to **map[string]string** | Map of custom field ID to value. A &#x60;create&#x60; review must carry every required custom field. An &#x60;edit&#x60; review is validated against the current test case, so send only the fields the proposal changes. | [optional] 

## Methods

### NewReviewCaseData

`func NewReviewCaseData() *ReviewCaseData`

NewReviewCaseData instantiates a new ReviewCaseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewCaseDataWithDefaults

`func NewReviewCaseDataWithDefaults() *ReviewCaseData`

NewReviewCaseDataWithDefaults instantiates a new ReviewCaseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ReviewCaseData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReviewCaseData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReviewCaseData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReviewCaseData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ReviewCaseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReviewCaseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReviewCaseData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReviewCaseData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ReviewCaseData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ReviewCaseData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPreconditions

`func (o *ReviewCaseData) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *ReviewCaseData) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *ReviewCaseData) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *ReviewCaseData) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### SetPreconditionsNil

`func (o *ReviewCaseData) SetPreconditionsNil(b bool)`

 SetPreconditionsNil sets the value for Preconditions to be an explicit nil

### UnsetPreconditions
`func (o *ReviewCaseData) UnsetPreconditions()`

UnsetPreconditions ensures that no value is present for Preconditions, not even an explicit nil
### GetPostconditions

`func (o *ReviewCaseData) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *ReviewCaseData) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *ReviewCaseData) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *ReviewCaseData) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### SetPostconditionsNil

`func (o *ReviewCaseData) SetPostconditionsNil(b bool)`

 SetPostconditionsNil sets the value for Postconditions to be an explicit nil

### UnsetPostconditions
`func (o *ReviewCaseData) UnsetPostconditions()`

UnsetPostconditions ensures that no value is present for Postconditions, not even an explicit nil
### GetSeverity

`func (o *ReviewCaseData) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ReviewCaseData) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ReviewCaseData) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ReviewCaseData) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *ReviewCaseData) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *ReviewCaseData) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetPriority

`func (o *ReviewCaseData) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ReviewCaseData) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ReviewCaseData) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ReviewCaseData) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ReviewCaseData) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ReviewCaseData) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetBehavior

`func (o *ReviewCaseData) GetBehavior() int32`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *ReviewCaseData) GetBehaviorOk() (*int32, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *ReviewCaseData) SetBehavior(v int32)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *ReviewCaseData) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### SetBehaviorNil

`func (o *ReviewCaseData) SetBehaviorNil(b bool)`

 SetBehaviorNil sets the value for Behavior to be an explicit nil

### UnsetBehavior
`func (o *ReviewCaseData) UnsetBehavior()`

UnsetBehavior ensures that no value is present for Behavior, not even an explicit nil
### GetType

`func (o *ReviewCaseData) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReviewCaseData) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReviewCaseData) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ReviewCaseData) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ReviewCaseData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ReviewCaseData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetLayer

`func (o *ReviewCaseData) GetLayer() int32`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *ReviewCaseData) GetLayerOk() (*int32, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *ReviewCaseData) SetLayer(v int32)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *ReviewCaseData) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### SetLayerNil

`func (o *ReviewCaseData) SetLayerNil(b bool)`

 SetLayerNil sets the value for Layer to be an explicit nil

### UnsetLayer
`func (o *ReviewCaseData) UnsetLayer()`

UnsetLayer ensures that no value is present for Layer, not even an explicit nil
### GetIsFlaky

`func (o *ReviewCaseData) GetIsFlaky() int32`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *ReviewCaseData) GetIsFlakyOk() (*int32, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *ReviewCaseData) SetIsFlaky(v int32)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *ReviewCaseData) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *ReviewCaseData) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *ReviewCaseData) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetIsMuted

`func (o *ReviewCaseData) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *ReviewCaseData) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *ReviewCaseData) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *ReviewCaseData) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetSuiteId

`func (o *ReviewCaseData) GetSuiteId() int64`

GetSuiteId returns the SuiteId field if non-nil, zero value otherwise.

### GetSuiteIdOk

`func (o *ReviewCaseData) GetSuiteIdOk() (*int64, bool)`

GetSuiteIdOk returns a tuple with the SuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteId

`func (o *ReviewCaseData) SetSuiteId(v int64)`

SetSuiteId sets SuiteId field to given value.

### HasSuiteId

`func (o *ReviewCaseData) HasSuiteId() bool`

HasSuiteId returns a boolean if a field has been set.

### SetSuiteIdNil

`func (o *ReviewCaseData) SetSuiteIdNil(b bool)`

 SetSuiteIdNil sets the value for SuiteId to be an explicit nil

### UnsetSuiteId
`func (o *ReviewCaseData) UnsetSuiteId()`

UnsetSuiteId ensures that no value is present for SuiteId, not even an explicit nil
### GetMilestoneId

`func (o *ReviewCaseData) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *ReviewCaseData) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *ReviewCaseData) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *ReviewCaseData) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *ReviewCaseData) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *ReviewCaseData) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetIsManual

`func (o *ReviewCaseData) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *ReviewCaseData) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *ReviewCaseData) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *ReviewCaseData) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### GetIsToBeAutomated

`func (o *ReviewCaseData) GetIsToBeAutomated() bool`

GetIsToBeAutomated returns the IsToBeAutomated field if non-nil, zero value otherwise.

### GetIsToBeAutomatedOk

`func (o *ReviewCaseData) GetIsToBeAutomatedOk() (*bool, bool)`

GetIsToBeAutomatedOk returns a tuple with the IsToBeAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsToBeAutomated

`func (o *ReviewCaseData) SetIsToBeAutomated(v bool)`

SetIsToBeAutomated sets IsToBeAutomated field to given value.

### HasIsToBeAutomated

`func (o *ReviewCaseData) HasIsToBeAutomated() bool`

HasIsToBeAutomated returns a boolean if a field has been set.

### GetStatus

`func (o *ReviewCaseData) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReviewCaseData) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReviewCaseData) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReviewCaseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ReviewCaseData) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReviewCaseData) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStepsType

`func (o *ReviewCaseData) GetStepsType() string`

GetStepsType returns the StepsType field if non-nil, zero value otherwise.

### GetStepsTypeOk

`func (o *ReviewCaseData) GetStepsTypeOk() (*string, bool)`

GetStepsTypeOk returns a tuple with the StepsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsType

`func (o *ReviewCaseData) SetStepsType(v string)`

SetStepsType sets StepsType field to given value.

### HasStepsType

`func (o *ReviewCaseData) HasStepsType() bool`

HasStepsType returns a boolean if a field has been set.

### GetAttachments

`func (o *ReviewCaseData) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ReviewCaseData) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ReviewCaseData) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ReviewCaseData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *ReviewCaseData) GetSteps() []ReviewStepData`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ReviewCaseData) GetStepsOk() (*[]ReviewStepData, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ReviewCaseData) SetSteps(v []ReviewStepData)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ReviewCaseData) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *ReviewCaseData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReviewCaseData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReviewCaseData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReviewCaseData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ReviewCaseData) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ReviewCaseData) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetParameters

`func (o *ReviewCaseData) GetParameters() []TestCaseParameterCreate`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ReviewCaseData) GetParametersOk() (*[]TestCaseParameterCreate, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ReviewCaseData) SetParameters(v []TestCaseParameterCreate)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ReviewCaseData) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *ReviewCaseData) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *ReviewCaseData) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetCustomField

`func (o *ReviewCaseData) GetCustomField() map[string]string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *ReviewCaseData) GetCustomFieldOk() (*map[string]string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *ReviewCaseData) SetCustomField(v map[string]string)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *ReviewCaseData) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


