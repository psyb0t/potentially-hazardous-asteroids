package types

// Asteroid is the struct that holds data about an asteroid
type Asteroid struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	NASAJPLURL        string `json:"nasa_jpl_url"`
	EstimatedDiameter struct {
		Kilometers struct {
			Min float64 `json:"min"`
			Max float64 `json:"max"`
		} `json:"kilometers"`
		Miles struct {
			Min float64 `json:"min"`
			Max float64 `json:"max"`
		} `json:"miles"`
	} `json:"estimated_diameter"`
	CloseApproachTimestamp int `json:"close_approach_timestamp"`
	RelativeVelocity       struct {
		KilometersPerHour float64 `json:"kilometers_per_hour"`
		MilesPerHour      float64 `json:"miles_per_hour"`
	} `json:"relative_velocity"`
	MissDistance struct {
		Kilometers float64 `json:"kilometers"`
		Miles      float64 `json:"miles"`
	} `json:"miss_distance"`
	IsSentryObject bool `json:"is_sentry_object"`
}

// NewAsteroid initializes an Asteroid struct and returns a pointer to it
func NewAsteroid() *Asteroid {
	return &Asteroid{}
}
