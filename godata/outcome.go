package godata

import (
	"fmt"
	"strings"
)

type Outcome string

const (
	Active           Outcome = "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_ACTIVE"
	Deceased         Outcome = "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_DECEASED"
	Recovered        Outcome = "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_RECOVERED"
	DeceasedNonCovid Outcome = "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_DECEASED_NONCOVID"
)

type OutcomeErr struct {
	Reason string
	Inner  error
}

func (e OutcomeErr) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("outcome error: %s: %v", e.Reason, e.Inner)
	}
	return fmt.Sprintf("outcome error: %s", e.Reason)
}

func (e OutcomeErr) Unwrap() error {
	return e.Inner
}

func toOutcome(o string) (Outcome, error) {
	switch strings.ToLower(o) {
	case "active":
		return Active, nil
	case "deceased":
		return Deceased, nil
	case "recovered":
		return Recovered, nil
	case "deceased - noncovid":
		return DeceasedNonCovid, nil
	default:
		return "", OutcomeErr{
			Reason: fmt.Sprintf("outcome: %s is not a valid value", o),
			Inner:  nil,
		}
	}
}
