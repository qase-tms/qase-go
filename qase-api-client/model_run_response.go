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

// checks if the RunResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunResponse{}

// RunResponse struct for RunResponse
type RunResponse struct {
	Status *bool `json:"status,omitempty"`
	Result *Run  `json:"result,omitempty"`
}

// NewRunResponse instantiates a new RunResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunResponse() *RunResponse {
	this := RunResponse{}
	return &this
}

// NewRunResponseWithDefaults instantiates a new RunResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunResponseWithDefaults() *RunResponse {
	this := RunResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RunResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RunResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *RunResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RunResponse) GetResult() Run {
	if o == nil || IsNil(o.Result) {
		var ret Run
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunResponse) GetResultOk() (*Run, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RunResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Run and assigns it to the Result field.
func (o *RunResponse) SetResult(v Run) {
	o.Result = &v
}

func (o RunResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableRunResponse struct {
	value *RunResponse
	isSet bool
}

func (v NullableRunResponse) Get() *RunResponse {
	return v.value
}

func (v *NullableRunResponse) Set(val *RunResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRunResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRunResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunResponse(val *RunResponse) *NullableRunResponse {
	return &NullableRunResponse{value: val, isSet: true}
}

func (v NullableRunResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
