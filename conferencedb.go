package confelo

import "context"

// ConferenceDB defines conference db interface (repository)
type ConferenceDB interface {
	Save(context.Context, *Conference) (ConferenceID, error)
	FindByName(context.Context, string) (*Conference, error)
}
