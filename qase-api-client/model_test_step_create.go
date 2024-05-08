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

// checks if the TestStepCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestStepCreate{}

// TestStepCreate struct for TestStepCreate
type TestStepCreate struct {
	Action         *string `json:"action,omitempty"`
	ExpectedResult *string `json:"expected_result,omitempty"`
	Data           *string `json:"data,omitempty"`
	// Deprecated
	Position *int32 `json:"position,omitempty"`
	// A list of Attachment hashes.
	Attachments []string `json:"attachments,omitempty"`
	// Nested steps may be passed here. Use same structure for them.
	Steps []map[string]interface{} `json:"steps,omitempty"`
}

// NewTestStepCreate instantiates a new TestStepCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestStepCreate() *TestStepCreate {
	this := TestStepCreate{}
	return &this
}

// NewTestStepCreateWithDefaults instantiates a new TestStepCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestStepCreateWithDefaults() *TestStepCreate {
	this := TestStepCreate{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TestStepCreate) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStepCreate) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TestStepCreate) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *TestStepCreate) SetAction(v string) {
	o.Action = &v
}

// GetExpectedResult returns the ExpectedResult field value if set, zero value otherwise.
func (o *TestStepCreate) GetExpectedResult() string {
	if o == nil || IsNil(o.ExpectedResult) {
		var ret string
		return ret
	}
	return *o.ExpectedResult
}

// GetExpectedResultOk returns a tuple with the ExpectedResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStepCreate) GetExpectedResultOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedResult) {
		return nil, false
	}
	return o.ExpectedResult, true
}

// HasExpectedResult returns a boolean if a field has been set.
func (o *TestStepCreate) HasExpectedResult() bool {
	if o != nil && !IsNil(o.ExpectedResult) {
		return true
	}

	return false
}

// SetExpectedResult gets a reference to the given string and assigns it to the ExpectedResult field.
func (o *TestStepCreate) SetExpectedResult(v string) {
	o.ExpectedResult = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TestStepCreate) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStepCreate) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TestStepCreate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *TestStepCreate) SetData(v string) {
	o.Data = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
// Deprecated
func (o *TestStepCreate) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TestStepCreate) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *TestStepCreate) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
// Deprecated
func (o *TestStepCreate) SetPosition(v int32) {
	o.Position = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TestStepCreate) GetAttachments() []string {
	if o == nil || IsNil(o.Attachments) {
		var ret []string
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStepCreate) GetAttachmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestStepCreate) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *TestStepCreate) SetAttachments(v []string) {
	o.Attachments = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *TestStepCreate) GetSteps() []map[string]interface{} {
	if o == nil || IsNil(o.Steps) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStepCreate) GetStepsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *TestStepCreate) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []map[string]interface{} and assigns it to the Steps field.
func (o *TestStepCreate) SetSteps(v []map[string]interface{}) {
	o.Steps = v
}

func (o TestStepCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestStepCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ExpectedResult) {
		toSerialize["expected_result"] = o.ExpectedResult
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	return toSerialize, nil
}

type NullableTestStepCreate struct {
	value *TestStepCreate
	isSet bool
}

func (v NullableTestStepCreate) Get() *TestStepCreate {
	return v.value
}

func (v *NullableTestStepCreate) Set(val *TestStepCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTestStepCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTestStepCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestStepCreate(val *TestStepCreate) *NullableTestStepCreate {
	return &NullableTestStepCreate{value: val, isSet: true}
}

func (v NullableTestStepCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestStepCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}