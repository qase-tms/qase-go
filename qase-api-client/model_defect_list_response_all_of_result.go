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

// checks if the DefectListResponseAllOfResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefectListResponseAllOfResult{}

// DefectListResponseAllOfResult struct for DefectListResponseAllOfResult
type DefectListResponseAllOfResult struct {
	Total    *int32   `json:"total,omitempty"`
	Filtered *int32   `json:"filtered,omitempty"`
	Count    *int32   `json:"count,omitempty"`
	Entities []Defect `json:"entities,omitempty"`
}

// NewDefectListResponseAllOfResult instantiates a new DefectListResponseAllOfResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefectListResponseAllOfResult() *DefectListResponseAllOfResult {
	this := DefectListResponseAllOfResult{}
	return &this
}

// NewDefectListResponseAllOfResultWithDefaults instantiates a new DefectListResponseAllOfResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefectListResponseAllOfResultWithDefaults() *DefectListResponseAllOfResult {
	this := DefectListResponseAllOfResult{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DefectListResponseAllOfResult) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectListResponseAllOfResult) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DefectListResponseAllOfResult) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *DefectListResponseAllOfResult) SetTotal(v int32) {
	o.Total = &v
}

// GetFiltered returns the Filtered field value if set, zero value otherwise.
func (o *DefectListResponseAllOfResult) GetFiltered() int32 {
	if o == nil || IsNil(o.Filtered) {
		var ret int32
		return ret
	}
	return *o.Filtered
}

// GetFilteredOk returns a tuple with the Filtered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectListResponseAllOfResult) GetFilteredOk() (*int32, bool) {
	if o == nil || IsNil(o.Filtered) {
		return nil, false
	}
	return o.Filtered, true
}

// HasFiltered returns a boolean if a field has been set.
func (o *DefectListResponseAllOfResult) HasFiltered() bool {
	if o != nil && !IsNil(o.Filtered) {
		return true
	}

	return false
}

// SetFiltered gets a reference to the given int32 and assigns it to the Filtered field.
func (o *DefectListResponseAllOfResult) SetFiltered(v int32) {
	o.Filtered = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DefectListResponseAllOfResult) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectListResponseAllOfResult) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DefectListResponseAllOfResult) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *DefectListResponseAllOfResult) SetCount(v int32) {
	o.Count = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *DefectListResponseAllOfResult) GetEntities() []Defect {
	if o == nil || IsNil(o.Entities) {
		var ret []Defect
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectListResponseAllOfResult) GetEntitiesOk() ([]Defect, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *DefectListResponseAllOfResult) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []Defect and assigns it to the Entities field.
func (o *DefectListResponseAllOfResult) SetEntities(v []Defect) {
	o.Entities = v
}

func (o DefectListResponseAllOfResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefectListResponseAllOfResult) ToMap() (map[string]interface{}, error) {
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

type NullableDefectListResponseAllOfResult struct {
	value *DefectListResponseAllOfResult
	isSet bool
}

func (v NullableDefectListResponseAllOfResult) Get() *DefectListResponseAllOfResult {
	return v.value
}

func (v *NullableDefectListResponseAllOfResult) Set(val *DefectListResponseAllOfResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDefectListResponseAllOfResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDefectListResponseAllOfResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefectListResponseAllOfResult(val *DefectListResponseAllOfResult) *NullableDefectListResponseAllOfResult {
	return &NullableDefectListResponseAllOfResult{value: val, isSet: true}
}

func (v NullableDefectListResponseAllOfResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefectListResponseAllOfResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}