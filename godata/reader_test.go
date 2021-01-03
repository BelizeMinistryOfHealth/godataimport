package godata

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

const username = "epidemiologyunit@health.gov.bz"
const password = "4ycu6VKWzAvAUCS"

func TestRead(t *testing.T) {
	// Get GoData Token
	authResp, err := GetToken(username, password)
	if err != nil {
		t.Fatalf("auth failed: %+v", err)
	}
	locs, err := GetLocations(authResp.Response.AccessToken)
	if err != nil {
		t.Fatalf("error retrieving locations %+v", err)
	}

	csvFile, err := os.Open("COVID19_godata_1.csv")
	if err != nil {
		t.Fatalf("failed to open csv file: %+v", err)
	}
	r := csv.NewReader(csvFile)
	tests, err := Read(r, locs)
	if err != nil {
		t.Fatalf("failed to parse csv file: %+v", err)
	}
	//t.Log(tests)
	//jsonFile, _ := os.OpenFile("20201203.json", os.O_CREATE, os.ModePerm)
	//defer jsonFile.Close()
	//err = json.NewEncoder(jsonFile).Encode(tests)
	//if err != nil {
	//	t.Fatalf("failed to write to json: %+v", err)
	//}
	jsonFile, _ := json.MarshalIndent(tests, "", "")
	err = ioutil.WriteFile("2020-12-03.json", jsonFile, os.ModePerm)
	if err != nil {
		t.Fatalf("failed to write to json: %+v", err)
	}
}
