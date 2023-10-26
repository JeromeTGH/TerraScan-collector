package config

import (
	"os"

	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
	"gopkg.in/yaml.v3"
)

const configDataFile = "./config/private/private.yaml"

type Config struct {
	Lcd struct {
		Url 							string `yaml:"url"`
		GetTimeoutInSeconds 			int `yaml:"get_timeout_in_seconds"`
		NbOfAttemptsIfFailure 			int `yaml:"nb_of_attempts_if_failure"`
		NbMinutesOfBreakBetweenAttempts int `yaml:"nb_minutes_of_break_between_attempts"`
		PathForTotalSupplies 			string `yaml:"path_for_total_supplies"`
		PathForLuncStaking	 			string `yaml:"path_for_lunc_staking"`
		PathForCommunityPoolContent		string `yaml:"path_for_community_pool_content"`
		PathForOraclePoolContent 		string `yaml:"path_for_oracle_pool_content"`
	}
	Bdd struct {
		HostName  				string `yaml:"host_name"`
		DbName    				string `yaml:"db_name"`
		UserName  				string `yaml:"user_name"`
		Password  				string `yaml:"password"`
		Port      				int `yaml:"port"`
		TblTotalSuppliesName	string `yaml:"tbl_TotalSupplies_name"`
		TblLuncStaking 			string `yaml:"tbl_LuncStaking_name"`
		TblCommunityPoolContent string `yaml:"tbl_CommunityPoolContent_name"`
		TblOraclePoolContent 	string `yaml:"tbl_OraclePoolContent_name"`
	}
	Email struct {
		HostName  string `yaml:"host_name"`
		SmtpPort  int `yaml:"smtp_port"`
		From      string `yaml:"from"`
		Pwd       string `yaml:"pwd"`
		To        string `yaml:"to"`
	}
}

var AppConfig *Config

func LoadConfig() {

	// Teste la présence du fichier de configuration
	_, errStat := os.Stat(configDataFile)
	if errStat != nil {
		if os.IsNotExist(errStat) {
			logger.WriteLog("[config] error : config file not found")
			os.Exit(500)
		} else {
			logger.WriteLog("[config] error : " + errStat.Error())
			panic(errStat)
		}
	}

	// Lit le fichier contenant les données de config
	privateData, errOsReadFile := os.ReadFile(configDataFile)
	if errOsReadFile != nil {
		logger.WriteLog("[config] failed to read config file")
		panic(errOsReadFile)
	}

	// Mesure la quantité de données, pour savoir le fichier de config contient un "minimum" de données
	if len(privateData) < 300 {
		logger.WriteLog("[config] error : not enough data, in config file")
		os.Exit(500)
	}

	// Parse le fichier lu, pour le mettre dans la structure de config
	errYamlUnmarshal := yaml.Unmarshal([]byte(privateData), &AppConfig)
	if errYamlUnmarshal != nil {
		logger.WriteLog("[config] failed to unmarshal config data")
		panic(errYamlUnmarshal)
	}

}