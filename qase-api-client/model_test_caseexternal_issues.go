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

// checks if the TestCaseexternalIssues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestCaseexternalIssues{}

// TestCaseexternalIssues struct for TestCaseexternalIssues
type TestCaseexternalIssues struct {
	Type  string                             `json:"type"`
	Links []TestCaseExternalIssuesLinksInner `json:"links"`
}

type _TestCaseexternalIssues TestCaseexternalIssues

// NewTestCaseexternalIssues instantiates a new TestCaseexternalIssues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestCaseexternalIssues(type_ string, links []TestCaseExternalIssuesLinksInner) *TestCaseexternalIssues {
	this := TestCaseexternalIssues{}
	this.Type = type_
	this.Links = links
	return &this
}

// NewTestCaseexternalIssuesWithDefaults instantiates a new TestCaseexternalIssues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestCaseexternalIssuesWithDefaults() *TestCaseexternalIssues {
	this := TestCaseexternalIssues{}
	return &this
}

// GetType returns the Type field value
func (o *TestCaseexternalIssues) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TestCaseexternalIssues) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TestCaseexternalIssues) SetType(v string) {
	o.Type = v
}

// GetLinks returns the Links field value
func (o *TestCaseexternalIssues) GetLinks() []TestCaseExternalIssuesLinksInner {
	if o == nil {
		var ret []TestCaseExternalIssuesLinksInner
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TestCaseexternalIssues) GetLinksOk() ([]TestCaseExternalIssuesLinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *TestCaseexternalIssues) SetLinks(v []TestCaseExternalIssuesLinksInner) {
	o.Links = v
}

func (o TestCaseexternalIssues) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestCaseexternalIssues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *TestCaseexternalIssues) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"links",
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

	varTestCaseexternalIssues := _TestCaseexternalIssues{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestCaseexternalIssues)

	if err != nil {
		return err
	}

	*o = TestCaseexternalIssues(varTestCaseexternalIssues)

	return err
}

type NullableTestCaseexternalIssues struct {
	value *TestCaseexternalIssues
	isSet bool
}

func (v NullableTestCaseexternalIssues) Get() *TestCaseexternalIssues {
	return v.value
}

func (v *NullableTestCaseexternalIssues) Set(val *TestCaseexternalIssues) {
	v.value = val
	v.isSet = true
}

func (v NullableTestCaseexternalIssues) IsSet() bool {
	return v.isSet
}

func (v *NullableTestCaseexternalIssues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestCaseexternalIssues(val *TestCaseexternalIssues) *NullableTestCaseexternalIssues {
	return &NullableTestCaseexternalIssues{value: val, isSet: true}
}

func (v NullableTestCaseexternalIssues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestCaseexternalIssues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
