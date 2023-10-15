package main

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/asyncroutines"
	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
)



func main() {

	// Chargement des données de configuration, dans la variable "AppConfig"
	config.LoadConfig()
	
	// Enregistrement des date/heure de démarrage, dans le fichier log
	logger.WriteLogWithoutPrinting("main", "script started")

	// Lancement asynchrone de fonctions, via des channels
	channelForTotalSupplies := asyncroutines.AsyncLoadAndSaveTotalSupplies()

	// Résultats
	fmt.Println("Retour 'channelForTotalSupplies' =", <- channelForTotalSupplies)

	// Inscription dans le log de la fin de ce script
	logger.WriteLogWithoutPrinting("main", "script done")
	
}