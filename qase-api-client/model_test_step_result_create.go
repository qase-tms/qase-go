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

// checks if the TestStepResultCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestStepResultCreate{}

// TestStepResultCreate struct for TestStepResultCreate
type TestStepResultCreate struct {
	// Deprecated
	Position       *int32         `json:"position,omitempty"`
	Status         string         `json:"status"`
	Comment        NullableString `json:"comment,omitempty"`
	Attachments    []string       `json:"attachments,omitempty"`
	Action         *string        `json:"action,omitempty"`
	ExpectedResult NullableString `json:"expected_result,omitempty"`
	Data           NullableString `json:"data,omitempty"`
	// Nested steps results may be passed here. Use same structure for them.
	Steps []map[string]interface{} `json:"steps,omitempty"`
}

type _TestStepResultCreate TestStepResultCreate

// NewTestStepResultCreate instantiates a new TestStepResultCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestStepResultCreate(status string) *TestStepResultCreate {
	this := TestStepResultCreate{}
	this.Status = status
	return &this
}

// NewTestStepResultCreateWithDefaults instantiates a new TestStepResultCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestStepResultCreateWithDefaults() *TestStepResultCreate {
	this := TestStepResultCreate{}
	return &this
}

// GetPosition returns the Position field value if set, zero value otherwise.
// Deprecated
func (o *TestStepResultCreate) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TestStepResultCreate) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *TestStepResultCreate) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
// Deprecated
func (o *TestStepResultCreate) SetPosition(v int32) {
	o.Position = &v
}

// GetStatus returns the Status field value
func (o *TestStepResultCreate) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TestStepResultCreate) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TestStepResultCreate) SetStatus(v string) {
	o.Status = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStepResultCreate) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStepResultCreate) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *TestStepResultCreate) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *TestStepResultCreate) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *TestStepResultCreate) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *TestStepResultCreate) UnsetComment() {
	o.Comment.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStepResultCreate) GetAttachments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStepResultCreate) GetAttachmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestStepResultCreate) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *TestStepResultCreate) SetAttachments(v []string) {
	o.Attachments = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TestStepResultCreate) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStepResultCreate) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TestStepResultCreate) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *TestStepResultCreate) SetAction(v string) {
	o.Action = &v
}

// GetExpectedResult returns the ExpectedResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStepResultCreate) GetExpectedResult() string {
	if o == nil || IsNil(o.ExpectedResult.Get()) {
		var ret string
		return ret
	}
	return *o.ExpectedResult.Get()
}

// GetExpectedResultOk returns a tuple with the ExpectedResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStepResultCreate) GetExpectedResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedResult.Get(), o.ExpectedResult.IsSet()
}

// HasExpectedResult returns a boolean if a field has been set.
func (o *TestStepResultCreate) HasExpectedResult() bool {
	if o != nil && o.ExpectedResult.IsSet() {
		return true
	}

	return false
}

// SetExpectedResult gets a reference to the given NullableString and assigns it to the ExpectedResult field.
func (o *TestStepResultCreate) SetExpectedResult(v string) {
	o.ExpectedResult.Set(&v)
}

// SetExpectedResultNil sets the value for ExpectedResult to be an explicit nil
func (o *TestStepResultCreate) SetExpectedResultNil() {
	o.ExpectedResult.Set(nil)
}

// UnsetExpectedResult ensures that no value is present for ExpectedResult, not even an explicit nil
func (o *TestStepResultCreate) UnsetExpectedResult() {
	o.ExpectedResult.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStepResultCreate) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStepResultCreate) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *TestStepResultCreate) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *TestStepResultCreate) SetData(v string) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil
func (o *TestStepResultCreate) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *TestStepResultCreate) UnsetData() {
	o.Data.Unset()
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *TestStepResultCreate) GetSteps() []map[string]interface{} {
	if o == nil || IsNil(o.Steps) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestStepResultCreate) GetStepsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *TestStepResultCreate) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []map[string]interface{} and assigns it to the Steps field.
func (o *TestStepResultCreate) SetSteps(v []map[string]interface{}) {
	o.Steps = v
}

func (o TestStepResultCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestStepResultCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	toSerialize["status"] = o.Status
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
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
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	return toSerialize, nil
}

func (o *TestStepResultCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varTestStepResultCreate := _TestStepResultCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestStepResultCreate)

	if err != nil {
		return err
	}

	*o = TestStepResultCreate(varTestStepResultCreate)

	return err
}

type NullableTestStepResultCreate struct {
	value *TestStepResultCreate
	isSet bool
}

func (v NullableTestStepResultCreate) Get() *TestStepResultCreate {
	return v.value
}

func (v *NullableTestStepResultCreate) Set(val *TestStepResultCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTestStepResultCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTestStepResultCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestStepResultCreate(val *TestStepResultCreate) *NullableTestStepResultCreate {
	return &NullableTestStepResultCreate{value: val, isSet: true}
}

func (v NullableTestStepResultCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestStepResultCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
