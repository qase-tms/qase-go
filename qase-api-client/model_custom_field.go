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

// checks if the CustomField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomField{}

// CustomField struct for CustomField
type CustomField struct {
	Id                      *int64         `json:"id,omitempty"`
	Title                   *string        `json:"title,omitempty"`
	Entity                  *string        `json:"entity,omitempty"`
	Type                    *string        `json:"type,omitempty"`
	Placeholder             NullableString `json:"placeholder,omitempty"`
	DefaultValue            NullableString `json:"default_value,omitempty"`
	Value                   NullableString `json:"value,omitempty"`
	IsRequired              *bool          `json:"is_required,omitempty"`
	IsVisible               *bool          `json:"is_visible,omitempty"`
	IsFilterable            *bool          `json:"is_filterable,omitempty"`
	IsEnabledForAllProjects *bool          `json:"is_enabled_for_all_projects,omitempty"`
	// Deprecated, use the `created_at` property instead.
	// Deprecated
	Created *string `json:"created,omitempty"`
	// Deprecated, use the `updated_at` property instead.
	// Deprecated
	Updated       NullableString `json:"updated,omitempty"`
	CreatedAt     *time.Time     `json:"created_at,omitempty"`
	UpdatedAt     *time.Time     `json:"updated_at,omitempty"`
	ProjectsCodes []string       `json:"projects_codes,omitempty"`
}

// NewCustomField instantiates a new CustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomField() *CustomField {
	this := CustomField{}
	return &this
}

// NewCustomFieldWithDefaults instantiates a new CustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldWithDefaults() *CustomField {
	this := CustomField{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomField) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomField) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CustomField) SetId(v int64) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CustomField) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CustomField) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CustomField) SetTitle(v string) {
	o.Title = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *CustomField) GetEntity() string {
	if o == nil || IsNil(o.Entity) {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetEntityOk() (*string, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *CustomField) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *CustomField) SetEntity(v string) {
	o.Entity = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomField) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomField) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomField) SetType(v string) {
	o.Type = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomField) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomField) GetPlaceholderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *CustomField) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *CustomField) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}

// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *CustomField) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *CustomField) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomField) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomField) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *CustomField) HasDefaultValue() bool {
	if o != nil && o.DefaultValue.IsSet() {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given NullableString and assigns it to the DefaultValue field.
func (o *CustomField) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}

// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil
func (o *CustomField) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
func (o *CustomField) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomField) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomField) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *CustomField) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *CustomField) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *CustomField) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *CustomField) UnsetValue() {
	o.Value.Unset()
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *CustomField) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *CustomField) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *CustomField) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetIsVisible returns the IsVisible field value if set, zero value otherwise.
func (o *CustomField) GetIsVisible() bool {
	if o == nil || IsNil(o.IsVisible) {
		var ret bool
		return ret
	}
	return *o.IsVisible
}

// GetIsVisibleOk returns a tuple with the IsVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetIsVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVisible) {
		return nil, false
	}
	return o.IsVisible, true
}

// HasIsVisible returns a boolean if a field has been set.
func (o *CustomField) HasIsVisible() bool {
	if o != nil && !IsNil(o.IsVisible) {
		return true
	}

	return false
}

// SetIsVisible gets a reference to the given bool and assigns it to the IsVisible field.
func (o *CustomField) SetIsVisible(v bool) {
	o.IsVisible = &v
}

// GetIsFilterable returns the IsFilterable field value if set, zero value otherwise.
func (o *CustomField) GetIsFilterable() bool {
	if o == nil || IsNil(o.IsFilterable) {
		var ret bool
		return ret
	}
	return *o.IsFilterable
}

// GetIsFilterableOk returns a tuple with the IsFilterable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetIsFilterableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFilterable) {
		return nil, false
	}
	return o.IsFilterable, true
}

// HasIsFilterable returns a boolean if a field has been set.
func (o *CustomField) HasIsFilterable() bool {
	if o != nil && !IsNil(o.IsFilterable) {
		return true
	}

	return false
}

// SetIsFilterable gets a reference to the given bool and assigns it to the IsFilterable field.
func (o *CustomField) SetIsFilterable(v bool) {
	o.IsFilterable = &v
}

