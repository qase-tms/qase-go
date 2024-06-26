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

// checks if the HashResponseAllOfResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HashResponseAllOfResult{}

// HashResponseAllOfResult struct for HashResponseAllOfResult
type HashResponseAllOfResult struct {
	Hash *string `json:"hash,omitempty"`
}

// NewHashResponseAllOfResult instantiates a new HashResponseAllOfResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashResponseAllOfResult() *HashResponseAllOfResult {
	this := HashResponseAllOfResult{}
	return &this
}

// NewHashResponseAllOfResultWithDefaults instantiates a new HashResponseAllOfResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashResponseAllOfResultWithDefaults() *HashResponseAllOfResult {
	this := HashResponseAllOfResult{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *HashResponseAllOfResult) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashResponseAllOfResult) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *HashResponseAllOfResult) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *HashResponseAllOfResult) SetHash(v string) {
	o.Hash = &v
}

func (o HashResponseAllOfResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashResponseAllOfResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	return toSerialize, nil
}

type NullableHashResponseAllOfResult struct {
	value *HashResponseAllOfResult
	isSet bool
}

func (v NullableHashResponseAllOfResult) Get() *HashResponseAllOfResult {
	return v.value
}

func (v *NullableHashResponseAllOfResult) Set(val *HashResponseAllOfResult) {
	v.value = val
	v.isSet = true
}

func (v NullableHashResponseAllOfResult) IsSet() bool {
	return v.isSet
}

func (v *NullableHashResponseAllOfResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashResponseAllOfResult(val *HashResponseAllOfResult) *NullableHashResponseAllOfResult {
	return &NullableHashResponseAllOfResult{value: val, isSet: true}
}

func (v NullableHashResponseAllOfResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashResponseAllOfResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
