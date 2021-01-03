package godata

// toQuestionnaire formats a row into a Questionnaire
func toQuestionnaire(r []string) GoDataQuestionnaire {
	caseForms := []GoDataCaseForm{{Value: []string{
		"Form A0: Minimum data reporting form – for suspected and probable cases",
		"Form A2: Case follow-up form – for confirmed cases (Day 14-21)",
	}},
	}
	collectorName := []QuestionnaireAnswer{{Value: r[38]}}
	country := []QuestionnaireAnswer{{Value: r[39]}}
	showsSymptoms := []QuestionnaireAnswer{{Value: r[41]}}
	fever := []QuestionnaireAnswer{{Value: r[42]}}
	cough := []QuestionnaireAnswer{{Value: r[43]}}
	soreThroat := []QuestionnaireAnswer{{Value: r[44]}}
	shortnessBreath := []QuestionnaireAnswer{{Value: r[45]}}
	runnyNose := []QuestionnaireAnswer{{Value: r[46]}}
	chills := []QuestionnaireAnswer{{Value: r[47]}}
	headache := []QuestionnaireAnswer{{Value: r[48]}}
	musclePain := []QuestionnaireAnswer{{Value: r[49]}}
	vomit := []QuestionnaireAnswer{{Value: r[50]}}
	diarrhea := []QuestionnaireAnswer{{Value: r[51]}}
	anosmia := []QuestionnaireAnswer{{Value: r[52]}}
	aguesia := []QuestionnaireAnswer{{Value: r[53]}}
	respSample := []QuestionnaireAnswer{{Value: r[54]}}
	mechVent := []QuestionnaireAnswer{{Value: r[55]}}
	q := GoDataQuestionnaire{
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
		EyeFacialPain:         []QuestionnaireAnswer{},
		GeneralizedRash:       []QuestionnaireAnswer{},
		BlurredVision:         []QuestionnaireAnswer{},
		AbdominalPain:         []QuestionnaireAnswer{},
		CaseType:              "",
		PriorXdayExposureTravelledInternationally:     []QuestionnaireAnswer{{Value: r[56]}},
		PriorXdayExposureContactWithCase:              []QuestionnaireAnswer{{Value: r[59]}},
		PriorXDayexposureContactWithCaseDate:          []QuestionnaireAnswer{},
		PriorXdayExposureInternationalDateTravelFrom:  []QuestionnaireAnswer{},
		PriorXdayExposureInternationalDatetravelTo:    []QuestionnaireAnswer{{Value: r[57]}},
		PriorXdayexposureInternationaltravelcountries: []QuestionnaireAnswer{{Value: r[58]}},
		PriorXdayExposureInternationalTravelCities:    []QuestionnaireAnswer{},
		TypeOfTraveller:      []QuestionnaireAnswer{},
		PurposeOfTravel:      []QuestionnaireAnswer{},
		FlightNumber:         []QuestionnaireAnswer{},
		PcrTestInPast72Hours: []QuestionnaireAnswer{},
		DeathContrib:         []QuestionnaireAnswer{{Value: r[61]}},
		PostMortem:           []QuestionnaireAnswer{{Value: r[62]}},
		CauseOfDeath:         []QuestionnaireAnswer{{Value: r[63]}},
	}

	return q
}
