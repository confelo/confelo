package confelo

import "time"

// ConferenceID represents Conference primary key
type ConferenceID int64

// CreateConfCmd defines fields needed for
// conference creation via app service
type CreateConfCmd struct {
	Name          string
	Description   string
	WebsiteURL    string
	GetTicketsURL string
	StartTime     time.Time
	EndTime       time.Time
}

// Conference represents conference entity
type Conference struct {
	Name          string
	Description   string
	WebsiteURL    string
	GetTicketsURL string
	StartTime     time.Time
	EndTime       time.Time
}
