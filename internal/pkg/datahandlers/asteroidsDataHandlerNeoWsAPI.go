package datahandlers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/psyb0t/potentially-hazardous-asteroids/internal/pkg/cache"
	phatypes "github.com/psyb0t/potentially-hazardous-asteroids/internal/pkg/types"
)

// AsteroidsDataHandlerNeoWsAPI is the struct used to handle asteroids data
// via the NeoWs API
type AsteroidsDataHandlerNeoWsAPI struct {
	apiKey string
	cache  cache.Cache
}

// NewAsteroidsDataHandlerNeoWsAPI instantiates an AsteroidsDataHandlerNeoWsAPI struct
// with the given apiKey assigned to its field with the same name
func NewAsteroidsDataHandlerNeoWsAPI(apiKey string) *AsteroidsDataHandlerNeoWsAPI {
	asteroidsDataHandlerNeoWsAPI := &AsteroidsDataHandlerNeoWsAPI{apiKey: apiKey}
	asteroidsDataHandlerNeoWsAPI.cache = make(cache.Cache)

	return asteroidsDataHandlerNeoWsAPI
}

// GetAll returns the entire asteroid list
func (h *AsteroidsDataHandlerNeoWsAPI) GetAll() ([]*phatypes.Asteroid, error) {
	var err error
	var responseData []byte

	// Try getting the response data from the cache, if it's set and the cache item is not expired
	cachedResponseItem, ok := h.cache["asteroids_response_data"]
	if !ok || cachedResponseItem.IsExpired() {
		log.Debug("cached response item not set or expired. retrieving asteroid data from NASA's API")

		apiURL := fmt.Sprintf(
			"https://api.nasa.gov/neo/rest/v1/feed?start_date=%s&api_key=%s",
			time.Now().Format("2006-01-02"), h.apiKey)

		client := http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		request, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			return nil, err
		}

		response, err := client.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return nil, ErrStatusNotOK
		}

		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		cacheItem := cache.NewItem(time.Now().Add(time.Minute*10), responseData)
		h.cache["asteroids_response_data"] = cacheItem
	} else {
		log.Debug("retrieving asteroid data from cache")

		responseData, ok = cachedResponseItem.Data.([]byte)
		if !ok {
			return nil, ErrCacheItemDataUnexpectedDataType
		}
	}

	asteroids, err := getAsteroidsFromResponseData(responseData)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(asteroids, func(i, j int) bool {
		return asteroids[i].CloseApproachTimestamp < asteroids[j].CloseApproachTimestamp
	})

	return asteroids, nil
}

func getAsteroidsFromResponseData(responseData []byte) ([]*phatypes.Asteroid, error) {
	var err error

	asteroids := make([]*phatypes.Asteroid, 0)

	responseDataStruct := struct {
		NearEarthObjects map[string][]struct {
			ID                string `json:"id"`
			Name              string `json:"name"`
			NASAJPLURL        string `json:"nasa_jpl_url"`
			EstimatedDiameter map[string]struct {
				EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
			} `json:"estimated_diameter"`
			IsPotentiallyHazardousAsteroid bool `json:"is_potentially_hazardous_asteroid"`
			CloseApproachData              []struct {
				EpochDateCloseApproach int `json:"epoch_date_close_approach"`
				RelativeVelocity       struct {
					KilometersPerHour string `json:"kilometers_per_hour"`
					MilesPerHour      string `json:"miles_per_hour"`
				} `json:"relative_velocity"`
				MissDistance struct {
					Kilometers string `json:"kilometers"`
					Miles      string `json:"miles"`
				} `json:"miss_distance"`
			} `json:"close_approach_data"`
			IsSentryObject bool `json:"is_sentry_object"`
		} `json:"near_earth_objects"`
	}{}

	if err := json.Unmarshal(responseData, &responseDataStruct); err != nil {
		return nil, err
	}

	for _, nearEarthObjectsForDate := range responseDataStruct.NearEarthObjects {
		for _, nearEarthObjectForDate := range nearEarthObjectsForDate {
			// Only interested in hazardous ones
			if !nearEarthObjectForDate.IsPotentiallyHazardousAsteroid {
				continue
			}

			asteroid := phatypes.NewAsteroid()

			asteroid.ID = nearEarthObjectForDate.ID
			asteroid.Name = nearEarthObjectForDate.Name
			asteroid.NASAJPLURL = nearEarthObjectForDate.NASAJPLURL

			asteroid.EstimatedDiameter.Kilometers.Min = nearEarthObjectForDate.EstimatedDiameter["kilometers"].EstimatedDiameterMin
			asteroid.EstimatedDiameter.Kilometers.Max = nearEarthObjectForDate.EstimatedDiameter["kilometers"].EstimatedDiameterMax

			asteroid.EstimatedDiameter.Miles.Min = nearEarthObjectForDate.EstimatedDiameter["miles"].EstimatedDiameterMin
			asteroid.EstimatedDiameter.Miles.Max = nearEarthObjectForDate.EstimatedDiameter["miles"].EstimatedDiameterMax

			asteroid.CloseApproachTimestamp = nearEarthObjectForDate.CloseApproachData[0].EpochDateCloseApproach

			asteroid.RelativeVelocity.KilometersPerHour, err = strconv.ParseFloat(
				nearEarthObjectForDate.CloseApproachData[0].RelativeVelocity.KilometersPerHour, 64)

			if err != nil {
				return nil, err
			}

			asteroid.RelativeVelocity.MilesPerHour, err = strconv.ParseFloat(
				nearEarthObjectForDate.CloseApproachData[0].RelativeVelocity.MilesPerHour, 64)

			if err != nil {
				return nil, err
			}

			asteroid.MissDistance.Kilometers, err = strconv.ParseFloat(
				nearEarthObjectForDate.CloseApproachData[0].MissDistance.Kilometers, 64)

			if err != nil {
				return nil, err
			}

			asteroid.MissDistance.Miles, err = strconv.ParseFloat(
				nearEarthObjectForDate.CloseApproachData[0].MissDistance.Miles, 64)

			if err != nil {
				return nil, err
			}

			asteroid.IsSentryObject = nearEarthObjectForDate.IsSentryObject

			asteroids = append(asteroids, asteroid)
		}
	}

	return asteroids, nil
}
