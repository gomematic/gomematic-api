package service

import (
	"github.com/gomematic/gomematic-api/pkg/service/teams"
	"github.com/gomematic/gomematic-api/pkg/service/users"
)

// Registry just stores references to all available services.
type Registry struct {
	Teams teams.Service
	Users users.Service
}

// New just initializes an empty registry so far.
func New() *Registry {
	return &Registry{}
}
