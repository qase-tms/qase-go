/*
Qase.io TestOps API v2

Qase TestOps API v2 Specification.

API version: 2.0.0
Contact: support@qase.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ResultStepData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultStepData{}

// ResultStepData struct for ResultStepData
type ResultStepData struct {
	Action         string   `json:"action"`
	ExpectedResult *string  `json:"expected_result,omitempty"`
	InputData      *string  `json:"input_data,omitempty"`
	Attachments    []string `json:"attachments,omitempty"`
}

type _ResultStepData ResultStepData

// NewResultStepData instantiates a new ResultStepData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultStepData(action string) *ResultStepData {
	this := ResultStepData{}
	this.Action = action
	return &this
}

// NewResultStepDataWithDefaults instantiates a new ResultStepData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultStepDataWithDefaults() *ResultStepData {
	this := ResultStepData{}
	return &this
}

// GetAction returns the Action field value
func (o *ResultStepData) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ResultStepData) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ResultStepData) SetAction(v string) {
	o.Action = v
}

// GetExpectedResult returns the ExpectedResult field value if set, zero value otherwise.
func (o *ResultStepData) GetExpectedResult() string {
	if o == nil || IsNil(o.ExpectedResult) {
		var ret string
		return ret
	}
	return *o.ExpectedResult
}

// GetExpectedResultOk returns a tuple with the ExpectedResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultStepData) GetExpectedResultOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedResult) {
		return nil, false
	}
	return o.ExpectedResult, true
}

// HasExpectedResult returns a boolean if a field has been set.
func (o *ResultStepData) HasExpectedResult() bool {
	if o != nil && !IsNil(o.ExpectedResult) {
		return true
	}

	return false
}

// SetExpectedResult gets a reference to the given string and assigns it to the ExpectedResult field.
func (o *ResultStepData) SetExpectedResult(v string) {
	o.ExpectedResult = &v
}

// GetInputData returns the InputData field value if set, zero value otherwise.
func (o *ResultStepData) GetInputData() string {
	if o == nil || IsNil(o.InputData) {
		var ret string
		return ret
	}
	return *o.InputData
}

// GetInputDataOk returns a tuple with the InputData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultStepData) GetInputDataOk() (*string, bool) {
	if o == nil || IsNil(o.InputData) {
		return nil, false
	}
	return o.InputData, true
}

// HasInputData returns a boolean if a field has been set.
func (o *ResultStepData) HasInputData() bool {
	if o != nil && !IsNil(o.InputData) {
		return true
	}

	return false
}

// SetInputData gets a reference to the given string and assigns it to the InputData field.
func (o *ResultStepData) SetInputData(v string) {
	o.InputData = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ResultStepData) GetAttachments() []string {
	if o == nil || IsNil(o.Attachments) {
		var ret []string
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultStepData) GetAttachmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ResultStepData) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *ResultStepData) SetAttachments(v []string) {
	o.Attachments = v
}

func (o ResultStepData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultStepData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.ExpectedResult) {
		toSerialize["expected_result"] = o.ExpectedResult
	}
	if !IsNil(o.InputData) {
		toSerialize["input_data"] = o.InputData
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

func (o *ResultStepData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varResultStepData := _ResultStepData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultStepData)

	if err != nil {
		return err
	}

	*o = ResultStepData(varResultStepData)

	return err
}

type NullableResultStepData struct {
	value *ResultStepData
	isSet bool
}

func (v NullableResultStepData) Get() *ResultStepData {
	return v.value
}

func (v *NullableResultStepData) Set(val *ResultStepData) {
	v.value = val
	v.isSet = true
}

func (v NullableResultStepData) IsSet() bool {
	return v.isSet
}

func (v *NullableResultStepData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultStepData(val *ResultStepData) *NullableResultStepData {
	return &NullableResultStepData{value: val, isSet: true}
}

func (v NullableResultStepData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultStepData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
