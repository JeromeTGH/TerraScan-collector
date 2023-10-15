package main

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations"
	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
)



func main() {

	// Chargement des données de configuration, dans la variable "AppConfig"
	config.LoadConfig()
	
	// Enregistrement des date/heure de démarrage, dans le fichier log
	logger.WriteLogWithoutPrinting("main", "script called")

	// Lancement asynchrone de fonctions, via des channels
	channelForTotalSupplies := asyncLoadAndSaveTotalSupplies()

	// Résultats
	fmt.Println("Retour 'channelForTotalSupplies' =", <- channelForTotalSupplies)

	// Inscription dans le log de la fin de ce script
	logger.WriteLogWithoutPrinting("main", "script done")
	

}




func asyncLoadAndSaveTotalSupplies() chan int {

	// Création d'un channel de retour, pour cette fonction asynchrone
	r := make(chan int)

	go func() {
		// Chargement des données, en faisant appel au LCD
		dataFromLcd := dataloader.LoadTotalSupplies()

		// Écriture en base de données
		dboperations.WriteTotalSuppliesInDb(dataFromLcd)

		// Et signalement de fin, via le channel de cette fonction
		r <- 1
	}()

	return r

}