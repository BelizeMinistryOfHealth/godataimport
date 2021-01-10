package godata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

var godataApiUrl = "https://godata-dev.epi.openstep.bz/api"

const GEO_ADMIN_LEVEL_1 = "LNG_REFERENCE_DATA_CATEGORY_LOCATION_GEOGRAPHICAL_LEVEL_ADMIN_LEVEL_1"
const BZ_COUNTRY_ID = "69a4f321-2f56-4cc3-8b36-e5c7ff18b86b"

type AddressLocation struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Active            bool   `json:"active"`
	GeographicLevelId string `json:"geographicalLevelId"`
	ParentLocationId  string `json:"parentLocationId"`
}

func FindLocation(name string, locs []AddressLocation) *AddressLocation {
	for _, l := range locs {
		if strings.ToLower(l.Name) == strings.Trim(strings.ToLower(name), "") && l.Active {
			return &l
		}
	}
	return nil
}

// GetLocations retrieves all the locations from GoData
// It filters out all the districts that are assigned to the wrong parent country.
func GetLocations(token string) ([]AddressLocation, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/locations", godataApiUrl), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error fetching locations: %+v", err)
	}
	var a []AddressLocation

	if err := json.NewDecoder(resp.Body).Decode(&a); err != nil {
		return nil, fmt.Errorf("error reading location response from godata: %w", err)
	}

	// Filter out all the invalid addresses.
	// Invalid addresses are those Districts that have the wrong parent location id.
	var addresses []AddressLocation
	for _, i := range a {
		if i.GeographicLevelId == GEO_ADMIN_LEVEL_1 && i.ParentLocationId == BZ_COUNTRY_ID {
			addresses = append(addresses, i)
		}
		if i.GeographicLevelId != GEO_ADMIN_LEVEL_1 && i.Active {
			addresses = append(addresses, i)
		}
	}
	return addresses, nil
}

type GoDataAuthResponse struct {
	Response goDataAuthResponseBody `json:"response"`
}

type goDataAuthResponseBody struct {
	AccessToken string `json:"access_token"`
}

func GetToken(username, password string) (*GoDataAuthResponse, error) {
	reqBody, err := json.Marshal(map[string]string{"username": username, "password": password})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error authenticating with GoData")
		return nil, err
	}
	req, err := http.Post(fmt.Sprintf("%s/oauth/token", godataApiUrl), "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error retrieving token from GoData")
		return nil, err
	}

	var authResp *GoDataAuthResponse

	if err := json.NewDecoder(req.Body).Decode(&authResp); err != nil {
		log.WithFields(log.Fields{"error": err, "response": req}).Error("failed to decode oauth token from godata")
		return nil, err
	}
	if req.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{"responseFromGoData": req}).Error("auth with godata failed")
		return nil, fmt.Errorf("auth with godata failed")
	}
	return authResp, nil

}
