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

// checks if the ResultCreateBulk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultCreateBulk{}

// ResultCreateBulk struct for ResultCreateBulk
type ResultCreateBulk struct {
	Results []ResultCreate `json:"results"`
}

type _ResultCreateBulk ResultCreateBulk

// NewResultCreateBulk instantiates a new ResultCreateBulk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultCreateBulk(results []ResultCreate) *ResultCreateBulk {
	this := ResultCreateBulk{}
	this.Results = results
	return &this
}

// NewResultCreateBulkWithDefaults instantiates a new ResultCreateBulk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultCreateBulkWithDefaults() *ResultCreateBulk {
	this := ResultCreateBulk{}
	return &this
}

// GetResults returns the Results field value
func (o *ResultCreateBulk) GetResults() []ResultCreate {
	if o == nil {
		var ret []ResultCreate
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResultCreateBulk) GetResultsOk() ([]ResultCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResultCreateBulk) SetResults(v []ResultCreate) {
	o.Results = v
}

func (o ResultCreateBulk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultCreateBulk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *ResultCreateBulk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varResultCreateBulk := _ResultCreateBulk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultCreateBulk)

	if err != nil {
		return err
	}

	*o = ResultCreateBulk(varResultCreateBulk)

	return err
}

type NullableResultCreateBulk struct {
	value *ResultCreateBulk
	isSet bool
}

func (v NullableResultCreateBulk) Get() *ResultCreateBulk {
	return v.value
}

func (v *NullableResultCreateBulk) Set(val *ResultCreateBulk) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCreateBulk) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCreateBulk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCreateBulk(val *ResultCreateBulk) *NullableResultCreateBulk {
	return &NullableResultCreateBulk{value: val, isSet: true}
}

func (v NullableResultCreateBulk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCreateBulk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}