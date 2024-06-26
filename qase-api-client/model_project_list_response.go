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

// checks if the ProjectListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectListResponse{}

// ProjectListResponse struct for ProjectListResponse
type ProjectListResponse struct {
	Status *bool                           `json:"status,omitempty"`
	Result *ProjectListResponseAllOfResult `json:"result,omitempty"`
}

// NewProjectListResponse instantiates a new ProjectListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectListResponse() *ProjectListResponse {
	this := ProjectListResponse{}
	return &this
}

// NewProjectListResponseWithDefaults instantiates a new ProjectListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListResponseWithDefaults() *ProjectListResponse {
	this := ProjectListResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProjectListResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProjectListResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *ProjectListResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ProjectListResponse) GetResult() ProjectListResponseAllOfResult {
	if o == nil || IsNil(o.Result) {
		var ret ProjectListResponseAllOfResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResponse) GetResultOk() (*ProjectListResponseAllOfResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ProjectListResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ProjectListResponseAllOfResult and assigns it to the Result field.
func (o *ProjectListResponse) SetResult(v ProjectListResponseAllOfResult) {
	o.Result = &v
}

func (o ProjectListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableProjectListResponse struct {
	value *ProjectListResponse
	isSet bool
}

func (v NullableProjectListResponse) Get() *ProjectListResponse {
	return v.value
}

func (v *NullableProjectListResponse) Set(val *ProjectListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectListResponse(val *ProjectListResponse) *NullableProjectListResponse {
	return &NullableProjectListResponse{value: val, isSet: true}
}

func (v NullableProjectListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
