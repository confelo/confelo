package mock

import "github.com/confelo/confelo"

type pubs struct {
	pub confelo.Publisher
	err error
}

type IAM struct {
	Pubs map[string]pubs
}

func (i *IAM) AddPub(e string, p confelo.Publisher, err error) {
	if i.Pubs == nil {
		i.Pubs = make(map[string]pubs)
	}

	i.Pubs[e] = pubs{p, err}
}

func (i *IAM) AsPublisher(email string) (confelo.Publisher, error) {
	p := i.Pubs[email]
	return p.pub, p.err
}
