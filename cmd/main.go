package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	csv2 "bz.epi.covid/munging/csv"
	"bz.epi.covid/munging/godata"
)

func main() {
	var username string
	var password string
	var fileName string
	var destFile string
	flag.StringVar(&username, "u", "", "Specify godata username file.")
	flag.StringVar(&password, "p", "", "Specify godata password file.")
	flag.StringVar(&fileName, "f", "", "Specify csv file.")
	flag.StringVar(&destFile, "d", "", "Specify destination file.")
	flag.Parse()

	if len(username) == 0 {
		panic(errors.New("godata username should not be empty"))
	}
	if len(password) == 0 {
		panic(errors.New("godata password should not be empty"))
	}
	if len(fileName) == 0 {
		panic(errors.New("filename should not be empty"))
	}
	if len(destFile) == 0 {
		panic(errors.New("destination file should not be empty"))
	}

	authResp, err := godata.GetToken(username, password)
	if err != nil {
		panic(fmt.Errorf("auth failed: %w", err))
	}
	locs, err := godata.GetLocations(authResp.Response.AccessToken)
	if err != nil {
		panic(fmt.Errorf("error retrieving locations %w", err))
	}

	csvFile, err := os.Open("COVID19_godata_1.csv")
	if err != nil {
		panic(fmt.Errorf("failed to open csv file: %w", err))
	}

	r := csv.NewReader(csvFile)
	tests, err := csv2.Read(r, locs)
	jsonFile, _ := json.MarshalIndent(tests, "", "    ")
	err = ioutil.WriteFile(destFile, jsonFile, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("failed to write to json: %w", err))
	}

}
