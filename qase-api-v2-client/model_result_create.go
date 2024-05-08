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

// checks if the ResultCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultCreate{}

// ResultCreate struct for ResultCreate
type ResultCreate struct {
	// If passed, used as an idempotency key
	Id          *string                 `json:"id,omitempty"`
	Title       string                  `json:"title"`
	Signature   *string                 `json:"signature,omitempty"`
	TestopsId   NullableInt64           `json:"testops_id,omitempty"`
	Execution   ResultExecution         `json:"execution"`
	Fields      *map[string]string      `json:"fields,omitempty"`
	Attachments []string                `json:"attachments,omitempty"`
	Steps       []ResultStep            `json:"steps,omitempty"`
	StepsType   NullableResultStepsType `json:"steps_type,omitempty"`
	Params      *map[string]string      `json:"params,omitempty"`
	Author      *string                 `json:"author,omitempty"`
	Relations   NullableResultRelations `json:"relations,omitempty"`
	Muted       *bool                   `json:"muted,omitempty"`
	Message     NullableString          `json:"message,omitempty"`
	CreatedAt   NullableFloat64         `json:"created_at,omitempty"`
}

type _ResultCreate ResultCreate

// NewResultCreate instantiates a new ResultCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultCreate(title string, execution ResultExecution) *ResultCreate {
	this := ResultCreate{}
	this.Title = title
	this.Execution = execution
	return &this
}

// NewResultCreateWithDefaults instantiates a new ResultCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultCreateWithDefaults() *ResultCreate {
	this := ResultCreate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResultCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResultCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResultCreate) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value
func (o *ResultCreate) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ResultCreate) SetTitle(v string) {
	o.Title = v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ResultCreate) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ResultCreate) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *ResultCreate) SetSignature(v string) {
	o.Signature = &v
}

// GetTestopsId returns the TestopsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetTestopsId() int64 {
	if o == nil || IsNil(o.TestopsId.Get()) {
		var ret int64
		return ret
	}
	return *o.TestopsId.Get()
}

// GetTestopsIdOk returns a tuple with the TestopsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetTestopsIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestopsId.Get(), o.TestopsId.IsSet()
}

// HasTestopsId returns a boolean if a field has been set.
func (o *ResultCreate) HasTestopsId() bool {
	if o != nil && o.TestopsId.IsSet() {
		return true
	}

	return false
}

// SetTestopsId gets a reference to the given NullableInt64 and assigns it to the TestopsId field.
func (o *ResultCreate) SetTestopsId(v int64) {
	o.TestopsId.Set(&v)
}

// SetTestopsIdNil sets the value for TestopsId to be an explicit nil
func (o *ResultCreate) SetTestopsIdNil() {
	o.TestopsId.Set(nil)
}

// UnsetTestopsId ensures that no value is present for TestopsId, not even an explicit nil
func (o *ResultCreate) UnsetTestopsId() {
	o.TestopsId.Unset()
}

// GetExecution returns the Execution field value
func (o *ResultCreate) GetExecution() ResultExecution {
	if o == nil {
		var ret ResultExecution
		return ret
	}

	return o.Execution
}

// GetExecutionOk returns a tuple with the Execution field value
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetExecutionOk() (*ResultExecution, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Execution, true
}

