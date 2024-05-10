package cobra

import (
	listTask "goTaskManager/argumentsFunctions/ListTask"

	cobra "github.com/spf13/cobra"
)

var listtask = &cobra.Command{
	Use:   "listtask",
	Short: "Allows users to list tasks",
	Long: `
Allows users to list all the tasks.
Syntax: listtask`,
	Run: func(cmd *cobra.Command, args []string) {
		callingListTaskFunction := &listTask.Cobratest{
			Command: cmd,
		}
		callingListTaskFunction.ListTask()
	},
}

func init() {
	rootCmd.AddCommand(listtask)
}
