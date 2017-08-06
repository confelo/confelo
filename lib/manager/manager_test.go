package manager_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/confelo/confelo"
	"github.com/confelo/confelo/lib/manager"
	"github.com/confelo/confelo/lib/mock"
	m "github.com/stretchr/testify/mock"
)

func TestCanCreateConference(t *testing.T) {
	confStart := time.Now()
	confEnd := time.Now()

	cases := map[string]struct {
		cmd           *confelo.CreateConfCmd
		ctx           context.Context
		expectedID    confelo.ConferenceID
		expectedError error
		assertCalls   func(*mock.ConferenceDB)
		getRepo       func() *mock.ConferenceDB
	}{
		"create conference": {
			cmd: &confelo.CreateConfCmd{Name: "Foo Conference", Description: "Foo conf desc", WebsiteURL: "http://www.foo.com",
				GetTicketsURL: "https://www.foo.com/get_tickets", StartTime: confStart, EndTime: confEnd},
			ctx:           getValidContext(),
			expectedID:    1,
			expectedError: nil,
			getRepo: func() *mock.ConferenceDB {
				repo := new(mock.ConferenceDB)
				repo.On("FindByName", m.Anything, m.Anything).Return(nil, nil)
				repo.On("Save", m.Anything, &confelo.Conference{Name: "Foo Conference", Description: "Foo conf desc", WebsiteURL: "http://www.foo.com",
					GetTicketsURL: "https://www.foo.com/get_tickets", StartTime: confStart, EndTime: confEnd}).Return(confelo.ConferenceID(1), nil)
				return repo
			},
		},

		"create second conference": {
			cmd:           &confelo.CreateConfCmd{Name: "Bar Conference"},
			ctx:           getValidContext(),
			expectedID:    2,
			expectedError: nil,
			getRepo: func() *mock.ConferenceDB {
				repo := new(mock.ConferenceDB)
				repo.On("FindByName", m.Anything, m.Anything).Return(nil, nil)
				repo.On("Save", m.Anything, &confelo.Conference{Name: "Bar Conference"}).Return(confelo.ConferenceID(2), nil)
				return repo
			},
		},

		"conference exists error": {
			cmd:           &confelo.CreateConfCmd{Name: "Baz Conference"},
			ctx:           getValidContext(),
			expectedID:    0,
			expectedError: fmt.Errorf("conference already exists"),
			getRepo: func() *mock.ConferenceDB {
				repo := new(mock.ConferenceDB)
				repo.On("FindByName", m.Anything, m.Anything).Return(&confelo.Conference{}, nil)
				return repo
			},
		},

		"not in publisher role": {
			cmd:           &confelo.CreateConfCmd{Name: "Baz Conference"},
			ctx:           getInvalidContext(),
			expectedID:    0,
			expectedError: fmt.Errorf("you are not authorized to perform this action"),
			getRepo: func() *mock.ConferenceDB {
				return nil
			},
		},
	}

	i := getMockIAM()

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			repo := c.getRepo()

			svc := manager.New(repo, i)

			id, err := svc.CreateConference(c.ctx, c.cmd)

			assert.Equal(t, c.expectedID, id)
			assert.Equal(t, c.expectedError, err)

			if c.assertCalls != nil {
				c.assertCalls(repo)
			}
		})
	}
}

func getValidContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "email", "publisher@mail.com")
	return ctx
}

func getInvalidContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "email", "intruder@mail.com")
	return ctx
}

func getMockIAM() *mock.IAM {
	var p confelo.Publisher
	iam := mock.IAM{}

	iam.AddPub("", p, fmt.Errorf("you are not authorized to perform this action"))
	iam.AddPub("intruder@mail.com", p, fmt.Errorf("you are not authorized to perform this action"))
	p = confelo.NewPublisher(123, "publisher")
	iam.AddPub("publisher@mail.com", p, nil)

	return &iam
}
