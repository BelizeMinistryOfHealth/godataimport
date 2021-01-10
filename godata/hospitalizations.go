package godata

import (
	"fmt"
	"time"
)

type HospitalizationType string

const (
	AeHospitalizationId      HospitalizationType = "LNG_REFERENCE_DATA_CATEGORY_PERSON_DATE_TYPE_A_E_VISIT"
	HospitalizationId        HospitalizationType = "LNG_REFERENCE_DATA_CATEGORY_PERSON_DATE_TYPE_HOSPITALIZATION"
	OtherMedicalConditionsId HospitalizationType = "LNG_REFERENCE_DATA_CATEGORY_PERSON_DATE_TYPE_HOSPITALIZATION_FOR_OTHER_MEDICAL_CONDITIONS"
	IcuId                    HospitalizationType = "LNG_REFERENCE_DATA_CATEGORY_PERSON_DATE_TYPE_ICU_ADMISSION"
	IsolationId              HospitalizationType = "LNG_REFERENCE_DATA_CATEGORY_PERSON_DATE_TYPE_ISOLATION"
	OtherHospitalizationId   HospitalizationType = "LNG_REFERENCE_DATA_CATEGORY_PERSON_DATE_TYPE_OTHER"
	PrimaryHealthCareId      HospitalizationType = "LNG_REFERENCE_DATA_CATEGORY_PERSON_DATE_TYPE_PRIMARY_HEALTH_CARE_PHC_GP_ETC_VISIT"
)

type Hospitalization struct {
	TypeId     string     `json:"typeId"`
	StartDate  time.Time  `json:"startDate"`
	EndDate    *time.Time `json:"endDate,omitempty"`
	CenterName string     `json:"centerName"`
	LocationId string     `json:"locationId"`
	Comments   string     `json:"comments"`
}

func toHospitalization(r []string, col int, locs []AddressLocation) (Hospitalization, error) {
	if r[col] == "" {
		return Hospitalization{}, EmptyHospitalizationErr{
			Reason: "0 Hospitalization Found",
			Inner:  nil,
		}
	}
	startDate, err := time.Parse(layoutISO, r[col+1])
	if err != nil {
		return Hospitalization{},
			HospitalizationErr{
				Reason: fmt.Sprintf("wrong date format (%s) for hospitalization starting at column %d", r[col+1], col+1),
				Inner:  err,
			}
	}
	endDate, err := time.Parse(layoutISO, r[col+2])
	if err != nil {
		return Hospitalization{},
			HospitalizationErr{
				Reason: fmt.Sprintf("wrong date format (%s) for hospitalization end date starting at column %d", r[col+2], col+2),
				Inner:  err,
			}
	}
	hospitalization := Hospitalization{
		TypeId:     r[col],
		StartDate:  startDate,
		EndDate:    &endDate,
		CenterName: r[col+3],
		Comments:   r[col+8],
	}

	location := FindLocation(r[col+6], locs)
	if location != nil {
		hospitalization.LocationId = location.Id
	}

	return hospitalization, nil

}

type HospitalizationErr struct {
	Reason string
	Inner  error
}

func (e HospitalizationErr) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("hospitalization error: %s: %v", e.Reason, e.Inner)
	}
	return fmt.Sprintf("hospitalization error: %s", e.Reason)
}

func (e HospitalizationErr) Unwrap() error {
	return e.Inner
}

type EmptyHospitalizationErr struct {
	Reason string
	Inner  error
}

func (e EmptyHospitalizationErr) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("empty hospitalization error: %s: %v", e.Reason, e.Inner)
	}
	return fmt.Sprintf("empty hospitalization error: %s", e.Reason)
}

func (e EmptyHospitalizationErr) Unwrap() error {
	return e.Inner
}
