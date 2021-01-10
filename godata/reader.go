package godata

import (
	"encoding/csv"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"strconv"
	"time"
)

//const (
//	layoutISO = "2006-01-02 00:00:00-07"
//)

const (
	layoutISO = "2006-01-02"
)

type QuestionnaireAnswer struct {
	Value string `json:"value"`
}

type CaseForm struct {
	Value []string `json:"value"`
}

// GoDataQuestionnaire represents the GoData questionnaire. GoData stores these as a flat list.
// The CaseForm identifies the forms, and GoData uses this to extract the fields for each form from the
// flat list of questions.
type Questionnaire struct {
	CaseForm                                      []CaseForm            `json:"Case_WhichForm"`
	DataCollectorName                             []QuestionnaireAnswer `json:"FA0_datacollector_name"`
	CountryResidence                              []QuestionnaireAnswer `json:"FA0_case_countryresidence"`
	ShowsSymptoms                                 []QuestionnaireAnswer `json:"FA0_symptoms_caseshowssymptoms"`
	Fever                                         []QuestionnaireAnswer `json:"FA0_symptom_fever"`
	SoreThroat                                    []QuestionnaireAnswer `json:"FA0_symptom_sorethroat"`
	RunnyNose                                     []QuestionnaireAnswer `json:"FA0_symptom_runnynose"`
	Cough                                         []QuestionnaireAnswer `json:"FA0_symptom_cough"`
	Vomiting                                      []QuestionnaireAnswer `json:"FA0_symptom_vomiting"`
	Nausea                                        []QuestionnaireAnswer `json:"FA0_symptom_nausea"`
	Diarrhea                                      []QuestionnaireAnswer `json:"FA0_symptom_diarrhea"`
	ShortnessOfBreath                             []QuestionnaireAnswer `json:"FA0_symptom_shortnessofbreath"`
	DifficultyBreathing                           []QuestionnaireAnswer `json:"FA0_symptom_difficulty_breathing"`
	SymptomsChills                                []QuestionnaireAnswer `json:"FA0_symptom_chills"`
	Headache                                      []QuestionnaireAnswer `json:"FA0_symptom_headache"`
	Malaise                                       []QuestionnaireAnswer `json:"FA0_symptom_malaise"`
	Anosmia                                       []QuestionnaireAnswer `json:"FA0_symptom_anosmia"`
	Aguesia                                       []QuestionnaireAnswer `json:"FA0_symptom_aguesia"`
	Bleeding                                      []QuestionnaireAnswer `json:"FA0_symptom_bleeding"`
	JointMusclePain                               []QuestionnaireAnswer `json:"FA0_symptom_joint_muscle_pain"`
	EyeFacialPain                                 []QuestionnaireAnswer `json:"FA0_symptom_eye_facial_pain"`
	GeneralizedRash                               []QuestionnaireAnswer `json:"FA0_symptom_generalized_rash"`
	BlurredVision                                 []QuestionnaireAnswer `json:"FA0_symptom_blurred_vision"`
	AbdominalPain                                 []QuestionnaireAnswer `json:"FA0_symptom_abdominal_pain"`
	CaseType                                      string                `json:"case_type"`
	Ssn                                           []QuestionnaireAnswer `json:"FA0_caseidentifier_socialnumber"`
	PriorXdayExposureTravelledInternationally     []QuestionnaireAnswer `json:"FA0_priorXdayexposure_travelledinternationally"`
	PriorXdayExposureContactWithCase              []QuestionnaireAnswer `json:"FA0_priorXdayexposure_contactwithcase"`
	PriorXDayexposureContactWithCaseDate          []QuestionnaireAnswer `json:"FA0_priorXdayexposure_contactwithcasedate"`
	PriorXdayExposureInternationalDateTravelFrom  []QuestionnaireAnswer `json:"FA0_priorXdayexposure_internationaldatetravelfrom"`
	PriorXdayExposureInternationalDatetravelTo    []QuestionnaireAnswer `json:"FA0_priorXdayexposure_internationaldatetravelto"`
	PriorXdayexposureInternationaltravelcountries []QuestionnaireAnswer `json:"FA0_priorXdayexposure_internationaltravelcountries"`
	PriorXdayExposureInternationalTravelCities    []QuestionnaireAnswer `json:"FA0_priorXdayexposure_internationaltravelcities"`
	TypeOfTraveller                               []QuestionnaireAnswer `json:"FA0_priorXdayexposure_typeoftraveler"`
	PurposeOfTravel                               []QuestionnaireAnswer `json:"FA0_priorXdayexposure_purposeoftravel"`
	FlightNumber                                  []QuestionnaireAnswer `json:"FA0_priorXdayexposure_flightnumber"`
	PcrTestInPast72Hours                          []QuestionnaireAnswer `json:"FA0_priorXdayexposure_tookpcrtest_past72hours"`
	DeathContrib                                  []QuestionnaireAnswer `json:"FA2_outcome_deathnCoVcontribution"`
	PostMortem                                    []QuestionnaireAnswer `json:"FA2_outcome_postmortemperformed"`
	CauseOfDeath                                  []QuestionnaireAnswer `json:"FA2_symptoms_causeofdeath"`
	RespSampleCollected                           []QuestionnaireAnswer `json:"FA0_respiratorysample_collectedYN"`
	MechanicalVentilation                         []QuestionnaireAnswer `json:"FA0_clinicalcomplications_mechanicalventilation"`
}

