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

// checks if the SharedStepUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedStepUpdate{}

// SharedStepUpdate struct for SharedStepUpdate
type SharedStepUpdate struct {
	Title string `json:"title"`
	// Deprecated, use the `steps` property instead.
	// Deprecated
	Action *string `json:"action,omitempty"`
	// Deprecated, use the `steps` property instead.
	// Deprecated
	ExpectedResult *string `json:"expected_result,omitempty"`
	// Deprecated, use the `steps` property instead.
	// Deprecated
	Data  *string                   `json:"data,omitempty"`
	Steps []SharedStepContentCreate `json:"steps,omitempty"`
}

type _SharedStepUpdate SharedStepUpdate

// NewSharedStepUpdate instantiates a new SharedStepUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedStepUpdate(title string) *SharedStepUpdate {
	this := SharedStepUpdate{}
	this.Title = title
	return &this
}

// NewSharedStepUpdateWithDefaults instantiates a new SharedStepUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedStepUpdateWithDefaults() *SharedStepUpdate {
	this := SharedStepUpdate{}
	return &this
}

// GetTitle returns the Title field value
func (o *SharedStepUpdate) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SharedStepUpdate) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SharedStepUpdate) SetTitle(v string) {
	o.Title = v
}

// GetAction returns the Action field value if set, zero value otherwise.
// Deprecated
func (o *SharedStepUpdate) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SharedStepUpdate) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SharedStepUpdate) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
// Deprecated
func (o *SharedStepUpdate) SetAction(v string) {
	o.Action = &v
}

// GetExpectedResult returns the ExpectedResult field value if set, zero value otherwise.
// Deprecated
func (o *SharedStepUpdate) GetExpectedResult() string {
	if o == nil || IsNil(o.ExpectedResult) {
		var ret string
		return ret
	}
	return *o.ExpectedResult
}

// GetExpectedResultOk returns a tuple with the ExpectedResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SharedStepUpdate) GetExpectedResultOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedResult) {
		return nil, false
	}
	return o.ExpectedResult, true
}

// HasExpectedResult returns a boolean if a field has been set.
func (o *SharedStepUpdate) HasExpectedResult() bool {
	if o != nil && !IsNil(o.ExpectedResult) {
		return true
	}

	return false
}

// SetExpectedResult gets a reference to the given string and assigns it to the ExpectedResult field.
// Deprecated
func (o *SharedStepUpdate) SetExpectedResult(v string) {
	o.ExpectedResult = &v
}

// GetData returns the Data field value if set, zero value otherwise.
// Deprecated
func (o *SharedStepUpdate) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SharedStepUpdate) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SharedStepUpdate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
// Deprecated
func (o *SharedStepUpdate) SetData(v string) {
	o.Data = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *SharedStepUpdate) GetSteps() []SharedStepContentCreate {
	if o == nil || IsNil(o.Steps) {
		var ret []SharedStepContentCreate
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepUpdate) GetStepsOk() ([]SharedStepContentCreate, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SharedStepUpdate) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []SharedStepContentCreate and assigns it to the Steps field.
func (o *SharedStepUpdate) SetSteps(v []SharedStepContentCreate) {
	o.Steps = v
}

func (o SharedStepUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedStepUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ExpectedResult) {
		toSerialize["expected_result"] = o.ExpectedResult
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	return toSerialize, nil
}

func (o *SharedStepUpdate) UnmarshalJSON(data []byte) (err error) {
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

	varSharedStepUpdate := _SharedStepUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSharedStepUpdate)

	if err != nil {
		return err
	}

	*o = SharedStepUpdate(varSharedStepUpdate)

	return err
}

type NullableSharedStepUpdate struct {
	value *SharedStepUpdate
	isSet bool
}

func (v NullableSharedStepUpdate) Get() *SharedStepUpdate {
	return v.value
}

func (v *NullableSharedStepUpdate) Set(val *SharedStepUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedStepUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedStepUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedStepUpdate(val *SharedStepUpdate) *NullableSharedStepUpdate {
	return &NullableSharedStepUpdate{value: val, isSet: true}
}

func (v NullableSharedStepUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedStepUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}