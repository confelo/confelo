package service

import (
	"log"
	"net/http"

	h "github.com/tonto/kit/http"
)

// NewAH creates new ah service
func NewAH(l *log.Logger) *AH {
	svc := &AH{
		logger: l,
	}

	svc.RegisterEndpoint("start", http.HandlerFunc(svc.startHandler), "GET")
	svc.RegisterEndpoint("stop", http.HandlerFunc(svc.stopHandler), "GET")
	svc.RegisterEndpoint("health", http.HandlerFunc(svc.healthHandler), "GET")

	return svc
}

// AH represents AH status http service
type AH struct {
	h.BaseService

	logger *log.Logger
}

// Prefix returns ah routing prefix
func (a *AH) Prefix() string {
	return "_ah"
}

func (ah *AH) startHandler(w http.ResponseWriter, r *http.Request) {
	ah.logger.Println("_ah/start request received")
	w.WriteHeader(http.StatusOK)
}

func (ah *AH) stopHandler(w http.ResponseWriter, r *http.Request) {
	ah.logger.Println("_ah/stop request received")
	w.WriteHeader(http.StatusOK)
}

func (*AH) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
