package config

import (
	"os"

	"github.com/JeromeTGH/TerraScan-collector/utils/logger"
	"gopkg.in/yaml.v3"
)

const configDataFile = "config/private/private.yaml"

type Config struct {
	Lcd struct {
		Url string `yaml:"url"`
		GetTimeoutInSeconds int `yaml:"get_timeout_in_seconds"`
		NbOfAttemptsIfFailure int `yaml:"nb_of_attempts_if_failure"`
		NbMinutesOfBreakBetweenAttempts int `yaml:"nb_minutes_of_break_between_attempts"`
	}
	Bdd struct {
		HostName  string `yaml:"host_name"`
		BddName   string `yaml:"bdd_name"`
		UserName  string `yaml:"user_name"`
		Password  string `yaml:"password"`
		Port      int `yaml:"port"`
	}
	Email struct {
		HostName  string `yaml:"host_name"`
		SmtpPort  int `yaml:"smtp_port"`
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
			logger.WriteLog("config", "config file not found")
			os.Exit(500)
		} else {
			logger.WriteLog("config", errStat.Error())
			panic(errStat)
		}
	}

	// Lit le fichier contenant les données de config
	privateData, errOsReadFile := os.ReadFile(configDataFile)
	if errOsReadFile != nil {
		logger.WriteLog("config", "failed to read config file")
		panic(errOsReadFile)
	}

	// Mesure la quantité de données, pour savoir le fichier de config contient un "minimum" de données
	if len(privateData) < 300 {
		logger.WriteLog("config", "not enough data, in config file")
		os.Exit(500)
	}

	// Parse le fichier lu, pour le mettre dans la structure de config
	errYamlUnmarshal := yaml.Unmarshal([]byte(privateData), &appConfig)
	if errYamlUnmarshal != nil {
		logger.WriteLog("config", "failed to unmarshal config data")
		panic(errYamlUnmarshal)
	}

}