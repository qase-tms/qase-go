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

// checks if the TestCasebulk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestCasebulk{}

// TestCasebulk struct for TestCasebulk
type TestCasebulk struct {
	Cases []TestCasebulkCasesInner `json:"cases"`
}

type _TestCasebulk TestCasebulk

// NewTestCasebulk instantiates a new TestCasebulk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestCasebulk(cases []TestCasebulkCasesInner) *TestCasebulk {
	this := TestCasebulk{}
	this.Cases = cases
	return &this
}

// NewTestCasebulkWithDefaults instantiates a new TestCasebulk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestCasebulkWithDefaults() *TestCasebulk {
	this := TestCasebulk{}
	return &this
}

// GetCases returns the Cases field value
func (o *TestCasebulk) GetCases() []TestCasebulkCasesInner {
	if o == nil {
		var ret []TestCasebulkCasesInner
		return ret
	}

	return o.Cases
}

// GetCasesOk returns a tuple with the Cases field value
// and a boolean to check if the value has been set.
func (o *TestCasebulk) GetCasesOk() ([]TestCasebulkCasesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cases, true
}

// SetCases sets field value
func (o *TestCasebulk) SetCases(v []TestCasebulkCasesInner) {
	o.Cases = v
}

func (o TestCasebulk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestCasebulk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cases"] = o.Cases
	return toSerialize, nil
}

func (o *TestCasebulk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cases",
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

	varTestCasebulk := _TestCasebulk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestCasebulk)

	if err != nil {
		return err
	}

	*o = TestCasebulk(varTestCasebulk)

	return err
}

type NullableTestCasebulk struct {
	value *TestCasebulk
	isSet bool
}

func (v NullableTestCasebulk) Get() *TestCasebulk {
	return v.value
}

func (v *NullableTestCasebulk) Set(val *TestCasebulk) {
	v.value = val
	v.isSet = true
}

func (v NullableTestCasebulk) IsSet() bool {
	return v.isSet
}

func (v *NullableTestCasebulk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestCasebulk(val *TestCasebulk) *NullableTestCasebulk {
	return &NullableTestCasebulk{value: val, isSet: true}
}

func (v NullableTestCasebulk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestCasebulk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
