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

// checks if the ResultCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultCreateResponse{}

// ResultCreateResponse struct for ResultCreateResponse
type ResultCreateResponse struct {
	Status *bool                            `json:"status,omitempty"`
	Result *ResultCreateResponseAllOfResult `json:"result,omitempty"`
}

// NewResultCreateResponse instantiates a new ResultCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultCreateResponse() *ResultCreateResponse {
	this := ResultCreateResponse{}
	return &this
}

// NewResultCreateResponseWithDefaults instantiates a new ResultCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultCreateResponseWithDefaults() *ResultCreateResponse {
	this := ResultCreateResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResultCreateResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResultCreateResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *ResultCreateResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ResultCreateResponse) GetResult() ResultCreateResponseAllOfResult {
	if o == nil || IsNil(o.Result) {
		var ret ResultCreateResponseAllOfResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateResponse) GetResultOk() (*ResultCreateResponseAllOfResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ResultCreateResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ResultCreateResponseAllOfResult and assigns it to the Result field.
func (o *ResultCreateResponse) SetResult(v ResultCreateResponseAllOfResult) {
	o.Result = &v
}

func (o ResultCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableResultCreateResponse struct {
	value *ResultCreateResponse
	isSet bool
}

func (v NullableResultCreateResponse) Get() *ResultCreateResponse {
	return v.value
}

func (v *NullableResultCreateResponse) Set(val *ResultCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCreateResponse(val *ResultCreateResponse) *NullableResultCreateResponse {
	return &NullableResultCreateResponse{value: val, isSet: true}
}

func (v NullableResultCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
