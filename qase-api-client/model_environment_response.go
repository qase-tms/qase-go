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

// checks if the EnvironmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentResponse{}

// EnvironmentResponse struct for EnvironmentResponse
type EnvironmentResponse struct {
	Status *bool        `json:"status,omitempty"`
	Result *Environment `json:"result,omitempty"`
}

// NewEnvironmentResponse instantiates a new EnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentResponse() *EnvironmentResponse {
	this := EnvironmentResponse{}
	return &this
}

// NewEnvironmentResponseWithDefaults instantiates a new EnvironmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentResponseWithDefaults() *EnvironmentResponse {
	this := EnvironmentResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnvironmentResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnvironmentResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *EnvironmentResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *EnvironmentResponse) GetResult() Environment {
	if o == nil || IsNil(o.Result) {
		var ret Environment
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetResultOk() (*Environment, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *EnvironmentResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Environment and assigns it to the Result field.
func (o *EnvironmentResponse) SetResult(v Environment) {
	o.Result = &v
}

func (o EnvironmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableEnvironmentResponse struct {
	value *EnvironmentResponse
	isSet bool
}

func (v NullableEnvironmentResponse) Get() *EnvironmentResponse {
	return v.value
}

func (v *NullableEnvironmentResponse) Set(val *EnvironmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentResponse(val *EnvironmentResponse) *NullableEnvironmentResponse {
	return &NullableEnvironmentResponse{value: val, isSet: true}
}

func (v NullableEnvironmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
