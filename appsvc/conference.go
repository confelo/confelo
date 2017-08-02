package appsvc

import (
	"context"
	"fmt"

	"github.com/confelo/confelo"
)

func NewConference(repo confelo.ConferenceRepo) *ConferenceService {
	return &ConferenceService{repo}
}

type ConferenceService struct {
	repo confelo.ConferenceRepo
}

func (cs *ConferenceService) Create(c context.Context, cmd *confelo.CreateConferenceCommand) (confelo.ConferenceID, error) {
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

func (cs *ConferenceService) conferenceFromCmd(cmd *confelo.CreateConferenceCommand) *confelo.Conference {
	return &confelo.Conference{
		Name: cmd.Name,
	}
}
