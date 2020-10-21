package values

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/sheets/v4"
)

type Request struct {
	SpreadsheetID string
	Range         string
	RequestBody   struct {
		Range          string
		MajorDimension string
		Values         [][]string
	}
}
type body struct {
	Data struct {
		Range  string     `json:"range"`
		Values [][]string `json:"values"`
	} `json:"data"`
	ValueInputOption string `json:"valueInputOption"`
}

// Update adds values to a Google Sheet
func Update(client *http.Client, spreadsheetID string, r string, values []interface{}) {

	/* ctx := context.Background() */

	// Authenticate
	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to create sheets Client %v", err)
	}

	var vr sheets.ValueRange
	/* values := []interface{}{"One", "Two"} */
	vr.Values = append(vr.Values, values)
	_, err = sheetsService.Spreadsheets.Values.Update(spreadsheetID, r, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to post data to sheet. %v", err)
	}
}

// https://stackoverflow.com/questions/46230624/google-sheets-api-golang-batchupdatevaluesrequest
func BatchUpdate(client *http.Client, spreadsheetID string, rangeData string, values [][]interface{}) {
	/* client := auth.Authorize() */
	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	/* spreadsheetId := "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk" */
	/* rangeData := "sheet2!I1:J3" */
	/* values := [][]interface{}{{"sample_A1", "sample_B1"}, {"sample_A2", "sample_B2"}, {"sample_A3", "sample_A3"}} */
	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  rangeData,
		Values: values,
	})
	_, err = sheetsService.Spreadsheets.Values.BatchUpdate(spreadsheetID, rb).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done.")
}
