package argumentsFunctions

import (
	"fmt"
	preRequisite "goTaskManager/argumentsFunctions"
	"log"
	"strconv"
	"time"

	excelize "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cobra"
)

type Cobratest struct {
	*cobra.Command
}

// Used to Generated Report in Excel Format
func (cmd Cobratest) GenerateReport() {

	data := preRequisite.GenerateCSVReport()

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"
	index := xlsx.NewSheet("Sheet1")

	var ID, JiraID, Title, Status, DaysLeft, DaysRemaining, JiraLink string
	var Data []preRequisite.Details

	// Set header row
	headers := []string{"ID", "JiraID", "Title", "Status", "DaysLeft", "Days_Remaining", "JiraLink"}
	for colIndex, header := range headers {
		// cell will be taking the index of each headers list and convert integer to Excel sheet column title. Example: convert 1 to column title A:
		cell := excelize.ToAlphaString(colIndex) + "1" // cell: A B C D E
		xlsx.SetCellValue(sheetName, cell, header)
	}

	for data.Next() {
		err := data.Scan(&ID, &JiraID, &Title, &Status, &DaysLeft, &DaysRemaining, &JiraLink)
		if err != nil {
			log.Print("Error scanning row:", err)
			continue
		}
		// Appending the values from the database to a slice of the dataAdd struct
		Data = append(Data, preRequisite.Details{
			ID:     ID,
			JiraID: JiraID,
			Title:  Title,
			Status: Status,
			DaysLeft: func() float64 {
				output, _ := strconv.ParseFloat(DaysLeft, 64)
				return output
			}(),
			DaysRemaining: func() time.Time {
				output, _ := time.Parse("2024-05-08", DaysRemaining)
				return output
			}(),
			Link: JiraLink,
		})
	}
	for rowIndex, rowValue := range Data {
		rowNumber := rowIndex + 2 // Excel rows start from 1, so increment by 2
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNumber), rowValue.ID)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", rowNumber), rowValue.JiraID)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", rowNumber), rowValue.Title)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", rowNumber), rowValue.Status)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", rowNumber), rowValue.DaysLeft)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", rowNumber), rowValue.DaysRemaining)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", rowNumber), rowValue.Link)
	}
	xlsx.SetActiveSheet(index)
	xlsx.SaveAs("Report.xlsx")
	log.Print("Report generated successfully!")
}
