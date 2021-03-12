package extract_end_sessions

import "time"

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

// extract the session end details for the calculations , Initial Timestamp and Final timestamp
func EndSessions(parts []string) {
	if IsInitial == false {
		InitialTimeStamp := parts[0]
		InitialTime, _ = time.Parse("15:04:05", InitialTimeStamp)
		IsInitial = true
	}
	if parts[2] == `End` {
		ts, _ := time.Parse("15:04:05", parts[0])
		SessionEndDetails = append(SessionEndDetails, Details{
			Name:        parts[1],
			Timestamp:   ts,
			IsCompleted: false,
		})
	}
	FinalTimestamp := parts[0]
	FinalTime, _ = time.Parse("15:04:05", FinalTimestamp)

}
