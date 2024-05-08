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

// checks if the RunPublicResponseAllOfResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunPublicResponseAllOfResult{}

// RunPublicResponseAllOfResult struct for RunPublicResponseAllOfResult
type RunPublicResponseAllOfResult struct {
	Url *string `json:"url,omitempty"`
}

// NewRunPublicResponseAllOfResult instantiates a new RunPublicResponseAllOfResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunPublicResponseAllOfResult() *RunPublicResponseAllOfResult {
	this := RunPublicResponseAllOfResult{}
	return &this
}

// NewRunPublicResponseAllOfResultWithDefaults instantiates a new RunPublicResponseAllOfResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunPublicResponseAllOfResultWithDefaults() *RunPublicResponseAllOfResult {
	this := RunPublicResponseAllOfResult{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RunPublicResponseAllOfResult) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunPublicResponseAllOfResult) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RunPublicResponseAllOfResult) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RunPublicResponseAllOfResult) SetUrl(v string) {
	o.Url = &v
}

func (o RunPublicResponseAllOfResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunPublicResponseAllOfResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableRunPublicResponseAllOfResult struct {
	value *RunPublicResponseAllOfResult
	isSet bool
}

func (v NullableRunPublicResponseAllOfResult) Get() *RunPublicResponseAllOfResult {
	return v.value
}

func (v *NullableRunPublicResponseAllOfResult) Set(val *RunPublicResponseAllOfResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRunPublicResponseAllOfResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRunPublicResponseAllOfResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunPublicResponseAllOfResult(val *RunPublicResponseAllOfResult) *NullableRunPublicResponseAllOfResult {
	return &NullableRunPublicResponseAllOfResult{value: val, isSet: true}
}

func (v NullableRunPublicResponseAllOfResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunPublicResponseAllOfResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
