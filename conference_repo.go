package confelo

import "context"

type ConferenceRepo interface {
	Save(context.Context, *Conference) (ConferenceID, error)
	FindByName(context.Context, string) (*Conference, error)
}
