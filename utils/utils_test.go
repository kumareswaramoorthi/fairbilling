package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	FairBillingUtils = NewFairBilllingUtils()
)

func TestValidateLine(t *testing.T) {
	input := "14:02:03 ALICE99 Start"
	expectedCondition := true
	expectedParts := []string{"14:02:03", "ALICE99", "Start"}
	actualCondition, actualParts := FairBillingUtils.ValidateLine(input)
	assert.Equal(t, expectedParts, actualParts)
	assert.Equal(t, expectedCondition, actualCondition)
}

func TestValidateLineFailIfTimeStampNotValid(t *testing.T) {
	input := "14:A2:037 ALICE99 Start"
	expectedCondition := false
	expectedParts := []string(nil)
	actualCondition, actualParts := FairBillingUtils.ValidateLine(input)
	assert.Equal(t, expectedParts, actualParts)
	assert.Equal(t, expectedCondition, actualCondition)
}

func TestValidateLineFailIfNameNotValid(t *testing.T) {
	input := "14:02:03 ALICE99&* Start"
	expectedCondition := false
	expectedParts := []string(nil)
	actualCondition, actualParts := FairBillingUtils.ValidateLine(input)
	assert.Equal(t, expectedParts, actualParts)
	assert.Equal(t, expectedCondition, actualCondition)
}

func TestValidateLineFailIfFormatNotValid(t *testing.T) {
	input := "14:02:03ALICE99 Start"
	expectedCondition := false
	expectedParts := []string(nil)
	actualCondition, actualParts := FairBillingUtils.ValidateLine(input)
	assert.Equal(t, expectedParts, actualParts)
	assert.Equal(t, expectedCondition, actualCondition)
}

func TestOpenFile(t *testing.T) {
	input := "../test_files/test_input.txt"
	actualOutput := FairBillingUtils.OpenFile(input)
	assert.NotNil(t, actualOutput)
}

func TestOpenFileReturnsNilIfNoFile(t *testing.T) {
	input := "../test_files/test_test.txt"
	actualOutput := FairBillingUtils.OpenFile(input)
	assert.Nil(t, actualOutput)

}

func TestParseTime(t *testing.T) {
	input := "14:03:37"
	actualtime := FairBillingUtils.ParseTime(input)
	actualOutput := fmt.Sprint(actualtime)
	expectedOutput := "0000-01-01 14:03:37 +0000 UTC"
	assert.NotNil(t, actualOutput)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestCaptureSessionEndSkipsIfStatusIsStart(t *testing.T) {
	input := []string{"14:03:33", "ALICE99", "Start"}
	FairBillingUtils.CaptureSessionEndInfo(input)
	assert.Equal(t, "0000-01-01 14:03:33 +0000 UTC", fmt.Sprint(InitialTime))
	assert.Nil(t, SessionEndDetails)
}
func TestCaptureSessionEndInfo(t *testing.T) {
	input := []string{"14:03:33", "ALICE99", "End"}
	FairBillingUtils.CaptureSessionEndInfo(input)
	assert.Equal(t, "0000-01-01 14:03:33 +0000 UTC", fmt.Sprint(InitialTime))
	assert.NotNil(t, SessionEndDetails)
	assert.Equal(t, false, SessionEndDetails[0].IsCompleted)
	assert.Equal(t, "ALICE99", SessionEndDetails[0].Name)
	assert.Equal(t, "0000-01-01 14:03:33 +0000 UTC", fmt.Sprint(SessionEndDetails[0].Timestamp))
}
