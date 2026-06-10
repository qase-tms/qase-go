# TestCaseCreate

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
**SuiteId** | Pointer to **int64** |  | [optional] 
**MilestoneId** | Pointer to **int64** |  | [optional] 
**Automation** | Pointer to **int32** | Deprecated, use &#x60;isManual&#x60; and &#x60;isToBeAutomated&#x60; instead. Encodes the test case automation state as a single integer: &#x60;0&#x60; &#x3D; manual, &#x60;1&#x60; &#x3D; manual planned to be automated, &#x60;2&#x60; &#x3D; automated. If both &#x60;automation&#x60; and &#x60;isManual&#x60;/&#x60;isToBeAutomated&#x60; are provided, &#x60;isManual&#x60; and &#x60;isToBeAutomated&#x60; take precedence. | [optional] 
**IsManual** | Pointer to **int32** | &#x60;1&#x60; if the case is manual, &#x60;0&#x60; if it is automated. Combined with &#x60;isToBeAutomated&#x60;, replaces the deprecated &#x60;automation&#x60; field. | [optional] 
**IsToBeAutomated** | Pointer to **int32** | &#x60;1&#x60; if a manual case is planned to be automated, &#x60;0&#x60; otherwise. Only meaningful when &#x60;isManual &#x3D; 1&#x60;; ignored when &#x60;isManual &#x3D; 0&#x60;. | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StepsType** | Pointer to **string** | Determines the format of the steps field. When \&quot;classic\&quot;, steps use the standard action/expected_result/data format. When \&quot;gherkin\&quot;, steps use the {value: \&quot;Given...\\nWhen...\\nThen...\&quot;} format. | [optional] [default to "classic"]
**Attachments** | Pointer to **[]string** | A list of Attachment hashes. | [optional] 
**Steps** | Pointer to [**[]TestStepCreate**](TestStepCreate.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Params** | Pointer to **map[string][]string** | Deprecated, use &#x60;parameters&#x60; instead. | [optional] 
**Parameters** | Pointer to [**[]TestCaseParameterCreate**](TestCaseParameterCreate.md) |  | [optional] 
**CustomField** | Pointer to **map[string]string** | Custom field values keyed by the field&#39;s project-scoped &#x60;internal_id&#x60; (see &#x60;GET /custom_field&#x60;). Values are always **scalar strings**; arrays, objects or non-scalars are rejected.  | Field type           | Value format                              | Example                 | |----------------------|-------------------------------------------|-------------------------| | &#x60;string&#x60;, &#x60;text&#x60;     | Plain string                              | &#x60;\&quot;hello\&quot;&#x60;               | | &#x60;number&#x60;             | Numeric string                            | &#x60;\&quot;42\&quot;&#x60;                  | | &#x60;url&#x60;                | Valid URL                                 | &#x60;\&quot;https://qase.io\&quot;&#x60;     | | &#x60;datetime&#x60;           | Absolute date (ISO 8601 recommended)      | &#x60;\&quot;2026-04-29T15:00:00Z\&quot;&#x60;| | &#x60;selectbox&#x60;, &#x60;radio&#x60; | Option &#x60;id&#x60; as string                     | &#x60;\&quot;1\&quot;&#x60;                   | | &#x60;multiselect&#x60;        | Comma-separated option &#x60;id&#x60;s (no spaces)  | &#x60;\&quot;1,2,3\&quot;&#x60;               | | &#x60;checkbox&#x60;           | &#x60;\&quot;1\&quot;&#x60; to check, &#x60;\&quot;\&quot;&#x60; to uncheck           | &#x60;\&quot;1\&quot;&#x60;                   | | &#x60;user&#x60;               | Team member &#x60;internal_id&#x60; as string       | &#x60;\&quot;42\&quot;&#x60;                  |  Validation: all required fields without a default value must be present and non-empty; unknown &#x60;internal_id&#x60;s are rejected; option-based values must reference an existing option.  Note: a &#x60;required&#x60; checkbox without a default cannot be unchecked via the API — set a default or clear &#x60;required&#x60; in workspace settings.  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewTestCaseCreate

`func NewTestCaseCreate(title string, ) *TestCaseCreate`

NewTestCaseCreate instantiates a new TestCaseCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCaseCreateWithDefaults

`func NewTestCaseCreateWithDefaults() *TestCaseCreate`

NewTestCaseCreateWithDefaults instantiates a new TestCaseCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TestCaseCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestCaseCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestCaseCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestCaseCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPreconditions

`func (o *TestCaseCreate) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *TestCaseCreate) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *TestCaseCreate) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *TestCaseCreate) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### GetPostconditions

`func (o *TestCaseCreate) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *TestCaseCreate) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *TestCaseCreate) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *TestCaseCreate) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### GetTitle

`func (o *TestCaseCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestCaseCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestCaseCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSeverity

`func (o *TestCaseCreate) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TestCaseCreate) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TestCaseCreate) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *TestCaseCreate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetPriority

