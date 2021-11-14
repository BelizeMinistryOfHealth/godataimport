package godata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

var (
	ErrUserAccessDenied = errors.New("you do not have access to the requested resource")
	ErrNotFound         = errors.New("the requested resource not found")
	ErrBadRequest       = errors.New("the request was not accepted by the server")
)

type getResponse struct {
	ID string `json:"id"`
}

func getCaseByVisualId(visualId, token, baseUrl string) (getResponse, error) {
	// We need the id, so we should query for it.
	filter := fmt.Sprintf("{\"where\":{\"visualId\": \"%s\"}}", visualId)
	getUrl := fmt.Sprintf("%s?filter=%s", baseUrl, url.QueryEscape(filter))
	getReq, _ := http.NewRequest(http.MethodGet, getUrl, nil)
	getReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	getReq.Header.Set("Content-Type", "application/json")
	getClient := &http.Client{}
	getResp, err := getClient.Do(getReq)
	if err != nil {
		return getResponse{}, fmt.Errorf("could not retrieve case with visualId: %s", visualId)
	}

	var resps []getResponse
	defer getResp.Body.Close()
	if err := json.NewDecoder(getResp.Body).Decode(&resps); err != nil {
		log.WithFields(log.Fields{
			"body": getResp.Body,
		}).WithError(err).Info("raw body")
		return getResponse{}, fmt.Errorf("failed to decode case data: status %d | godataApiUrl: %s : %w", getResp.StatusCode, getUrl, err)
	}
	if len(resps) == 0 {
		return getResponse{}, fmt.Errorf("did not find a matching case with visualid: %s", visualId)
	}
	ocase := resps[0]
	return ocase, nil
}

func PostTest(token, outbreak string, test CovidTest) error {
	client := &http.Client{}

	body, err := json.Marshal(test)
	if err != nil {
		return fmt.Errorf("could not encode test %s", test.VisualID)
	}

	baseUrl := fmt.Sprintf("%s/outbreaks/%s/cases", godataApiUrl, outbreak)
	req, _ := http.NewRequest(http.MethodPost, baseUrl, bytes.NewReader(body))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not post case to GoData: %s", test.VisualID)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return ErrUserAccessDenied
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusBadRequest:
		log.WithFields(log.Fields{
			"response":     resp.Body,
			"test":         test,
			"godataApiUrl": baseUrl,
			"statusCode":   resp.StatusCode,
		}).Error("server could not read the request")
		return ErrBadRequest
	case http.StatusConflict:
		// We need to update the record
		// We need the id, so we should query for it.
		ocase, err := getCaseByVisualId(test.VisualID, token, baseUrl)
		if err != nil {
			return fmt.Errorf("failed to retrieve case by visualId: %s", test.VisualID)
		}
		putBody, _ := json.Marshal(test)
		putReq, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", baseUrl, ocase.ID), bytes.NewReader(putBody))
		putReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		putReq.Header.Set("Content-Type", "application/json")
		putClient := &http.Client{}
		putResp, err := putClient.Do(putReq)
		defer putResp.Body.Close()
		if err != nil {
			return fmt.Errorf("could not post updated case to GoData: %s", test.VisualID)
		}
		if putResp.StatusCode != http.StatusOK {
			log.WithFields(log.Fields{
				"statusCode": putResp.StatusCode,
				"case":       test.VisualID,
			}).Info("failure when updating")
		}

		return nil
	case http.StatusOK, http.StatusCreated:
		return nil
	default:
		var b interface{}
		json.NewDecoder(resp.Body).Decode(b)
		log.WithFields(log.Fields{
			"response":     b,
			"test":         test,
			"godataApiUrl": baseUrl,
			"statusCode":   resp.StatusCode,
		}).Error("server rejected the request")
		return errors.New("unknown error")
	}

}
