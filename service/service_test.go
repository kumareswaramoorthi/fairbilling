package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSession(t *testing.T) {
	fbService := NewFairBillingService("../test_files/test_input.txt")
	actualOutput := fbService.CalculateSession()
	expectedOutput := map[string]Session{
		"ALICE99": {SessionCount: 4, Duration: 240},
		"CHARLIE": {SessionCount: 3, Duration: 37},
	}
	assert.NotNil(t, actualOutput)
	assert.Equal(t, &expectedOutput, actualOutput)
}

func TestCalculateSessionHavingPreperEndStatus(t *testing.T) {
	fbService := NewFairBillingService("../test_files/test_proper_end_status.txt")
	actualOutput := fbService.CalculateSession()
	expectedOutput := map[string]Session{
		"ALICE99": {SessionCount: 2, Duration: 33},
		"CHARLIE": {SessionCount: 2, Duration: 121},
	}
	assert.NotNil(t, actualOutput)
	assert.Equal(t, &expectedOutput, actualOutput)
}

func TestCalculateSessionInputHavingError(t *testing.T) {
	fbService := NewFairBillingService("../test_files/test_error.txt")
	actualOutput := fbService.CalculateSession()
	expectedOutput := map[string]Session{}
	assert.Equal(t, &expectedOutput, actualOutput)
}
