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

// checks if the TestCaseListResponseAllOfResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestCaseListResponseAllOfResult{}

// TestCaseListResponseAllOfResult struct for TestCaseListResponseAllOfResult
type TestCaseListResponseAllOfResult struct {
	Total    *int32     `json:"total,omitempty"`
	Filtered *int32     `json:"filtered,omitempty"`
	Count    *int32     `json:"count,omitempty"`
	Entities []TestCase `json:"entities,omitempty"`
}

// NewTestCaseListResponseAllOfResult instantiates a new TestCaseListResponseAllOfResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestCaseListResponseAllOfResult() *TestCaseListResponseAllOfResult {
	this := TestCaseListResponseAllOfResult{}
	return &this
}

// NewTestCaseListResponseAllOfResultWithDefaults instantiates a new TestCaseListResponseAllOfResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestCaseListResponseAllOfResultWithDefaults() *TestCaseListResponseAllOfResult {
	this := TestCaseListResponseAllOfResult{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TestCaseListResponseAllOfResult) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseListResponseAllOfResult) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TestCaseListResponseAllOfResult) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *TestCaseListResponseAllOfResult) SetTotal(v int32) {
	o.Total = &v
}

// GetFiltered returns the Filtered field value if set, zero value otherwise.
func (o *TestCaseListResponseAllOfResult) GetFiltered() int32 {
	if o == nil || IsNil(o.Filtered) {
		var ret int32
		return ret
	}
	return *o.Filtered
}

// GetFilteredOk returns a tuple with the Filtered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseListResponseAllOfResult) GetFilteredOk() (*int32, bool) {
	if o == nil || IsNil(o.Filtered) {
		return nil, false
	}
	return o.Filtered, true
}

// HasFiltered returns a boolean if a field has been set.
func (o *TestCaseListResponseAllOfResult) HasFiltered() bool {
	if o != nil && !IsNil(o.Filtered) {
		return true
	}

	return false
}

// SetFiltered gets a reference to the given int32 and assigns it to the Filtered field.
func (o *TestCaseListResponseAllOfResult) SetFiltered(v int32) {
	o.Filtered = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TestCaseListResponseAllOfResult) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseListResponseAllOfResult) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TestCaseListResponseAllOfResult) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *TestCaseListResponseAllOfResult) SetCount(v int32) {
	o.Count = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *TestCaseListResponseAllOfResult) GetEntities() []TestCase {
	if o == nil || IsNil(o.Entities) {
		var ret []TestCase
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCaseListResponseAllOfResult) GetEntitiesOk() ([]TestCase, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *TestCaseListResponseAllOfResult) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []TestCase and assigns it to the Entities field.
func (o *TestCaseListResponseAllOfResult) SetEntities(v []TestCase) {
	o.Entities = v
}

func (o TestCaseListResponseAllOfResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestCaseListResponseAllOfResult) ToMap() (map[string]interface{}, error) {
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

type NullableTestCaseListResponseAllOfResult struct {
	value *TestCaseListResponseAllOfResult
	isSet bool
}

func (v NullableTestCaseListResponseAllOfResult) Get() *TestCaseListResponseAllOfResult {
	return v.value
}

func (v *NullableTestCaseListResponseAllOfResult) Set(val *TestCaseListResponseAllOfResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestCaseListResponseAllOfResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestCaseListResponseAllOfResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestCaseListResponseAllOfResult(val *TestCaseListResponseAllOfResult) *NullableTestCaseListResponseAllOfResult {
	return &NullableTestCaseListResponseAllOfResult{value: val, isSet: true}
}

func (v NullableTestCaseListResponseAllOfResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestCaseListResponseAllOfResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
