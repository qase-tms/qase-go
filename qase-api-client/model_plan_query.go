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

// checks if the PlanQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanQuery{}

// PlanQuery struct for PlanQuery
type PlanQuery struct {
	Id          *int64         `json:"id,omitempty"`
	PlanId      int64          `json:"plan_id"`
	Title       *string        `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	CasesCount  *int32         `json:"cases_count,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
}

type _PlanQuery PlanQuery

// NewPlanQuery instantiates a new PlanQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanQuery(planId int64) *PlanQuery {
	this := PlanQuery{}
	this.PlanId = planId
	return &this
}

// NewPlanQueryWithDefaults instantiates a new PlanQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanQueryWithDefaults() *PlanQuery {
	this := PlanQuery{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlanQuery) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanQuery) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlanQuery) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PlanQuery) SetId(v int64) {
	o.Id = &v
}

// GetPlanId returns the PlanId field value
func (o *PlanQuery) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *PlanQuery) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *PlanQuery) SetPlanId(v int64) {
	o.PlanId = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PlanQuery) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanQuery) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PlanQuery) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PlanQuery) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlanQuery) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanQuery) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PlanQuery) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PlanQuery) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PlanQuery) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PlanQuery) UnsetDescription() {
	o.Description.Unset()
}

// GetCasesCount returns the CasesCount field value if set, zero value otherwise.
func (o *PlanQuery) GetCasesCount() int32 {
	if o == nil || IsNil(o.CasesCount) {
		var ret int32
		return ret
	}
	return *o.CasesCount
}

// GetCasesCountOk returns a tuple with the CasesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanQuery) GetCasesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CasesCount) {
		return nil, false
	}
	return o.CasesCount, true
}

// HasCasesCount returns a boolean if a field has been set.
func (o *PlanQuery) HasCasesCount() bool {
	if o != nil && !IsNil(o.CasesCount) {
		return true
	}

	return false
}

// SetCasesCount gets a reference to the given int32 and assigns it to the CasesCount field.
func (o *PlanQuery) SetCasesCount(v int32) {
	o.CasesCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PlanQuery) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanQuery) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PlanQuery) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PlanQuery) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PlanQuery) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanQuery) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PlanQuery) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PlanQuery) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PlanQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["plan_id"] = o.PlanId
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

func (o *PlanQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plan_id",
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

	varPlanQuery := _PlanQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlanQuery)

	if err != nil {
		return err
	}

	*o = PlanQuery(varPlanQuery)

	return err
}

type NullablePlanQuery struct {
	value *PlanQuery
	isSet bool
}

func (v NullablePlanQuery) Get() *PlanQuery {
	return v.value
}

func (v *NullablePlanQuery) Set(val *PlanQuery) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanQuery) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanQuery(val *PlanQuery) *NullablePlanQuery {
	return &NullablePlanQuery{value: val, isSet: true}
}

func (v NullablePlanQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
