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

// checks if the SharedStepListResponseAllOfResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedStepListResponseAllOfResult{}

// SharedStepListResponseAllOfResult struct for SharedStepListResponseAllOfResult
type SharedStepListResponseAllOfResult struct {
	Total    *int32       `json:"total,omitempty"`
	Filtered *int32       `json:"filtered,omitempty"`
	Count    *int32       `json:"count,omitempty"`
	Entities []SharedStep `json:"entities,omitempty"`
}

// NewSharedStepListResponseAllOfResult instantiates a new SharedStepListResponseAllOfResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedStepListResponseAllOfResult() *SharedStepListResponseAllOfResult {
	this := SharedStepListResponseAllOfResult{}
	return &this
}

// NewSharedStepListResponseAllOfResultWithDefaults instantiates a new SharedStepListResponseAllOfResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedStepListResponseAllOfResultWithDefaults() *SharedStepListResponseAllOfResult {
	this := SharedStepListResponseAllOfResult{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SharedStepListResponseAllOfResult) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepListResponseAllOfResult) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SharedStepListResponseAllOfResult) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *SharedStepListResponseAllOfResult) SetTotal(v int32) {
	o.Total = &v
}

// GetFiltered returns the Filtered field value if set, zero value otherwise.
func (o *SharedStepListResponseAllOfResult) GetFiltered() int32 {
	if o == nil || IsNil(o.Filtered) {
		var ret int32
		return ret
	}
	return *o.Filtered
}

// GetFilteredOk returns a tuple with the Filtered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepListResponseAllOfResult) GetFilteredOk() (*int32, bool) {
	if o == nil || IsNil(o.Filtered) {
		return nil, false
	}
	return o.Filtered, true
}

// HasFiltered returns a boolean if a field has been set.
func (o *SharedStepListResponseAllOfResult) HasFiltered() bool {
	if o != nil && !IsNil(o.Filtered) {
		return true
	}

	return false
}

// SetFiltered gets a reference to the given int32 and assigns it to the Filtered field.
func (o *SharedStepListResponseAllOfResult) SetFiltered(v int32) {
	o.Filtered = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SharedStepListResponseAllOfResult) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepListResponseAllOfResult) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SharedStepListResponseAllOfResult) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SharedStepListResponseAllOfResult) SetCount(v int32) {
	o.Count = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *SharedStepListResponseAllOfResult) GetEntities() []SharedStep {
	if o == nil || IsNil(o.Entities) {
		var ret []SharedStep
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepListResponseAllOfResult) GetEntitiesOk() ([]SharedStep, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *SharedStepListResponseAllOfResult) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []SharedStep and assigns it to the Entities field.
func (o *SharedStepListResponseAllOfResult) SetEntities(v []SharedStep) {
	o.Entities = v
}

func (o SharedStepListResponseAllOfResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedStepListResponseAllOfResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Filtered) {
		toSerialize["filtered"] = o.Filtered
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	return toSerialize, nil
}

type NullableSharedStepListResponseAllOfResult struct {
	value *SharedStepListResponseAllOfResult
	isSet bool
}

func (v NullableSharedStepListResponseAllOfResult) Get() *SharedStepListResponseAllOfResult {
	return v.value
}

func (v *NullableSharedStepListResponseAllOfResult) Set(val *SharedStepListResponseAllOfResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedStepListResponseAllOfResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedStepListResponseAllOfResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedStepListResponseAllOfResult(val *SharedStepListResponseAllOfResult) *NullableSharedStepListResponseAllOfResult {
	return &NullableSharedStepListResponseAllOfResult{value: val, isSet: true}
}

func (v NullableSharedStepListResponseAllOfResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedStepListResponseAllOfResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
