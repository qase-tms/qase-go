package domain

// Package domain provides Go domain models for Qase test results
// that match the JavaScript interfaces for cross-language compatibility.

const (
	// Version is the current version of the domain package
	Version = "1.0.0"
	
	// PackageName is the name of this package
	PackageName = "qase-go-domain"
)

// ValidateStatus validates if the given status is valid
func ValidateStatus(status string) error {
	if !TestResultStatus(status).IsValid() {
		return ErrInvalidStatus
	}
	return nil
}

// ValidateTitle validates if the given title is valid
func ValidateTitle(title string) error {
	if title == "" {
		return ErrEmptyTitle
	}
	return nil
}