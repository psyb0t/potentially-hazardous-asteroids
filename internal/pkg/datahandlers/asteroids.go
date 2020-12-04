package datahandlers

import (
	phatypes "github.com/psyb0t/potentially-hazardous-asteroids/internal/pkg/types"
)

// GetAllAsteroids gets all of the asteroid data using the provided data handler
func GetAllAsteroids(asteroidsDataHandler AsteroidsDataHandler) ([]*phatypes.Asteroid, error) {
	asteroids, err := asteroidsDataHandler.GetAll()
	if err != nil {
		return nil, err
	}

	return asteroids, nil
}
