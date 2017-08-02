package service

import (
	"context"
	"fmt"

	"github.com/confelo/confelo"
)

// NewConference constructs Conference Application service
func NewConference(repo confelo.ConferenceRepo) *Conference {
	return &Conference{repo}
}

// Conference represents Conference application service
type Conference struct {
	repo confelo.ConferenceRepo
}

// Create creates conference given CreateConferenceCommand
func (cs *Conference) Create(c context.Context, cmd *confelo.CreateConferenceCommand) (confelo.ConferenceID, error) {
	conf, err := cs.repo.FindByName(c, cmd.Name)
	if err != nil {
		return 0, err
	}

	if conf != nil {
		return 0, fmt.Errorf("conference already exists")
	}

	conf = cs.conferenceFromCmd(cmd)

	return cs.repo.Save(c, conf)
}

func (cs *Conference) conferenceFromCmd(cmd *confelo.CreateConferenceCommand) *confelo.Conference {
	return &confelo.Conference{
		Name: cmd.Name,
	}
}
