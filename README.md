# GoJiraCLI
Developed a Go-based command-line interface (CLI) integrating with Jira's REST API, enabling seamless task management and data manipulation directly from the terminal.

## Overview
GoCLIForJira is a Go-based Command-Line Interface (CLI) tool designed to seamlessly integrate with Jira's REST API for efficient task management. This tool allows users to fetch and display tasks from Jira, apply custom filters for task organization, and export task data to Excel for detailed analysis and reporting.

## Database Configuration
To use GoCLIForJira, you need to configure a MySQL database with the following table structure:

CREATE TABLE jiratask (
    ID INT PRIMARY KEY,
    JiraID VARCHAR(255),
    Title VARCHAR(255),
    Status VARCHAR(255),
    DaysLeft FLOAT,
    DaysRemaining DATE,
    JiraLink VARCHAR(255)
);

Make sure to provide the database credentials in the argumentsFunctions/sqlDB.go file:
- Username: "root"
- Password: "1234"
- DatabaseName: "gotaskmanager"
- TableName: "jiratask"

## Jira Credentials
You also need to provide Jira credentials for authentication. You can use either basic authentication (username and password) or a token key.

## Installation and Usage
1. Clone the repository: git clone https://github.com/yourusername/GoCLIForJira.git
2. Navigate to the project directory: cd GoCLIForJira
3. Install dependencies: go mod tidy
4. Build the CLI tool: go build -o jira main.go

### Supported Commands
- addjiraID: Add a Jira ID to the database.
  Syntax: addjiraID --jiraQuery ["Jira Query"] --filePathofPassword ["/etc/passwd"]
  Options:
    - --jiraQuery: Specify the Jira query for adding the ID. (default "assignee = currentUser() AND resolution = Unresolved order by updated DESC")
    - --filePathofPassword: Specify the file location where the password is stored.

Usage Example:
./jira addjiraID --filePathofPassword=credentials.json

- listtask: Allows users to list tasks.
  Syntax: listtask
  Usage Example:
./jira listtask

- generatereport: Allows users to generate a report based on the data present in the database.
  Syntax: generatereport
  Usage Example:
./jira generatereport

Feel free to contribute to the project by opening issues or submitting pull requests. Happy task managing with GoCLIForJira!

#GoCLIForJira #JiraIntegration #CLI #TaskManagement #Productivity #DeveloperTools #GitHub #OpenSource
