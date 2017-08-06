package service

import (
	"log"
	"net/http"

	h "github.com/tonto/kit/http"
)

func NewAH(l *log.Logger) *AH {
	svc := &AH{
		logger: l,
	}

	svc.RegisterEndpoint("start", http.HandlerFunc(svc.startHandler), "GET")
	svc.RegisterEndpoint("stop", http.HandlerFunc(svc.stopHandler), "GET")
	svc.RegisterEndpoint("health", http.HandlerFunc(svc.healthHandler), "GET")

	return svc
}

type AH struct {
	h.BaseService

	logger *log.Logger
}

func (a *AH) Prefix() string {
	return "_ah"
}

func (s *AH) startHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("_ah/start request received")
	w.WriteHeader(http.StatusOK)
}

func (s *AH) stopHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("_ah/stop request received")
	w.WriteHeader(http.StatusOK)
}

func (s *AH) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
