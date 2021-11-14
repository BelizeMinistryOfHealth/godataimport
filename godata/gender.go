package godata

import (
	"fmt"
	"strings"
)

type Gender string

const (
	Male   Gender = "LNG_REFERENCE_DATA_CATEGORY_GENDER_MALE"
	Female Gender = "LNG_REFERENCE_DATA_CATEGORY_GENDER_FEMALE"
)

type GenderErr struct {
	Reason string
	Inner  error
}

func (e GenderErr) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("gender error: %s: %v", e.Reason, e.Inner)
	}
	return fmt.Sprintf("gender error: %s", e.Reason)
}

func (e GenderErr) Unwrap() error {
	return e.Inner
}

func toGender(g string) (Gender, error) {
	switch strings.ToLower(g) {
	case "male":
		return Male, nil
	case "female":
		return Female, nil
	default:
		return "", GenderErr{
			Reason: fmt.Sprintf("gender: %s is not a valid value", g),
			Inner:  nil,
		}
	}
}
