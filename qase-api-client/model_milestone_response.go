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

// checks if the MilestoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MilestoneResponse{}

// MilestoneResponse struct for MilestoneResponse
type MilestoneResponse struct {
	Status *bool      `json:"status,omitempty"`
	Result *Milestone `json:"result,omitempty"`
}

// NewMilestoneResponse instantiates a new MilestoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMilestoneResponse() *MilestoneResponse {
	this := MilestoneResponse{}
	return &this
}

// NewMilestoneResponseWithDefaults instantiates a new MilestoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMilestoneResponseWithDefaults() *MilestoneResponse {
	this := MilestoneResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MilestoneResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilestoneResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MilestoneResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *MilestoneResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *MilestoneResponse) GetResult() Milestone {
	if o == nil || IsNil(o.Result) {
		var ret Milestone
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilestoneResponse) GetResultOk() (*Milestone, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *MilestoneResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Milestone and assigns it to the Result field.
func (o *MilestoneResponse) SetResult(v Milestone) {
	o.Result = &v
}

func (o MilestoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MilestoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableMilestoneResponse struct {
	value *MilestoneResponse
	isSet bool
}

func (v NullableMilestoneResponse) Get() *MilestoneResponse {
	return v.value
}

func (v *NullableMilestoneResponse) Set(val *MilestoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMilestoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMilestoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMilestoneResponse(val *MilestoneResponse) *NullableMilestoneResponse {
	return &NullableMilestoneResponse{value: val, isSet: true}
}

func (v NullableMilestoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMilestoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
