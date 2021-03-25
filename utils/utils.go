package utils

import (
	"fairBilling/constants"
	"os"
	"regexp"
	"strings"
	"time"
)

type Details struct {
	Name        string
	Timestamp   time.Time
	IsCompleted bool
}

var (
	IsInitial         = false
	InitialTime       time.Time
	FinalTime         time.Time
	SessionEndDetails []Details
)

type FairBilllingUtils interface {
	ValidateLine(text string) (bool, []string)
	CaptureSessionEndInfo(parts []string)
	OpenFile(filePath string) *os.File
	ParseTime(timeStr string) time.Time
}

type fairBillingUtils struct {
}

func NewFairBilllingUtils() FairBilllingUtils {
	return fairBillingUtils{}
}

// Validate the line for 2 spaces, 8 chars in timestamp with proper regex and alphanumeric value in user name.
func (f fairBillingUtils) ValidateLine(text string) (bool, []string) {
	spaces := strings.Count(text, " ")
	if spaces == constants.LineSpaces {
		parts := strings.Split(text, " ")
		lenTimestamp := len(parts[0])
		timeStampValid := regexp.MustCompile(constants.TimestampRegex).MatchString(parts[0]) //regex validation
		nameValid := regexp.MustCompile(constants.NameRegex).MatchString(parts[1])
		if lenTimestamp == constants.LenTimestamp && timeStampValid && nameValid {
			return true, parts
		}
	}
	return false, nil
}

//CaptureSessionEndInfo
func (f fairBillingUtils) CaptureSessionEndInfo(parts []string) {
	if !IsInitial {
		InitialTimeStamp := parts[0]
		InitialTime = f.ParseTime(InitialTimeStamp)
		IsInitial = true
	}
	if parts[2] == constants.StatusEnd {
		ts := f.ParseTime(parts[0])
		SessionEndDetails = append(SessionEndDetails, Details{
			Name:        parts[1],
			Timestamp:   ts,
			IsCompleted: false,
		})
	}
	FinalTimestamp := parts[0]
	FinalTime = f.ParseTime(FinalTimestamp)
}

//open File
func (f fairBillingUtils) OpenFile(filePath string) *os.File {
	file, _ := os.Open(filePath)
	return file
}

//Parse the time according to the format present in input file
func (f fairBillingUtils) ParseTime(timeStr string) time.Time {
	tm, _ := time.Parse(constants.TimeFormat, timeStr)
	return tm
}
