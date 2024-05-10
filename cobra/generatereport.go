package cobra

import (
	generateReport "goTaskManager/argumentsFunctions/GenerateReport"

	cobra "github.com/spf13/cobra"
)

var generatereport = &cobra.Command{
	Use:   "generatereport",
	Short: "Allows users to generate report",
	Long: `
Allows users to generate report based on the data present in database.
Syntax: generateReport
Options:
  --generateReport: Generates a report.`,
	Run: func(cmd *cobra.Command, args []string) {
		callingGenerateReportFunction := &generateReport.Cobratest{
			Command: cmd,
		}
		callingGenerateReportFunction.GenerateReport()
	},
}

func init() {
	rootCmd.AddCommand(generatereport)
}
