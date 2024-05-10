package argumentsFunctions

import (
	"database/sql"
	"fmt"
	data "goTaskManager/argumentsFunctions"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type Cobratest struct {
	*cobra.Command
}

// ListTask retrieves jira data from a MySQL database table
func (cmd Cobratest) ListTask() {
	credentials := fmt.Sprintf("%s:%s@/%s", data.Username, data.Password, data.DatabaseName)
	db, _ := sql.Open("mysql", credentials)
	defer db.Close()

	query := fmt.Sprintf(`
		SELECT * FROM %s;
	`, data.TableName)
	resp, err := db.Query(query)
	if err != nil {
		log.Print("Error executing query:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	var ID, JiraID, Title, Status, DaysLeft, DaysRemaining, JiraLink string
	table.SetHeader([]string{"ID", "JiraID", "Title", "Status", "DaysLeft", "Days_Remaining", "JiraLink"})
	for resp.Next() {
		err := resp.Scan(&ID, &JiraID, &Title, &Status, &DaysLeft, &DaysRemaining, &JiraLink)
		if err != nil {
			log.Print("Error scanning row:", err)
			continue
		}
		temp := []string{ID, JiraID, Title, Status, DaysLeft, DaysRemaining, JiraLink}
		table.Append(temp)
	}
	table.SetRowLine(true) // Enable row line
	table.SetRowSeparator("-")
	table.Render()
}
