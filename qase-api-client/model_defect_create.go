/*
Qase.io TestOps API v1

Qase TestOps API v1 Specification.

API version: 1.0.0
Contact: support@qase.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v1_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DefectCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefectCreate{}

// DefectCreate struct for DefectCreate
type DefectCreate struct {
	Title        string        `json:"title"`
	ActualResult string        `json:"actual_result"`
	Severity     int32         `json:"severity"`
	MilestoneId  NullableInt64 `json:"milestone_id,omitempty"`
	Attachments  []string      `json:"attachments,omitempty"`
	// A map of custom fields values (id => value)
	CustomField *map[string]string `json:"custom_field,omitempty"`
	Tags        []string           `json:"tags,omitempty"`
}

type _DefectCreate DefectCreate

// NewDefectCreate instantiates a new DefectCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefectCreate(title string, actualResult string, severity int32) *DefectCreate {
	this := DefectCreate{}
	this.Title = title
	this.ActualResult = actualResult
	this.Severity = severity
	return &this
}

// NewDefectCreateWithDefaults instantiates a new DefectCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefectCreateWithDefaults() *DefectCreate {
	this := DefectCreate{}
	return &this
}

// GetTitle returns the Title field value
func (o *DefectCreate) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DefectCreate) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DefectCreate) SetTitle(v string) {
	o.Title = v
}

// GetActualResult returns the ActualResult field value
func (o *DefectCreate) GetActualResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActualResult
}

// GetActualResultOk returns a tuple with the ActualResult field value
// and a boolean to check if the value has been set.
func (o *DefectCreate) GetActualResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActualResult, true
}

// SetActualResult sets field value
func (o *DefectCreate) SetActualResult(v string) {
	o.ActualResult = v
}

// GetSeverity returns the Severity field value
func (o *DefectCreate) GetSeverity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *DefectCreate) GetSeverityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *DefectCreate) SetSeverity(v int32) {
	o.Severity = v
}

// GetMilestoneId returns the MilestoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefectCreate) GetMilestoneId() int64 {
	if o == nil || IsNil(o.MilestoneId.Get()) {
		var ret int64
		return ret
	}
	return *o.MilestoneId.Get()
}

// GetMilestoneIdOk returns a tuple with the MilestoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefectCreate) GetMilestoneIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MilestoneId.Get(), o.MilestoneId.IsSet()
}

// HasMilestoneId returns a boolean if a field has been set.
func (o *DefectCreate) HasMilestoneId() bool {
	if o != nil && o.MilestoneId.IsSet() {
		return true
	}

	return false
}

// SetMilestoneId gets a reference to the given NullableInt64 and assigns it to the MilestoneId field.
func (o *DefectCreate) SetMilestoneId(v int64) {
	o.MilestoneId.Set(&v)
}

// SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil
func (o *DefectCreate) SetMilestoneIdNil() {
	o.MilestoneId.Set(nil)
}

// UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
func (o *DefectCreate) UnsetMilestoneId() {
	o.MilestoneId.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *DefectCreate) GetAttachments() []string {
	if o == nil || IsNil(o.Attachments) {
		var ret []string
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectCreate) GetAttachmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *DefectCreate) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *DefectCreate) SetAttachments(v []string) {
	o.Attachments = v
}

// GetCustomField returns the CustomField field value if set, zero value otherwise.
func (o *DefectCreate) GetCustomField() map[string]string {
	if o == nil || IsNil(o.CustomField) {
		var ret map[string]string
		return ret
	}
	return *o.CustomField
}

// GetCustomFieldOk returns a tuple with the CustomField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectCreate) GetCustomFieldOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomField) {
		return nil, false
	}
	return o.CustomField, true
}

// HasCustomField returns a boolean if a field has been set.
func (o *DefectCreate) HasCustomField() bool {
	if o != nil && !IsNil(o.CustomField) {
		return true
	}

	return false
}

// SetCustomField gets a reference to the given map[string]string and assigns it to the CustomField field.
func (o *DefectCreate) SetCustomField(v map[string]string) {
	o.CustomField = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DefectCreate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectCreate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DefectCreate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DefectCreate) SetTags(v []string) {
	o.Tags = v
}

func (o DefectCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefectCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["actual_result"] = o.ActualResult
	toSerialize["severity"] = o.Severity
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

func (o *DefectCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"actual_result",
		"severity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDefectCreate := _DefectCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDefectCreate)

	if err != nil {
		return err
	}

	*o = DefectCreate(varDefectCreate)

	return err
}

type NullableDefectCreate struct {
	value *DefectCreate
	isSet bool
}

func (v NullableDefectCreate) Get() *DefectCreate {
	return v.value
}

func (v *NullableDefectCreate) Set(val *DefectCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDefectCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDefectCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefectCreate(val *DefectCreate) *NullableDefectCreate {
	return &NullableDefectCreate{value: val, isSet: true}
}

func (v NullableDefectCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefectCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
