package argumentsFunctions

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

// AddTasktoDB add jira data to a MySQL database table
func AddTasktoDB(Data []Details) {
	credentials := fmt.Sprintf("%s:%s@/%s", Username, Password, DatabaseName)
	db, _ := sql.Open("mysql", credentials)

	defer db.Close()
	// Delete records that are not present in the input data
	deleteQuery := fmt.Sprintf("DELETE FROM %s;", TableName)
	_, err := db.Exec(deleteQuery)
	if err != nil {
		log.Print("Error while deleting the records ", err)
	}
	var rowsAdded int
	for _, value := range Data {
		query := fmt.Sprintf(`
	        INSERT INTO %s (
				ID, JiraID, Title, Status, DaysLeft, DaysRemaining, JiraLink
	        )
	        VALUES (?, ?, ?, ?, ?, ?, ?);
	    `, TableName)
		resp, err := db.Exec(query, value.ID, value.JiraID, value.Title, value.Status, value.DaysLeft, value.DaysRemaining, value.Link)
		if err != nil {
			fmt.Println(value.Title)
			log.Print("Error executing query: ", err)
			return
		}
		if count, _ := resp.RowsAffected(); count == 1 {
			rowsAdded++
		}
	}
	log.Print("Data inserted successfully!")
	log.Print("RowsAffected: ", rowsAdded)
}

// Return all the data from DB to GenerateReport Function
func GenerateCSVReport() *sql.Rows {
	credentials := fmt.Sprintf("%s:%s@/%s", Username, Password, DatabaseName)
	db, _ := sql.Open("mysql", credentials)
	defer db.Close()

	query := fmt.Sprintf(`
		SELECT * FROM %s;
	`, TableName)
	data, err := db.Query(query)
	if err != nil {
		log.Print("Error executing query:", err)
		return nil
	}
	return data
}
