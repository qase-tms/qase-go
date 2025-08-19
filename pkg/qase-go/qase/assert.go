package qase

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"math"

	"github.com/stretchr/testify/assert"
)

// Basic assertions
func True(t *testing.T, value bool, msgAndArgs ...interface{}) {
	success := assert.True(t, value, msgAndArgs...)
	if !success {
		// Try to capture detailed error information
		if currentResult := getCurrentTestResult(); currentResult != nil {
			// Set error details with expected/actual values
			currentResult.SetErrorDetailsWithValues(
				"Should be true",
				"true",
				fmt.Sprintf("%v", value),
				"Value should be true but was false",
			)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}

		AddMessage(fmt.Sprintf("Should be true, but got: %v", value))
		t.Fail()
	}
}

func False(t *testing.T, value bool, msgAndArgs ...interface{}) {
	success := assert.False(t, value, msgAndArgs...)
	if !success {
		// Try to capture detailed error information
		if currentResult := getCurrentTestResult(); currentResult != nil {
			// Set error details with expected/actual values
			currentResult.SetErrorDetailsWithValues(
				"Should be false",
				"false",
				fmt.Sprintf("%v", value),
				"Value should be false but was true",
			)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}

		AddMessage(fmt.Sprintf("Should be false, but got: %v", value))
		t.Fail()
	}
}

func Equal(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	success := assert.Equal(t, expected, actual, msgAndArgs...)
	if !success {
		// Try to capture detailed error information
		if currentResult := getCurrentTestResult(); currentResult != nil {
			// Set error details with expected/actual values
			currentResult.SetErrorDetailsWithValues(
				"Not equal",
				fmt.Sprintf("%v", expected),
				fmt.Sprintf("%v", actual),
				fmt.Sprintf("Expected %v but got %v", expected, actual),
			)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}

		AddMessage(fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected, actual))
		t.Fail()
	}
}

func NotEqual(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	success := assert.NotEqual(t, expected, actual, msgAndArgs...)
	if !success {
		// Try to capture detailed error information
		if currentResult := getCurrentTestResult(); currentResult != nil {
			// Set error details with expected/actual values
			currentResult.SetErrorDetailsWithValues(
				"Should not be",
				fmt.Sprintf("not %v", expected),
				fmt.Sprintf("%v", actual),
				fmt.Sprintf("Value should not be %v", expected),
			)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}

		AddMessage(fmt.Sprintf("Should not be equal to %v, but got: %v", expected, actual))
		t.Fail()
	}
}

func EqualValues(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	success := assert.EqualValues(t, expected, actual, msgAndArgs...)
	if !success {
		// Try to capture detailed error information
		if currentResult := getCurrentTestResult(); currentResult != nil {
			// Set error details with expected/actual values
			currentResult.SetErrorDetailsWithValues(
				"Not equal values",
				fmt.Sprintf("%v", expected),
				fmt.Sprintf("%v", actual),
				fmt.Sprintf("Expected values %v but got %v", expected, actual),
			)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}

		AddMessage(fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected, actual))
		t.Fail()
	}
}

func NotEqualValues(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	success := assert.NotEqualValues(t, expected, actual, msgAndArgs...)
	if !success {
		// Try to capture detailed error information
		if currentResult := getCurrentTestResult(); currentResult != nil {
			// Set error details with expected/actual values
			currentResult.SetErrorDetailsWithValues(
				"Should not be equal values",
				fmt.Sprintf("not %v", expected),
				fmt.Sprintf("%v", actual),
				fmt.Sprintf("Values should not be equal to %v", expected),
			)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}

		AddMessage(fmt.Sprintf("Should not be equal values, but got: %v", actual))
		t.Fail()
	}
}

// Error assertions
func Error(t *testing.T, err error, msgAndArgs ...interface{}) {
	success := assert.Error(t, err, msgAndArgs...)
	if !success {
		// Try to capture detailed error information
		if currentResult := getCurrentTestResult(); currentResult != nil {
			// Set error details with expected/actual values
			currentResult.SetErrorDetailsWithValues(
				"Expected error but got nil",
				"error",
				"nil",
				fmt.Sprintf("An error is expected but got nil. Received: %v", err),
			)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}

		AddMessage(fmt.Sprintf("An error is expected but got nil. Received: %v", err))
		t.Fail()
	}
}

func NoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	success := assert.NoError(t, err, msgAndArgs...)
	if !success {
		// Try to capture detailed error information
		if currentResult := getCurrentTestResult(); currentResult != nil {
			// Set error details with expected/actual values
			currentResult.SetErrorDetailsWithValues(
				"Unexpected error",
				"nil",
				fmt.Sprintf("%v", err),
				fmt.Sprintf("Received unexpected error: %v", err),
			)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}

		AddMessage(fmt.Sprintf("Received unexpected error:\n%+v", err))
		t.Fail()
	}
}

func EqualError(t *testing.T, theError error, errString string, msgAndArgs ...interface{}) {
	var actualString string

	if theError != nil {
		actualString = theError.Error()
	}
	success := assert.EqualError(t, theError, errString, msgAndArgs...)

	if !success {
		// Try to capture detailed error information
		if currentResult := getCurrentTestResult(); currentResult != nil {
			// Set error details with expected/actual values
			currentResult.SetErrorDetailsWithValues(
				"Error message not equal",
				errString,
				actualString,
				fmt.Sprintf("Expected error message '%s' but got '%s'", errString, actualString),
			)

			// Try to capture file and line information
			if file, line := getCallerInfo(); file != "" && line > 0 {
				currentResult.SetErrorLocation(file, line)
			}
		}

		AddMessage(fmt.Sprintf("Error message not equal:\n"+
			"expected: %q\n"+
			"actual  : %q", errString, actualString))
		t.Fail()
	}
}

func ErrorIs(t *testing.T, err error, target error, msgAndArgs ...interface{}) {
	var (
		actualString string
		targetString string
	)

	if target != nil {
		targetString = target.Error()
	}

	if err != nil {
		actualString = err.Error()
	}
	success := assert.ErrorIs(t, err, target, msgAndArgs...)

	if !success {
		AddMessage(fmt.Sprintf("Target error should be in err chain:\n"+
			"expected: %s\n"+
			"in chain: %s\n"+
			"Target error was not found in the error chain", targetString, actualString))
		t.Fail()
	}
}

func ErrorAs(t *testing.T, err error, target interface{}, msgAndArgs ...interface{}) {
	success := assert.ErrorAs(t, err, target, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Error should be of type %T, but was %T", target, err))
		t.Fail()
	}
}

// Nil assertions
func Nil(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	success := assert.Nil(t, object, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Expected nil, but got: %#v", object))
		t.Fail()
	}
}

func NotNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	success := assert.NotNil(t, object, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Expected value not to be nil, but got: %v", object))
		t.Fail()
	}
}

// Collection assertions
func Len(t *testing.T, object interface{}, length int, msgAndArgs ...interface{}) {
	success := assert.Len(t, object, length, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Object should have %d item(s), but has %d", length, getLength(object)))
		t.Fail()
	}
}

func Contains(t *testing.T, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	success := assert.Contains(t, s, contains, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("%#v does not contain %#v", s, contains))
		t.Fail()
	}
}

func NotContains(t *testing.T, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	success := assert.NotContains(t, s, contains, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("%#v should not contain %#v, but it does", s, contains))
		t.Fail()
	}
}

func ElementsMatch(t *testing.T, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	success := assert.ElementsMatch(t, listA, listB, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Elements do not match:\nlistA: %v\nlistB: %v", listA, listB))
		t.Fail()
	}
}

func Same(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	success := assert.Same(t, expected, actual, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Expected same pointer, but got different pointers:\nexpected: %p\nactual  : %p", expected, actual))
		t.Fail()
	}
}

func NotSame(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	success := assert.NotSame(t, expected, actual, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Expected different pointers, but got same pointer: %p", expected))
		t.Fail()
	}
}

func Subset(t *testing.T, list, subset interface{}, msgAndArgs ...interface{}) {
	success := assert.Subset(t, list, subset, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("%v is not a subset of %v. Check if all elements from subset exist in the main list", subset, list))
		t.Fail()
	}
}

func NotSubset(t *testing.T, list, subset interface{}, msgAndArgs ...interface{}) {
	success := assert.NotSubset(t, list, subset, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("%v should not be a subset of %v, but it is", subset, list))
		t.Fail()
	}
}

// Numeric assertions
func Greater(t *testing.T, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	success := assert.Greater(t, e1, e2, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("%v is not greater than %v. %v <= %v", e1, e2, e1, e2))
		t.Fail()
	}
}

func GreaterOrEqual(t *testing.T, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	success := assert.GreaterOrEqual(t, e1, e2, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("%v is not greater than or equal to %v. %v < %v", e1, e2, e1, e2))
		t.Fail()
	}
}

func Less(t *testing.T, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	success := assert.Less(t, e1, e2, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("%v is not less than %v. %v >= %v", e1, e2, e1, e2))
		t.Fail()
	}
}

