/*
Qase.io TestOps API v2

Qase TestOps API v2 Specification.

API version: 2.0.0
Contact: support@qase.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RelationSuiteItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelationSuiteItem{}

// RelationSuiteItem struct for RelationSuiteItem
type RelationSuiteItem struct {
	Title    string        `json:"title"`
	PublicId NullableInt64 `json:"public_id,omitempty"`
}

type _RelationSuiteItem RelationSuiteItem

// NewRelationSuiteItem instantiates a new RelationSuiteItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationSuiteItem(title string) *RelationSuiteItem {
	this := RelationSuiteItem{}
	this.Title = title
	return &this
}

// NewRelationSuiteItemWithDefaults instantiates a new RelationSuiteItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationSuiteItemWithDefaults() *RelationSuiteItem {
	this := RelationSuiteItem{}
	return &this
}

// GetTitle returns the Title field value
func (o *RelationSuiteItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *RelationSuiteItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *RelationSuiteItem) SetTitle(v string) {
	o.Title = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RelationSuiteItem) GetPublicId() int64 {
	if o == nil || IsNil(o.PublicId.Get()) {
		var ret int64
		return ret
	}
	return *o.PublicId.Get()
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RelationSuiteItem) GetPublicIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicId.Get(), o.PublicId.IsSet()
}

// HasPublicId returns a boolean if a field has been set.
func (o *RelationSuiteItem) HasPublicId() bool {
	if o != nil && o.PublicId.IsSet() {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given NullableInt64 and assigns it to the PublicId field.
func (o *RelationSuiteItem) SetPublicId(v int64) {
	o.PublicId.Set(&v)
}

// SetPublicIdNil sets the value for PublicId to be an explicit nil
func (o *RelationSuiteItem) SetPublicIdNil() {
	o.PublicId.Set(nil)
}

// UnsetPublicId ensures that no value is present for PublicId, not even an explicit nil
func (o *RelationSuiteItem) UnsetPublicId() {
	o.PublicId.Unset()
}

func (o RelationSuiteItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelationSuiteItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if o.PublicId.IsSet() {
		toSerialize["public_id"] = o.PublicId.Get()
	}
	return toSerialize, nil
}

func (o *RelationSuiteItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varRelationSuiteItem := _RelationSuiteItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRelationSuiteItem)

	if err != nil {
		return err
	}

	*o = RelationSuiteItem(varRelationSuiteItem)

	return err
}

type NullableRelationSuiteItem struct {
	value *RelationSuiteItem
	isSet bool
}

func (v NullableRelationSuiteItem) Get() *RelationSuiteItem {
	return v.value
}

func (v *NullableRelationSuiteItem) Set(val *RelationSuiteItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationSuiteItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationSuiteItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationSuiteItem(val *RelationSuiteItem) *NullableRelationSuiteItem {
	return &NullableRelationSuiteItem{value: val, isSet: true}
}

func (v NullableRelationSuiteItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationSuiteItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
