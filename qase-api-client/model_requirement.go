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
	"time"
)

// checks if the Requirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Requirement{}

// Requirement struct for Requirement
type Requirement struct {
	Id            *int64         `json:"id,omitempty"`
	RequirementId *int64         `json:"requirement_id,omitempty"`
	ParentId      NullableInt64  `json:"parent_id,omitempty"`
	MemberId      *int64         `json:"member_id,omitempty"`
	Title         *string        `json:"title,omitempty"`
	Description   NullableString `json:"description,omitempty"`
	Status        *string        `json:"status,omitempty"`
	Type          *string        `json:"type,omitempty"`
	CreatedAt     *time.Time     `json:"created_at,omitempty"`
	UpdatedAt     NullableTime   `json:"updated_at,omitempty"`
}

// NewRequirement instantiates a new Requirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequirement() *Requirement {
	this := Requirement{}
	return &this
}

// NewRequirementWithDefaults instantiates a new Requirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequirementWithDefaults() *Requirement {
	this := Requirement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Requirement) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requirement) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Requirement) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Requirement) SetId(v int64) {
	o.Id = &v
}

// GetRequirementId returns the RequirementId field value if set, zero value otherwise.
func (o *Requirement) GetRequirementId() int64 {
	if o == nil || IsNil(o.RequirementId) {
		var ret int64
		return ret
	}
	return *o.RequirementId
}

// GetRequirementIdOk returns a tuple with the RequirementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requirement) GetRequirementIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RequirementId) {
		return nil, false
	}
	return o.RequirementId, true
}

// HasRequirementId returns a boolean if a field has been set.
func (o *Requirement) HasRequirementId() bool {
	if o != nil && !IsNil(o.RequirementId) {
		return true
	}

	return false
}

// SetRequirementId gets a reference to the given int64 and assigns it to the RequirementId field.
func (o *Requirement) SetRequirementId(v int64) {
	o.RequirementId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Requirement) GetParentId() int64 {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret int64
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Requirement) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *Requirement) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableInt64 and assigns it to the ParentId field.
func (o *Requirement) SetParentId(v int64) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *Requirement) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *Requirement) UnsetParentId() {
	o.ParentId.Unset()
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *Requirement) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requirement) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *Requirement) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
func (o *Requirement) SetMemberId(v int64) {
	o.MemberId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Requirement) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requirement) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Requirement) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Requirement) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Requirement) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Requirement) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Requirement) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Requirement) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Requirement) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Requirement) UnsetDescription() {
	o.Description.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Requirement) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requirement) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Requirement) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Requirement) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Requirement) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requirement) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Requirement) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Requirement) SetType(v string) {
	o.Type = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Requirement) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Requirement) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Requirement) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Requirement) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Requirement) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Requirement) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Requirement) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *Requirement) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Requirement) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Requirement) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o Requirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Requirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RequirementId) {
		toSerialize["requirement_id"] = o.RequirementId
	}
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}
	if !IsNil(o.MemberId) {
		toSerialize["member_id"] = o.MemberId
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	return toSerialize, nil
}

type NullableRequirement struct {
	value *Requirement
	isSet bool
}

func (v NullableRequirement) Get() *Requirement {
	return v.value
}

func (v *NullableRequirement) Set(val *Requirement) {
	v.value = val
	v.isSet = true
}

func (v NullableRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequirement(val *Requirement) *NullableRequirement {
	return &NullableRequirement{value: val, isSet: true}
}

func (v NullableRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