// GetIsEnabledForAllProjects returns the IsEnabledForAllProjects field value if set, zero value otherwise.
func (o *CustomField) GetIsEnabledForAllProjects() bool {
	if o == nil || IsNil(o.IsEnabledForAllProjects) {
		var ret bool
		return ret
	}
	return *o.IsEnabledForAllProjects
}

// GetIsEnabledForAllProjectsOk returns a tuple with the IsEnabledForAllProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetIsEnabledForAllProjectsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabledForAllProjects) {
		return nil, false
	}
	return o.IsEnabledForAllProjects, true
}

// HasIsEnabledForAllProjects returns a boolean if a field has been set.
func (o *CustomField) HasIsEnabledForAllProjects() bool {
	if o != nil && !IsNil(o.IsEnabledForAllProjects) {
		return true
	}

	return false
}

// SetIsEnabledForAllProjects gets a reference to the given bool and assigns it to the IsEnabledForAllProjects field.
func (o *CustomField) SetIsEnabledForAllProjects(v bool) {
	o.IsEnabledForAllProjects = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
// Deprecated
func (o *CustomField) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CustomField) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CustomField) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
// Deprecated
func (o *CustomField) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *CustomField) GetUpdated() string {
	if o == nil || IsNil(o.Updated.Get()) {
		var ret string
		return ret
	}
	return *o.Updated.Get()
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *CustomField) GetUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated.Get(), o.Updated.IsSet()
}

// HasUpdated returns a boolean if a field has been set.
func (o *CustomField) HasUpdated() bool {
	if o != nil && o.Updated.IsSet() {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given NullableString and assigns it to the Updated field.
// Deprecated
func (o *CustomField) SetUpdated(v string) {
	o.Updated.Set(&v)
}

// SetUpdatedNil sets the value for Updated to be an explicit nil
func (o *CustomField) SetUpdatedNil() {
	o.Updated.Set(nil)
}

// UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
func (o *CustomField) UnsetUpdated() {
	o.Updated.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomField) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomField) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomField) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomField) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomField) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomField) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetProjectsCodes returns the ProjectsCodes field value if set, zero value otherwise.
func (o *CustomField) GetProjectsCodes() []string {
	if o == nil || IsNil(o.ProjectsCodes) {
		var ret []string
		return ret
	}
	return o.ProjectsCodes
}

// GetProjectsCodesOk returns a tuple with the ProjectsCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetProjectsCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectsCodes) {
		return nil, false
	}
	return o.ProjectsCodes, true
}

// HasProjectsCodes returns a boolean if a field has been set.
func (o *CustomField) HasProjectsCodes() bool {
	if o != nil && !IsNil(o.ProjectsCodes) {
		return true
	}

	return false
}

// SetProjectsCodes gets a reference to the given []string and assigns it to the ProjectsCodes field.
func (o *CustomField) SetProjectsCodes(v []string) {
	o.ProjectsCodes = v
}

func (o CustomField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if o.DefaultValue.IsSet() {
		toSerialize["default_value"] = o.DefaultValue.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if !IsNil(o.IsRequired) {
		toSerialize["is_required"] = o.IsRequired
	}
	if !IsNil(o.IsVisible) {
		toSerialize["is_visible"] = o.IsVisible
	}
	if !IsNil(o.IsFilterable) {
		toSerialize["is_filterable"] = o.IsFilterable
	}
	if !IsNil(o.IsEnabledForAllProjects) {
		toSerialize["is_enabled_for_all_projects"] = o.IsEnabledForAllProjects
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if o.Updated.IsSet() {
		toSerialize["updated"] = o.Updated.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ProjectsCodes) {
		toSerialize["projects_codes"] = o.ProjectsCodes
	}
	return toSerialize, nil
}

type NullableCustomField struct {
	value *CustomField
	isSet bool
}

func (v NullableCustomField) Get() *CustomField {
	return v.value
}

func (v *NullableCustomField) Set(val *CustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomField(val *CustomField) *NullableCustomField {
	return &NullableCustomField{value: val, isSet: true}
}

func (v NullableCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}