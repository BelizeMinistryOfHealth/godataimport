package godata

import "fmt"

type PregnancyStatus string

const (
	FirstTrimester         PregnancyStatus = "LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_YES_FIRST_TRIMESTER"
	SecondTrimester        PregnancyStatus = "LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_YES_SECOND_TRIMESTER"
	ThirdTrimester         PregnancyStatus = "LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_YES_THIRD_TRIMESTER"
	UnknownTrimester       PregnancyStatus = "LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_YES_TRIMESTER_UNKNOWN"
	NotPregnant            PregnancyStatus = "LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_NOT_PREGNANT"
	PregnancyNotApplicable PregnancyStatus = "LNG_REFERENCE_DATA_CATEGORY_PREGNANCY_STATUS_NONE"
)

type PregnancyStatusErr struct {
	Reason string
	Inner  error
}

func (e PregnancyStatusErr) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("pregnancy status error: %s: %v", e.Reason, e.Inner)
	}
	return fmt.Sprintf("pregnancy error: %s", e.Reason)
}

func (e PregnancyStatusErr) Unwrap() error {
	return e.Inner
}

func ToPregnancy(s string) (PregnancyStatus, error) {
	switch s {
	case string(FirstTrimester):
		return FirstTrimester, nil
	case string(SecondTrimester):
		return SecondTrimester, nil
	case string(ThirdTrimester):
		return ThirdTrimester, nil
	case string(UnknownTrimester):
		return UnknownTrimester, nil
	case string(NotPregnant):
		return NotPregnant, nil
	case string(PregnancyNotApplicable):
		return PregnancyNotApplicable, nil
	case "":
		return NotPregnant, nil
	default:
		return "", PregnancyStatusErr{
			Reason: fmt.Sprintf("invalid pregnancy status: %s", s),
			Inner:  nil,
		}
	}
}
