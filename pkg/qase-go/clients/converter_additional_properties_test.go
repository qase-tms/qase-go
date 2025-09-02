package clients

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestV2Converter_AdditionalProperties(t *testing.T) {
	converter := NewV2Converter()

	// Create a test result with unknown fields
	domainResult := domain.NewTestResult("Test with Additional Properties")
	domainResult.Execution.Status = domain.StatusPassed

	// Add known fields
	domainResult.SetField("description", "Test description")
	domainResult.SetField("severity", "high")
	domainResult.SetField("priority", "critical")

	// Add unknown fields that should go to AdditionalProperties
	domainResult.SetField("custom_field_1", "custom_value_1")
	domainResult.SetField("custom_field_2", "custom_value_2")
	domainResult.SetField("unknown_field", "should_be_in_additional_properties")

	apiResult, err := converter.ConvertTestResult(domainResult)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify that known fields are set correctly
	if apiResult.Fields.Description == nil || *apiResult.Fields.Description != "Test description" {
		t.Error("Description should be set correctly")
	}
	if apiResult.Fields.Severity == nil || *apiResult.Fields.Severity != "high" {
		t.Error("Severity should be set correctly")
	}
	if apiResult.Fields.Priority == nil || *apiResult.Fields.Priority != "critical" {
		t.Error("Priority should be set correctly")
	}

	// Verify that unknown fields are in AdditionalProperties
	if apiResult.Fields.AdditionalProperties == nil {
		t.Error("AdditionalProperties should not be nil")
	}

	additionalProps := apiResult.Fields.AdditionalProperties
	if len(additionalProps) != 3 {
		t.Errorf("Expected 3 additional properties, got %d", len(additionalProps))
	}

	// Check specific additional properties
	if val, ok := additionalProps["custom_field_1"]; !ok || val != "custom_value_1" {
		t.Error("custom_field_1 should be in AdditionalProperties")
	}
	if val, ok := additionalProps["custom_field_2"]; !ok || val != "custom_value_2" {
		t.Error("custom_field_2 should be in AdditionalProperties")
	}
	if val, ok := additionalProps["unknown_field"]; !ok || val != "should_be_in_additional_properties" {
		t.Error("unknown_field should be in AdditionalProperties")
	}
}
