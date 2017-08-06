// Package iam provides iam domain service
package iam

import "github.com/confelo/confelo"

type userDB interface {
}

// New instantiates IAM service
func New(userDB userDB) *IAM {
	iam := IAM{userDB}
	return &iam
}

// IAM represents IAM domain service
type IAM struct {
	// or it might require something like
	// (if going over RPC)
	userDB userDB
}

// AsPublisher checks wether or not the user associated
// with the provided email is in Publisher role or not
func (i *IAM) AsPublisher(string) (confelo.Publisher, error) {
	panic("not implemented")
}
