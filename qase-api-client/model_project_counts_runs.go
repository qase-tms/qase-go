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

// checks if the ProjectCountsRuns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCountsRuns{}

// ProjectCountsRuns struct for ProjectCountsRuns
type ProjectCountsRuns struct {
	Total  *int32 `json:"total,omitempty"`
	Active *int32 `json:"active,omitempty"`
}

// NewProjectCountsRuns instantiates a new ProjectCountsRuns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCountsRuns() *ProjectCountsRuns {
	this := ProjectCountsRuns{}
	return &this
}

// NewProjectCountsRunsWithDefaults instantiates a new ProjectCountsRuns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCountsRunsWithDefaults() *ProjectCountsRuns {
	this := ProjectCountsRuns{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ProjectCountsRuns) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCountsRuns) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ProjectCountsRuns) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ProjectCountsRuns) SetTotal(v int32) {
	o.Total = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ProjectCountsRuns) GetActive() int32 {
	if o == nil || IsNil(o.Active) {
		var ret int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCountsRuns) GetActiveOk() (*int32, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ProjectCountsRuns) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given int32 and assigns it to the Active field.
func (o *ProjectCountsRuns) SetActive(v int32) {
	o.Active = &v
}

func (o ProjectCountsRuns) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCountsRuns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableProjectCountsRuns struct {
	value *ProjectCountsRuns
	isSet bool
}

func (v NullableProjectCountsRuns) Get() *ProjectCountsRuns {
	return v.value
}

func (v *NullableProjectCountsRuns) Set(val *ProjectCountsRuns) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCountsRuns) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCountsRuns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCountsRuns(val *ProjectCountsRuns) *NullableProjectCountsRuns {
	return &NullableProjectCountsRuns{value: val, isSet: true}
}

func (v NullableProjectCountsRuns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCountsRuns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
