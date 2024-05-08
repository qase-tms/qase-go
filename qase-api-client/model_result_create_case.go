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

// checks if the ResultCreateCase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultCreateCase{}

// ResultCreateCase Could be used instead of `case_id`.
type ResultCreateCase struct {
	Title *string `json:"title,omitempty"`
	// Nested suites should be separated with `TAB` symbol.
	SuiteTitle     NullableString `json:"suite_title,omitempty"`
	Description    NullableString `json:"description,omitempty"`
	Preconditions  NullableString `json:"preconditions,omitempty"`
	Postconditions NullableString `json:"postconditions,omitempty"`
	// Slug of the layer. You can get it in the System Field settings.
	Layer *string `json:"layer,omitempty"`
	// Slug of the severity. You can get it in the System Field settings.
	Severity *string `json:"severity,omitempty"`
	// Slug of the priority. You can get it in the System Field settings.
	Priority *string `json:"priority,omitempty"`
}

// NewResultCreateCase instantiates a new ResultCreateCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultCreateCase() *ResultCreateCase {
	this := ResultCreateCase{}
	return &this
}

// NewResultCreateCaseWithDefaults instantiates a new ResultCreateCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultCreateCaseWithDefaults() *ResultCreateCase {
	this := ResultCreateCase{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ResultCreateCase) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateCase) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ResultCreateCase) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ResultCreateCase) SetTitle(v string) {
	o.Title = &v
}

// GetSuiteTitle returns the SuiteTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreateCase) GetSuiteTitle() string {
	if o == nil || IsNil(o.SuiteTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SuiteTitle.Get()
}

// GetSuiteTitleOk returns a tuple with the SuiteTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreateCase) GetSuiteTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuiteTitle.Get(), o.SuiteTitle.IsSet()
}

// HasSuiteTitle returns a boolean if a field has been set.
func (o *ResultCreateCase) HasSuiteTitle() bool {
	if o != nil && o.SuiteTitle.IsSet() {
		return true
	}

	return false
}

// SetSuiteTitle gets a reference to the given NullableString and assigns it to the SuiteTitle field.
func (o *ResultCreateCase) SetSuiteTitle(v string) {
	o.SuiteTitle.Set(&v)
}

// SetSuiteTitleNil sets the value for SuiteTitle to be an explicit nil
func (o *ResultCreateCase) SetSuiteTitleNil() {
	o.SuiteTitle.Set(nil)
}

// UnsetSuiteTitle ensures that no value is present for SuiteTitle, not even an explicit nil
func (o *ResultCreateCase) UnsetSuiteTitle() {
	o.SuiteTitle.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreateCase) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreateCase) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ResultCreateCase) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ResultCreateCase) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ResultCreateCase) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ResultCreateCase) UnsetDescription() {
	o.Description.Unset()
}

// GetPreconditions returns the Preconditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreateCase) GetPreconditions() string {
	if o == nil || IsNil(o.Preconditions.Get()) {
		var ret string
		return ret
	}
	return *o.Preconditions.Get()
}

// GetPreconditionsOk returns a tuple with the Preconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreateCase) GetPreconditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preconditions.Get(), o.Preconditions.IsSet()
}

// HasPreconditions returns a boolean if a field has been set.
func (o *ResultCreateCase) HasPreconditions() bool {
	if o != nil && o.Preconditions.IsSet() {
		return true
	}

	return false
}

// SetPreconditions gets a reference to the given NullableString and assigns it to the Preconditions field.
func (o *ResultCreateCase) SetPreconditions(v string) {
	o.Preconditions.Set(&v)
}

// SetPreconditionsNil sets the value for Preconditions to be an explicit nil
func (o *ResultCreateCase) SetPreconditionsNil() {
	o.Preconditions.Set(nil)
}

// UnsetPreconditions ensures that no value is present for Preconditions, not even an explicit nil
func (o *ResultCreateCase) UnsetPreconditions() {
	o.Preconditions.Unset()
}

// GetPostconditions returns the Postconditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreateCase) GetPostconditions() string {
	if o == nil || IsNil(o.Postconditions.Get()) {
		var ret string
		return ret
	}
	return *o.Postconditions.Get()
}

// GetPostconditionsOk returns a tuple with the Postconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreateCase) GetPostconditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postconditions.Get(), o.Postconditions.IsSet()
}

// HasPostconditions returns a boolean if a field has been set.
func (o *ResultCreateCase) HasPostconditions() bool {
	if o != nil && o.Postconditions.IsSet() {
		return true
	}

	return false
}

// SetPostconditions gets a reference to the given NullableString and assigns it to the Postconditions field.
func (o *ResultCreateCase) SetPostconditions(v string) {
	o.Postconditions.Set(&v)
}

// SetPostconditionsNil sets the value for Postconditions to be an explicit nil
func (o *ResultCreateCase) SetPostconditionsNil() {
	o.Postconditions.Set(nil)
}

// UnsetPostconditions ensures that no value is present for Postconditions, not even an explicit nil
func (o *ResultCreateCase) UnsetPostconditions() {
	o.Postconditions.Unset()
}

// GetLayer returns the Layer field value if set, zero value otherwise.
func (o *ResultCreateCase) GetLayer() string {
	if o == nil || IsNil(o.Layer) {
		var ret string
		return ret
	}
	return *o.Layer
}

// GetLayerOk returns a tuple with the Layer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateCase) GetLayerOk() (*string, bool) {
	if o == nil || IsNil(o.Layer) {
		return nil, false
	}
	return o.Layer, true
}

// HasLayer returns a boolean if a field has been set.
func (o *ResultCreateCase) HasLayer() bool {
	if o != nil && !IsNil(o.Layer) {
		return true
	}

	return false
}

// SetLayer gets a reference to the given string and assigns it to the Layer field.
func (o *ResultCreateCase) SetLayer(v string) {
	o.Layer = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ResultCreateCase) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateCase) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ResultCreateCase) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ResultCreateCase) SetSeverity(v string) {
	o.Severity = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ResultCreateCase) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateCase) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ResultCreateCase) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *ResultCreateCase) SetPriority(v string) {
	o.Priority = &v
}

func (o ResultCreateCase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultCreateCase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.SuiteTitle.IsSet() {
		toSerialize["suite_title"] = o.SuiteTitle.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Preconditions.IsSet() {
		toSerialize["preconditions"] = o.Preconditions.Get()
	}
	if o.Postconditions.IsSet() {
		toSerialize["postconditions"] = o.Postconditions.Get()
	}
	if !IsNil(o.Layer) {
		toSerialize["layer"] = o.Layer
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return toSerialize, nil
}

type NullableResultCreateCase struct {
	value *ResultCreateCase
	isSet bool
}

func (v NullableResultCreateCase) Get() *ResultCreateCase {
	return v.value
}

func (v *NullableResultCreateCase) Set(val *ResultCreateCase) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCreateCase) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCreateCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCreateCase(val *ResultCreateCase) *NullableResultCreateCase {
	return &NullableResultCreateCase{value: val, isSet: true}
}

func (v NullableResultCreateCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCreateCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
