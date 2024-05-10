package cobra

import (
	Addjira "goTaskManager/argumentsFunctions/Addjira"

	cobra "github.com/spf13/cobra"
)

var addjiraID = &cobra.Command{
	Use:   "addjiraID",
	Short: "Add a Jira ID to the database.",
	Long: `
Add a Jira ID to the database.
Syntax: addjiraID --jiraQuery ["Jira Query"] --filePathofPassword ["/etc/passwd"]"
Options:
  --jiraQuery: 				Specify the Jira query for adding the ID.
  --filePathofPassword: 	Specify the file location where password is stored.`,
	Run: func(cmd *cobra.Command, args []string) {
		callingAddTaskFunction := &Addjira.Cobratest{
			Command: cmd,
		}
		callingAddTaskFunction.Addjira()
	},
}

func init() {
	initializingAddTaskFlags()
	rootCmd.AddCommand(addjiraID)
}

func initializingAddTaskFlags() {
	addjiraID.Flags().StringP("jiraQuery", "q", "assignee = currentUser() AND resolution = Unresolved order by updated DESC", "Specify the Jira query for adding the ID.")
	addjiraID.Flags().StringP("filePathofPassword", "f", "", "Specify the file path for authentication.")
}
