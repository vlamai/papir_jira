package config

type GRPC struct {
	ConnectionSting string
}

func NewGRPC() GRPC {
	return GRPC{
		ConnectionSting: getEnvVar("PAPIR_JIRA_GRPC_SERVER"),
	}
}
