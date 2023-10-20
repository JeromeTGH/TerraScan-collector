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

	// Channel qui contiendra toutes les lignes "intermédiaires" à stocker dans le fichier log
	channelForLogMsgs := make(chan string, 1000)
	
	// Channels de LOADING data
	channelForTotalSuppliesLoading := make(chan lcd.StructReponseTotalSupplies, 1)
	channelForNbStakedLunc := make(chan lcd.StructReponseNbStakedLunc, 1)

	// Channels de SAVING data
	channelForTotalSuppliesAndStakingInfosSaving := make(chan bool, 1)

	// Clôture de tous channels à la sortie de cette fonction
	defer close(channelForLogMsgs)
	defer close(channelForTotalSuppliesLoading)
	defer close(channelForNbStakedLunc)
	defer close(channelForTotalSuppliesAndStakingInfosSaving)
	
	// Chargement asynchrone de données auprès du LCD
	go dataloader.LoadTotalSupplies(channelForTotalSuppliesLoading, channelForLogMsgs)
	go dataloader.LoadNbStakedLunc(channelForNbStakedLunc, channelForLogMsgs)

	// Enregistrements asynchrone en BDD
	go dboperations.SaveTotalSuppliesAndStakingInfos(channelForTotalSuppliesLoading, channelForNbStakedLunc, channelForTotalSuppliesAndStakingInfosSaving, channelForLogMsgs)

	// Attente que tous les "saving chanels" aient fini leur tâche
	<- channelForTotalSuppliesAndStakingInfosSaving

	// Enregistrement de tous les messages "intermédiaire" à inscrire dans le log
	nbMsgs := len(channelForLogMsgs)
	for idxMsg := 0; idxMsg < nbMsgs ; idxMsg++ {
		msg := <- channelForLogMsgs
		logger.WriteLog(msg)
	}

	// Et fin de ce script (avec mention dans le fichier log)
	logger.WriteLogWithoutPrinting("[main] script done")
	
}