// SetExecution sets field value
func (o *ResultCreate) SetExecution(v ResultExecution) {
	o.Execution = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ResultCreate) GetFields() map[string]string {
	if o == nil || IsNil(o.Fields) {
		var ret map[string]string
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ResultCreate) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]string and assigns it to the Fields field.
func (o *ResultCreate) SetFields(v map[string]string) {
	o.Fields = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ResultCreate) GetAttachments() []string {
	if o == nil || IsNil(o.Attachments) {
		var ret []string
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetAttachmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ResultCreate) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []string and assigns it to the Attachments field.
func (o *ResultCreate) SetAttachments(v []string) {
	o.Attachments = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ResultCreate) GetSteps() []ResultStep {
	if o == nil || IsNil(o.Steps) {
		var ret []ResultStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetStepsOk() ([]ResultStep, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ResultCreate) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []ResultStep and assigns it to the Steps field.
func (o *ResultCreate) SetSteps(v []ResultStep) {
	o.Steps = v
}

// GetStepsType returns the StepsType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetStepsType() ResultStepsType {
	if o == nil || IsNil(o.StepsType.Get()) {
		var ret ResultStepsType
		return ret
	}
	return *o.StepsType.Get()
}

// GetStepsTypeOk returns a tuple with the StepsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetStepsTypeOk() (*ResultStepsType, bool) {
	if o == nil {
		return nil, false
	}
	return o.StepsType.Get(), o.StepsType.IsSet()
}

// HasStepsType returns a boolean if a field has been set.
func (o *ResultCreate) HasStepsType() bool {
	if o != nil && o.StepsType.IsSet() {
		return true
	}

	return false
}

// SetStepsType gets a reference to the given NullableResultStepsType and assigns it to the StepsType field.
func (o *ResultCreate) SetStepsType(v ResultStepsType) {
	o.StepsType.Set(&v)
}

// SetStepsTypeNil sets the value for StepsType to be an explicit nil
func (o *ResultCreate) SetStepsTypeNil() {
	o.StepsType.Set(nil)
}

// UnsetStepsType ensures that no value is present for StepsType, not even an explicit nil
func (o *ResultCreate) UnsetStepsType() {
	o.StepsType.Unset()
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ResultCreate) GetParams() map[string]string {
	if o == nil || IsNil(o.Params) {
		var ret map[string]string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetParamsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ResultCreate) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *ResultCreate) SetParams(v map[string]string) {
	o.Params = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ResultCreate) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *ResultCreate) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *ResultCreate) SetAuthor(v string) {
	o.Author = &v
}

// GetRelations returns the Relations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetRelations() ResultRelations {
	if o == nil || IsNil(o.Relations.Get()) {
		var ret ResultRelations
		return ret
	}
	return *o.Relations.Get()
}

// GetRelationsOk returns a tuple with the Relations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetRelationsOk() (*ResultRelations, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relations.Get(), o.Relations.IsSet()
}

// HasRelations returns a boolean if a field has been set.
func (o *ResultCreate) HasRelations() bool {
	if o != nil && o.Relations.IsSet() {
		return true
	}

	return false
}

// SetRelations gets a reference to the given NullableResultRelations and assigns it to the Relations field.
func (o *ResultCreate) SetRelations(v ResultRelations) {
	o.Relations.Set(&v)
}

// SetRelationsNil sets the value for Relations to be an explicit nil
func (o *ResultCreate) SetRelationsNil() {
	o.Relations.Set(nil)
}

// UnsetRelations ensures that no value is present for Relations, not even an explicit nil
func (o *ResultCreate) UnsetRelations() {
	o.Relations.Unset()
}

// GetMuted returns the Muted field value if set, zero value otherwise.
func (o *ResultCreate) GetMuted() bool {
	if o == nil || IsNil(o.Muted) {
		var ret bool
		return ret
	}
	return *o.Muted
}

// GetMutedOk returns a tuple with the Muted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.Muted) {
		return nil, false
	}
	return o.Muted, true
}

// HasMuted returns a boolean if a field has been set.
func (o *ResultCreate) HasMuted() bool {
	if o != nil && !IsNil(o.Muted) {
		return true
	}

	return false
}

// SetMuted gets a reference to the given bool and assigns it to the Muted field.
func (o *ResultCreate) SetMuted(v bool) {
	o.Muted = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ResultCreate) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ResultCreate) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *ResultCreate) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ResultCreate) UnsetMessage() {
	o.Message.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetCreatedAt() float64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret float64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetCreatedAtOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ResultCreate) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableFloat64 and assigns it to the CreatedAt field.
func (o *ResultCreate) SetCreatedAt(v float64) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ResultCreate) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ResultCreate) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

func (o ResultCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["title"] = o.Title
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if o.TestopsId.IsSet() {
		toSerialize["testops_id"] = o.TestopsId.Get()
	}
	toSerialize["execution"] = o.Execution
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if o.StepsType.IsSet() {
		toSerialize["steps_type"] = o.StepsType.Get()
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if o.Relations.IsSet() {
		toSerialize["relations"] = o.Relations.Get()
	}
	if !IsNil(o.Muted) {
		toSerialize["muted"] = o.Muted
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	return toSerialize, nil
}

func (o *ResultCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"execution",
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

	varResultCreate := _ResultCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultCreate)

	if err != nil {
		return err
	}

	*o = ResultCreate(varResultCreate)

	return err
}

type NullableResultCreate struct {
	value *ResultCreate
	isSet bool
}

func (v NullableResultCreate) Get() *ResultCreate {
	return v.value
}

func (v *NullableResultCreate) Set(val *ResultCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCreate(val *ResultCreate) *NullableResultCreate {
	return &NullableResultCreate{value: val, isSet: true}
}

func (v NullableResultCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
