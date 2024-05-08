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

// checks if the HashResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HashResponse{}

// HashResponse struct for HashResponse
type HashResponse struct {
	Status *bool                    `json:"status,omitempty"`
	Result *HashResponseAllOfResult `json:"result,omitempty"`
}

// NewHashResponse instantiates a new HashResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashResponse() *HashResponse {
	this := HashResponse{}
	return &this
}

// NewHashResponseWithDefaults instantiates a new HashResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashResponseWithDefaults() *HashResponse {
	this := HashResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HashResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HashResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *HashResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *HashResponse) GetResult() HashResponseAllOfResult {
	if o == nil || IsNil(o.Result) {
		var ret HashResponseAllOfResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashResponse) GetResultOk() (*HashResponseAllOfResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *HashResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given HashResponseAllOfResult and assigns it to the Result field.
func (o *HashResponse) SetResult(v HashResponseAllOfResult) {
	o.Result = &v
}

func (o HashResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableHashResponse struct {
	value *HashResponse
	isSet bool
}

func (v NullableHashResponse) Get() *HashResponse {
	return v.value
}

func (v *NullableHashResponse) Set(val *HashResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHashResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHashResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashResponse(val *HashResponse) *NullableHashResponse {
	return &NullableHashResponse{value: val, isSet: true}
}

func (v NullableHashResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
