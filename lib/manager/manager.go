// Package manager provides manager application service functionality
package manager

import (
	"context"
	"fmt"

	"github.com/confelo/confelo"
)

// New constructs Manager application service
func New(repo confelo.ConferenceDB) *Manager {
	return &Manager{repo}
}

// Manager represents Manager application service
type Manager struct {
	confDB confelo.ConferenceDB
}

// Create creates conference given CreateConferenceCommand
func (cs *Manager) Create(c context.Context, cmd *confelo.CreateConferenceCommand) (confelo.ConferenceID, error) {
	conf, err := cs.confDB.FindByName(c, cmd.Name)
	if err != nil {
		return 0, err
	}

	if conf != nil {
		return 0, fmt.Errorf("conference already exists")
	}

	conf = cs.conferenceFromCmd(cmd)

	return cs.confDB.Save(c, conf)
}

func (cs *Manager) conferenceFromCmd(cmd *confelo.CreateConferenceCommand) *confelo.Conference {
	return &confelo.Conference{
		Name: cmd.Name,
	}
}
