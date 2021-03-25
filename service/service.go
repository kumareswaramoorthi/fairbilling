package service

import (
	"bufio"
	"fairBilling/constants"
	"fairBilling/utils"
)

type Session struct {
	SessionCount int
	Duration     float64
}

type FairBillingService interface {
	CalculateSession() *map[string]Session
	MapSessions(name string, sessionTime float64)
}

type fairBillingService struct {
	file    string
	userMap map[string]Session
}

func NewFairBillingService(file string) FairBillingService {
	return &fairBillingService{
		file:    file,
		userMap: map[string]Session{},
	}
}

func (s fairBillingService) CalculateSession() *map[string]Session {
	fbUtils := utils.NewFairBilllingUtils()
	// Create new scanner to scan the .txt file
	scan := bufio.NewScanner(fbUtils.OpenFile(s.file))
	for scan.Scan() {
		isValid, parts := fbUtils.ValidateLine(scan.Text())
		if isValid {
			fbUtils.CaptureSessionEndInfo(parts)
		}
	}
	scan = bufio.NewScanner(fbUtils.OpenFile(s.file))
	for scan.Scan() {
		isValid, parts := fbUtils.ValidateLine(scan.Text())
		if isValid {
			if parts[2] == constants.StatusStart { // Check for the start of the session for a user
				var foundEndSession bool
				for id, elem := range utils.SessionEndDetails { // Range the end session list to calculate time between a particular session
					if elem.Name == parts[1] && !elem.IsCompleted { // check if the particular end session is already taken to calculation or not
						tn := fbUtils.ParseTime(parts[0])
						sessionTime := elem.Timestamp.Sub(tn).Seconds()
						if sessionTime > 0 {
							s.MapSessions(parts[1], sessionTime)           // increment session by 1 and add the session time
							utils.SessionEndDetails[id].IsCompleted = true // mark the particular value of end session as completed.
							foundEndSession = true
							break
						}
					}
				}
				// If no end session found in the list ie.,(extract_end_sessions), then the end time is the last timestamp value
				if !foundEndSession {
					tn := fbUtils.ParseTime(parts[0])
					sessionTime := utils.FinalTime.Sub(tn).Seconds()
					s.MapSessions(parts[1], sessionTime)
				}
				foundEndSession = false
			}
			// check if there is a new user, and if he dont have any particular end session details, then consider last timestamp as end time
			if _, ok := s.userMap[parts[1]]; !ok && parts[2] == constants.StatusStart {
				tn := fbUtils.ParseTime(parts[0])
				sessionTime := utils.FinalTime.Sub(tn).Seconds()
				s.MapSessions(parts[1], sessionTime)
			}

		}
	}
	//calculate the session for the user where the start time has not been mentioned, so consider the initial timestamp for calculations
	for id, elem := range utils.SessionEndDetails {
		if !elem.IsCompleted {
			sessionTime := elem.Timestamp.Sub(utils.InitialTime).Seconds()
			s.MapSessions(elem.Name, sessionTime)
			utils.SessionEndDetails[id].IsCompleted = true
		}
	}
	return &s.userMap
}

func (s fairBillingService) MapSessions(name string, sessionTime float64) {
	s.userMap[name] = Session{
		SessionCount: s.userMap[name].SessionCount + 1,
		Duration:     s.userMap[name].Duration + sessionTime,
	}
}
