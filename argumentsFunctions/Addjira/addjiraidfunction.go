package argumentsFunctions

import (
	"fmt"
	data "goTaskManager/argumentsFunctions"
	_ "io"
	"log"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
)

type Cobratest struct {
	*cobra.Command
}

func (cmd Cobratest) Addjira() {
	filePath, err := cmd.Flags().GetString("filePathofPassword")
	if err != nil {
		log.Print("Error while retrieving the file path.", err)
	}
	//Used to return the client after connecting to jira based on the password its provided in the credentials.json file
	client := data.RetrieveClientDetails(filePath)
	query, _ := cmd.Flags().GetString("jiraQuery")
	resp, _, err := client.Issue.Search(query, nil)
	if err != nil {
		fmt.Println(err)
	}
	var Data []data.Details
	for _, value := range resp {
		temp := data.Details{
			ID:       value.ID,
			JiraID:   value.Key,
			Title:    value.Fields.Summary,
			Status:   value.Fields.Status.Name,
			DaysLeft: returnTime(client, value.Key) / 6, // 1 Day = 6 hrs
			DaysRemaining: func() time.Time {
				timeinDays := returnTime(client, value.Key)
				return time.Now().Add(time.Hour * 24 * time.Duration(timeinDays) / 6)
			}(),
			Link: fmt.Sprintf("https://jiradc2.ext.net.nokia.com/browse/%s", value.Key),
		}
		Data = append(Data, temp)
	}
	data.AddTasktoDB(Data)
}
func returnTime(client *jira.Client, ID string) float64 {
	resp, _, err := client.Issue.Get(ID, nil)
	if err != nil {
		log.Printf("Error while retrieving details of JIRAID: %s", ID)
	}
	return float64(resp.Fields.TimeTracking.RemainingEstimateSeconds) / 3600
}
