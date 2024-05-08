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

// checks if the SuiteListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuiteListResponse{}

// SuiteListResponse struct for SuiteListResponse
type SuiteListResponse struct {
	Status *bool                         `json:"status,omitempty"`
	Result *SuiteListResponseAllOfResult `json:"result,omitempty"`
}

// NewSuiteListResponse instantiates a new SuiteListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuiteListResponse() *SuiteListResponse {
	this := SuiteListResponse{}
	return &this
}

// NewSuiteListResponseWithDefaults instantiates a new SuiteListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuiteListResponseWithDefaults() *SuiteListResponse {
	this := SuiteListResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SuiteListResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuiteListResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SuiteListResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *SuiteListResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SuiteListResponse) GetResult() SuiteListResponseAllOfResult {
	if o == nil || IsNil(o.Result) {
		var ret SuiteListResponseAllOfResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuiteListResponse) GetResultOk() (*SuiteListResponseAllOfResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SuiteListResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given SuiteListResponseAllOfResult and assigns it to the Result field.
func (o *SuiteListResponse) SetResult(v SuiteListResponseAllOfResult) {
	o.Result = &v
}

func (o SuiteListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuiteListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableSuiteListResponse struct {
	value *SuiteListResponse
	isSet bool
}

func (v NullableSuiteListResponse) Get() *SuiteListResponse {
	return v.value
}

func (v *NullableSuiteListResponse) Set(val *SuiteListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuiteListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuiteListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuiteListResponse(val *SuiteListResponse) *NullableSuiteListResponse {
	return &NullableSuiteListResponse{value: val, isSet: true}
}

func (v NullableSuiteListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuiteListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
