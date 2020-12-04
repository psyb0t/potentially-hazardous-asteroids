package datahandlers

import phatypes "github.com/psyb0t/potentially-hazardous-asteroids/internal/pkg/types"

// AsteroidsDataHandler is the interface defining a data handler for Asteroids
type AsteroidsDataHandler interface {
	GetAll() ([]*phatypes.Asteroid, error)
}
