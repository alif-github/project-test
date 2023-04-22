package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/tkanos/gonfig"
	"os"
)

var ApplicationConfiguration Configuration

type Configuration interface {
	GetPostgresql() Postgresql
	GetLogFile() []string
	GetServer() Server
	GetExternalAPI() ExternalApi
}

type Server struct {
	Protocol    string
	Host        string
	Port        string
	Version     string
	PrefixPath  string
	Application string
}

type Postgresql struct {
	Driver            string
	Address           string
	DefaultSchema     string
	MaxOpenConnection int
	MaxIdleConnection int
}

type ExternalApi struct {
	Url  string
	Path struct {
		List string
		View string
	}
}

func GenerateConfiguration(argument string) {
	var (
		err      error
		enviName string
	)

	if argument == "development" {
		err = godotenv.Load()
		if err != nil {
			fmt.Print("Error load file env -> ", err)
			os.Exit(2)
		}

		enviName = os.Getenv("PROJECTKEYDEV")
		temp := DevelopmentConfig{}
		err = gonfig.GetConf(enviName+"/config_development.json", &temp)
		if err != nil {
			fmt.Print("Error get config development -> ", err)
			os.Exit(2)
		}
		ApplicationConfiguration = &temp
	} else if argument == "docker" {
		enviName = os.Getenv("PROJECT-KEY-DOCKER")
		temp := DockerConfig{}
		err = gonfig.GetConf(enviName+"/config_docker.json", &temp)
		if err != nil {
			fmt.Println("Error get config docker -> ", err)
			os.Exit(2)
		}
		err = envconfig.Process(enviName+"/config_docker.json", &temp)
		if err != nil {
			fmt.Print("Error get config 2-> ", err)
			os.Exit(2)
		}
		ApplicationConfiguration = &temp
	}
}
