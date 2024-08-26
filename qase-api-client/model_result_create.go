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

// checks if the ResultCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultCreate{}

// ResultCreate struct for ResultCreate
type ResultCreate struct {
	CaseId *int64            `json:"case_id,omitempty"`
	Case   *ResultCreateCase `json:"case,omitempty"`
	// Can have the following values `passed`, `failed`, `blocked`, `skipped`, `invalid` + custom statuses
	Status      string         `json:"status"`
	StartTime   NullableInt32  `json:"start_time,omitempty"`
	Time        NullableInt64  `json:"time,omitempty"`
	TimeMs      NullableInt64  `json:"time_ms,omitempty"`
	Defect      NullableBool   `json:"defect,omitempty"`
	Attachments []string       `json:"attachments,omitempty"`
	Stacktrace  NullableString `json:"stacktrace,omitempty"`
	Comment     NullableString `json:"comment,omitempty"`
	// A map of parameters (name => value)
	Param map[string]string `json:"param,omitempty"`
	// A list of parameter groups
	ParamGroups [][]string             `json:"param_groups,omitempty"`
	Steps       []TestStepResultCreate `json:"steps,omitempty"`
	AuthorId    NullableInt64          `json:"author_id,omitempty"`
}

type _ResultCreate ResultCreate

// NewResultCreate instantiates a new ResultCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultCreate(status string) *ResultCreate {
	this := ResultCreate{}
	this.Status = status
	return &this
}

// NewResultCreateWithDefaults instantiates a new ResultCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultCreateWithDefaults() *ResultCreate {
	this := ResultCreate{}
	return &this
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *ResultCreate) GetCaseId() int64 {
	if o == nil || IsNil(o.CaseId) {
		var ret int64
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetCaseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CaseId) {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *ResultCreate) HasCaseId() bool {
	if o != nil && !IsNil(o.CaseId) {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given int64 and assigns it to the CaseId field.
func (o *ResultCreate) SetCaseId(v int64) {
	o.CaseId = &v
}

// GetCase returns the Case field value if set, zero value otherwise.
func (o *ResultCreate) GetCase() ResultCreateCase {
	if o == nil || IsNil(o.Case) {
		var ret ResultCreateCase
		return ret
	}
	return *o.Case
}

// GetCaseOk returns a tuple with the Case field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetCaseOk() (*ResultCreateCase, bool) {
	if o == nil || IsNil(o.Case) {
		return nil, false
	}
	return o.Case, true
}

// HasCase returns a boolean if a field has been set.
func (o *ResultCreate) HasCase() bool {
	if o != nil && !IsNil(o.Case) {
		return true
	}

	return false
}

// SetCase gets a reference to the given ResultCreateCase and assigns it to the Case field.
func (o *ResultCreate) SetCase(v ResultCreateCase) {
	o.Case = &v
}

// GetStatus returns the Status field value
func (o *ResultCreate) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResultCreate) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResultCreate) SetStatus(v string) {
	o.Status = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetStartTime() int32 {
	if o == nil || IsNil(o.StartTime.Get()) {
		var ret int32
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetStartTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *ResultCreate) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableInt32 and assigns it to the StartTime field.
func (o *ResultCreate) SetStartTime(v int32) {
	o.StartTime.Set(&v)
}

// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *ResultCreate) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *ResultCreate) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetTime returns the Time field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetTime() int64 {
	if o == nil || IsNil(o.Time.Get()) {
		var ret int64
		return ret
	}
	return *o.Time.Get()
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Time.Get(), o.Time.IsSet()
}

// HasTime returns a boolean if a field has been set.
func (o *ResultCreate) HasTime() bool {
	if o != nil && o.Time.IsSet() {
		return true
	}

	return false
}

// SetTime gets a reference to the given NullableInt64 and assigns it to the Time field.
func (o *ResultCreate) SetTime(v int64) {
	o.Time.Set(&v)
}

// SetTimeNil sets the value for Time to be an explicit nil
func (o *ResultCreate) SetTimeNil() {
	o.Time.Set(nil)
}

// UnsetTime ensures that no value is present for Time, not even an explicit nil
func (o *ResultCreate) UnsetTime() {
	o.Time.Unset()
}

// GetTimeMs returns the TimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetTimeMs() int64 {
	if o == nil || IsNil(o.TimeMs.Get()) {
		var ret int64
		return ret
	}
	return *o.TimeMs.Get()
}

// GetTimeMsOk returns a tuple with the TimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetTimeMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeMs.Get(), o.TimeMs.IsSet()
}

// HasTimeMs returns a boolean if a field has been set.
func (o *ResultCreate) HasTimeMs() bool {
	if o != nil && o.TimeMs.IsSet() {
		return true
	}

	return false
}

// SetTimeMs gets a reference to the given NullableInt64 and assigns it to the TimeMs field.
func (o *ResultCreate) SetTimeMs(v int64) {
	o.TimeMs.Set(&v)
}

// SetTimeMsNil sets the value for TimeMs to be an explicit nil
func (o *ResultCreate) SetTimeMsNil() {
	o.TimeMs.Set(nil)
}

// UnsetTimeMs ensures that no value is present for TimeMs, not even an explicit nil
func (o *ResultCreate) UnsetTimeMs() {
	o.TimeMs.Unset()
}

// GetDefect returns the Defect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetDefect() bool {
	if o == nil || IsNil(o.Defect.Get()) {
		var ret bool
		return ret
	}
	return *o.Defect.Get()
}

