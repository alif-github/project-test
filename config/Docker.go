package config

type DockerConfig struct {
	Configuration
	Server struct {
		Protocol    string `json:"protocol"`
		Host        string `json:"host"`
		Port        string `envconfig:"PROJECT_WEB_PORT"`
		Version     string `json:"version"`
		PrefixPath  string `json:"prefix_path"`
		Application string `json:"application"`
	}
	Postgresql struct {
		Driver            string `json:"driver"`
		Address           string `envconfig:"PROJECT_ADDRESS_PSQL"`
		DefaultSchema     string `envconfig:"PROJECT_SCHEMA_PSQL"`
		MaxOpenConnection int    `json:"max_open_connection"`
		MaxIdleConnection int    `json:"max_idle_connection"`
	} `json:"postgresql"`
	LogFile []string `json:"log_file"`
}

func (input DockerConfig) GetServer() Server {
	return Server{
		Protocol:    input.Server.Protocol,
		Host:        input.Server.Host,
		Port:        input.Server.Port,
		Version:     input.Server.Version,
		PrefixPath:  input.Server.PrefixPath,
		Application: input.Server.Application,
	}
}

func (input DockerConfig) GetPostgresql() Postgresql {
	return Postgresql{
		Driver:            input.Postgresql.Driver,
		Address:           input.Postgresql.Address,
		DefaultSchema:     input.Postgresql.DefaultSchema,
		MaxOpenConnection: input.Postgresql.MaxOpenConnection,
		MaxIdleConnection: input.Postgresql.MaxIdleConnection,
	}
}

func (input DockerConfig) GetLogFile() []string {
	return input.LogFile
}
