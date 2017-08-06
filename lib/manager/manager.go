// Package manager provides conference manager use cases
// It is used by admins and other alike users to handle
// conference lifecycle (creation, deleteion, etc...)
package manager

import (
	"context"
	"fmt"

	"github.com/confelo/confelo"
)

// New constructs Manager application service
func New(repo confDB, iam confelo.IAM) *Service {
	return &Service{repo, iam}
}

// Service represents manager application service
type Service struct {
	confDB confDB
	iam    confelo.IAM
}

// CreateConference creates conference given CreateConfCmd
func (s *Service) CreateConference(c context.Context, cmd *confelo.CreateConfCmd) (confelo.ConferenceID, error) {
	_, err := s.iam.AsPublisher(emailFromCtx(c))
	if err != nil {
		return 0, err
	}

	conf, err := s.confDB.FindByName(c, cmd.Name)
	if err != nil {
		return 0, err
	}

	if conf != nil {
		return 0, fmt.Errorf("conference already exists")
	}

	conf = s.conferenceFromCmd(cmd)

	return s.confDB.Save(c, conf)
}

func (*Service) conferenceFromCmd(cmd *confelo.CreateConfCmd) *confelo.Conference {
	conf := confelo.Conference{
		Name:          cmd.Name,
		Description:   cmd.Description,
		WebsiteURL:    cmd.WebsiteURL,
		GetTicketsURL: cmd.GetTicketsURL,
		StartTime:     cmd.StartTime,
		EndTime:       cmd.EndTime,
	}
	return &conf
}

func emailFromCtx(c context.Context) string {
	e, ok := c.Value("email").(string)
	if !ok {
		return ""
	}
	return e
}