const AddressType = "LNG_REFERENCE_DATA_CATEGORY_ADDRESS_TYPE_USUAL_PLACE_OF_RESIDENCE"
const OtherAddressType = "LNG_REFERENCE_DATA_CATEGORY_ADDRESS_TYPE_OTHER"
const ReportingFacility = "LNG_REFERENCE_DATA_CATEGORY_ADDRESS_TYPE_ACCOMMODATION_NAME"

type GeoLocation struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type Address struct {
	TypeId       string       `json:"typeId"`
	Country      string       `json:"country"`
	City         string       `json:"city"`
	AddressLine1 string       `json:"addressLine1"`
	AddressLine2 string       `json:"addressLine2"`
	Date         string       `json:"date"`
	PhoneNumber  string       `json:"phoneNumber"`
	LocationId   string       `json:"locationId"`
	GeoLocation  *GeoLocation `json:"geoLocation,omitempty"`
}

type PersonAge struct {
	Years  int `json:"years"`
	Months int `josn:"months"`
}

const CaseSuspectClassification = "LNG_REFERENCE_DATA_CATEGORY_CASE_CLASSIFICATION_SUSPECT"
const CaseConfirmedClassification = "LNG_REFERENCE_DATA_CATEGORY_CASE_CLASSIFICATION_CONFIRMED"

// toClassification converts a string to the corresponding classification string that GoData expects.
// Returns an empty string if the input is invalid.
func toClassification(s string) (string, error) {
	switch s {
	case "Confirmed":
		return CaseConfirmedClassification, nil
	case "Suspect":
		return CaseSuspectClassification, nil
	default:
		return "", fmt.Errorf("wrong case classification")
	}
}

const LowRisk = "LNG_REFERENCE_DATA_CATEGORY_RISK_LEVEL_1_LOW"
const MediumRisk = "LNG_REFERENCE_DATA_CATEGORY_RISK_LEVEL_2_MEDIUM"
const HighRisk = "LNG_REFERENCE_DATA_CATEGORY_RISK_LEVEL_3_HIGH"

func toRiskLevel(s string) string {
	switch s {
	case "3 - High":
		return HighRisk
	case "2 - Medium":
		return MediumRisk
	case "3 - Low":
		return LowRisk
	default:
		return ""
	}
}

func toOutcome(s string) (string, error) {
	switch s {
	case "Active":
		return "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_ACTIVE", nil
	case "Alive":
		return "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_ALIVE", nil
	case "Deceased":
		return "", nil
	case "Recovered":
		return "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_RECOVERED", nil
	default:
		return "", fmt.Errorf("invalid classification")

	}
}

type CovidTest struct {
	VisualID         string            `json:"visualId"`
	Bhis             int               `json:"bhis"`
	ReportingDate    time.Time         `json:"dateOfReporting"`
	CreatedAt        time.Time         `json:"createdAt"`
	CreatedBy        string            `json:"createdBy"`
	FirstName        string            `json:"firstName"`
	LastName         string            `json:"lastName"`
	Gender           string            `json:"gender"`
	Occupation       string            `json:"occupation"`
	Age              PersonAge         `json:"age"`
	Dob              time.Time         `json:"dob"`
	Classification   string            `json:"classification"`
	DateBecameCase   *time.Time        `json:"dateBecomeCase"`
	DateOfOnset      *time.Time        `json:"dateOfOnset"`
	RiskLevel        string            `json:"riskLevel"`
	RiskReason       string            `json:"riskReason"`
	Outcome          string            `json:"outcome"`
	PregnancyStatus  string            `json:"pregnancyStatus"`
	DateOfOutcome    *time.Time        `json:"dateOfOutCome"`
	Addresses        []Address         `json:"addresses"`
	Questionnaire    Questionnaire     `json:"questionnaireAnswers,omitempty"`
	Hospitalizations []Hospitalization `json:"dateRanges,omitempty"`
}

