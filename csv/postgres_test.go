package csv

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	csvFile, err := os.Open("COVID19_godata_1.csv")
	if err != nil {
		t.Fatalf("failed to open csv file: %+v", err)
	}
	r := csv.NewReader(csvFile)
	Read(r)
}
