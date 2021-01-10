package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"bz.epi.covid/munging/godata"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	log.SetOutput(os.Stdout)

	var username string
	var password string
	var fileName string
	var destFile string
	var outbreak string
	checkFile := false
	flag.StringVar(&username, "u", "", "Specify godata username.")
	flag.StringVar(&password, "p", "", "Specify godata password.")
	flag.StringVar(&fileName, "f", "", "Specify csv file.")
	flag.StringVar(&destFile, "d", "", "Specify destination file.")
	flag.StringVar(&outbreak, "o", "", "Specify outbreak id.")
	flag.BoolVar(&checkFile, "c", false, "set to true if you want to see the invalid records")
	flag.Parse()

	helpMsg := `

-u : The godata username
-p : The godata password
-f : the csv file with the data you want to import
-d : the destination file for dumping the json file. Useful when you want to troubleshoot. Not Required
-o : the outbreak id
-c : check the validity of the file. Set this to true if you want to see the records that might not be
     exportable. Not required.

`

	if len(username) == 0 {
		fmt.Printf("%s %s", "godata username should not be empty", helpMsg)
		os.Exit(-1)
	}
	if len(password) == 0 {
		fmt.Printf("%s %s", "godata password should not be empty", helpMsg)
		os.Exit(-1)
	}
	if len(fileName) == 0 {
		fmt.Printf("%s %s", "filename should not be empty", helpMsg)
		os.Exit(-1)
	}
	if len(destFile) == 0 {
		fmt.Printf("%s %s", "destination file should not be empty")
		os.Exit(-1)
	}

	authResp, err := godata.GetToken(username, password)
	if err != nil {
		panic(fmt.Errorf("auth failed: %w", err))
	}

	tests, err := parseFile(fileName, destFile, authResp)
	if err != nil {
		panic(fmt.Errorf("failed parsing the csv file: %w", err))
	}
	if checkFile {
		os.Exit(0)
	}

	log.WithFields(log.Fields{"numberOfTests": len(tests)}).Info("importing cases....")
	testBatches := godata.SplitTests(tests, 1000)
	// Post To GoData
	for _, tb := range testBatches {
		for _, t := range tb {
			err := godata.PostTest(authResp.Response.AccessToken, outbreak, t)
			if err != nil {
				log.WithFields(log.Fields{
					"bhisNumber": t.Bhis,
					"outbreak":   outbreak,
				}).WithError(err).Error("failed to post new outbreak")
			}
		}
		// Pausing in between batches so we don't cause a DOS
		time.Sleep(5 * time.Second)
		log.WithFields(log.Fields{
			"total": len(tb),
		}).Info("Imported cases")
	}
	log.WithFields(log.Fields{
		"total": len(tests),
	}).Info("Done.")
}

func parseFile(fileName, destFile string, authResp *godata.GoDataAuthResponse) ([]godata.CovidTest, error) {
	locs, err := godata.GetLocations(authResp.Response.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("error retrieving locations %w", err)
	}

	log.Infof("Opening file: %s", fileName)
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open csv file: %w", err)
	}
	r := csv.NewReader(csvFile)
	tests, err := godata.Read(r, locs)
	if err != nil {
		log.WithFields(log.Fields{
			"file": fileName,
		}).WithError(err).Error("error reading file")
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	jsonFile, _ := json.MarshalIndent(tests, "", "    ")
	err = ioutil.WriteFile(destFile, jsonFile, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to write to json: %w", err)
	}
	return tests, nil
}
