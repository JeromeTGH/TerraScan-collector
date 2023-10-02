package main

import (
	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations"
	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
)


func main() {

	// Chargement des données de configuration, dans la variable "AppConfig"
	config.LoadConfig()
	
	// Inscription dans le log de l'appel de ce script
	logger.WriteLogWithoutPrinting("main", "script called")

	// Chargement des données, en faisant appel au LCD
	dataFromLcd := dataloader.LoadTotalSupplies()

	// Écriture en base de données
	dboperations.WriteTotalSuppliesInDb(dataFromLcd)

}