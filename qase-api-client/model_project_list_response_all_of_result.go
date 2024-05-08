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

// checks if the ProjectListResponseAllOfResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectListResponseAllOfResult{}

// ProjectListResponseAllOfResult struct for ProjectListResponseAllOfResult
type ProjectListResponseAllOfResult struct {
	Total    *int32    `json:"total,omitempty"`
	Filtered *int32    `json:"filtered,omitempty"`
	Count    *int32    `json:"count,omitempty"`
	Entities []Project `json:"entities,omitempty"`
}

// NewProjectListResponseAllOfResult instantiates a new ProjectListResponseAllOfResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectListResponseAllOfResult() *ProjectListResponseAllOfResult {
	this := ProjectListResponseAllOfResult{}
	return &this
}

// NewProjectListResponseAllOfResultWithDefaults instantiates a new ProjectListResponseAllOfResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListResponseAllOfResultWithDefaults() *ProjectListResponseAllOfResult {
	this := ProjectListResponseAllOfResult{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ProjectListResponseAllOfResult) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResponseAllOfResult) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ProjectListResponseAllOfResult) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ProjectListResponseAllOfResult) SetTotal(v int32) {
	o.Total = &v
}

// GetFiltered returns the Filtered field value if set, zero value otherwise.
func (o *ProjectListResponseAllOfResult) GetFiltered() int32 {
	if o == nil || IsNil(o.Filtered) {
		var ret int32
		return ret
	}
	return *o.Filtered
}

// GetFilteredOk returns a tuple with the Filtered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResponseAllOfResult) GetFilteredOk() (*int32, bool) {
	if o == nil || IsNil(o.Filtered) {
		return nil, false
	}
	return o.Filtered, true
}

// HasFiltered returns a boolean if a field has been set.
func (o *ProjectListResponseAllOfResult) HasFiltered() bool {
	if o != nil && !IsNil(o.Filtered) {
		return true
	}

	return false
}

// SetFiltered gets a reference to the given int32 and assigns it to the Filtered field.
func (o *ProjectListResponseAllOfResult) SetFiltered(v int32) {
	o.Filtered = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ProjectListResponseAllOfResult) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResponseAllOfResult) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ProjectListResponseAllOfResult) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ProjectListResponseAllOfResult) SetCount(v int32) {
	o.Count = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *ProjectListResponseAllOfResult) GetEntities() []Project {
	if o == nil || IsNil(o.Entities) {
		var ret []Project
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResponseAllOfResult) GetEntitiesOk() ([]Project, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *ProjectListResponseAllOfResult) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []Project and assigns it to the Entities field.
func (o *ProjectListResponseAllOfResult) SetEntities(v []Project) {
	o.Entities = v
}

func (o ProjectListResponseAllOfResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectListResponseAllOfResult) ToMap() (map[string]interface{}, error) {
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

type NullableProjectListResponseAllOfResult struct {
	value *ProjectListResponseAllOfResult
	isSet bool
}

func (v NullableProjectListResponseAllOfResult) Get() *ProjectListResponseAllOfResult {
	return v.value
}

func (v *NullableProjectListResponseAllOfResult) Set(val *ProjectListResponseAllOfResult) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectListResponseAllOfResult) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectListResponseAllOfResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectListResponseAllOfResult(val *ProjectListResponseAllOfResult) *NullableProjectListResponseAllOfResult {
	return &NullableProjectListResponseAllOfResult{value: val, isSet: true}
}

func (v NullableProjectListResponseAllOfResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectListResponseAllOfResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}