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

// checks if the CreateResultsRequestV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateResultsRequestV2{}

// CreateResultsRequestV2 struct for CreateResultsRequestV2
type CreateResultsRequestV2 struct {
	Results []ResultCreate `json:"results,omitempty"`
}

// NewCreateResultsRequestV2 instantiates a new CreateResultsRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateResultsRequestV2() *CreateResultsRequestV2 {
	this := CreateResultsRequestV2{}
	return &this
}

// NewCreateResultsRequestV2WithDefaults instantiates a new CreateResultsRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateResultsRequestV2WithDefaults() *CreateResultsRequestV2 {
	this := CreateResultsRequestV2{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CreateResultsRequestV2) GetResults() []ResultCreate {
	if o == nil || IsNil(o.Results) {
		var ret []ResultCreate
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResultsRequestV2) GetResultsOk() ([]ResultCreate, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CreateResultsRequestV2) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ResultCreate and assigns it to the Results field.
func (o *CreateResultsRequestV2) SetResults(v []ResultCreate) {
	o.Results = v
}

func (o CreateResultsRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateResultsRequestV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableCreateResultsRequestV2 struct {
	value *CreateResultsRequestV2
	isSet bool
}

func (v NullableCreateResultsRequestV2) Get() *CreateResultsRequestV2 {
	return v.value
}

func (v *NullableCreateResultsRequestV2) Set(val *CreateResultsRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateResultsRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateResultsRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateResultsRequestV2(val *CreateResultsRequestV2) *NullableCreateResultsRequestV2 {
	return &NullableCreateResultsRequestV2{value: val, isSet: true}
}

func (v NullableCreateResultsRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateResultsRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}