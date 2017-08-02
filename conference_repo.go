package confelo

import "context"

// ConferenceRepo defines conference repository interface
type ConferenceRepo interface {
	Save(context.Context, *Conference) (ConferenceID, error)
	FindByName(context.Context, string) (*Conference, error)
}
