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
	"fmt"
)

// SearchResponseAllOfResultEntities - struct for SearchResponseAllOfResultEntities
type SearchResponseAllOfResultEntities struct {
	DefectQuery   *DefectQuery
	PlanQuery     *PlanQuery
	Requirement   *Requirement
	Result        *Result
	Run           *Run
	TestCaseQuery *TestCaseQuery
}

// DefectQueryAsSearchResponseAllOfResultEntities is a convenience function that returns DefectQuery wrapped in SearchResponseAllOfResultEntities
func DefectQueryAsSearchResponseAllOfResultEntities(v *DefectQuery) SearchResponseAllOfResultEntities {
	return SearchResponseAllOfResultEntities{
		DefectQuery: v,
	}
}

// PlanQueryAsSearchResponseAllOfResultEntities is a convenience function that returns PlanQuery wrapped in SearchResponseAllOfResultEntities
func PlanQueryAsSearchResponseAllOfResultEntities(v *PlanQuery) SearchResponseAllOfResultEntities {
	return SearchResponseAllOfResultEntities{
		PlanQuery: v,
	}
}

// RequirementAsSearchResponseAllOfResultEntities is a convenience function that returns Requirement wrapped in SearchResponseAllOfResultEntities
func RequirementAsSearchResponseAllOfResultEntities(v *Requirement) SearchResponseAllOfResultEntities {
	return SearchResponseAllOfResultEntities{
		Requirement: v,
	}
}

// ResultAsSearchResponseAllOfResultEntities is a convenience function that returns Result wrapped in SearchResponseAllOfResultEntities
func ResultAsSearchResponseAllOfResultEntities(v *Result) SearchResponseAllOfResultEntities {
	return SearchResponseAllOfResultEntities{
		Result: v,
	}
}

// RunAsSearchResponseAllOfResultEntities is a convenience function that returns Run wrapped in SearchResponseAllOfResultEntities
func RunAsSearchResponseAllOfResultEntities(v *Run) SearchResponseAllOfResultEntities {
	return SearchResponseAllOfResultEntities{
		Run: v,
	}
}

// TestCaseQueryAsSearchResponseAllOfResultEntities is a convenience function that returns TestCaseQuery wrapped in SearchResponseAllOfResultEntities
func TestCaseQueryAsSearchResponseAllOfResultEntities(v *TestCaseQuery) SearchResponseAllOfResultEntities {
	return SearchResponseAllOfResultEntities{
		TestCaseQuery: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SearchResponseAllOfResultEntities) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DefectQuery
	err = newStrictDecoder(data).Decode(&dst.DefectQuery)
	if err == nil {
		jsonDefectQuery, _ := json.Marshal(dst.DefectQuery)
		if string(jsonDefectQuery) == "{}" { // empty struct
			dst.DefectQuery = nil
		} else {
			match++
		}
	} else {
		dst.DefectQuery = nil
	}

	// try to unmarshal data into PlanQuery
	err = newStrictDecoder(data).Decode(&dst.PlanQuery)
	if err == nil {
		jsonPlanQuery, _ := json.Marshal(dst.PlanQuery)
		if string(jsonPlanQuery) == "{}" { // empty struct
			dst.PlanQuery = nil
		} else {
			match++
		}
	} else {
		dst.PlanQuery = nil
	}

	// try to unmarshal data into Requirement
	err = newStrictDecoder(data).Decode(&dst.Requirement)
	if err == nil {
		jsonRequirement, _ := json.Marshal(dst.Requirement)
		if string(jsonRequirement) == "{}" { // empty struct
			dst.Requirement = nil
		} else {
			match++
		}
	} else {
		dst.Requirement = nil
	}

	// try to unmarshal data into Result
	err = newStrictDecoder(data).Decode(&dst.Result)
	if err == nil {
		jsonResult, _ := json.Marshal(dst.Result)
		if string(jsonResult) == "{}" { // empty struct
			dst.Result = nil
		} else {
			match++
		}
	} else {
		dst.Result = nil
	}

	// try to unmarshal data into Run
	err = newStrictDecoder(data).Decode(&dst.Run)
	if err == nil {
		jsonRun, _ := json.Marshal(dst.Run)
		if string(jsonRun) == "{}" { // empty struct
			dst.Run = nil
		} else {
			match++
		}
	} else {
		dst.Run = nil
	}

	// try to unmarshal data into TestCaseQuery
	err = newStrictDecoder(data).Decode(&dst.TestCaseQuery)
	if err == nil {
		jsonTestCaseQuery, _ := json.Marshal(dst.TestCaseQuery)
		if string(jsonTestCaseQuery) == "{}" { // empty struct
			dst.TestCaseQuery = nil
		} else {
			match++
		}
	} else {
		dst.TestCaseQuery = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DefectQuery = nil
		dst.PlanQuery = nil
		dst.Requirement = nil
		dst.Result = nil
		dst.Run = nil
		dst.TestCaseQuery = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SearchResponseAllOfResultEntities)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SearchResponseAllOfResultEntities)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SearchResponseAllOfResultEntities) MarshalJSON() ([]byte, error) {
	if src.DefectQuery != nil {
		return json.Marshal(&src.DefectQuery)
	}

	if src.PlanQuery != nil {
		return json.Marshal(&src.PlanQuery)
	}

	if src.Requirement != nil {
		return json.Marshal(&src.Requirement)
	}

	if src.Result != nil {
		return json.Marshal(&src.Result)
	}

	if src.Run != nil {
		return json.Marshal(&src.Run)
	}

	if src.TestCaseQuery != nil {
		return json.Marshal(&src.TestCaseQuery)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SearchResponseAllOfResultEntities) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DefectQuery != nil {
		return obj.DefectQuery
	}

	if obj.PlanQuery != nil {
		return obj.PlanQuery
	}

	if obj.Requirement != nil {
		return obj.Requirement
	}

	if obj.Result != nil {
		return obj.Result
	}

	if obj.Run != nil {
		return obj.Run
	}

	if obj.TestCaseQuery != nil {
		return obj.TestCaseQuery
	}

	// all schemas are nil
	return nil
}

type NullableSearchResponseAllOfResultEntities struct {
	value *SearchResponseAllOfResultEntities
	isSet bool
}

func (v NullableSearchResponseAllOfResultEntities) Get() *SearchResponseAllOfResultEntities {
	return v.value
}

func (v *NullableSearchResponseAllOfResultEntities) Set(val *SearchResponseAllOfResultEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResponseAllOfResultEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResponseAllOfResultEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResponseAllOfResultEntities(val *SearchResponseAllOfResultEntities) *NullableSearchResponseAllOfResultEntities {
	return &NullableSearchResponseAllOfResultEntities{value: val, isSet: true}
}

func (v NullableSearchResponseAllOfResultEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResponseAllOfResultEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}