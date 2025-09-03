package domain

import "errors"

var (
	// ErrInvalidStatus indicates an invalid test result status
	ErrInvalidStatus = errors.New("invalid test result status")
	
	// ErrEmptyTitle indicates an empty or missing title
	ErrEmptyTitle = errors.New("title cannot be empty")
	
	// ErrInvalidAttachment indicates an invalid attachment
	ErrInvalidAttachment = errors.New("invalid attachment")
	
	// ErrInvalidTestCase indicates an invalid test case
	ErrInvalidTestCase = errors.New("invalid test case")
)