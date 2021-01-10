package godata

import "testing"

var tests = []struct {
	status        string
	want          PregnancyStatus
	expectedError bool
}{
	{"LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_YES_FIRST_TRIMESTER",
		FirstTrimester,
		false,
	},
	{
		"LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_YES_SECOND_TRIMESTER",
		SecondTrimester,
		false,
	},
	{
		"LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_YES_THIRD_TRIMESTER",
		ThirdTrimester,
		false,
	},
	{
		"LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_YES_TRIMESTER_UNKNOWN",
		UnknownTrimester,
		false,
	},
	{
		"LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_NOT_PREGNANT",
		NotPregnant,
		false,
	},
	{
		"LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_NONE",
		PregnancyNotApplicable,
		false,
	},
	{
		"Wrong value",
		"",
		true,
	},
	{
		"",
		NotPregnant,
		false,
	},
}

func TestToPregnancy(t *testing.T) {
	for _, e := range tests {
		got, err := ToPregnancy(e.status)
		if err != nil && !e.expectedError {
			t.Errorf("ToPregnancy(%s) = unexpected error: %v", e.status, err)
		}
		if err == nil && e.expectedError {
			t.Errorf("ToPregnancy(%s) = %v, expected error", e.status, err)
		}
		if got != e.want {
			t.Errorf("ToPregnancy(%s) = %v, expected %v", e.status, got, e.want)
		}
	}
	//status := "LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_YES_FIRST_TRIMESTER"
	//got, err := ToPregnancy(status)
	//if err != nil {
	//	t.Errorf("failed pregnancy parsing, want: %s, got: %s", status, string(got))
	//}
	//t.Logf("got: %s", got)
}
