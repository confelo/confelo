package manager_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	m "github.com/stretchr/testify/mock"

	"github.com/confelo/confelo"
	"github.com/confelo/confelo/lib/manager"
	"github.com/confelo/confelo/lib/mock"
)

func TestCanCreateConference(t *testing.T) {
	cases := map[string]struct {
		cmd           *confelo.CreateConferenceCommand
		expectedID    confelo.ConferenceID
		expectedError error
		assertCalls   func(*mock.ConferenceDB)
		getRepo       func() *mock.ConferenceDB
	}{
		"create conference": {
			cmd: &confelo.CreateConferenceCommand{Name: "Foo Conference"},
			/*
				cmd: &confelo.CreateConferenceCommand{Name: "Foo Conference", Description: "Foo conf desc", WebsiteURL: "http://www.foo.com",
					GetTicketsURL: "https://www.foo.com/get_tickets"},
			*/
			expectedID:    1,
			expectedError: nil,
			getRepo: func() *mock.ConferenceDB {
				repo := new(mock.ConferenceDB)
				repo.On("FindByName", m.Anything, m.Anything).Return(nil, nil)
				repo.On("Save", context.Background(), &confelo.Conference{Name: "Foo Conference"}).Return(confelo.ConferenceID(1), nil)
				return repo
			},
		},

		"create second conference": {
			cmd:           &confelo.CreateConferenceCommand{Name: "Bar Conference"},
			expectedID:    2,
			expectedError: nil,
			getRepo: func() *mock.ConferenceDB {
				repo := new(mock.ConferenceDB)
				repo.On("FindByName", m.Anything, m.Anything).Return(nil, nil)
				repo.On("Save", context.Background(), &confelo.Conference{Name: "Bar Conference"}).Return(confelo.ConferenceID(2), nil)
				return repo
			},
		},

		"conference exists error": {
			cmd:           &confelo.CreateConferenceCommand{Name: "Baz Conference"},
			expectedID:    0,
			expectedError: fmt.Errorf("conference already exists"),
			getRepo: func() *mock.ConferenceDB {
				repo := new(mock.ConferenceDB)
				repo.On("FindByName", m.Anything, m.Anything).Return(&confelo.Conference{}, nil)
				return repo
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			repo := c.getRepo()

			svc := manager.New(repo)

			id, err := svc.Create(context.Background(), c.cmd)

			assert.Equal(t, c.expectedID, id)
			assert.Equal(t, c.expectedError, err)

			if c.assertCalls != nil {
				c.assertCalls(repo)
			}
		})
	}
}
