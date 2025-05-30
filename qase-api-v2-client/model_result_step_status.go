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
	"fmt"
)

// ResultStepStatus the model 'ResultStepStatus'
type ResultStepStatus string

// List of ResultStepStatus
const (
	PASSED      ResultStepStatus = "passed"
	FAILED      ResultStepStatus = "failed"
	BLOCKED     ResultStepStatus = "blocked"
	SKIPPED     ResultStepStatus = "skipped"
	IN_PROGRESS ResultStepStatus = "in_progress"
)

// All allowed values of ResultStepStatus enum
var AllowedResultStepStatusEnumValues = []ResultStepStatus{
	"passed",
	"failed",
	"blocked",
	"skipped",
	"in_progress",
}

func (v *ResultStepStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResultStepStatus(value)
	for _, existing := range AllowedResultStepStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResultStepStatus", value)
}

// NewResultStepStatusFromValue returns a pointer to a valid ResultStepStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResultStepStatusFromValue(v string) (*ResultStepStatus, error) {
	ev := ResultStepStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResultStepStatus: valid values are %v", v, AllowedResultStepStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResultStepStatus) IsValid() bool {
	for _, existing := range AllowedResultStepStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResultStepStatus value
func (v ResultStepStatus) Ptr() *ResultStepStatus {
	return &v
}

type NullableResultStepStatus struct {
	value *ResultStepStatus
	isSet bool
}

func (v NullableResultStepStatus) Get() *ResultStepStatus {
	return v.value
}

func (v *NullableResultStepStatus) Set(val *ResultStepStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableResultStepStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableResultStepStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultStepStatus(val *ResultStepStatus) *NullableResultStepStatus {
	return &NullableResultStepStatus{value: val, isSet: true}
}

func (v NullableResultStepStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultStepStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
