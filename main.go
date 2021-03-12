package main

import (
	"bufio"
	"fairBilling/extract_end_sessions"
	"fairBilling/validate"
	"fmt"
	"log"
	"os"
	"time"
)

type Session struct {
	SessionCount int
	Duration     float64
}

var (
	UserMap         = map[string]Session{}
	foundEndSession bool
)

func main() {
	// Open the session file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Create new scanner to scan the .txt file
	scanner := bufio.NewScanner(file)
	// Scan the .txt file line by line
	for scanner.Scan() {
		// Validate each line for the timestamp, user name
		isValid, parts := validate.ValidateLine(scanner.Text())
		if isValid {
			// Store the session End details, Initial Timestamp & Final Timestamp in structs for further calculations
			extract_end_sessions.EndSessions(parts)
		}
	}
	// Calculate all the start sessions in the first step
	CalculateStartSession()

	// Calculate all the end sessions in the next step
	CalculateEndSession()

	// Print the result
	for k, v := range UserMap {
		fmt.Printf("%s %d %v \n", k, v.SessionCount, v.Duration)
	}

}

func CalculateStartSession() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanners := bufio.NewScanner(file)
	for scanners.Scan() {
		isValid, parts := validate.ValidateLine(scanners.Text())
		if isValid {
			// Check for the start of the session for a user
			if parts[2] == `Start` {
				// Range the end session list to calculate time between a particular session
				for id, elem := range extract_end_sessions.SessionEndDetails {
					// check if the particular end session is already taken to calculation or not
					if elem.Name == parts[1] && elem.IsCompleted == false {
						tn, _ := time.Parse("15:04:05", parts[0])
						sessionTime := elem.Timestamp.Sub(tn).Seconds()
						if sessionTime > 0 {
							// increment session by 1 and add the session time
							UserMap[parts[1]] = Session{
								SessionCount: UserMap[parts[1]].SessionCount + 1,
								Duration:     UserMap[parts[1]].Duration + sessionTime,
							}
							// mark the particular value of end session as completed ie., the session has ended.
							extract_end_sessions.SessionEndDetails[id].IsCompleted = true
							foundEndSession = true
							break
						}
					}
				}
				// If no end session found in the list ie.,(extract_end_sessions), then the end time is the last timestamp value
				if !foundEndSession {
					tn, _ := time.Parse("15:04:05", parts[0])
					sessionTime := extract_end_sessions.FinalTime.Sub(tn).Seconds()
					UserMap[parts[1]] = Session{
						SessionCount: UserMap[parts[1]].SessionCount + 1,
						Duration:     UserMap[parts[1]].Duration + sessionTime,
					}
				}
				foundEndSession = false
			}
			// check if there is a new user, and if he dont have any particular end session details, then consider last timestamp as end time
			if _, ok := UserMap[parts[1]]; !ok && parts[2] == `Start` {
				tn, _ := time.Parse("15:04:05", parts[0])
				sessionTime := extract_end_sessions.FinalTime.Sub(tn).Seconds()
				UserMap[parts[1]] = Session{
					SessionCount: UserMap[parts[1]].SessionCount + 1,
					Duration:     UserMap[parts[1]].Duration + sessionTime,
				}
			}
		}

	}

}

//CalculateEndSession calculates the session for the user where the start time has not been mentioned, so consider the initial timestamp for calculations
func CalculateEndSession() {
	for _, elem := range extract_end_sessions.SessionEndDetails {
		if elem.IsCompleted == false {
			sessionTime := elem.Timestamp.Sub(extract_end_sessions.InitialTime).Seconds()
			UserMap[elem.Name] = Session{
				SessionCount: UserMap[elem.Name].SessionCount + 1,
				Duration:     UserMap[elem.Name].Duration + sessionTime,
			}
		}
	}
}
