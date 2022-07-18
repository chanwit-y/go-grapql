package gopher

import "github.com/graphql-go/graphql"

// type Resolver interface {
// 	// ResolveGophers should return a list of all gophers in the repository
// 	ResolveGophers(p graphql.ResolveParams) (interface{}, error)
// 	// ResolveGopher is used to respond to single queries for gophers
// 	ResolveGopher(p graphql.ResolveParams) (interface{}, error)
// }
type GopherService struct {
	gophers Repository
}

// NewService is a factory that creates a new GopherService
func NewService(repo Repository) GopherService {
	return GopherService{
		gophers: repo,
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
