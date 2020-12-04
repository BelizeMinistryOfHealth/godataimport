package godata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var baseUrl = "https://godata-dev.epi.openstep.bz/api"

type AddressLocation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func FindLocation(name string, locs []AddressLocation) *AddressLocation {
	for _, l := range locs {
		if l.Name == name {
			return &l
		}
	}
	return nil
}

func GetLocations(token string) ([]AddressLocation, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/locations", baseUrl), nil)
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
	return a, nil
}

type goDataAuthResponse struct {
	Response goDataAuthResponseBody `json:"response"`
}

type goDataAuthResponseBody struct {
	AccessToken string `json:"access_token"`
}

func GetToken(username, password string) (*goDataAuthResponse, error) {
	reqBody, err := json.Marshal(map[string]string{"username": username, "password": password})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error authenticating with GoData")
		return nil, err
	}
	req, err := http.Post(fmt.Sprintf("%s/oauth/token", baseUrl), "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error retrieving token from GoData")
		return nil, err
	}

	var authResp *goDataAuthResponse

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
