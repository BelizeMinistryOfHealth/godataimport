package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"time"
)

type QuestionnaireAnswer struct {
	Value string `json:"value"`
}

type GoDataQuestionnaire struct {
	CaseForm                                      []QuestionnaireAnswer `json:"Case_WhichForm"`
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
	DeathContrib string `json:"FA2_outcome_deathnCoVcontribution"`
	PostMortem   string `json:"FA2_outcome_portmortemperformed"`
	CauseOfDeath string `json:"FA2_symptoms_causeofdeath"`
	PuiId        string `json:"pui_id"`
	InterviewKey string `json:"interviewKey"`
	CaseNo       int    `json:"case_no"`
	ID2          int    `json:"ID2"`
}

const AddressType = "LNG_REFERENCE_DATA_CATEGORY_ADDRESS_TYPE_USUAL_PLACE_OF_RESIDENCE"

type Address struct {
	TypeId       string `json:"typeId"`
	Country      string `json:"country"`
	City         string `json:"city"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	Date         string `json:"date"`
	PhoneNumber  string `json:"phoneNumber"`
	LocationId   string `json:"locationId"`
}

type PersonAge struct {
	Years  int `json:"years"`
	Months int `josn:"months"`
}

const CaseClassification = "LNG_REFERENCE_DATA_CATEGORY_CASE_CLASSIFICATION_SUSPECT"

type CaseOutcome string

const (
	Active    CaseOutcome = "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_ACTIVE"
	Alive     CaseOutcome = "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_ALIVE"
	Deceased  CaseOutcome = "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_DECEASED"
	Recovered CaseOutcome = "LNG_REFERENCE_DATA_CATEGORY_OUTCOME_RECOVERED"
)

type QuestionnaireAnswers struct {

}

type CovidTest struct {
	ID             string      `json:"id"`
	Bhis           int         `json:"bhis"`
	ReportingDate  time.Time   `json:"dateOfReporting"`
	CreatedAt      time.Time   `json:"createdAt"`
	CreatedBy      string      `json:"createdBy"`
	FirstName      string      `json:"firstName"`
	LastName       string      `json:"lastName"`
	Gender         string      `json:"gender"`
	Occupation     string      `json:"occupation"`
	Age            PersonAge   `json:"age"`
	Classification string      `json:"classification"`
	WasCase        int         `json:"wasCase"`
	DateBecameCase *time.Time  `json:"dateBecomeCase"`
	DateOfOnset    *time.Time  `json:"dateOfOnset"`
	RiskLevel      string      `json:"riskLevel"`
	RiskReason     string      `json:"riskReason"`
	Outcome        CaseOutcome `json:"outcome"`
	DateOfOutcome  *time.Time  `json:"dateOfOutCome"`
	Addresses      []Address   `json:"addresses"`
	Questionnaire GoDataQuestionnaire `json:"questionnaireAnswers"`

}

// Read files from a csv file generated from a postgres table
func Read(r *csv.Reader) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
