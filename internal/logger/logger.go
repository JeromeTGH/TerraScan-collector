package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func WriteLog(textToAppend string) {

	// Écriture dans le fichier log
	WriteLogWithoutPrinting(textToAppend)

	// Et écriture dans la console aussi
	fmt.Println(textToAppend)
	
}

func WriteLogWithoutPrinting(textToAppend string) {

	// ***********************************************************
	// Enregistrement de l'activité dans le fichier "activity.log"
	// ***********************************************************

	// Ouverture du fichier activity.log
	activityLogFile, errOpenActivityLogFile := os.OpenFile("./logs/activity.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)	// 6 = rw pour le créateur, 4=read only pour les autres
	if errOpenActivityLogFile != nil {
		panic(errOpenActivityLogFile)
	}
	defer activityLogFile.Close()

	// Écriture d'une nouvelle ligne
	logger := log.New(activityLogFile, "", log.LstdFlags)
	logger.Println(textToAppend)		// Nota : exemple de ligne écrite :      2023/10/17 21:46:18 [main] script called
	
	// *******************************************************
	// Enregistrement des erreurs dans le fichier "errors.log"
	// *******************************************************
	message := strings.ToLower(textToAppend)
	if (strings.Contains(message, "error") || strings.Contains(message, "failed") || strings.Contains(message, "impossible")) {

		// Ouverture du fichier errors.log
		errorsLogFile, errOpenErrorsLogFile := os.OpenFile("./logs/errors.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)	// 6 = rw pour le créateur, 4=read only pour les autres
		if errOpenErrorsLogFile != nil {
			panic(errOpenErrorsLogFile)
		}
		defer errorsLogFile.Close()

		// Écriture d'une nouvelle ligne
		logger := log.New(errorsLogFile, "", log.LstdFlags)
		logger.Println(textToAppend)		// Nota : exemple de ligne écrite :      2023/10/17 21:46:18 [main] script called

	}


	
}