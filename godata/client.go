package godata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	ErrUserAccessDenied = errors.New("you do not have access to the requested resource")
	ErrNotFound         = errors.New("the requested resource not found")
	ErrBadRequest       = errors.New("the request was not accepted by the server")
)

func PostTest(token, outbreak string, test CovidTest) error {
	client := &http.Client{}

	body, err := json.Marshal(test)
	if err != nil {
		return fmt.Errorf("could not encode test %s", test.VisualID)
	}

	url := fmt.Sprintf("%s/outbreaks/%s/cases", baseUrl, outbreak)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("could not post case to GoData: %s", test.VisualID)
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return ErrUserAccessDenied
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusBadRequest:
		log.WithFields(log.Fields{
			"response":   resp.Body,
			"test":       test,
			"url":        url,
			"statusCode": resp.StatusCode,
		}).Error("server could not read the request")
		return ErrBadRequest
	case http.StatusOK, http.StatusCreated:
		return nil
	default:
		var b interface{}
		json.NewDecoder(resp.Body).Decode(b)
		log.WithFields(log.Fields{
			"response":   b,
			"test":       test,
			"url":        url,
			"statusCode": resp.StatusCode,
		}).Error("server rejected the request")
		return errors.New("unknown error")
	}

}
