package godata

// toQuestionnaire formats a row into a Questionnaire
func toQuestionnaire(r []string) Questionnaire {
	caseForms := []CaseForm{{Value: []string{
		"Form A0: Minimum data reporting form – for suspected and probable cases",
		"Form A2: Case follow-up form – for confirmed cases (Day 14-21)",
	}},
	}
	collectorName := []QuestionnaireAnswer{{Value: r[38]}}
	country := []QuestionnaireAnswer{{Value: r[37]}}
	showsSymptoms := []QuestionnaireAnswer{{Value: r[39]}}
	fever := []QuestionnaireAnswer{{Value: r[40]}}
	cough := []QuestionnaireAnswer{{Value: r[41]}}
	soreThroat := []QuestionnaireAnswer{{Value: r[42]}}
	shortnessBreath := []QuestionnaireAnswer{{Value: r[43]}}
	runnyNose := []QuestionnaireAnswer{{Value: r[44]}}
	chills := []QuestionnaireAnswer{{Value: r[45]}}
	headache := []QuestionnaireAnswer{{Value: r[46]}}
	musclePain := []QuestionnaireAnswer{{Value: r[47]}}
	vomit := []QuestionnaireAnswer{{Value: r[48]}}
	diarrhea := []QuestionnaireAnswer{{Value: r[49]}}
	anosmia := []QuestionnaireAnswer{{Value: r[50]}}
	aguesia := []QuestionnaireAnswer{{Value: r[51]}}
	respSample := []QuestionnaireAnswer{{Value: r[52]}}
	mechVent := []QuestionnaireAnswer{{Value: r[53]}}
	ssn := []QuestionnaireAnswer{{Value: r[36]}}
	q := Questionnaire{
		CaseForm:              caseForms,
		DataCollectorName:     collectorName,
		CountryResidence:      country,
		ShowsSymptoms:         showsSymptoms,
		Fever:                 fever,
		SoreThroat:            soreThroat,
		RunnyNose:             runnyNose,
		Cough:                 cough,
		Vomiting:              vomit,
		Nausea:                []QuestionnaireAnswer{},
		Diarrhea:              diarrhea,
		ShortnessOfBreath:     shortnessBreath,
		DifficultyBreathing:   []QuestionnaireAnswer{},
		SymptomsChills:        chills,
		Headache:              headache,
		Malaise:               []QuestionnaireAnswer{},
		Anosmia:               anosmia,
		Aguesia:               aguesia,
		Bleeding:              []QuestionnaireAnswer{},
		JointMusclePain:       musclePain,
		RespSampleCollected:   respSample,
		MechanicalVentilation: mechVent,
		Ssn:                   ssn,
		EyeFacialPain:         []QuestionnaireAnswer{},
		GeneralizedRash:       []QuestionnaireAnswer{},
		BlurredVision:         []QuestionnaireAnswer{},
		AbdominalPain:         []QuestionnaireAnswer{},
		CaseType:              "",
		PriorXdayExposureTravelledInternationally:     []QuestionnaireAnswer{{Value: r[54]}},
		PriorXdayExposureContactWithCase:              []QuestionnaireAnswer{{Value: r[57]}},
		PriorXDayexposureContactWithCaseDate:          []QuestionnaireAnswer{},
		PriorXdayExposureInternationalDateTravelFrom:  []QuestionnaireAnswer{},
		PriorXdayExposureInternationalDatetravelTo:    []QuestionnaireAnswer{{Value: r[55]}},
		PriorXdayexposureInternationaltravelcountries: []QuestionnaireAnswer{{Value: r[56]}},
		PriorXdayExposureInternationalTravelCities:    []QuestionnaireAnswer{},
		TypeOfTraveller:                []QuestionnaireAnswer{},
		PurposeOfTravel:                []QuestionnaireAnswer{},
		FlightNumber:                   []QuestionnaireAnswer{},
		PcrTestInPast72Hours:           []QuestionnaireAnswer{},
		DeathContrib:                   []QuestionnaireAnswer{{Value: r[59]}},
		PostMortem:                     []QuestionnaireAnswer{{Value: r[60]}},
		CauseOfDeath:                   []QuestionnaireAnswer{{Value: r[61]}},
		RespiratorySampleDateCollected: []QuestionnaireAnswer{{Value: r[105]}},
	}

	return q
}
