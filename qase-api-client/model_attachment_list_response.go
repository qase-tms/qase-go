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

// checks if the AttachmentListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentListResponse{}

// AttachmentListResponse struct for AttachmentListResponse
type AttachmentListResponse struct {
	Status *bool                              `json:"status,omitempty"`
	Result *AttachmentListResponseAllOfResult `json:"result,omitempty"`
}

// NewAttachmentListResponse instantiates a new AttachmentListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentListResponse() *AttachmentListResponse {
	this := AttachmentListResponse{}
	return &this
}

// NewAttachmentListResponseWithDefaults instantiates a new AttachmentListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentListResponseWithDefaults() *AttachmentListResponse {
	this := AttachmentListResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AttachmentListResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentListResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AttachmentListResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *AttachmentListResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AttachmentListResponse) GetResult() AttachmentListResponseAllOfResult {
	if o == nil || IsNil(o.Result) {
		var ret AttachmentListResponseAllOfResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentListResponse) GetResultOk() (*AttachmentListResponseAllOfResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AttachmentListResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AttachmentListResponseAllOfResult and assigns it to the Result field.
func (o *AttachmentListResponse) SetResult(v AttachmentListResponseAllOfResult) {
	o.Result = &v
}

func (o AttachmentListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableAttachmentListResponse struct {
	value *AttachmentListResponse
	isSet bool
}

func (v NullableAttachmentListResponse) Get() *AttachmentListResponse {
	return v.value
}

func (v *NullableAttachmentListResponse) Set(val *AttachmentListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentListResponse(val *AttachmentListResponse) *NullableAttachmentListResponse {
	return &NullableAttachmentListResponse{value: val, isSet: true}
}

func (v NullableAttachmentListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