// min finds the smallest number
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// SplitTests batches the CovidTests. This allows us to import the data in batches.
// GoData will restart if it gets too much traffic. Batching the upload prevents us
// from DOSing the GoData application.
func SplitTests(t []CovidTest, batchSize int) [][]CovidTest {
	var batch [][]CovidTest
	for i := 0; i < len(t); i += batchSize {
		bs := t[i:min(i+batchSize, len(t))]
		batch = append(batch, bs)
	}
	return batch
}

func toCurrentAddress(r []string, locs []AddressLocation) (*Address, error) {
	loc := FindLocation(r[24], locs)
	if loc == nil {
		return nil, fmt.Errorf("invalid address %s", r[24])
	}
	lat, _ := strconv.ParseFloat(r[28], 32)
	lng, _ := strconv.ParseFloat(r[29], 32)
	addr := Address{
		TypeId:       AddressType,
		Country:      "Belize",
		City:         loc.Name,
		AddressLine1: r[23],
		AddressLine2: "",
		Date:         "",
		PhoneNumber:  r[11],
		LocationId:   loc.Id,
		GeoLocation: &GeoLocation{
			Lat: float32(lat),
			Lng: float32(lng),
		},
	}
	return &addr, nil
}

func toReportingFacility(r []string, locs []AddressLocation) (*Address, error) {
	loc := FindLocation(r[65], locs)
	if loc == nil {
		return nil, fmt.Errorf("invalid address %s", r[24])
	}

	addr := Address{
		TypeId:       ReportingFacility,
		Country:      "Belize",
		City:         loc.Name,
		AddressLine1: r[63],
		AddressLine2: r[64],
		Date:         "",
		LocationId:   loc.Id,
	}
	return &addr, nil
}

func toOtherAddress(r []string, locs []AddressLocation) (*Address, error) {
	loc := FindLocation(r[33], locs)
	if loc == nil {
		return nil, fmt.Errorf("invalid address %s", r[33])
	}

	addr := Address{
		TypeId:       OtherAddressType,
		Country:      "Belize",
		City:         loc.Name,
		AddressLine1: r[32],
		AddressLine2: "",
		Date:         "",
		PhoneNumber:  "",
		LocationId:   loc.Id,
	}
	return &addr, nil
}

