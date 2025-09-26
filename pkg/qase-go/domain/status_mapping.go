package domain

import (
	"fmt"
	"strings"
)

// StatusMapping represents a mapping from one status to another
type StatusMapping map[TestResultStatus]TestResultStatus

// NewStatusMapping creates a new status mapping from a map
func NewStatusMapping(mapping map[string]string) (StatusMapping, error) {
	result := make(StatusMapping)
	
	for from, to := range mapping {
		fromStatus := TestResultStatus(strings.ToLower(strings.TrimSpace(from)))
		toStatus := TestResultStatus(strings.ToLower(strings.TrimSpace(to)))
		
		// Validate source status
		if !fromStatus.IsValid() {
			return nil, fmt.Errorf("invalid source status '%s' in mapping", from)
		}
		
		// Validate target status
		if !toStatus.IsValid() {
			return nil, fmt.Errorf("invalid target status '%s' in mapping", to)
		}
		
		// Prevent mapping to the same status
		if fromStatus == toStatus {
			return nil, fmt.Errorf("cannot map status '%s' to itself", from)
		}
		
		result[fromStatus] = toStatus
	}
	
	return result, nil
}

// NewStatusMappingFromEnv creates a status mapping from environment variable format
// Format: "invalid=failed,skipped=passed"
func NewStatusMappingFromEnv(envValue string) (StatusMapping, error) {
	if envValue == "" {
		return nil, nil
	}
	
	mapping := make(map[string]string)
	
	// Split by comma
	pairs := strings.Split(envValue, ",")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		
		// Split by equals sign
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid mapping format '%s', expected 'from=to'", pair)
		}
		
		from := strings.TrimSpace(parts[0])
		to := strings.TrimSpace(parts[1])
		
		if from == "" || to == "" {
			return nil, fmt.Errorf("empty status in mapping '%s'", pair)
		}
		
		mapping[from] = to
	}
	
	return NewStatusMapping(mapping)
}

// ApplyMapping applies the status mapping to a test result status
// Returns the mapped status if mapping exists, otherwise returns the original status
func (sm StatusMapping) ApplyMapping(status TestResultStatus) TestResultStatus {
	if mappedStatus, exists := sm[status]; exists {
		return mappedStatus
	}
	return status
}

// ApplyMappingToResult applies the status mapping to a test result
// Modifies the result in place and returns true if mapping was applied
func (sm StatusMapping) ApplyMappingToResult(result *TestResult) bool {
	if result == nil {
		return false
	}
	
	originalStatus := result.Execution.Status
	mappedStatus := sm.ApplyMapping(originalStatus)
	
	if mappedStatus != originalStatus {
		result.Execution.Status = mappedStatus
		return true
	}
	
	return false
}

// IsEmpty returns true if the mapping is empty
func (sm StatusMapping) IsEmpty() bool {
	return len(sm) == 0
}

// GetMappings returns a copy of the current mappings
func (sm StatusMapping) GetMappings() map[TestResultStatus]TestResultStatus {
	result := make(map[TestResultStatus]TestResultStatus)
	for from, to := range sm {
		result[from] = to
	}
	return result
}

// String returns a string representation of the mapping
func (sm StatusMapping) String() string {
	if sm.IsEmpty() {
		return "{}"
	}
	
	var parts []string
	for from, to := range sm {
		parts = append(parts, fmt.Sprintf("%s->%s", from, to))
	}
	
	return fmt.Sprintf("{%s}", strings.Join(parts, ", "))
}

// ValidateMapping validates that all statuses in the mapping are valid
func ValidateMapping(mapping map[string]string) error {
	validStatuses := []string{"passed", "failed", "blocked", "skipped", "in_progress", "invalid"}
	
	for from, to := range mapping {
		from = strings.ToLower(strings.TrimSpace(from))
		to = strings.ToLower(strings.TrimSpace(to))
		
		// Check source status
		if !contains(validStatuses, from) {
			return fmt.Errorf("invalid source status '%s', must be one of: %s", from, strings.Join(validStatuses, ", "))
		}
		
		// Check target status
		if !contains(validStatuses, to) {
			return fmt.Errorf("invalid target status '%s', must be one of: %s", to, strings.Join(validStatuses, ", "))
		}
		
		// Prevent mapping to the same status
		if from == to {
			return fmt.Errorf("cannot map status '%s' to itself", from)
		}
	}
	
	return nil
}

// helper function to check if slice contains a value
func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
