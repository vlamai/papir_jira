package config

type Jira struct {
	BaseUrl string
	Name    string
	Token   string
}

func NewJira() Jira {
	return Jira{
		BaseUrl: getEnvVar("PAPIR_JIRA_BASE_URL"),
		Name:    getEnvVar("PAPIR_JIRA_LOGIN"),
		Token:   getEnvVar("PAPIR_JIRA_TOKEN"),
	}
}
