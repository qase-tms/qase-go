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

// checks if the ConfigurationListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationListResponse{}

// ConfigurationListResponse struct for ConfigurationListResponse
type ConfigurationListResponse struct {
	Status *bool                                 `json:"status,omitempty"`
	Result *ConfigurationListResponseAllOfResult `json:"result,omitempty"`
}

// NewConfigurationListResponse instantiates a new ConfigurationListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationListResponse() *ConfigurationListResponse {
	this := ConfigurationListResponse{}
	return &this
}

// NewConfigurationListResponseWithDefaults instantiates a new ConfigurationListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationListResponseWithDefaults() *ConfigurationListResponse {
	this := ConfigurationListResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConfigurationListResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationListResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConfigurationListResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *ConfigurationListResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigurationListResponse) GetResult() ConfigurationListResponseAllOfResult {
	if o == nil || IsNil(o.Result) {
		var ret ConfigurationListResponseAllOfResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationListResponse) GetResultOk() (*ConfigurationListResponseAllOfResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigurationListResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigurationListResponseAllOfResult and assigns it to the Result field.
func (o *ConfigurationListResponse) SetResult(v ConfigurationListResponseAllOfResult) {
	o.Result = &v
}

func (o ConfigurationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigurationListResponse struct {
	value *ConfigurationListResponse
	isSet bool
}

func (v NullableConfigurationListResponse) Get() *ConfigurationListResponse {
	return v.value
}

func (v *NullableConfigurationListResponse) Set(val *ConfigurationListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationListResponse(val *ConfigurationListResponse) *NullableConfigurationListResponse {
	return &NullableConfigurationListResponse{value: val, isSet: true}
}

func (v NullableConfigurationListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
