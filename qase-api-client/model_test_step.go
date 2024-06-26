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
)

// checks if the TestStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestStep{}

// TestStep struct for TestStep
type TestStep struct {
	// Deprecated
	Hash                 *string        `json:"hash,omitempty"`
	SharedStepHash       NullableString `json:"shared_step_hash,omitempty"`
	SharedStepNestedHash NullableString `json:"shared_step_nested_hash,omitempty"`
	// Deprecated
	Position       *int32         `json:"position,omitempty"`
	Action         *string        `json:"action,omitempty"`
	ExpectedResult NullableString `json:"expected_result,omitempty"`
	Data           NullableString `json:"data,omitempty"`
	Attachments    []Attachment   `json:"attachments,omitempty"`
	// Nested steps will be here. The same structure is used for them.
	Steps []map[string]interface{} `json:"steps,omitempty"`
}

// NewTestStep instantiates a new TestStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestStep() *TestStep {
	this := TestStep{}
	return &this
}

// NewTestStepWithDefaults instantiates a new TestStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestStepWithDefaults() *TestStep {
	this := TestStep{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
// Deprecated
func (o *TestStep) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TestStep) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *TestStep) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
// Deprecated
func (o *TestStep) SetHash(v string) {
	o.Hash = &v
}

// GetSharedStepHash returns the SharedStepHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStep) GetSharedStepHash() string {
	if o == nil || IsNil(o.SharedStepHash.Get()) {
		var ret string
		return ret
	}
	return *o.SharedStepHash.Get()
}

// GetSharedStepHashOk returns a tuple with the SharedStepHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStep) GetSharedStepHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedStepHash.Get(), o.SharedStepHash.IsSet()
}

// HasSharedStepHash returns a boolean if a field has been set.
func (o *TestStep) HasSharedStepHash() bool {
	if o != nil && o.SharedStepHash.IsSet() {
		return true
	}

	return false
}

// SetSharedStepHash gets a reference to the given NullableString and assigns it to the SharedStepHash field.
func (o *TestStep) SetSharedStepHash(v string) {
	o.SharedStepHash.Set(&v)
}

// SetSharedStepHashNil sets the value for SharedStepHash to be an explicit nil
func (o *TestStep) SetSharedStepHashNil() {
	o.SharedStepHash.Set(nil)
}

// UnsetSharedStepHash ensures that no value is present for SharedStepHash, not even an explicit nil
func (o *TestStep) UnsetSharedStepHash() {
	o.SharedStepHash.Unset()
}

// GetSharedStepNestedHash returns the SharedStepNestedHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStep) GetSharedStepNestedHash() string {
	if o == nil || IsNil(o.SharedStepNestedHash.Get()) {
		var ret string
		return ret
	}
	return *o.SharedStepNestedHash.Get()
}

// GetSharedStepNestedHashOk returns a tuple with the SharedStepNestedHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStep) GetSharedStepNestedHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedStepNestedHash.Get(), o.SharedStepNestedHash.IsSet()
}

// HasSharedStepNestedHash returns a boolean if a field has been set.
func (o *TestStep) HasSharedStepNestedHash() bool {
	if o != nil && o.SharedStepNestedHash.IsSet() {
		return true
	}

	return false
}

// SetSharedStepNestedHash gets a reference to the given NullableString and assigns it to the SharedStepNestedHash field.
func (o *TestStep) SetSharedStepNestedHash(v string) {
	o.SharedStepNestedHash.Set(&v)
}

// SetSharedStepNestedHashNil sets the value for SharedStepNestedHash to be an explicit nil
func (o *TestStep) SetSharedStepNestedHashNil() {
	o.SharedStepNestedHash.Set(nil)
}

// UnsetSharedStepNestedHash ensures that no value is present for SharedStepNestedHash, not even an explicit nil
func (o *TestStep) UnsetSharedStepNestedHash() {
	o.SharedStepNestedHash.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise.
// Deprecated
func (o *TestStep) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TestStep) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *TestStep) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
// Deprecated
func (o *TestStep) SetPosition(v int32) {
	o.Position = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TestStep) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStep) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TestStep) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *TestStep) SetAction(v string) {
	o.Action = &v
}

// GetExpectedResult returns the ExpectedResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStep) GetExpectedResult() string {
	if o == nil || IsNil(o.ExpectedResult.Get()) {
		var ret string
		return ret
	}
	return *o.ExpectedResult.Get()
}

// GetExpectedResultOk returns a tuple with the ExpectedResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStep) GetExpectedResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedResult.Get(), o.ExpectedResult.IsSet()
}

// HasExpectedResult returns a boolean if a field has been set.
func (o *TestStep) HasExpectedResult() bool {
	if o != nil && o.ExpectedResult.IsSet() {
		return true
	}

	return false
}

// SetExpectedResult gets a reference to the given NullableString and assigns it to the ExpectedResult field.
func (o *TestStep) SetExpectedResult(v string) {
	o.ExpectedResult.Set(&v)
}

// SetExpectedResultNil sets the value for ExpectedResult to be an explicit nil
func (o *TestStep) SetExpectedResultNil() {
	o.ExpectedResult.Set(nil)
}

// UnsetExpectedResult ensures that no value is present for ExpectedResult, not even an explicit nil
func (o *TestStep) UnsetExpectedResult() {
	o.ExpectedResult.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStep) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStep) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *TestStep) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *TestStep) SetData(v string) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil
func (o *TestStep) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *TestStep) UnsetData() {
	o.Data.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TestStep) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStep) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestStep) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *TestStep) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *TestStep) GetSteps() []map[string]interface{} {
	if o == nil || IsNil(o.Steps) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStep) GetStepsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *TestStep) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []map[string]interface{} and assigns it to the Steps field.
func (o *TestStep) SetSteps(v []map[string]interface{}) {
	o.Steps = v
}

func (o TestStep) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if o.SharedStepHash.IsSet() {
		toSerialize["shared_step_hash"] = o.SharedStepHash.Get()
	}
	if o.SharedStepNestedHash.IsSet() {
		toSerialize["shared_step_nested_hash"] = o.SharedStepNestedHash.Get()
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if o.ExpectedResult.IsSet() {
		toSerialize["expected_result"] = o.ExpectedResult.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	return toSerialize, nil
}

type NullableTestStep struct {
	value *TestStep
	isSet bool
}

func (v NullableTestStep) Get() *TestStep {
	return v.value
}

func (v *NullableTestStep) Set(val *TestStep) {
	v.value = val
	v.isSet = true
}

func (v NullableTestStep) IsSet() bool {
	return v.isSet
}

func (v *NullableTestStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestStep(val *TestStep) *NullableTestStep {
	return &NullableTestStep{value: val, isSet: true}
}

func (v NullableTestStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
