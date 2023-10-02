package main

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/utils/dataloader"
	"github.com/JeromeTGH/TerraScan-collector/utils/dboperations"
	"github.com/JeromeTGH/TerraScan-collector/utils/logger"
)


func main() {

	// Chargement des données de configuration, dans la variable "AppConfig"
	config.LoadConfig()
	
	// Inscription dans le log de l'appel de ce script
	logger.WriteLogWithoutPrinting("main", "script called")

	// Chargement des données, en faisant appel au LCD
	dataFromLcd := dataloader.LoadTotalSupplies()

	// Écriture en base de données
	dboperations.DropTotalSuppliesTable()
	fmt.Println(dataFromLcd)
	// dboperations.WriteTotalSuppliesInDb(dataFromLcd)

}