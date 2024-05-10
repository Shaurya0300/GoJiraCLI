package argumentsFunctions

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/andygrunwald/go-jira"
)

const (
	Username     string = "root"
	Password     string = "1234"
	DatabaseName string = "gotaskmanager"
	TableName    string = "jiratask"
)

type Details struct {
	ID            string
	JiraID        string
	Title         string
	Status        string
	DaysLeft      float64
	DaysRemaining time.Time
	Link          string
}
type Authentication struct {
	Credentials Credentials `json:"credentials"`
}
type Credentials struct {
	BearerTokenAuthentication BearerTokenAuthentication `json:"bearerToken"`
	BasicAuthentication       BasicAuthentication       `json:"basicAuth"`
}
type BearerTokenAuthentication struct {
	Token string `json:"value"`
}
type BasicAuthentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginWithBasicAuthentication() *jira.BasicAuthTransport {
	return &jira.BasicAuthTransport{ // For Basic Authentication
		Username: "",
		Password: "",
	}
}

func LoginWithBearerToken() *jira.BearerAuthTransport {
	return &jira.BearerAuthTransport{
		Token: "Mzc0ODkzNjIzOTE0Ogvf0sPV9kZPeeQj/29nmay1vNKo",
	}
}
func RetrieveClientDetails(filePath string) *jira.Client {

	file, err := os.Open(filePath)
	if err != nil {
		log.Panic("Failed to open the password file", err)
	}
	defer file.Close()
	output, err := io.ReadAll(file)
	if err != nil {
		log.Panic("Failed to read the password file", err)
	}

	var temp Authentication
	err = json.Unmarshal(output, &temp)
	if err != nil {
		log.Panic("Failed to unmarshal JSON data", err)
	}
	// Check if any password is provided in the crediantials.json file or not
	if temp.Credentials.BearerTokenAuthentication.Token == "" && (temp.Credentials.BasicAuthentication.Username == "" && temp.Credentials.BasicAuthentication.Password == "") {
		log.Panic("No Password is provided in crediantials.json")
	}
	// If Token key is provided it will take that and return the client
	if temp.Credentials.BearerTokenAuthentication.Token != "" {
		client, err := jira.NewClient(LoginThroughToken(temp).Client(), "https://jiradc2.ext.net.nokia.com")
		if err != nil {
			log.Println("Error while connecting to the jira.", err)
		}
		return client
	}
	// Else if username & password is being used then it will take that and return the client
	client, err := jira.NewClient(LoginThroughBasicAuth(temp).Client(), "https://jiradc2.ext.net.nokia.com")
	if err != nil {
		log.Panic("Error while connecting to the jira.", err)
	}
	return client
}
func LoginThroughToken(tokenValue Authentication) *jira.BearerAuthTransport {
	return &jira.BearerAuthTransport{Token: tokenValue.Credentials.BearerTokenAuthentication.Token}
}
func LoginThroughBasicAuth(tokenValue Authentication) *jira.BasicAuthTransport {
	return &jira.BasicAuthTransport{
		Username: tokenValue.Credentials.BasicAuthentication.Username,
		Password: tokenValue.Credentials.BasicAuthentication.Password,
	}
}
