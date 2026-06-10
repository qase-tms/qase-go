# TestCaseUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Preconditions** | Pointer to **string** |  | [optional] 
**Postconditions** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
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
**CustomField** | Pointer to **map[string]string** | Custom field values keyed by the field&#39;s project-scoped &#x60;internal_id&#x60; (see &#x60;GET /custom_field&#x60;). Values are always **scalar strings**; arrays, objects or non-scalars are rejected.  | Field type           | Value format                              | Example                 | |----------------------|-------------------------------------------|-------------------------| | &#x60;string&#x60;, &#x60;text&#x60;     | Plain string                              | &#x60;\&quot;hello\&quot;&#x60;               | | &#x60;number&#x60;             | Numeric string                            | &#x60;\&quot;42\&quot;&#x60;                  | | &#x60;url&#x60;                | Valid URL                                 | &#x60;\&quot;https://qase.io\&quot;&#x60;     | | &#x60;datetime&#x60;           | Absolute date (ISO 8601 recommended)      | &#x60;\&quot;2026-04-29T15:00:00Z\&quot;&#x60;| | &#x60;selectbox&#x60;, &#x60;radio&#x60; | Option &#x60;id&#x60; as string                     | &#x60;\&quot;1\&quot;&#x60;                   | | &#x60;multiselect&#x60;        | Comma-separated option &#x60;id&#x60;s (no spaces)  | &#x60;\&quot;1,2,3\&quot;&#x60;               | | &#x60;checkbox&#x60;           | &#x60;\&quot;1\&quot;&#x60; to check, &#x60;\&quot;\&quot;&#x60; to uncheck           | &#x60;\&quot;1\&quot;&#x60;                   | | &#x60;user&#x60;               | Team member &#x60;internal_id&#x60; as string       | &#x60;\&quot;42\&quot;&#x60;                  |  Partial update: only fields present in the payload are validated; required fields not included are not enforced. Send &#x60;\&quot;\&quot;&#x60; to clear a value. Unknown &#x60;internal_id&#x60;s are rejected; option-based values must reference an existing option.  Note: a &#x60;required&#x60; checkbox without a default cannot be unchecked via the API — set a default or clear &#x60;required&#x60; in workspace settings.  | [optional] 

## Methods

### NewTestCaseUpdate

`func NewTestCaseUpdate() *TestCaseUpdate`

NewTestCaseUpdate instantiates a new TestCaseUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCaseUpdateWithDefaults

`func NewTestCaseUpdateWithDefaults() *TestCaseUpdate`

NewTestCaseUpdateWithDefaults instantiates a new TestCaseUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TestCaseUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestCaseUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestCaseUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestCaseUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPreconditions

`func (o *TestCaseUpdate) GetPreconditions() string`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *TestCaseUpdate) GetPreconditionsOk() (*string, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *TestCaseUpdate) SetPreconditions(v string)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *TestCaseUpdate) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### GetPostconditions

`func (o *TestCaseUpdate) GetPostconditions() string`

GetPostconditions returns the Postconditions field if non-nil, zero value otherwise.

### GetPostconditionsOk

`func (o *TestCaseUpdate) GetPostconditionsOk() (*string, bool)`

GetPostconditionsOk returns a tuple with the Postconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditions

`func (o *TestCaseUpdate) SetPostconditions(v string)`

SetPostconditions sets Postconditions field to given value.

### HasPostconditions

`func (o *TestCaseUpdate) HasPostconditions() bool`

HasPostconditions returns a boolean if a field has been set.

### GetTitle

`func (o *TestCaseUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestCaseUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestCaseUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TestCaseUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSeverity

`func (o *TestCaseUpdate) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TestCaseUpdate) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TestCaseUpdate) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *TestCaseUpdate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetPriority

`func (o *TestCaseUpdate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestCaseUpdate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestCaseUpdate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TestCaseUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetBehavior

`func (o *TestCaseUpdate) GetBehavior() int32`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *TestCaseUpdate) GetBehaviorOk() (*int32, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *TestCaseUpdate) SetBehavior(v int32)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *TestCaseUpdate) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetType

`func (o *TestCaseUpdate) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestCaseUpdate) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestCaseUpdate) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *TestCaseUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayer

`func (o *TestCaseUpdate) GetLayer() int32`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *TestCaseUpdate) GetLayerOk() (*int32, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *TestCaseUpdate) SetLayer(v int32)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *TestCaseUpdate) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetIsFlaky

