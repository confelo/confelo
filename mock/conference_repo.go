package mock

import (
	"context"

	"github.com/confelo/confelo"
	m "github.com/stretchr/testify/mock"
)

type ConferenceRepo struct {
	m.Mock
}

func (cr *ConferenceRepo) Save(
	c context.Context,
	conf *confelo.Conference) (confelo.ConferenceID, error) {

	args := cr.Called(c, conf)

	return args.Get(0).(confelo.ConferenceID), args.Error(1)
}

func (cr *ConferenceRepo) FindByName(
	c context.Context,
	name string) (*confelo.Conference, error) {

	args := cr.Called(c, name)

	if conf, ok := args.Get(0).(*confelo.Conference); ok {
		return conf, nil
	}

	return nil, args.Error(1)
}
