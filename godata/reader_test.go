package godata

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	username := os.Getenv("GODATA_USER")
	password := os.Getenv("GODATA_PASSWORD")
	// Get GoData Token
	authResp, err := GetToken(username, password)
	if err != nil {
		t.Fatalf("auth failed: %+v", err)
	}
	locs, err := GetLocations(authResp.AccessToken)
	if err != nil {
		t.Fatalf("error retrieving locations %+v", err)
	}

	csvFile, err := os.Open("../COVID19_godata_1.csv")
	if err != nil {
		t.Fatalf("failed to open csv file: %+v", err)
	}
	r := csv.NewReader(csvFile)
	tests, err := Read(r, locs)
	if err != nil {
		t.Fatalf("failed to parse csv file: %+v", err)
	}

	jsonFile, _ := json.MarshalIndent(tests, "", "")
	err = ioutil.WriteFile("2021-01-09.json", jsonFile, os.ModePerm)
	if err != nil {
		t.Fatalf("failed to write to json: %+v", err)
	}
}