// Read files from a csv file generated from a postgres table
func Read(r *csv.Reader, locs []AddressLocation) ([]CovidTest, error) {
	var tests []CovidTest
	row := 0
	for {
		record, err := r.Read()
		// ignore the header
		if row == 0 {
			row = row + 1
			record, err = r.Read()
		}
		if err == io.EOF {
			break
		}
		row = row + 1
		if err != nil {
			return nil, fmt.Errorf("error reading the csv file: %w", err)
		}
		//fmt.Println(record)
		bhisNumber, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing the bhis number for: %s (%w)", record[1], err)
		}
		repDate, err := time.Parse(layoutISO, record[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing the reporting date(%s) for id %s: %w", record[2], record[0], err)
		}
		createDate, err := time.Parse(layoutISO, record[3])
		if err != nil {
			return nil, fmt.Errorf("error parsing the create date(%s) for id %s: %w", record[2], record[0], err)
		}
		ageYears, err := strconv.Atoi(record[12])
		if err != nil {
			return nil, fmt.Errorf("error parsing the age in years(%s) for id %s: %w", record[12], record[0], err)
		}
		ageMonths, err := strconv.Atoi(record[13])
		if err != nil {
			return nil, fmt.Errorf("error parsing the age in months(%s) for id %s: %w", record[13], record[0], err)
		}
		classification, err := toClassification(record[14])
		if err != nil {
			return nil, fmt.Errorf("error parsing the case clssification (%s) for id %s: %w", record[14], record[0], err)
		}
		dateBecameCase, _ := time.Parse(layoutISO, record[16])
		var dateOfOnset *time.Time
		dOfOnset, err := time.Parse(layoutISO, record[17])
		if err != nil {
			dateOfOnset = nil
		} else {
			dateOfOnset = &dOfOnset
		}
		outcome, err := toOutcome(record[20])
		if err != nil {
			return nil, fmt.Errorf("could not parse outcome(%s) for id %s: %w", record[20], record[0], err)
		}
		var dateOfOutcome *time.Time
		if record[21] == "" {
			dateOfOutcome = nil
		} else {
			dOutcome, err := time.Parse(layoutISO, record[21])
			if err != nil {
				return nil, fmt.Errorf("wrong date format for date of outcome (%s) for id %s: %w", record[21], record[0], err)
			}
			dateOfOutcome = &dOutcome
		}

		currentAddress, currentAddrErr := toCurrentAddress(record, locs)
		//if err != nil {
		//	return nil, fmt.Errorf("could not parse the current address for id %s: %w", record[0], err)
		//}
		otherAddress, otherAddrErr := toOtherAddress(record, locs)
		//if err != nil {
		//	return nil, fmt.Errorf("could not parse the secondary address for id %s: %w", record[0], err)
		//}
		var addresses []Address
		if otherAddrErr == nil && currentAddrErr == nil {
			addresses = []Address{*currentAddress, *otherAddress}
		}
		reportingFacilityAddress, err := toReportingFacility(record, locs)
		if err == nil && reportingFacilityAddress != nil {
			addresses = append(addresses, *reportingFacilityAddress)
		}
		var hospitalizations []Hospitalization
		h, err := toHospitalization(record, 68, locs)
		var emptyHospErr1 *EmptyHospitalizationErr
		if err != nil && errors.As(err, &emptyHospErr1) && emptyHospErr1 == nil {
			return nil, fmt.Errorf("error parsing hospitalization(1): %w", err)
		}
		if len(h.TypeId) != 0 {
			hospitalizations = append(hospitalizations, h)
		}
		h, err = toHospitalization(record, 77, locs)
		var emptyHospErr2 *EmptyHospitalizationErr
		if err != nil && errors.As(err, &emptyHospErr2) && emptyHospErr2 == nil {
			return nil, fmt.Errorf("error parsing hospitalization(2): %w", err)
		}
		if len(h.TypeId) != 0 {
			hospitalizations = append(hospitalizations, h)
		}
		h, err = toHospitalization(record, 86, locs)
		var emptyHospErr3 *EmptyHospitalizationErr
		if err != nil && errors.As(err, &emptyHospErr3) && emptyHospErr3 == nil {
			return nil, fmt.Errorf("error parsing hospitalization(3): %w", err)
		}
		if len(h.TypeId) != 0 {
			hospitalizations = append(hospitalizations, h)
		}

		var emptyHospErr4 *EmptyHospitalizationErr
		h, err = toHospitalization(record, 95, locs)
		if err != nil && errors.As(err, &emptyHospErr4) && emptyHospErr4 == nil {
			return nil, fmt.Errorf("error parsing hospitalization(4): %w", err)
		}
		if len(h.TypeId) != 0 {
			hospitalizations = append(hospitalizations, h)
		}

		dob, err := time.Parse(layoutISO, record[104])
		if err != nil {
			return nil, fmt.Errorf("record (bhis:%s | row:%d | value:%s) does not have a date of birth", record[0], row, record[104])
		}

		pregnancyStatus, pregnancyErr := ToPregnancy(record[38])
		if record[9] == "Male" {
			pregnancyStatus = PregnancyNotApplicable
		}

		test := CovidTest{
			VisualID:      record[0],
			Bhis:          bhisNumber,
			ReportingDate: repDate,
			CreatedAt:     createDate,
			FirstName:     record[7],
			LastName:      record[8],
			Gender:        record[9],
			Occupation:    record[10],
			Age: PersonAge{
				Years:  ageYears,
				Months: ageMonths,
			},
			Dob:              dob,
			Classification:   classification,
			DateBecameCase:   &dateBecameCase,
			DateOfOnset:      dateOfOnset,
			RiskLevel:        toRiskLevel(record[18]),
			RiskReason:       record[19],
			Outcome:          outcome,
			DateOfOutcome:    dateOfOutcome,
			PregnancyStatus:  string(pregnancyStatus),
			Addresses:        addresses,
			Questionnaire:    toQuestionnaire(record),
			Hospitalizations: hospitalizations,
		}
		if currentAddrErr != nil || otherAddrErr != nil || pregnancyErr != nil {
			log.WithFields(log.Fields{
				"error":          currentAddrErr,
				"otherAddr":      otherAddrErr,
				"pregnancyError": pregnancyErr,
				"row":            row,
			}).Error("no address")
		}
		if currentAddrErr == nil && otherAddrErr == nil && pregnancyErr == nil {
			tests = append(tests, test)
		}

	}
	log.WithFields(log.Fields{
		"numberRows":  row,
		"numberTests": len(tests),
	}).Info("Parsed")
	return tests, nil
}
