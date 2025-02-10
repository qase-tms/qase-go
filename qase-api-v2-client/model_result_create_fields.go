/*
Qase.io TestOps API v2

Qase TestOps API v2 Specification.

API version: 2.0.0
Contact: support@qase.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_v2_client

import (
	"encoding/json"
)

// checks if the ResultCreateFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultCreateFields{}

// ResultCreateFields struct for ResultCreateFields
type ResultCreateFields struct {
	// Author of the related test case (member id, name or email). If set and test case auto-creation is enabled, the author will be used to create the test case
	Author         *string `json:"author,omitempty"`
	Description    *string `json:"description,omitempty"`
	Preconditions  *string `json:"preconditions,omitempty"`
	Postconditions *string `json:"postconditions,omitempty"`
	Layer          *string `json:"layer,omitempty"`
	Severity       *string `json:"severity,omitempty"`
	Priority       *string `json:"priority,omitempty"`
	Behavior       *string `json:"behavior,omitempty"`
	Type           *string `json:"type,omitempty"`
	Muted          *string `json:"muted,omitempty"`
	IsFlaky        *string `json:"is_flaky,omitempty"`
	// User who executed the test (member id, name or email)
	ExecutedBy           *string `json:"executed_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResultCreateFields ResultCreateFields

// NewResultCreateFields instantiates a new ResultCreateFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultCreateFields() *ResultCreateFields {
	this := ResultCreateFields{}
	return &this
}

// NewResultCreateFieldsWithDefaults instantiates a new ResultCreateFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultCreateFieldsWithDefaults() *ResultCreateFields {
	this := ResultCreateFields{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ResultCreateFields) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *ResultCreateFields) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *ResultCreateFields) SetAuthor(v string) {
	o.Author = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResultCreateFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResultCreateFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResultCreateFields) SetDescription(v string) {
	o.Description = &v
}

// GetPreconditions returns the Preconditions field value if set, zero value otherwise.
func (o *ResultCreateFields) GetPreconditions() string {
	if o == nil || IsNil(o.Preconditions) {
		var ret string
		return ret
	}
	return *o.Preconditions
}

// GetPreconditionsOk returns a tuple with the Preconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetPreconditionsOk() (*string, bool) {
	if o == nil || IsNil(o.Preconditions) {
		return nil, false
	}
	return o.Preconditions, true
}

// HasPreconditions returns a boolean if a field has been set.
func (o *ResultCreateFields) HasPreconditions() bool {
	if o != nil && !IsNil(o.Preconditions) {
		return true
	}

	return false
}

// SetPreconditions gets a reference to the given string and assigns it to the Preconditions field.
func (o *ResultCreateFields) SetPreconditions(v string) {
	o.Preconditions = &v
}

// GetPostconditions returns the Postconditions field value if set, zero value otherwise.
func (o *ResultCreateFields) GetPostconditions() string {
	if o == nil || IsNil(o.Postconditions) {
		var ret string
		return ret
	}
	return *o.Postconditions
}

// GetPostconditionsOk returns a tuple with the Postconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetPostconditionsOk() (*string, bool) {
	if o == nil || IsNil(o.Postconditions) {
		return nil, false
	}
	return o.Postconditions, true
}

// HasPostconditions returns a boolean if a field has been set.
func (o *ResultCreateFields) HasPostconditions() bool {
	if o != nil && !IsNil(o.Postconditions) {
		return true
	}

	return false
}

// SetPostconditions gets a reference to the given string and assigns it to the Postconditions field.
func (o *ResultCreateFields) SetPostconditions(v string) {
	o.Postconditions = &v
}

// GetLayer returns the Layer field value if set, zero value otherwise.
func (o *ResultCreateFields) GetLayer() string {
	if o == nil || IsNil(o.Layer) {
		var ret string
		return ret
	}
	return *o.Layer
}

// GetLayerOk returns a tuple with the Layer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetLayerOk() (*string, bool) {
	if o == nil || IsNil(o.Layer) {
		return nil, false
	}
	return o.Layer, true
}

// HasLayer returns a boolean if a field has been set.
func (o *ResultCreateFields) HasLayer() bool {
	if o != nil && !IsNil(o.Layer) {
		return true
	}

	return false
}

// SetLayer gets a reference to the given string and assigns it to the Layer field.
func (o *ResultCreateFields) SetLayer(v string) {
	o.Layer = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ResultCreateFields) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ResultCreateFields) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ResultCreateFields) SetSeverity(v string) {
	o.Severity = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ResultCreateFields) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ResultCreateFields) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *ResultCreateFields) SetPriority(v string) {
	o.Priority = &v
}

// GetBehavior returns the Behavior field value if set, zero value otherwise.
func (o *ResultCreateFields) GetBehavior() string {
	if o == nil || IsNil(o.Behavior) {
		var ret string
		return ret
	}
	return *o.Behavior
}

// GetBehaviorOk returns a tuple with the Behavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.Behavior) {
		return nil, false
	}
	return o.Behavior, true
}

// HasBehavior returns a boolean if a field has been set.
func (o *ResultCreateFields) HasBehavior() bool {
	if o != nil && !IsNil(o.Behavior) {
		return true
	}

	return false
}

// SetBehavior gets a reference to the given string and assigns it to the Behavior field.
func (o *ResultCreateFields) SetBehavior(v string) {
	o.Behavior = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResultCreateFields) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResultCreateFields) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResultCreateFields) SetType(v string) {
	o.Type = &v
}

// GetMuted returns the Muted field value if set, zero value otherwise.
func (o *ResultCreateFields) GetMuted() string {
	if o == nil || IsNil(o.Muted) {
		var ret string
		return ret
	}
	return *o.Muted
}

// GetMutedOk returns a tuple with the Muted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetMutedOk() (*string, bool) {
	if o == nil || IsNil(o.Muted) {
		return nil, false
	}
	return o.Muted, true
}

// HasMuted returns a boolean if a field has been set.
func (o *ResultCreateFields) HasMuted() bool {
	if o != nil && !IsNil(o.Muted) {
		return true
	}

	return false
}

// SetMuted gets a reference to the given string and assigns it to the Muted field.
func (o *ResultCreateFields) SetMuted(v string) {
	o.Muted = &v
}

// GetIsFlaky returns the IsFlaky field value if set, zero value otherwise.
func (o *ResultCreateFields) GetIsFlaky() string {
	if o == nil || IsNil(o.IsFlaky) {
		var ret string
		return ret
	}
	return *o.IsFlaky
}

// GetIsFlakyOk returns a tuple with the IsFlaky field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetIsFlakyOk() (*string, bool) {
	if o == nil || IsNil(o.IsFlaky) {
		return nil, false
	}
	return o.IsFlaky, true
}

// HasIsFlaky returns a boolean if a field has been set.
func (o *ResultCreateFields) HasIsFlaky() bool {
	if o != nil && !IsNil(o.IsFlaky) {
		return true
	}

	return false
}

// SetIsFlaky gets a reference to the given string and assigns it to the IsFlaky field.
func (o *ResultCreateFields) SetIsFlaky(v string) {
	o.IsFlaky = &v
}

// GetExecutedBy returns the ExecutedBy field value if set, zero value otherwise.
func (o *ResultCreateFields) GetExecutedBy() string {
	if o == nil || IsNil(o.ExecutedBy) {
		var ret string
		return ret
	}
	return *o.ExecutedBy
}

// GetExecutedByOk returns a tuple with the ExecutedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreateFields) GetExecutedByOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedBy) {
		return nil, false
	}
	return o.ExecutedBy, true
}

// HasExecutedBy returns a boolean if a field has been set.
func (o *ResultCreateFields) HasExecutedBy() bool {
	if o != nil && !IsNil(o.ExecutedBy) {
		return true
	}

	return false
}

// SetExecutedBy gets a reference to the given string and assigns it to the ExecutedBy field.
func (o *ResultCreateFields) SetExecutedBy(v string) {
	o.ExecutedBy = &v
}

func (o ResultCreateFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultCreateFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Preconditions) {
		toSerialize["preconditions"] = o.Preconditions
	}
	if !IsNil(o.Postconditions) {
		toSerialize["postconditions"] = o.Postconditions
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
	if !IsNil(o.Behavior) {
		toSerialize["behavior"] = o.Behavior
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Muted) {
		toSerialize["muted"] = o.Muted
	}
	if !IsNil(o.IsFlaky) {
		toSerialize["is_flaky"] = o.IsFlaky
	}
	if !IsNil(o.ExecutedBy) {
		toSerialize["executed_by"] = o.ExecutedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResultCreateFields) UnmarshalJSON(data []byte) (err error) {
	varResultCreateFields := _ResultCreateFields{}

	err = json.Unmarshal(data, &varResultCreateFields)

	if err != nil {
		return err
	}

	*o = ResultCreateFields(varResultCreateFields)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "author")
		delete(additionalProperties, "description")
		delete(additionalProperties, "preconditions")
		delete(additionalProperties, "postconditions")
		delete(additionalProperties, "layer")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "behavior")
		delete(additionalProperties, "type")
		delete(additionalProperties, "muted")
		delete(additionalProperties, "is_flaky")
		delete(additionalProperties, "executed_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResultCreateFields struct {
	value *ResultCreateFields
	isSet bool
}

func (v NullableResultCreateFields) Get() *ResultCreateFields {
	return v.value
}

func (v *NullableResultCreateFields) Set(val *ResultCreateFields) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCreateFields) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCreateFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCreateFields(val *ResultCreateFields) *NullableResultCreateFields {
	return &NullableResultCreateFields{value: val, isSet: true}
}

func (v NullableResultCreateFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCreateFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
