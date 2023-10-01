package main

import (
	"fmt"
	"os"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/utils/dataloader"
	"github.com/JeromeTGH/TerraScan-collector/utils/logger"
)


func main() {

	// Chargement des données de configuration, dans la variable "AppConfig"
	config.LoadConfig()
	
	// Inscription dans le log de l'appel de ce script
	logger.WriteLogWithoutPrinting("main", "script called")

	// Chargement des données, en faisant appel au LCD
	dataFromLcd, errFromLcd := dataloader.LoadTotalSupplies()
	if errFromLcd != "" {
		logger.WriteLog("main", errFromLcd)
		os.Exit(500)
	}

	// Afichage dans la console (debug)
	fmt.Printf("LUNCtotalSupply = %d\n", dataFromLcd.LuncTotalSupply)
	fmt.Printf("USTCtotalSupply = %d\n", dataFromLcd.UstcTotalSupply)

}