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
	logger.WriteLogWithoutPrinting("[main] script called")

	// Channels
	channelForLogMsgs := make(chan string, 1000)											// Logs data
	channelForTotalSuppliesLoading := make(chan lcd.StructReponseTotalSupplies, 1)			// Loading data
	channelForNbStakedLuncLoading := make(chan lcd.StructReponseNbStakedLunc, 1)			// Loading data
	channelForCommunityPoolLoading := make(chan lcd.StructReponseCommunityPoolContent, 1)	// Loading data
	channelForOraclePoolLoading := make(chan lcd.StructReponseOraclePoolContent, 1)			// Loading data
	channelForTotalSuppliesAndStakingInfosSaving := make(chan bool, 1)						// Saving data
	channelForCommunityPoolInfosSaving := make(chan bool, 1)								// Saving data
	channelForOraclePoolInfosSaving := make(chan bool, 1)									// Saving data

	// Clôture de tous channels à la sortie de cette fonction (inutile dans fonction "main", mais pour la bonne forme, si déplacé plus tard)
	defer close(channelForLogMsgs)
	defer close(channelForTotalSuppliesLoading)
	defer close(channelForNbStakedLuncLoading)
	defer close(channelForCommunityPoolLoading)
	defer close(channelForOraclePoolLoading)
	defer close(channelForTotalSuppliesAndStakingInfosSaving)
	defer close(channelForCommunityPoolInfosSaving)
	defer close(channelForOraclePoolInfosSaving)
	
	// Chargement asynchrone de données auprès du LCD
	go dataloader.LoadTotalSupplies(channelForTotalSuppliesLoading, channelForLogMsgs)
	go dataloader.LoadNbStakedLunc(channelForNbStakedLuncLoading, channelForLogMsgs)
	go dataloader.LoadCommunityPoolContent(channelForCommunityPoolLoading, channelForLogMsgs)
	go dataloader.LoadOraclePoolContent(channelForOraclePoolLoading, channelForLogMsgs)

	// Enregistrements asynchrone en BDD
	go dboperations.SaveTotalSuppliesAndStakingInfos(channelForTotalSuppliesLoading, channelForNbStakedLuncLoading, channelForTotalSuppliesAndStakingInfosSaving, channelForLogMsgs)
	go dboperations.SaveCommunityPoolInfos(channelForCommunityPoolLoading, channelForCommunityPoolInfosSaving, channelForLogMsgs)
	go dboperations.SaveOraclePoolInfos(channelForOraclePoolLoading, channelForOraclePoolInfosSaving, channelForLogMsgs)

	// Attente que tous les "saving chanels" aient fini leur tâche
	<- channelForTotalSuppliesAndStakingInfosSaving
	<- channelForCommunityPoolInfosSaving
	<- channelForOraclePoolInfosSaving

	// Enregistrement de tous les messages "intermédiaire" à inscrire dans le log
	nbMsgs := len(channelForLogMsgs)
	for idxMsg := 0; idxMsg < nbMsgs ; idxMsg++ {
		msg := <- channelForLogMsgs
		logger.WriteLog(msg)
	}

	// Et fin de ce script (avec mention dans le fichier log)
	logger.WriteLogWithoutPrinting("[main] script done")
	
}