`func (o *TestCaseCreate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestCaseCreate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestCaseCreate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TestCaseCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetBehavior

`func (o *TestCaseCreate) GetBehavior() int32`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *TestCaseCreate) GetBehaviorOk() (*int32, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *TestCaseCreate) SetBehavior(v int32)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *TestCaseCreate) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetType

`func (o *TestCaseCreate) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestCaseCreate) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestCaseCreate) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *TestCaseCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayer

`func (o *TestCaseCreate) GetLayer() int32`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *TestCaseCreate) GetLayerOk() (*int32, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *TestCaseCreate) SetLayer(v int32)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *TestCaseCreate) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetIsFlaky

`func (o *TestCaseCreate) GetIsFlaky() int32`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *TestCaseCreate) GetIsFlakyOk() (*int32, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *TestCaseCreate) SetIsFlaky(v int32)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *TestCaseCreate) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### GetSuiteId

`func (o *TestCaseCreate) GetSuiteId() int64`

GetSuiteId returns the SuiteId field if non-nil, zero value otherwise.

### GetSuiteIdOk

`func (o *TestCaseCreate) GetSuiteIdOk() (*int64, bool)`

GetSuiteIdOk returns a tuple with the SuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteId

`func (o *TestCaseCreate) SetSuiteId(v int64)`

SetSuiteId sets SuiteId field to given value.

### HasSuiteId

`func (o *TestCaseCreate) HasSuiteId() bool`

HasSuiteId returns a boolean if a field has been set.

### GetMilestoneId

`func (o *TestCaseCreate) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *TestCaseCreate) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *TestCaseCreate) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *TestCaseCreate) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### GetAutomation

`func (o *TestCaseCreate) GetAutomation() int32`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *TestCaseCreate) GetAutomationOk() (*int32, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *TestCaseCreate) SetAutomation(v int32)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *TestCaseCreate) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetIsManual

`func (o *TestCaseCreate) GetIsManual() int32`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *TestCaseCreate) GetIsManualOk() (*int32, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *TestCaseCreate) SetIsManual(v int32)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *TestCaseCreate) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### GetIsToBeAutomated

`func (o *TestCaseCreate) GetIsToBeAutomated() int32`

GetIsToBeAutomated returns the IsToBeAutomated field if non-nil, zero value otherwise.

### GetIsToBeAutomatedOk

`func (o *TestCaseCreate) GetIsToBeAutomatedOk() (*int32, bool)`

GetIsToBeAutomatedOk returns a tuple with the IsToBeAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsToBeAutomated

`func (o *TestCaseCreate) SetIsToBeAutomated(v int32)`

SetIsToBeAutomated sets IsToBeAutomated field to given value.

### HasIsToBeAutomated

`func (o *TestCaseCreate) HasIsToBeAutomated() bool`

HasIsToBeAutomated returns a boolean if a field has been set.

### GetStatus

`func (o *TestCaseCreate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestCaseCreate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestCaseCreate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestCaseCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStepsType

`func (o *TestCaseCreate) GetStepsType() string`

GetStepsType returns the StepsType field if non-nil, zero value otherwise.

### GetStepsTypeOk

`func (o *TestCaseCreate) GetStepsTypeOk() (*string, bool)`

GetStepsTypeOk returns a tuple with the StepsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsType

`func (o *TestCaseCreate) SetStepsType(v string)`

SetStepsType sets StepsType field to given value.

### HasStepsType

`func (o *TestCaseCreate) HasStepsType() bool`

HasStepsType returns a boolean if a field has been set.

### GetAttachments

`func (o *TestCaseCreate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestCaseCreate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestCaseCreate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestCaseCreate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *TestCaseCreate) GetSteps() []TestStepCreate`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TestCaseCreate) GetStepsOk() (*[]TestStepCreate, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TestCaseCreate) SetSteps(v []TestStepCreate)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TestCaseCreate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *TestCaseCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestCaseCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestCaseCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestCaseCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParams

`func (o *TestCaseCreate) GetParams() map[string][]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TestCaseCreate) GetParamsOk() (*map[string][]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TestCaseCreate) SetParams(v map[string][]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *TestCaseCreate) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *TestCaseCreate) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *TestCaseCreate) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetParameters

`func (o *TestCaseCreate) GetParameters() []TestCaseParameterCreate`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestCaseCreate) GetParametersOk() (*[]TestCaseParameterCreate, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestCaseCreate) SetParameters(v []TestCaseParameterCreate)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestCaseCreate) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestCaseCreate) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestCaseCreate) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetCustomField

`func (o *TestCaseCreate) GetCustomField() map[string]string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *TestCaseCreate) GetCustomFieldOk() (*map[string]string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *TestCaseCreate) SetCustomField(v map[string]string)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *TestCaseCreate) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TestCaseCreate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestCaseCreate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestCaseCreate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TestCaseCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TestCaseCreate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TestCaseCreate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TestCaseCreate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TestCaseCreate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


