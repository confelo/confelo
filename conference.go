package confelo

// ConferenceID represents Conference primary key
type ConferenceID int64

// CreateConferenceCommand defines fields needed for
// conference creation via app service
type CreateConferenceCommand struct {
	Name string
}

// Conference represents conference entity
type Conference struct {
	Name string
}
