package csv

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
		Nausea:                nil,
		Diarrhea:              diarrhea,
		ShortnessOfBreath:     shortnessBreath,
		DifficultyBreathing:   nil,
		SymptomsChills:        chills,
		Headache:              headache,
		Malaise:               nil,
		Anosmia:               anosmia,
		Aguesia:               aguesia,
		Bleeding:              nil,
		JointMusclePain:       musclePain,
		RespSampleCollected:   respSample,
		MechanicalVentilation: mechVent,
		EyeFacialPain:         nil,
		GeneralizedRash:       nil,
		BlurredVision:         nil,
		AbdominalPain:         nil,
		CaseType:              "",
		PriorXdayExposureTravelledInternationally:     []QuestionnaireAnswer{{Value: r[56]}},
		PriorXdayExposureContactWithCase:              []QuestionnaireAnswer{{Value: r[59]}},
		PriorXDayexposureContactWithCaseDate:          nil,
		PriorXdayExposureInternationalDateTravelFrom:  nil,
		PriorXdayExposureInternationalDatetravelTo:    []QuestionnaireAnswer{{Value: r[57]}},
		PriorXdayexposureInternationaltravelcountries: []QuestionnaireAnswer{{Value: r[58]}},
		PriorXdayExposureInternationalTravelCities:    nil,
		TypeOfTraveller:      nil,
		PurposeOfTravel:      nil,
		FlightNumber:         nil,
		PcrTestInPast72Hours: nil,
		DeathContrib:         r[61],
		PostMortem:           r[62],
		CauseOfDeath:         r[63],
		PuiId:                r[64],
		InterviewKey:         r[65],
		CaseNo:               r[66],
		ID2:                  r[67],
	}

	return q
}
