package manager

import (
	"context"

	"github.com/confelo/confelo"
)

type confDB interface {
	Save(context.Context, *confelo.Conference) (confelo.ConferenceID, error)
	FindByName(context.Context, string) (*confelo.Conference, error)
}
