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
	"time"
)

// checks if the QqlDefect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QqlDefect{}

// QqlDefect struct for QqlDefect
type QqlDefect struct {
	Id           *int64             `json:"id,omitempty"`
	DefectId     int64              `json:"defect_id"`
	Title        *string            `json:"title,omitempty"`
	ActualResult *string            `json:"actual_result,omitempty"`
	Severity     *string            `json:"severity,omitempty"`
	Status       *string            `json:"status,omitempty"`
	MilestoneId  NullableInt64      `json:"milestone_id,omitempty"`
	CustomFields []CustomFieldValue `json:"custom_fields,omitempty"`
	Attachments  []Attachment       `json:"attachments,omitempty"`
	Resolved     NullableTime       `json:"resolved,omitempty"`
	// Deprecated, use `author_id` instead.
	// Deprecated
	MemberId     *int64     `json:"member_id,omitempty"`
	AuthorId     *int64     `json:"author_id,omitempty"`
	ExternalData *string    `json:"external_data,omitempty"`
	Tags         []TagValue `json:"tags,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

type _QqlDefect QqlDefect

// NewQqlDefect instantiates a new QqlDefect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQqlDefect(defectId int64) *QqlDefect {
	this := QqlDefect{}
	this.DefectId = defectId
	return &this
}

// NewQqlDefectWithDefaults instantiates a new QqlDefect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQqlDefectWithDefaults() *QqlDefect {
	this := QqlDefect{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QqlDefect) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QqlDefect) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *QqlDefect) SetId(v int64) {
	o.Id = &v
}

// GetDefectId returns the DefectId field value
func (o *QqlDefect) GetDefectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DefectId
}

// GetDefectIdOk returns a tuple with the DefectId field value
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetDefectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefectId, true
}

// SetDefectId sets field value
func (o *QqlDefect) SetDefectId(v int64) {
	o.DefectId = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *QqlDefect) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *QqlDefect) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *QqlDefect) SetTitle(v string) {
	o.Title = &v
}

// GetActualResult returns the ActualResult field value if set, zero value otherwise.
func (o *QqlDefect) GetActualResult() string {
	if o == nil || IsNil(o.ActualResult) {
		var ret string
		return ret
	}
	return *o.ActualResult
}

// GetActualResultOk returns a tuple with the ActualResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetActualResultOk() (*string, bool) {
	if o == nil || IsNil(o.ActualResult) {
		return nil, false
	}
	return o.ActualResult, true
}

// HasActualResult returns a boolean if a field has been set.
func (o *QqlDefect) HasActualResult() bool {
	if o != nil && !IsNil(o.ActualResult) {
		return true
	}

	return false
}

// SetActualResult gets a reference to the given string and assigns it to the ActualResult field.
func (o *QqlDefect) SetActualResult(v string) {
	o.ActualResult = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *QqlDefect) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *QqlDefect) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *QqlDefect) SetSeverity(v string) {
	o.Severity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QqlDefect) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QqlDefect) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QqlDefect) SetStatus(v string) {
	o.Status = &v
}

// GetMilestoneId returns the MilestoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QqlDefect) GetMilestoneId() int64 {
	if o == nil || IsNil(o.MilestoneId.Get()) {
		var ret int64
		return ret
	}
	return *o.MilestoneId.Get()
}

// GetMilestoneIdOk returns a tuple with the MilestoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QqlDefect) GetMilestoneIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MilestoneId.Get(), o.MilestoneId.IsSet()
}

// HasMilestoneId returns a boolean if a field has been set.
func (o *QqlDefect) HasMilestoneId() bool {
	if o != nil && o.MilestoneId.IsSet() {
		return true
	}

	return false
}

// SetMilestoneId gets a reference to the given NullableInt64 and assigns it to the MilestoneId field.
func (o *QqlDefect) SetMilestoneId(v int64) {
	o.MilestoneId.Set(&v)
}

// SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil
func (o *QqlDefect) SetMilestoneIdNil() {
	o.MilestoneId.Set(nil)
}

// UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
func (o *QqlDefect) UnsetMilestoneId() {
	o.MilestoneId.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *QqlDefect) GetCustomFields() []CustomFieldValue {
	if o == nil || IsNil(o.CustomFields) {
		var ret []CustomFieldValue
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetCustomFieldsOk() ([]CustomFieldValue, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *QqlDefect) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomFieldValue and assigns it to the CustomFields field.
func (o *QqlDefect) SetCustomFields(v []CustomFieldValue) {
	o.CustomFields = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *QqlDefect) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *QqlDefect) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *QqlDefect) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QqlDefect) GetResolved() time.Time {
	if o == nil || IsNil(o.Resolved.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QqlDefect) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *QqlDefect) HasResolved() bool {
	if o != nil && o.Resolved.IsSet() {
		return true
	}

	return false
}

// SetResolved gets a reference to the given NullableTime and assigns it to the Resolved field.
func (o *QqlDefect) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil
func (o *QqlDefect) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
func (o *QqlDefect) UnsetResolved() {
	o.Resolved.Unset()
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
// Deprecated
func (o *QqlDefect) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *QqlDefect) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *QqlDefect) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
// Deprecated
func (o *QqlDefect) SetMemberId(v int64) {
	o.MemberId = &v
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *QqlDefect) GetAuthorId() int64 {
	if o == nil || IsNil(o.AuthorId) {
		var ret int64
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetAuthorIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthorId) {
		return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *QqlDefect) HasAuthorId() bool {
	if o != nil && !IsNil(o.AuthorId) {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given int64 and assigns it to the AuthorId field.
func (o *QqlDefect) SetAuthorId(v int64) {
	o.AuthorId = &v
}

// GetExternalData returns the ExternalData field value if set, zero value otherwise.
func (o *QqlDefect) GetExternalData() string {
	if o == nil || IsNil(o.ExternalData) {
		var ret string
		return ret
	}
	return *o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetExternalDataOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalData) {
		return nil, false
	}
	return o.ExternalData, true
}

// HasExternalData returns a boolean if a field has been set.
func (o *QqlDefect) HasExternalData() bool {
	if o != nil && !IsNil(o.ExternalData) {
		return true
	}

	return false
}

// SetExternalData gets a reference to the given string and assigns it to the ExternalData field.
func (o *QqlDefect) SetExternalData(v string) {
	o.ExternalData = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *QqlDefect) GetTags() []TagValue {
	if o == nil || IsNil(o.Tags) {
		var ret []TagValue
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetTagsOk() ([]TagValue, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *QqlDefect) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagValue and assigns it to the Tags field.
func (o *QqlDefect) SetTags(v []TagValue) {
	o.Tags = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *QqlDefect) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *QqlDefect) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *QqlDefect) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *QqlDefect) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlDefect) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *QqlDefect) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *QqlDefect) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o QqlDefect) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QqlDefect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["defect_id"] = o.DefectId
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.ActualResult) {
		toSerialize["actual_result"] = o.ActualResult
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.MilestoneId.IsSet() {
		toSerialize["milestone_id"] = o.MilestoneId.Get()
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Resolved.IsSet() {
		toSerialize["resolved"] = o.Resolved.Get()
	}
	if !IsNil(o.MemberId) {
		toSerialize["member_id"] = o.MemberId
	}
	if !IsNil(o.AuthorId) {
		toSerialize["author_id"] = o.AuthorId
	}
	if !IsNil(o.ExternalData) {
		toSerialize["external_data"] = o.ExternalData
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *QqlDefect) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defect_id",
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

	varQqlDefect := _QqlDefect{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQqlDefect)

	if err != nil {
		return err
	}

	*o = QqlDefect(varQqlDefect)

	return err
}

type NullableQqlDefect struct {
	value *QqlDefect
	isSet bool
}

func (v NullableQqlDefect) Get() *QqlDefect {
	return v.value
}

func (v *NullableQqlDefect) Set(val *QqlDefect) {
	v.value = val
	v.isSet = true
}

func (v NullableQqlDefect) IsSet() bool {
	return v.isSet
}

func (v *NullableQqlDefect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQqlDefect(val *QqlDefect) *NullableQqlDefect {
	return &NullableQqlDefect{value: val, isSet: true}
}

func (v NullableQqlDefect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQqlDefect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
