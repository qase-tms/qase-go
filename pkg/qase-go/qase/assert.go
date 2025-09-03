package qase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Basic assertions - simplified versions that just delegate to testify
func True(t *testing.T, value bool, msgAndArgs ...interface{}) {
	assert.True(t, value, msgAndArgs...)
}

func False(t *testing.T, value bool, msgAndArgs ...interface{}) {
	assert.False(t, value, msgAndArgs...)
}

func Equal(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assert.Equal(t, expected, actual, msgAndArgs...)
}

func NotEqual(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assert.NotEqual(t, expected, actual, msgAndArgs...)
}

func EqualValues(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assert.EqualValues(t, expected, actual, msgAndArgs...)
}

func NotEqualValues(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assert.NotEqualValues(t, expected, actual, msgAndArgs...)
}

// Error assertions
func Error(t *testing.T, err error, msgAndArgs ...interface{}) {
	assert.Error(t, err, msgAndArgs...)
}

func NoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	assert.NoError(t, err, msgAndArgs...)
}

func EqualError(t *testing.T, theError error, errString string, msgAndArgs ...interface{}) {
	assert.EqualError(t, theError, errString, msgAndArgs...)
}

func ErrorIs(t *testing.T, err error, target error, msgAndArgs ...interface{}) {
	assert.ErrorIs(t, err, target, msgAndArgs...)
}

func ErrorAs(t *testing.T, err error, target interface{}, msgAndArgs ...interface{}) {
	assert.ErrorAs(t, err, target, msgAndArgs...)
}

// Nil assertions
func Nil(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	assert.Nil(t, object, msgAndArgs...)
}

func NotNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	assert.NotNil(t, object, msgAndArgs...)
}

// Collection assertions
func Len(t *testing.T, object interface{}, length int, msgAndArgs ...interface{}) {
	assert.Len(t, object, length, msgAndArgs...)
}

func Contains(t *testing.T, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	assert.Contains(t, s, contains, msgAndArgs...)
}

func NotContains(t *testing.T, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	assert.NotContains(t, s, contains, msgAndArgs...)
}

func ElementsMatch(t *testing.T, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	assert.ElementsMatch(t, listA, listB, msgAndArgs...)
}

func Same(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	assert.Same(t, expected, actual, msgAndArgs...)
}

func NotSame(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	assert.NotSame(t, expected, actual, msgAndArgs...)
}

func Subset(t *testing.T, list, subset interface{}, msgAndArgs ...interface{}) {
	assert.Subset(t, list, subset, msgAndArgs...)
}

func NotSubset(t *testing.T, list, subset interface{}, msgAndArgs ...interface{}) {
	assert.NotSubset(t, list, subset, msgAndArgs...)
}

// Numeric assertions
func Greater(t *testing.T, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assert.Greater(t, e1, e2, msgAndArgs...)
}

func GreaterOrEqual(t *testing.T, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assert.GreaterOrEqual(t, e1, e2, msgAndArgs...)
}

func Less(t *testing.T, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assert.Less(t, e1, e2, msgAndArgs...)
}

func LessOrEqual(t *testing.T, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assert.LessOrEqual(t, e1, e2, msgAndArgs...)
}

func InDelta(t *testing.T, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	assert.InDelta(t, expected, actual, delta, msgAndArgs...)
}

// Type assertions
func IsType(t *testing.T, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	assert.IsType(t, expectedType, object, msgAndArgs...)
}

func Implements(t *testing.T, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	assert.Implements(t, interfaceObject, object, msgAndArgs...)
}

// Empty/Zero assertions
func Empty(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	assert.Empty(t, object, msgAndArgs...)
}

func NotEmpty(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	assert.NotEmpty(t, object, msgAndArgs...)
}

func Zero(t *testing.T, i interface{}, msgAndArgs ...interface{}) {
	assert.Zero(t, i, msgAndArgs...)
}

func NotZero(t *testing.T, i interface{}, msgAndArgs ...interface{}) {
	assert.NotZero(t, i, msgAndArgs...)
}

// String assertions
func Regexp(t *testing.T, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	assert.Regexp(t, rx, str, msgAndArgs...)
}

func NotRegexp(t *testing.T, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	assert.NotRegexp(t, rx, str, msgAndArgs...)
}

// JSON assertions
func JSONEq(t *testing.T, expected, actual string, msgAndArgs ...interface{}) {
	assert.JSONEq(t, expected, actual, msgAndArgs...)
}

// Time assertions
func WithinDuration(t *testing.T, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	assert.WithinDuration(t, expected, actual, delta, msgAndArgs...)
}

// File system assertions
func DirExists(t *testing.T, path string, msgAndArgs ...interface{}) {
	assert.DirExists(t, path, msgAndArgs...)
}

func FileExists(t *testing.T, path string, msgAndArgs ...interface{}) {
	assert.FileExists(t, path, msgAndArgs...)
}

// Custom condition
func Condition(t *testing.T, condition assert.Comparison, msgAndArgs ...interface{}) {
	assert.Condition(t, condition, msgAndArgs...)
}