func LessOrEqual(t *testing.T, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	success := assert.LessOrEqual(t, e1, e2, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("%v is not less than or equal to %v. %v > %v", e1, e2, e1, e2))
		t.Fail()
	}
}

func InDelta(t *testing.T, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	success := assert.InDelta(t, expected, actual, delta, msgAndArgs...)
	if !success {
		// Calculate actual difference for better error message
		var actualDiff float64
		switch e := expected.(type) {
		case float64:
			if a, ok := actual.(float64); ok {
				actualDiff = math.Abs(e - a)
			}
		case float32:
			if a, ok := actual.(float32); ok {
				actualDiff = float64(math.Abs(float64(e - a)))
			}
		case int:
			if a, ok := actual.(int); ok {
				actualDiff = math.Abs(float64(e - a))
			}
		case int64:
			if a, ok := actual.(int64); ok {
				actualDiff = math.Abs(float64(e - a))
			}
		default:
			actualDiff = 0
		}

		AddMessage(fmt.Sprintf("Difference between %v and %v should be less than %v, but was %v", expected, actual, delta, actualDiff))
		t.Fail()
	}
}

// Type assertions
func IsType(t *testing.T, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	success := assert.IsType(t, expectedType, object, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Object expected to be of type %v, but was %v", reflect.TypeOf(expectedType), reflect.TypeOf(object)))
		t.Fail()
	}
}

func Implements(t *testing.T, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	success := assert.Implements(t, interfaceObject, object, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("%T must implement %v, but does not", object, interfaceObject))
		t.Fail()
	}
}

// Empty/Zero assertions
func Empty(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	success := assert.Empty(t, object, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Should be empty, but was %v", object))
		t.Fail()
	}
}

func NotEmpty(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	success := assert.NotEmpty(t, object, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Should NOT be empty, but was %v", object))
		t.Fail()
	}
}

func Zero(t *testing.T, i interface{}, msgAndArgs ...interface{}) {
	success := assert.Zero(t, i, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Should be zero, but was %v", i))
		t.Fail()
	}
}

func NotZero(t *testing.T, i interface{}, msgAndArgs ...interface{}) {
	success := assert.NotZero(t, i, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Should not be zero, but was %v", i))
		t.Fail()
	}
}

// String assertions
func Regexp(t *testing.T, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	success := assert.Regexp(t, rx, str, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("String \"%v\" does not match regexp \"%v\"", str, rx))
		t.Fail()
	}
}

func NotRegexp(t *testing.T, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	success := assert.NotRegexp(t, rx, str, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("String \"%v\" should not match regexp \"%v\", but it does", str, rx))
		t.Fail()
	}
}

// JSON assertions
func JSONEq(t *testing.T, expected, actual string, msgAndArgs ...interface{}) {
	success := assert.JSONEq(t, expected, actual, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("JSONs are not equal:\n"+
			"expected: %s\n"+
			"actual  : %s\n"+
			"Please check the differences above", expected, actual))
		t.Fail()
	}
}

// Time assertions
func WithinDuration(t *testing.T, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	success := assert.WithinDuration(t, expected, actual, delta, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Max difference between %v and %v allowed is %v, but was %v", expected, actual, delta, actual.Sub(expected).Abs()))
		t.Fail()
	}
}

// File system assertions
func DirExists(t *testing.T, path string, msgAndArgs ...interface{}) {
	success := assert.DirExists(t, path, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("Directory \"%v\" does not exist or is not a directory", path))
		t.Fail()
	}
}

func FileExists(t *testing.T, path string, msgAndArgs ...interface{}) {
	success := assert.FileExists(t, path, msgAndArgs...)
	if !success {
		AddMessage(fmt.Sprintf("File \"%v\" does not exist or is not a file", path))
		t.Fail()
	}
}

// Custom condition
func Condition(t *testing.T, condition assert.Comparison, msgAndArgs ...interface{}) {
	success := assert.Condition(t, condition, msgAndArgs...)
	if !success {
		AddMessage("Condition failed! The custom condition returned false.")
		t.Fail()
	}
}

// Helper function to get trace information
func getQaseTrace() string {
	traces := make([]string, 0)
	for _, v := range assert.CallerInfo() {
		if strings.Contains(v, "pkg/qase-go") || strings.Contains(v, "jtolds/gls") {
			continue
		}
		traces = append(traces, v)
	}

	return strings.Join(traces, "\n")
}

// Helper function to get the length of an object
func getLength(object interface{}) int {
	v := reflect.ValueOf(object)
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return v.Len()
	default:
		return 0
	}
}
