package application

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configDataFile = "application/private/private.yaml"

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

func LoadConfig(appConfig *Config) {

	// Teste la présence du fichier de configuration
	_, errStat := os.Stat(configDataFile)
	if errStat != nil {
		if os.IsNotExist(errStat) {
			fmt.Println("Config file not found.")
			os.Exit(500)
		} else {
			panic(errStat)
		}
	}

	// Lit le fichier contenant les données de config
	privateData, errOsReadFile := os.ReadFile(configDataFile)
	if errOsReadFile != nil {
		panic(errOsReadFile)
	}

	// Mesure la quantité de données, pour savoir le fichier de config contient un "minimum" de données
	if len(privateData) < 300 {
		fmt.Println("Not enough data, in config file.")
		os.Exit(500)
	}

	// Parse le fichier lu, pour le mettre dans la structure de config
	errYamlUnmarshal := yaml.Unmarshal([]byte(privateData), &appConfig)
	if errYamlUnmarshal != nil {
		panic(errYamlUnmarshal)
	}

}