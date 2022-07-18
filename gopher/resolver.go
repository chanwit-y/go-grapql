package gopher

import (
	"errors"

	"grapql/job"

	"github.com/graphql-go/graphql"
)

// type Resolver interface {
// 	// ResolveGophers should return a list of all gophers in the repository
// 	ResolveGophers(p graphql.ResolveParams) (interface{}, error)
// 	// ResolveGopher is used to respond to single queries for gophers
// 	ResolveGopher(p graphql.ResolveParams) (interface{}, error)
// }
type GopherService struct {
	gophers Repository
	jobs    job.Repository
}

// NewService is a factory that creates a new GopherService
func NewService(repo Repository, jobrepo job.Repository) GopherService {
	return GopherService{
		gophers: repo,
		jobs:    jobrepo,
	}
}

// ResolveGophers will be used to retrieve all available Gophers
func (gs GopherService) ResolveGophers(p graphql.ResolveParams) (interface{}, error) {
	// Fetch gophers from the Repository
	gophers, err := gs.gophers.GetGophers()
	if err != nil {
		return nil, err
	}
	return gophers, nil
}

func (gs *GopherService) ResolveJobs(p graphql.ResolveParams) (interface{}, error) {
	// Fetch Source Value
	g, ok := p.Source.(Gopher)

	if !ok {
		return nil, errors.New("source was not a Gopher")
	}
	// Find Jobs Based on the Gophers ID
	jobs, err := gs.jobs.GetJobs(g.ID)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
