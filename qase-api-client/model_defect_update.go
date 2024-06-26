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

// checks if the DefectUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefectUpdate{}

// DefectUpdate struct for DefectUpdate
type DefectUpdate struct {
	Title        *string       `json:"title,omitempty"`
	ActualResult *string       `json:"actual_result,omitempty"`
	Severity     *int32        `json:"severity,omitempty"`
	MilestoneId  NullableInt64 `json:"milestone_id,omitempty"`
	Attachments  []string      `json:"attachments,omitempty"`
	// A map of custom fields values (id => value)
	CustomField *map[string]string `json:"custom_field,omitempty"`
	Tags        []string           `json:"tags,omitempty"`
}

// NewDefectUpdate instantiates a new DefectUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefectUpdate() *DefectUpdate {
	this := DefectUpdate{}
	return &this
}

// NewDefectUpdateWithDefaults instantiates a new DefectUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefectUpdateWithDefaults() *DefectUpdate {
	this := DefectUpdate{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DefectUpdate) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectUpdate) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DefectUpdate) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DefectUpdate) SetTitle(v string) {
	o.Title = &v
}

// GetActualResult returns the ActualResult field value if set, zero value otherwise.
func (o *DefectUpdate) GetActualResult() string {
	if o == nil || IsNil(o.ActualResult) {
		var ret string
		return ret
	}
	return *o.ActualResult
}

// GetActualResultOk returns a tuple with the ActualResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectUpdate) GetActualResultOk() (*string, bool) {
	if o == nil || IsNil(o.ActualResult) {
		return nil, false
	}
	return o.ActualResult, true
}

// HasActualResult returns a boolean if a field has been set.
func (o *DefectUpdate) HasActualResult() bool {
	if o != nil && !IsNil(o.ActualResult) {
		return true
	}

	return false
}

// SetActualResult gets a reference to the given string and assigns it to the ActualResult field.
func (o *DefectUpdate) SetActualResult(v string) {
	o.ActualResult = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *DefectUpdate) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectUpdate) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *DefectUpdate) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *DefectUpdate) SetSeverity(v int32) {
	o.Severity = &v
}

// GetMilestoneId returns the MilestoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefectUpdate) GetMilestoneId() int64 {
	if o == nil || IsNil(o.MilestoneId.Get()) {
		var ret int64
		return ret
	}
	return *o.MilestoneId.Get()
}

// GetMilestoneIdOk returns a tuple with the MilestoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefectUpdate) GetMilestoneIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MilestoneId.Get(), o.MilestoneId.IsSet()
}

// HasMilestoneId returns a boolean if a field has been set.
func (o *DefectUpdate) HasMilestoneId() bool {
	if o != nil && o.MilestoneId.IsSet() {
		return true
	}

	return false
}

// SetMilestoneId gets a reference to the given NullableInt64 and assigns it to the MilestoneId field.
func (o *DefectUpdate) SetMilestoneId(v int64) {
	o.MilestoneId.Set(&v)
}

// SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil
func (o *DefectUpdate) SetMilestoneIdNil() {
	o.MilestoneId.Set(nil)
}

// UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
func (o *DefectUpdate) UnsetMilestoneId() {
	o.MilestoneId.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *DefectUpdate) GetAttachments() []string {
	if o == nil || IsNil(o.Attachments) {
		var ret []string
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectUpdate) GetAttachmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *DefectUpdate) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *DefectUpdate) SetAttachments(v []string) {
	o.Attachments = v
}

// GetCustomField returns the CustomField field value if set, zero value otherwise.
func (o *DefectUpdate) GetCustomField() map[string]string {
	if o == nil || IsNil(o.CustomField) {
		var ret map[string]string
		return ret
	}
	return *o.CustomField
}

// GetCustomFieldOk returns a tuple with the CustomField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectUpdate) GetCustomFieldOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomField) {
		return nil, false
	}
	return o.CustomField, true
}

// HasCustomField returns a boolean if a field has been set.
func (o *DefectUpdate) HasCustomField() bool {
	if o != nil && !IsNil(o.CustomField) {
		return true
	}

	return false
}

// SetCustomField gets a reference to the given map[string]string and assigns it to the CustomField field.
func (o *DefectUpdate) SetCustomField(v map[string]string) {
	o.CustomField = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DefectUpdate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectUpdate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DefectUpdate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DefectUpdate) SetTags(v []string) {
	o.Tags = v
}

func (o DefectUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefectUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.ActualResult) {
		toSerialize["actual_result"] = o.ActualResult
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if o.MilestoneId.IsSet() {
		toSerialize["milestone_id"] = o.MilestoneId.Get()
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.CustomField) {
		toSerialize["custom_field"] = o.CustomField
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableDefectUpdate struct {
	value *DefectUpdate
	isSet bool
}

func (v NullableDefectUpdate) Get() *DefectUpdate {
	return v.value
}

func (v *NullableDefectUpdate) Set(val *DefectUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDefectUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDefectUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefectUpdate(val *DefectUpdate) *NullableDefectUpdate {
	return &NullableDefectUpdate{value: val, isSet: true}
}

func (v NullableDefectUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefectUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
