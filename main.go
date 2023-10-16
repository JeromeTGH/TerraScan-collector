package main

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader"
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations"
	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
)



func main() {

	// Chargement des données de configuration, dans la variable "AppConfig"
	config.LoadConfig()
	
	// Enregistrement des date/heure de démarrage, dans le fichier log
	logger.WriteLogWithoutPrinting("main", "script started")
	
	// Channels de données
	channelForTotalSupplies := make(chan lcd.StructReponseTotalSupplies)
	// channelForNbStakedLunc := make(chan float64)
	
	// Chargement de données auprès du LCD
	go dataloader.LoadTotalSupplies(channelForTotalSupplies)
	// go dataloader.LoadNbStakedLunc(&channelForNbStakedLunc, &wg)

	// Récupération des données, via les différents channels
	totalSupplies := <- channelForTotalSupplies
	// nbStakedLunc := <- channelForNbStakedLunc


	// Enregistrement en base de données
	if(totalSupplies != lcd.StructReponseTotalSupplies{}) {

		// Enregistrement des total supplies		
		dboperations.WriteTotalSuppliesInDb(totalSupplies)

		// // Enregistrement du taux de staking
		// totalSupplyOfLunc := totalSupplies.LuncTotalSupply
		// nbStakedLunc := <- channelForNbStakedLunc
		// dboperations.WriteStakingPercentageInDb(nbStakedLunc, totalSupplyOfLunc)

	} else {
		fmt.Println("Total supplies error")
	}


	// Clôture des channels
	close(channelForTotalSupplies)

	// Inscription dans le log de la fin de ce script
	logger.WriteLogWithoutPrinting("main", "script done")
	
}