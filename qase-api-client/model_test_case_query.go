/*
Qase.io TestOps API v1

Qase TestOps API v1 Specification.

API version: 1.0.0
Contact: support@qase.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v1_client

import (
	"encoding/json"
	"time"
)

// checks if the TestCaseQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestCaseQuery{}

// TestCaseQuery struct for TestCaseQuery
type TestCaseQuery struct {
	Id             *int64             `json:"id,omitempty"`
	Position       *int32             `json:"position,omitempty"`
	Title          *string            `json:"title,omitempty"`
	Description    NullableString     `json:"description,omitempty"`
	Preconditions  NullableString     `json:"preconditions,omitempty"`
	Postconditions NullableString     `json:"postconditions,omitempty"`
	Severity       *int32             `json:"severity,omitempty"`
	Priority       *int32             `json:"priority,omitempty"`
	Type           *int32             `json:"type,omitempty"`
	Layer          *int32             `json:"layer,omitempty"`
	IsFlaky        *int32             `json:"is_flaky,omitempty"`
	Behavior       *int32             `json:"behavior,omitempty"`
	Automation     *int32             `json:"automation,omitempty"`
	Status         *int32             `json:"status,omitempty"`
	MilestoneId    NullableInt64      `json:"milestone_id,omitempty"`
	SuiteId        NullableInt64      `json:"suite_id,omitempty"`
	CustomFields   []CustomFieldValue `json:"custom_fields,omitempty"`
	Attachments    []Attachment       `json:"attachments,omitempty"`
	StepsType      NullableString     `json:"steps_type,omitempty"`
	Steps          []TestStep         `json:"steps,omitempty"`
	Params         *TestCaseParams    `json:"params,omitempty"`
	Tags           []TagValue         `json:"tags,omitempty"`
	// Deprecated, use `author_id` instead.
	// Deprecated
	MemberId  *int64     `json:"member_id,omitempty"`
	AuthorId  *int64     `json:"author_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewTestCaseQuery instantiates a new TestCaseQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestCaseQuery() *TestCaseQuery {
	this := TestCaseQuery{}
	return &this
}

// NewTestCaseQueryWithDefaults instantiates a new TestCaseQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestCaseQueryWithDefaults() *TestCaseQuery {
	this := TestCaseQuery{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestCaseQuery) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestCaseQuery) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TestCaseQuery) SetId(v int64) {
	o.Id = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *TestCaseQuery) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *TestCaseQuery) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *TestCaseQuery) SetPosition(v int32) {
	o.Position = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TestCaseQuery) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TestCaseQuery) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TestCaseQuery) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCaseQuery) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCaseQuery) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestCaseQuery) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestCaseQuery) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestCaseQuery) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestCaseQuery) UnsetDescription() {
	o.Description.Unset()
}

// GetPreconditions returns the Preconditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCaseQuery) GetPreconditions() string {
	if o == nil || IsNil(o.Preconditions.Get()) {
		var ret string
		return ret
	}
	return *o.Preconditions.Get()
}

// GetPreconditionsOk returns a tuple with the Preconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCaseQuery) GetPreconditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preconditions.Get(), o.Preconditions.IsSet()
}

// HasPreconditions returns a boolean if a field has been set.
func (o *TestCaseQuery) HasPreconditions() bool {
	if o != nil && o.Preconditions.IsSet() {
		return true
	}

	return false
}

// SetPreconditions gets a reference to the given NullableString and assigns it to the Preconditions field.
func (o *TestCaseQuery) SetPreconditions(v string) {
	o.Preconditions.Set(&v)
}

// SetPreconditionsNil sets the value for Preconditions to be an explicit nil
func (o *TestCaseQuery) SetPreconditionsNil() {
	o.Preconditions.Set(nil)
}

// UnsetPreconditions ensures that no value is present for Preconditions, not even an explicit nil
func (o *TestCaseQuery) UnsetPreconditions() {
	o.Preconditions.Unset()
}

// GetPostconditions returns the Postconditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCaseQuery) GetPostconditions() string {
	if o == nil || IsNil(o.Postconditions.Get()) {
		var ret string
		return ret
	}
	return *o.Postconditions.Get()
}

// GetPostconditionsOk returns a tuple with the Postconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCaseQuery) GetPostconditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postconditions.Get(), o.Postconditions.IsSet()
}

// HasPostconditions returns a boolean if a field has been set.
func (o *TestCaseQuery) HasPostconditions() bool {
	if o != nil && o.Postconditions.IsSet() {
		return true
	}

	return false
}

// SetPostconditions gets a reference to the given NullableString and assigns it to the Postconditions field.
func (o *TestCaseQuery) SetPostconditions(v string) {
	o.Postconditions.Set(&v)
}

// SetPostconditionsNil sets the value for Postconditions to be an explicit nil
func (o *TestCaseQuery) SetPostconditionsNil() {
	o.Postconditions.Set(nil)
}

// UnsetPostconditions ensures that no value is present for Postconditions, not even an explicit nil
func (o *TestCaseQuery) UnsetPostconditions() {
	o.Postconditions.Unset()
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *TestCaseQuery) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *TestCaseQuery) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *TestCaseQuery) SetSeverity(v int32) {
	o.Severity = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TestCaseQuery) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TestCaseQuery) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *TestCaseQuery) SetPriority(v int32) {
	o.Priority = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TestCaseQuery) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TestCaseQuery) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *TestCaseQuery) SetType(v int32) {
	o.Type = &v
}

// GetLayer returns the Layer field value if set, zero value otherwise.
func (o *TestCaseQuery) GetLayer() int32 {
	if o == nil || IsNil(o.Layer) {
		var ret int32
		return ret
	}
	return *o.Layer
}

// GetLayerOk returns a tuple with the Layer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetLayerOk() (*int32, bool) {
	if o == nil || IsNil(o.Layer) {
		return nil, false
	}
	return o.Layer, true
}

// HasLayer returns a boolean if a field has been set.
func (o *TestCaseQuery) HasLayer() bool {
	if o != nil && !IsNil(o.Layer) {
		return true
	}

	return false
}

// SetLayer gets a reference to the given int32 and assigns it to the Layer field.
func (o *TestCaseQuery) SetLayer(v int32) {
	o.Layer = &v
}

// GetIsFlaky returns the IsFlaky field value if set, zero value otherwise.
func (o *TestCaseQuery) GetIsFlaky() int32 {
	if o == nil || IsNil(o.IsFlaky) {
		var ret int32
		return ret
	}
	return *o.IsFlaky
}

// GetIsFlakyOk returns a tuple with the IsFlaky field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetIsFlakyOk() (*int32, bool) {
	if o == nil || IsNil(o.IsFlaky) {
		return nil, false
	}
	return o.IsFlaky, true
}

// HasIsFlaky returns a boolean if a field has been set.
func (o *TestCaseQuery) HasIsFlaky() bool {
	if o != nil && !IsNil(o.IsFlaky) {
		return true
	}

	return false
}

// SetIsFlaky gets a reference to the given int32 and assigns it to the IsFlaky field.
func (o *TestCaseQuery) SetIsFlaky(v int32) {
	o.IsFlaky = &v
}

// GetBehavior returns the Behavior field value if set, zero value otherwise.
func (o *TestCaseQuery) GetBehavior() int32 {
	if o == nil || IsNil(o.Behavior) {
		var ret int32
		return ret
	}
	return *o.Behavior
}

// GetBehaviorOk returns a tuple with the Behavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetBehaviorOk() (*int32, bool) {
	if o == nil || IsNil(o.Behavior) {
		return nil, false
	}
	return o.Behavior, true
}

// HasBehavior returns a boolean if a field has been set.
func (o *TestCaseQuery) HasBehavior() bool {
	if o != nil && !IsNil(o.Behavior) {
		return true
	}

	return false
}

// SetBehavior gets a reference to the given int32 and assigns it to the Behavior field.
func (o *TestCaseQuery) SetBehavior(v int32) {
	o.Behavior = &v
}

// GetAutomation returns the Automation field value if set, zero value otherwise.
func (o *TestCaseQuery) GetAutomation() int32 {
	if o == nil || IsNil(o.Automation) {
		var ret int32
		return ret
	}
	return *o.Automation
}

// GetAutomationOk returns a tuple with the Automation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetAutomationOk() (*int32, bool) {
	if o == nil || IsNil(o.Automation) {
		return nil, false
	}
	return o.Automation, true
}

// HasAutomation returns a boolean if a field has been set.
func (o *TestCaseQuery) HasAutomation() bool {
	if o != nil && !IsNil(o.Automation) {
		return true
	}

	return false
}

// SetAutomation gets a reference to the given int32 and assigns it to the Automation field.
func (o *TestCaseQuery) SetAutomation(v int32) {
	o.Automation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TestCaseQuery) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TestCaseQuery) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *TestCaseQuery) SetStatus(v int32) {
	o.Status = &v
}

// GetMilestoneId returns the MilestoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCaseQuery) GetMilestoneId() int64 {
	if o == nil || IsNil(o.MilestoneId.Get()) {
		var ret int64
		return ret
	}
	return *o.MilestoneId.Get()
}

// GetMilestoneIdOk returns a tuple with the MilestoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCaseQuery) GetMilestoneIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MilestoneId.Get(), o.MilestoneId.IsSet()
}

// HasMilestoneId returns a boolean if a field has been set.
func (o *TestCaseQuery) HasMilestoneId() bool {
	if o != nil && o.MilestoneId.IsSet() {
		return true
	}

	return false
}

// SetMilestoneId gets a reference to the given NullableInt64 and assigns it to the MilestoneId field.
func (o *TestCaseQuery) SetMilestoneId(v int64) {
	o.MilestoneId.Set(&v)
}

// SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil
func (o *TestCaseQuery) SetMilestoneIdNil() {
	o.MilestoneId.Set(nil)
}

// UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
func (o *TestCaseQuery) UnsetMilestoneId() {
	o.MilestoneId.Unset()
}

// GetSuiteId returns the SuiteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCaseQuery) GetSuiteId() int64 {
	if o == nil || IsNil(o.SuiteId.Get()) {
		var ret int64
		return ret
	}
	return *o.SuiteId.Get()
}

// GetSuiteIdOk returns a tuple with the SuiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCaseQuery) GetSuiteIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuiteId.Get(), o.SuiteId.IsSet()
}

// HasSuiteId returns a boolean if a field has been set.
func (o *TestCaseQuery) HasSuiteId() bool {
	if o != nil && o.SuiteId.IsSet() {
		return true
	}

	return false
}

// SetSuiteId gets a reference to the given NullableInt64 and assigns it to the SuiteId field.
func (o *TestCaseQuery) SetSuiteId(v int64) {
	o.SuiteId.Set(&v)
}

// SetSuiteIdNil sets the value for SuiteId to be an explicit nil
func (o *TestCaseQuery) SetSuiteIdNil() {
	o.SuiteId.Set(nil)
}

// UnsetSuiteId ensures that no value is present for SuiteId, not even an explicit nil
func (o *TestCaseQuery) UnsetSuiteId() {
	o.SuiteId.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *TestCaseQuery) GetCustomFields() []CustomFieldValue {
	if o == nil || IsNil(o.CustomFields) {
		var ret []CustomFieldValue
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetCustomFieldsOk() ([]CustomFieldValue, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TestCaseQuery) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomFieldValue and assigns it to the CustomFields field.
func (o *TestCaseQuery) SetCustomFields(v []CustomFieldValue) {
	o.CustomFields = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TestCaseQuery) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestCaseQuery) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *TestCaseQuery) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetStepsType returns the StepsType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCaseQuery) GetStepsType() string {
	if o == nil || IsNil(o.StepsType.Get()) {
		var ret string
		return ret
	}
	return *o.StepsType.Get()
}

// GetStepsTypeOk returns a tuple with the StepsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCaseQuery) GetStepsTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StepsType.Get(), o.StepsType.IsSet()
}

// HasStepsType returns a boolean if a field has been set.
func (o *TestCaseQuery) HasStepsType() bool {
	if o != nil && o.StepsType.IsSet() {
		return true
	}

	return false
}

// SetStepsType gets a reference to the given NullableString and assigns it to the StepsType field.
func (o *TestCaseQuery) SetStepsType(v string) {
	o.StepsType.Set(&v)
}

// SetStepsTypeNil sets the value for StepsType to be an explicit nil
func (o *TestCaseQuery) SetStepsTypeNil() {
	o.StepsType.Set(nil)
}

// UnsetStepsType ensures that no value is present for StepsType, not even an explicit nil
func (o *TestCaseQuery) UnsetStepsType() {
	o.StepsType.Unset()
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *TestCaseQuery) GetSteps() []TestStep {
	if o == nil || IsNil(o.Steps) {
		var ret []TestStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetStepsOk() ([]TestStep, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *TestCaseQuery) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []TestStep and assigns it to the Steps field.
func (o *TestCaseQuery) SetSteps(v []TestStep) {
	o.Steps = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *TestCaseQuery) GetParams() TestCaseParams {
	if o == nil || IsNil(o.Params) {
		var ret TestCaseParams
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetParamsOk() (*TestCaseParams, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *TestCaseQuery) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given TestCaseParams and assigns it to the Params field.
func (o *TestCaseQuery) SetParams(v TestCaseParams) {
	o.Params = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TestCaseQuery) GetTags() []TagValue {
	if o == nil || IsNil(o.Tags) {
		var ret []TagValue
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetTagsOk() ([]TagValue, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TestCaseQuery) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagValue and assigns it to the Tags field.
func (o *TestCaseQuery) SetTags(v []TagValue) {
	o.Tags = v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
// Deprecated
func (o *TestCaseQuery) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TestCaseQuery) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *TestCaseQuery) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
// Deprecated
func (o *TestCaseQuery) SetMemberId(v int64) {
	o.MemberId = &v
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *TestCaseQuery) GetAuthorId() int64 {
	if o == nil || IsNil(o.AuthorId) {
		var ret int64
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetAuthorIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthorId) {
		return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *TestCaseQuery) HasAuthorId() bool {
	if o != nil && !IsNil(o.AuthorId) {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given int64 and assigns it to the AuthorId field.
func (o *TestCaseQuery) SetAuthorId(v int64) {
	o.AuthorId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TestCaseQuery) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TestCaseQuery) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TestCaseQuery) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TestCaseQuery) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseQuery) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TestCaseQuery) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TestCaseQuery) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o TestCaseQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestCaseQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Preconditions.IsSet() {
		toSerialize["preconditions"] = o.Preconditions.Get()
	}
	if o.Postconditions.IsSet() {
		toSerialize["postconditions"] = o.Postconditions.Get()
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
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
	if !IsNil(o.Behavior) {
		toSerialize["behavior"] = o.Behavior
	}
	if !IsNil(o.Automation) {
		toSerialize["automation"] = o.Automation
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.MilestoneId.IsSet() {
		toSerialize["milestone_id"] = o.MilestoneId.Get()
	}
	if o.SuiteId.IsSet() {
		toSerialize["suite_id"] = o.SuiteId.Get()
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if o.StepsType.IsSet() {
		toSerialize["steps_type"] = o.StepsType.Get()
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.MemberId) {
		toSerialize["member_id"] = o.MemberId
	}
	if !IsNil(o.AuthorId) {
		toSerialize["author_id"] = o.AuthorId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableTestCaseQuery struct {
	value *TestCaseQuery
	isSet bool
}

func (v NullableTestCaseQuery) Get() *TestCaseQuery {
	return v.value
}

func (v *NullableTestCaseQuery) Set(val *TestCaseQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableTestCaseQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableTestCaseQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestCaseQuery(val *TestCaseQuery) *NullableTestCaseQuery {
	return &NullableTestCaseQuery{value: val, isSet: true}
}

func (v NullableTestCaseQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestCaseQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}