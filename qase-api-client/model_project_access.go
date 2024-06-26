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

// checks if the ProjectAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAccess{}

// ProjectAccess struct for ProjectAccess
type ProjectAccess struct {
	// Team member id title.
	MemberId *int64 `json:"member_id,omitempty"`
}

// NewProjectAccess instantiates a new ProjectAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAccess() *ProjectAccess {
	this := ProjectAccess{}
	return &this
}

// NewProjectAccessWithDefaults instantiates a new ProjectAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAccessWithDefaults() *ProjectAccess {
	this := ProjectAccess{}
	return &this
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *ProjectAccess) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAccess) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *ProjectAccess) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
func (o *ProjectAccess) SetMemberId(v int64) {
	o.MemberId = &v
}

func (o ProjectAccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberId) {
		toSerialize["member_id"] = o.MemberId
	}
	return toSerialize, nil
}

type NullableProjectAccess struct {
	value *ProjectAccess
	isSet bool
}

func (v NullableProjectAccess) Get() *ProjectAccess {
	return v.value
}

func (v *NullableProjectAccess) Set(val *ProjectAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAccess(val *ProjectAccess) *NullableProjectAccess {
	return &NullableProjectAccess{value: val, isSet: true}
}

func (v NullableProjectAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
