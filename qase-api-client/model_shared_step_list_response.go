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

// checks if the SharedStepListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedStepListResponse{}

// SharedStepListResponse struct for SharedStepListResponse
type SharedStepListResponse struct {
	Status *bool                              `json:"status,omitempty"`
	Result *SharedStepListResponseAllOfResult `json:"result,omitempty"`
}

// NewSharedStepListResponse instantiates a new SharedStepListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedStepListResponse() *SharedStepListResponse {
	this := SharedStepListResponse{}
	return &this
}

// NewSharedStepListResponseWithDefaults instantiates a new SharedStepListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedStepListResponseWithDefaults() *SharedStepListResponse {
	this := SharedStepListResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SharedStepListResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepListResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SharedStepListResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *SharedStepListResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SharedStepListResponse) GetResult() SharedStepListResponseAllOfResult {
	if o == nil || IsNil(o.Result) {
		var ret SharedStepListResponseAllOfResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepListResponse) GetResultOk() (*SharedStepListResponseAllOfResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SharedStepListResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given SharedStepListResponseAllOfResult and assigns it to the Result field.
func (o *SharedStepListResponse) SetResult(v SharedStepListResponseAllOfResult) {
	o.Result = &v
}

func (o SharedStepListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedStepListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableSharedStepListResponse struct {
	value *SharedStepListResponse
	isSet bool
}

func (v NullableSharedStepListResponse) Get() *SharedStepListResponse {
	return v.value
}

func (v *NullableSharedStepListResponse) Set(val *SharedStepListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedStepListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedStepListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedStepListResponse(val *SharedStepListResponse) *NullableSharedStepListResponse {
	return &NullableSharedStepListResponse{value: val, isSet: true}
}

func (v NullableSharedStepListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedStepListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
