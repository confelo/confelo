package service

import (
	"context"
	"net/http"
	"time"

	"github.com/confelo/confelo"
	"github.com/confelo/confelo/lib/manager"
	h "github.com/tonto/kit/http"
	"github.com/tonto/kit/http/respond"
)

// NewManager creates new manager http service
// TODO - Add rate limiting options to BaseService
// eg. WithTimeout(...) ...
func NewManager() *Manager {
	m := Manager{}
	m.RegisterEndpoint("create_conference", m.HandlerFromMethod(m.createConference), "POST")
	return &m
}

// Manager represents manager http service
type Manager struct {
	h.BaseService

	mgr manager.Service
}

// Prefix returns service routing prefix
func (*Manager) Prefix() string {
	return "manager"
}

type createReq struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	WebsiteURL    string    `json:"website_url"`
	GetTicketsURL string    `json:"get_tickets_url"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
}

type createResp struct {
	ID confelo.ConferenceID `json:"id"`
}

func (m *Manager) createConference(w http.ResponseWriter, r *http.Request, req *createReq) {
	c := context.Background()

	id, err := m.mgr.CreateConference(
		c,
		&confelo.CreateConfCmd{
			Name:          req.Name,
			Description:   req.Description,
			WebsiteURL:    req.WebsiteURL,
			GetTicketsURL: req.GetTicketsURL,
			StartTime:     req.StartTime,
			EndTime:       req.EndTime,
		},
	)
	if err != nil {
		respond.With(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.With(w, r, http.StatusOK, createResp{
		ID: id,
	})
}
