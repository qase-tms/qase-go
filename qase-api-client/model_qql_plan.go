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

// checks if the QqlPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QqlPlan{}

// QqlPlan struct for QqlPlan
type QqlPlan struct {
	Id          *int64         `json:"id,omitempty"`
	Title       *string        `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	CasesCount  *int32         `json:"cases_count,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
}

// NewQqlPlan instantiates a new QqlPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQqlPlan() *QqlPlan {
	this := QqlPlan{}
	return &this
}

// NewQqlPlanWithDefaults instantiates a new QqlPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQqlPlanWithDefaults() *QqlPlan {
	this := QqlPlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QqlPlan) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlPlan) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QqlPlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *QqlPlan) SetId(v int64) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *QqlPlan) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlPlan) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *QqlPlan) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *QqlPlan) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QqlPlan) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QqlPlan) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *QqlPlan) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *QqlPlan) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *QqlPlan) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *QqlPlan) UnsetDescription() {
	o.Description.Unset()
}

// GetCasesCount returns the CasesCount field value if set, zero value otherwise.
func (o *QqlPlan) GetCasesCount() int32 {
	if o == nil || IsNil(o.CasesCount) {
		var ret int32
		return ret
	}
	return *o.CasesCount
}

// GetCasesCountOk returns a tuple with the CasesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlPlan) GetCasesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CasesCount) {
		return nil, false
	}
	return o.CasesCount, true
}

// HasCasesCount returns a boolean if a field has been set.
func (o *QqlPlan) HasCasesCount() bool {
	if o != nil && !IsNil(o.CasesCount) {
		return true
	}

	return false
}

// SetCasesCount gets a reference to the given int32 and assigns it to the CasesCount field.
func (o *QqlPlan) SetCasesCount(v int32) {
	o.CasesCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *QqlPlan) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlPlan) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *QqlPlan) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *QqlPlan) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *QqlPlan) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QqlPlan) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *QqlPlan) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *QqlPlan) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o QqlPlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QqlPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.CasesCount) {
		toSerialize["cases_count"] = o.CasesCount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableQqlPlan struct {
	value *QqlPlan
	isSet bool
}

func (v NullableQqlPlan) Get() *QqlPlan {
	return v.value
}

func (v *NullableQqlPlan) Set(val *QqlPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableQqlPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableQqlPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQqlPlan(val *QqlPlan) *NullableQqlPlan {
	return &NullableQqlPlan{value: val, isSet: true}
}

func (v NullableQqlPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQqlPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}