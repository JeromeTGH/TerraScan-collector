package application

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Lcd struct {
		Url string `yaml:"url"`
	}
	Bdd struct {
		HostName string `yaml:"host_name"`
		BddName  string `yaml:"bdd_name"`
		UserName string `yaml:"user_name"`
		Password  string `yaml:"password"`
		Port      int16 `yaml:"port"`
	}
	Email struct {
		HostName string `yaml:"host_name"`
		SmtpPort int16 `yaml:"smtp_port"`
		From      string `yaml:"from"`
		Pwd       string `yaml:"pwd"`
		To        string `yaml:"to"`
	}
}

func LoadConfig(appConfig *Config)  {

	privateData, errOsReadFile := os.ReadFile("application/private/private.yaml")

	if errOsReadFile != nil {
		panic(errOsReadFile)
	}

	errYamlUnmarshal := yaml.Unmarshal([]byte(privateData), &appConfig)

	if errYamlUnmarshal != nil {
		panic(errYamlUnmarshal)
	}

}