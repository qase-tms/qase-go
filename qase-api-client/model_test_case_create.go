/*
Qase.io TestOps API v1

Qase TestOps API v1 Specification.

API version: 1.0.0
Contact: support@qase.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v1_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TestCaseCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestCaseCreate{}

// TestCaseCreate struct for TestCaseCreate
type TestCaseCreate struct {
	Description    *string `json:"description,omitempty"`
	Preconditions  *string `json:"preconditions,omitempty"`
	Postconditions *string `json:"postconditions,omitempty"`
	Title          string  `json:"title"`
	Severity       *int32  `json:"severity,omitempty"`
	Priority       *int32  `json:"priority,omitempty"`
	Behavior       *int32  `json:"behavior,omitempty"`
	Type           *int32  `json:"type,omitempty"`
	Layer          *int32  `json:"layer,omitempty"`
	IsFlaky        *int32  `json:"is_flaky,omitempty"`
	SuiteId        *int64  `json:"suite_id,omitempty"`
	MilestoneId    *int64  `json:"milestone_id,omitempty"`
	Automation     *int32  `json:"automation,omitempty"`
	Status         *int32  `json:"status,omitempty"`
	// A list of Attachment hashes.
	Attachments []string            `json:"attachments,omitempty"`
	Steps       []TestStepCreate    `json:"steps,omitempty"`
	Tags        []string            `json:"tags,omitempty"`
	Params      map[string][]string `json:"params,omitempty"`
	// A map of custom fields values (id => value)
	CustomField *map[string]string `json:"custom_field,omitempty"`
	CreatedAt   *string            `json:"created_at,omitempty"`
	UpdatedAt   *string            `json:"updated_at,omitempty"`
}

type _TestCaseCreate TestCaseCreate

// NewTestCaseCreate instantiates a new TestCaseCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestCaseCreate(title string) *TestCaseCreate {
	this := TestCaseCreate{}
	this.Title = title
	return &this
}

// NewTestCaseCreateWithDefaults instantiates a new TestCaseCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestCaseCreateWithDefaults() *TestCaseCreate {
	this := TestCaseCreate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TestCaseCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TestCaseCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TestCaseCreate) SetDescription(v string) {
	o.Description = &v
}

// GetPreconditions returns the Preconditions field value if set, zero value otherwise.
func (o *TestCaseCreate) GetPreconditions() string {
	if o == nil || IsNil(o.Preconditions) {
		var ret string
		return ret
	}
	return *o.Preconditions
}

// GetPreconditionsOk returns a tuple with the Preconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetPreconditionsOk() (*string, bool) {
	if o == nil || IsNil(o.Preconditions) {
		return nil, false
	}
	return o.Preconditions, true
}

// HasPreconditions returns a boolean if a field has been set.
func (o *TestCaseCreate) HasPreconditions() bool {
	if o != nil && !IsNil(o.Preconditions) {
		return true
	}

	return false
}

// SetPreconditions gets a reference to the given string and assigns it to the Preconditions field.
func (o *TestCaseCreate) SetPreconditions(v string) {
	o.Preconditions = &v
}

// GetPostconditions returns the Postconditions field value if set, zero value otherwise.
func (o *TestCaseCreate) GetPostconditions() string {
	if o == nil || IsNil(o.Postconditions) {
		var ret string
		return ret
	}
	return *o.Postconditions
}

// GetPostconditionsOk returns a tuple with the Postconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetPostconditionsOk() (*string, bool) {
	if o == nil || IsNil(o.Postconditions) {
		return nil, false
	}
	return o.Postconditions, true
}

// HasPostconditions returns a boolean if a field has been set.
func (o *TestCaseCreate) HasPostconditions() bool {
	if o != nil && !IsNil(o.Postconditions) {
		return true
	}

	return false
}

// SetPostconditions gets a reference to the given string and assigns it to the Postconditions field.
func (o *TestCaseCreate) SetPostconditions(v string) {
	o.Postconditions = &v
}

// GetTitle returns the Title field value
func (o *TestCaseCreate) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TestCaseCreate) SetTitle(v string) {
	o.Title = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *TestCaseCreate) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *TestCaseCreate) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *TestCaseCreate) SetSeverity(v int32) {
	o.Severity = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TestCaseCreate) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TestCaseCreate) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *TestCaseCreate) SetPriority(v int32) {
	o.Priority = &v
}

// GetBehavior returns the Behavior field value if set, zero value otherwise.
func (o *TestCaseCreate) GetBehavior() int32 {
	if o == nil || IsNil(o.Behavior) {
		var ret int32
		return ret
	}
	return *o.Behavior
}

// GetBehaviorOk returns a tuple with the Behavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetBehaviorOk() (*int32, bool) {
	if o == nil || IsNil(o.Behavior) {
		return nil, false
	}
	return o.Behavior, true
}

// HasBehavior returns a boolean if a field has been set.
func (o *TestCaseCreate) HasBehavior() bool {
	if o != nil && !IsNil(o.Behavior) {
		return true
	}

	return false
}

// SetBehavior gets a reference to the given int32 and assigns it to the Behavior field.
func (o *TestCaseCreate) SetBehavior(v int32) {
	o.Behavior = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TestCaseCreate) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TestCaseCreate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *TestCaseCreate) SetType(v int32) {
	o.Type = &v
}

// GetLayer returns the Layer field value if set, zero value otherwise.
func (o *TestCaseCreate) GetLayer() int32 {
	if o == nil || IsNil(o.Layer) {
		var ret int32
		return ret
	}
	return *o.Layer
}

// GetLayerOk returns a tuple with the Layer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetLayerOk() (*int32, bool) {
	if o == nil || IsNil(o.Layer) {
		return nil, false
	}
	return o.Layer, true
}

// HasLayer returns a boolean if a field has been set.
func (o *TestCaseCreate) HasLayer() bool {
	if o != nil && !IsNil(o.Layer) {
		return true
	}

	return false
}

// SetLayer gets a reference to the given int32 and assigns it to the Layer field.
func (o *TestCaseCreate) SetLayer(v int32) {
	o.Layer = &v
}

// GetIsFlaky returns the IsFlaky field value if set, zero value otherwise.
func (o *TestCaseCreate) GetIsFlaky() int32 {
	if o == nil || IsNil(o.IsFlaky) {
		var ret int32
		return ret
	}
	return *o.IsFlaky
}

// GetIsFlakyOk returns a tuple with the IsFlaky field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetIsFlakyOk() (*int32, bool) {
	if o == nil || IsNil(o.IsFlaky) {
		return nil, false
	}
	return o.IsFlaky, true
}

// HasIsFlaky returns a boolean if a field has been set.
func (o *TestCaseCreate) HasIsFlaky() bool {
	if o != nil && !IsNil(o.IsFlaky) {
		return true
	}

	return false
}

// SetIsFlaky gets a reference to the given int32 and assigns it to the IsFlaky field.
func (o *TestCaseCreate) SetIsFlaky(v int32) {
	o.IsFlaky = &v
}

// GetSuiteId returns the SuiteId field value if set, zero value otherwise.
func (o *TestCaseCreate) GetSuiteId() int64 {
	if o == nil || IsNil(o.SuiteId) {
		var ret int64
		return ret
	}
	return *o.SuiteId
}

// GetSuiteIdOk returns a tuple with the SuiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetSuiteIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SuiteId) {
		return nil, false
	}
	return o.SuiteId, true
}

// HasSuiteId returns a boolean if a field has been set.
func (o *TestCaseCreate) HasSuiteId() bool {
	if o != nil && !IsNil(o.SuiteId) {
		return true
	}

	return false
}

// SetSuiteId gets a reference to the given int64 and assigns it to the SuiteId field.
func (o *TestCaseCreate) SetSuiteId(v int64) {
	o.SuiteId = &v
}

// GetMilestoneId returns the MilestoneId field value if set, zero value otherwise.
func (o *TestCaseCreate) GetMilestoneId() int64 {
	if o == nil || IsNil(o.MilestoneId) {
		var ret int64
		return ret
	}
	return *o.MilestoneId
}

// GetMilestoneIdOk returns a tuple with the MilestoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetMilestoneIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MilestoneId) {
		return nil, false
	}
	return o.MilestoneId, true
}

// HasMilestoneId returns a boolean if a field has been set.
func (o *TestCaseCreate) HasMilestoneId() bool {
	if o != nil && !IsNil(o.MilestoneId) {
		return true
	}

	return false
}

// SetMilestoneId gets a reference to the given int64 and assigns it to the MilestoneId field.
func (o *TestCaseCreate) SetMilestoneId(v int64) {
	o.MilestoneId = &v
}

// GetAutomation returns the Automation field value if set, zero value otherwise.
func (o *TestCaseCreate) GetAutomation() int32 {
	if o == nil || IsNil(o.Automation) {
		var ret int32
		return ret
	}
	return *o.Automation
}

// GetAutomationOk returns a tuple with the Automation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetAutomationOk() (*int32, bool) {
	if o == nil || IsNil(o.Automation) {
		return nil, false
	}
	return o.Automation, true
}

// HasAutomation returns a boolean if a field has been set.
func (o *TestCaseCreate) HasAutomation() bool {
	if o != nil && !IsNil(o.Automation) {
		return true
	}

	return false
}

// SetAutomation gets a reference to the given int32 and assigns it to the Automation field.
func (o *TestCaseCreate) SetAutomation(v int32) {
	o.Automation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TestCaseCreate) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TestCaseCreate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *TestCaseCreate) SetStatus(v int32) {
	o.Status = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TestCaseCreate) GetAttachments() []string {
	if o == nil || IsNil(o.Attachments) {
		var ret []string
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetAttachmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestCaseCreate) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *TestCaseCreate) SetAttachments(v []string) {
	o.Attachments = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *TestCaseCreate) GetSteps() []TestStepCreate {
	if o == nil || IsNil(o.Steps) {
		var ret []TestStepCreate
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetStepsOk() ([]TestStepCreate, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *TestCaseCreate) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []TestStepCreate and assigns it to the Steps field.
func (o *TestCaseCreate) SetSteps(v []TestStepCreate) {
	o.Steps = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TestCaseCreate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TestCaseCreate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TestCaseCreate) SetTags(v []string) {
	o.Tags = v
}

// GetParams returns the Params field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCaseCreate) GetParams() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCaseCreate) GetParamsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *TestCaseCreate) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string][]string and assigns it to the Params field.
func (o *TestCaseCreate) SetParams(v map[string][]string) {
	o.Params = v
}

// GetCustomField returns the CustomField field value if set, zero value otherwise.
func (o *TestCaseCreate) GetCustomField() map[string]string {
	if o == nil || IsNil(o.CustomField) {
		var ret map[string]string
		return ret
	}
	return *o.CustomField
}

// GetCustomFieldOk returns a tuple with the CustomField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetCustomFieldOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomField) {
		return nil, false
	}
	return o.CustomField, true
}

// HasCustomField returns a boolean if a field has been set.
func (o *TestCaseCreate) HasCustomField() bool {
	if o != nil && !IsNil(o.CustomField) {
		return true
	}

	return false
}

// SetCustomField gets a reference to the given map[string]string and assigns it to the CustomField field.
func (o *TestCaseCreate) SetCustomField(v map[string]string) {
	o.CustomField = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TestCaseCreate) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TestCaseCreate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TestCaseCreate) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TestCaseCreate) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseCreate) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TestCaseCreate) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TestCaseCreate) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o TestCaseCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestCaseCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Preconditions) {
		toSerialize["preconditions"] = o.Preconditions
	}
	if !IsNil(o.Postconditions) {
		toSerialize["postconditions"] = o.Postconditions
	}
	toSerialize["title"] = o.Title
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Behavior) {
		toSerialize["behavior"] = o.Behavior
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Layer) {
		toSerialize["layer"] = o.Layer
	}
	if !IsNil(o.IsFlaky) {
		toSerialize["is_flaky"] = o.IsFlaky
	}
	if !IsNil(o.SuiteId) {
		toSerialize["suite_id"] = o.SuiteId
	}
	if !IsNil(o.MilestoneId) {
		toSerialize["milestone_id"] = o.MilestoneId
	}
	if !IsNil(o.Automation) {
		toSerialize["automation"] = o.Automation
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.CustomField) {
		toSerialize["custom_field"] = o.CustomField
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *TestCaseCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTestCaseCreate := _TestCaseCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestCaseCreate)

	if err != nil {
		return err
	}

	*o = TestCaseCreate(varTestCaseCreate)

	return err
}

type NullableTestCaseCreate struct {
	value *TestCaseCreate
	isSet bool
}

func (v NullableTestCaseCreate) Get() *TestCaseCreate {
	return v.value
}

func (v *NullableTestCaseCreate) Set(val *TestCaseCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTestCaseCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTestCaseCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestCaseCreate(val *TestCaseCreate) *NullableTestCaseCreate {
	return &NullableTestCaseCreate{value: val, isSet: true}
}

func (v NullableTestCaseCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestCaseCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
