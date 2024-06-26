/*
Qase.io TestOps API v2

Qase TestOps API v2 Specification.

API version: 2.0.0
Contact: support@qase.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2_client

import (
	"encoding/json"
)

// checks if the ResultRelations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultRelations{}

// ResultRelations struct for ResultRelations
type ResultRelations struct {
	Suite NullableRelationSuite `json:"suite,omitempty"`
}

// NewResultRelations instantiates a new ResultRelations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultRelations() *ResultRelations {
	this := ResultRelations{}
	return &this
}

// NewResultRelationsWithDefaults instantiates a new ResultRelations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultRelationsWithDefaults() *ResultRelations {
	this := ResultRelations{}
	return &this
}

// GetSuite returns the Suite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultRelations) GetSuite() RelationSuite {
	if o == nil || IsNil(o.Suite.Get()) {
		var ret RelationSuite
		return ret
	}
	return *o.Suite.Get()
}

// GetSuiteOk returns a tuple with the Suite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultRelations) GetSuiteOk() (*RelationSuite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suite.Get(), o.Suite.IsSet()
}

// HasSuite returns a boolean if a field has been set.
func (o *ResultRelations) HasSuite() bool {
	if o != nil && o.Suite.IsSet() {
		return true
	}

	return false
}

// SetSuite gets a reference to the given NullableRelationSuite and assigns it to the Suite field.
func (o *ResultRelations) SetSuite(v RelationSuite) {
	o.Suite.Set(&v)
}

// SetSuiteNil sets the value for Suite to be an explicit nil
func (o *ResultRelations) SetSuiteNil() {
	o.Suite.Set(nil)
}

// UnsetSuite ensures that no value is present for Suite, not even an explicit nil
func (o *ResultRelations) UnsetSuite() {
	o.Suite.Unset()
}

func (o ResultRelations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultRelations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Suite.IsSet() {
		toSerialize["suite"] = o.Suite.Get()
	}
	return toSerialize, nil
}

type NullableResultRelations struct {
	value *ResultRelations
	isSet bool
}

func (v NullableResultRelations) Get() *ResultRelations {
	return v.value
}

func (v *NullableResultRelations) Set(val *ResultRelations) {
	v.value = val
	v.isSet = true
}

func (v NullableResultRelations) IsSet() bool {
	return v.isSet
}

func (v *NullableResultRelations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultRelations(val *ResultRelations) *NullableResultRelations {
	return &NullableResultRelations{value: val, isSet: true}
}

func (v NullableResultRelations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultRelations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