`func (o *TestCaseUpdate) GetIsFlaky() int32`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *TestCaseUpdate) GetIsFlakyOk() (*int32, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *TestCaseUpdate) SetIsFlaky(v int32)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *TestCaseUpdate) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### GetSuiteId

`func (o *TestCaseUpdate) GetSuiteId() int64`

GetSuiteId returns the SuiteId field if non-nil, zero value otherwise.

### GetSuiteIdOk

`func (o *TestCaseUpdate) GetSuiteIdOk() (*int64, bool)`

GetSuiteIdOk returns a tuple with the SuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteId

`func (o *TestCaseUpdate) SetSuiteId(v int64)`

SetSuiteId sets SuiteId field to given value.

### HasSuiteId

`func (o *TestCaseUpdate) HasSuiteId() bool`

HasSuiteId returns a boolean if a field has been set.

### GetMilestoneId

`func (o *TestCaseUpdate) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *TestCaseUpdate) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *TestCaseUpdate) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *TestCaseUpdate) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### GetAutomation

`func (o *TestCaseUpdate) GetAutomation() int32`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *TestCaseUpdate) GetAutomationOk() (*int32, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *TestCaseUpdate) SetAutomation(v int32)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *TestCaseUpdate) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetIsManual

`func (o *TestCaseUpdate) GetIsManual() int32`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *TestCaseUpdate) GetIsManualOk() (*int32, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *TestCaseUpdate) SetIsManual(v int32)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *TestCaseUpdate) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### GetIsToBeAutomated

`func (o *TestCaseUpdate) GetIsToBeAutomated() int32`

GetIsToBeAutomated returns the IsToBeAutomated field if non-nil, zero value otherwise.

### GetIsToBeAutomatedOk

`func (o *TestCaseUpdate) GetIsToBeAutomatedOk() (*int32, bool)`

GetIsToBeAutomatedOk returns a tuple with the IsToBeAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsToBeAutomated

`func (o *TestCaseUpdate) SetIsToBeAutomated(v int32)`

SetIsToBeAutomated sets IsToBeAutomated field to given value.

### HasIsToBeAutomated

`func (o *TestCaseUpdate) HasIsToBeAutomated() bool`

HasIsToBeAutomated returns a boolean if a field has been set.

### GetStatus

`func (o *TestCaseUpdate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestCaseUpdate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestCaseUpdate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestCaseUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStepsType

`func (o *TestCaseUpdate) GetStepsType() string`

GetStepsType returns the StepsType field if non-nil, zero value otherwise.

### GetStepsTypeOk

`func (o *TestCaseUpdate) GetStepsTypeOk() (*string, bool)`

GetStepsTypeOk returns a tuple with the StepsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsType

`func (o *TestCaseUpdate) SetStepsType(v string)`

SetStepsType sets StepsType field to given value.

### HasStepsType

`func (o *TestCaseUpdate) HasStepsType() bool`

HasStepsType returns a boolean if a field has been set.

### GetAttachments

`func (o *TestCaseUpdate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestCaseUpdate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestCaseUpdate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestCaseUpdate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *TestCaseUpdate) GetSteps() []TestStepCreate`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TestCaseUpdate) GetStepsOk() (*[]TestStepCreate, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TestCaseUpdate) SetSteps(v []TestStepCreate)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TestCaseUpdate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *TestCaseUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestCaseUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestCaseUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestCaseUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetParams

`func (o *TestCaseUpdate) GetParams() map[string][]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TestCaseUpdate) GetParamsOk() (*map[string][]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TestCaseUpdate) SetParams(v map[string][]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *TestCaseUpdate) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *TestCaseUpdate) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *TestCaseUpdate) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetParameters

`func (o *TestCaseUpdate) GetParameters() []TestCaseParameterCreate`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestCaseUpdate) GetParametersOk() (*[]TestCaseParameterCreate, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestCaseUpdate) SetParameters(v []TestCaseParameterCreate)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestCaseUpdate) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestCaseUpdate) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestCaseUpdate) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetCustomField

`func (o *TestCaseUpdate) GetCustomField() map[string]string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *TestCaseUpdate) GetCustomFieldOk() (*map[string]string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *TestCaseUpdate) SetCustomField(v map[string]string)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *TestCaseUpdate) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


