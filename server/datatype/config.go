package datatype

//Config  defines config type
type Config struct {
	Environment        string
	DataSourceName     string
	TestDataSourceName string
	AwsSesKey          string
	AwsSesSecret       string
	AwsEmailFrom       string
	SlackWebhookURL    string
	DBMigrationPath    string
	ServerPort         string
	BaseURL            string
}
