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

// checks if the RunEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunEnvironment{}

// RunEnvironment struct for RunEnvironment
type RunEnvironment struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Host        *string `json:"host,omitempty"`
}

// NewRunEnvironment instantiates a new RunEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunEnvironment() *RunEnvironment {
	this := RunEnvironment{}
	return &this
}

// NewRunEnvironmentWithDefaults instantiates a new RunEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunEnvironmentWithDefaults() *RunEnvironment {
	this := RunEnvironment{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RunEnvironment) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunEnvironment) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RunEnvironment) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RunEnvironment) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RunEnvironment) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunEnvironment) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RunEnvironment) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RunEnvironment) SetDescription(v string) {
	o.Description = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *RunEnvironment) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunEnvironment) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *RunEnvironment) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *RunEnvironment) SetSlug(v string) {
	o.Slug = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *RunEnvironment) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunEnvironment) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *RunEnvironment) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *RunEnvironment) SetHost(v string) {
	o.Host = &v
}

func (o RunEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	return toSerialize, nil
}

type NullableRunEnvironment struct {
	value *RunEnvironment
	isSet bool
}

func (v NullableRunEnvironment) Get() *RunEnvironment {
	return v.value
}

func (v *NullableRunEnvironment) Set(val *RunEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableRunEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableRunEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunEnvironment(val *RunEnvironment) *NullableRunEnvironment {
	return &NullableRunEnvironment{value: val, isSet: true}
}

func (v NullableRunEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}