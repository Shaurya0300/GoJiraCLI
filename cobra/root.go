package cobra

import (
	"fmt"

	cobra "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "goJiraCLI",
	Long:  "Used for integrating golang based cli with JIRA.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
