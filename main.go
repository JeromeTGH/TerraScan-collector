package main

import (
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

	// Channel qui stockera les éventuelles erreurs à noter dans le fichier log
	// channelForErrorMsgs := make(chan string)
	
	// Channels de LOADING data
	channelForTotalSuppliesLoading := make(chan lcd.StructReponseTotalSupplies)
	channelForNbStakedLunc := make(chan lcd.StructReponseNbStakedLunc)

	// Channels de SAVING data
	channelForTotalSuppliesAndStakingInfosSaving := make(chan bool)

	// Clôture de tous channels à la sortie de cette fonction
	defer close(channelForTotalSuppliesLoading)
	defer close(channelForNbStakedLunc)
	defer close(channelForTotalSuppliesAndStakingInfosSaving)
	
	// Chargement de données auprès du LCD
	go dataloader.LoadTotalSupplies(channelForTotalSuppliesLoading)
	go dataloader.LoadNbStakedLunc(channelForNbStakedLunc)

	// Récupération des données via les différents channels, et enregistrement en BDD
	go dboperations.SaveTotalSuppliesAndStakingInfos(channelForTotalSuppliesLoading, channelForNbStakedLunc, channelForTotalSuppliesAndStakingInfosSaving)

	// Attente que tous les "saving chanels" aient répondu
	<- channelForTotalSuppliesAndStakingInfosSaving


	// Et fin de ce script (avec inscription dans le log)
	logger.WriteLogWithoutPrinting("main", "script done")
	
}