// GetDefectOk returns a tuple with the Defect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetDefectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Defect.Get(), o.Defect.IsSet()
}

// HasDefect returns a boolean if a field has been set.
func (o *ResultCreate) HasDefect() bool {
	if o != nil && o.Defect.IsSet() {
		return true
	}

	return false
}

// SetDefect gets a reference to the given NullableBool and assigns it to the Defect field.
func (o *ResultCreate) SetDefect(v bool) {
	o.Defect.Set(&v)
}

// SetDefectNil sets the value for Defect to be an explicit nil
func (o *ResultCreate) SetDefectNil() {
	o.Defect.Set(nil)
}

// UnsetDefect ensures that no value is present for Defect, not even an explicit nil
func (o *ResultCreate) UnsetDefect() {
	o.Defect.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetAttachments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
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

// GetStacktrace returns the Stacktrace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetStacktrace() string {
	if o == nil || IsNil(o.Stacktrace.Get()) {
		var ret string
		return ret
	}
	return *o.Stacktrace.Get()
}

// GetStacktraceOk returns a tuple with the Stacktrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetStacktraceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stacktrace.Get(), o.Stacktrace.IsSet()
}

// HasStacktrace returns a boolean if a field has been set.
func (o *ResultCreate) HasStacktrace() bool {
	if o != nil && o.Stacktrace.IsSet() {
		return true
	}

	return false
}

// SetStacktrace gets a reference to the given NullableString and assigns it to the Stacktrace field.
func (o *ResultCreate) SetStacktrace(v string) {
	o.Stacktrace.Set(&v)
}

// SetStacktraceNil sets the value for Stacktrace to be an explicit nil
func (o *ResultCreate) SetStacktraceNil() {
	o.Stacktrace.Set(nil)
}

// UnsetStacktrace ensures that no value is present for Stacktrace, not even an explicit nil
func (o *ResultCreate) UnsetStacktrace() {
	o.Stacktrace.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ResultCreate) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ResultCreate) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ResultCreate) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ResultCreate) UnsetComment() {
	o.Comment.Unset()
}

// GetParam returns the Param field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetParam() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetParamOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return &o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *ResultCreate) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given map[string]string and assigns it to the Param field.
func (o *ResultCreate) SetParam(v map[string]string) {
	o.Param = v
}

// GetParamGroups returns the ParamGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetParamGroups() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}
	return o.ParamGroups
}

// GetParamGroupsOk returns a tuple with the ParamGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetParamGroupsOk() ([][]string, bool) {
	if o == nil || IsNil(o.ParamGroups) {
		return nil, false
	}
	return o.ParamGroups, true
}

// HasParamGroups returns a boolean if a field has been set.
func (o *ResultCreate) HasParamGroups() bool {
	if o != nil && !IsNil(o.ParamGroups) {
		return true
	}

	return false
}

// SetParamGroups gets a reference to the given [][]string and assigns it to the ParamGroups field.
func (o *ResultCreate) SetParamGroups(v [][]string) {
	o.ParamGroups = v
}

// GetSteps returns the Steps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetSteps() []TestStepResultCreate {
	if o == nil {
		var ret []TestStepResultCreate
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetStepsOk() ([]TestStepResultCreate, bool) {
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

// SetSteps gets a reference to the given []TestStepResultCreate and assigns it to the Steps field.
func (o *ResultCreate) SetSteps(v []TestStepResultCreate) {
	o.Steps = v
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultCreate) GetAuthorId() int64 {
	if o == nil || IsNil(o.AuthorId.Get()) {
		var ret int64
		return ret
	}
	return *o.AuthorId.Get()
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultCreate) GetAuthorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorId.Get(), o.AuthorId.IsSet()
}

// HasAuthorId returns a boolean if a field has been set.
func (o *ResultCreate) HasAuthorId() bool {
	if o != nil && o.AuthorId.IsSet() {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given NullableInt64 and assigns it to the AuthorId field.
func (o *ResultCreate) SetAuthorId(v int64) {
	o.AuthorId.Set(&v)
}

// SetAuthorIdNil sets the value for AuthorId to be an explicit nil
func (o *ResultCreate) SetAuthorIdNil() {
	o.AuthorId.Set(nil)
}

// UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil
func (o *ResultCreate) UnsetAuthorId() {
	o.AuthorId.Unset()
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
	if !IsNil(o.CaseId) {
		toSerialize["case_id"] = o.CaseId
	}
	if !IsNil(o.Case) {
		toSerialize["case"] = o.Case
	}
	toSerialize["status"] = o.Status
	if o.StartTime.IsSet() {
		toSerialize["start_time"] = o.StartTime.Get()
	}
	if o.Time.IsSet() {
		toSerialize["time"] = o.Time.Get()
	}
	if o.TimeMs.IsSet() {
		toSerialize["time_ms"] = o.TimeMs.Get()
	}
	if o.Defect.IsSet() {
		toSerialize["defect"] = o.Defect.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Stacktrace.IsSet() {
		toSerialize["stacktrace"] = o.Stacktrace.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Param != nil {
		toSerialize["param"] = o.Param
	}
	if o.ParamGroups != nil {
		toSerialize["param_groups"] = o.ParamGroups
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.AuthorId.IsSet() {
		toSerialize["author_id"] = o.AuthorId.Get()
	}
	return toSerialize, nil
}

func (o *ResultCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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
