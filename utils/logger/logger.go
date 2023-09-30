package logger

import (
	"fmt"
	"log"
	"os"
)

func WriteLog(prefix string, textToAppend string) {

	// Écriture dans le fichier log
	WriteLogWithoutPrinting(prefix, textToAppend)

	// Et écriture dans la console aussi
	fmt.Println("[" + prefix + "] " + textToAppend)
	
}

func WriteLogWithoutPrinting(prefix string, textToAppend string) {

	// Ouverture du fichier log
	logFile, errOpenLogFile := os.OpenFile("logs/activity.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)	// 6 = rw pour le créateur, 4=read only pour les autres
	if errOpenLogFile != nil {
		panic(errOpenLogFile)
	}
	defer logFile.Close()

	// Écriture d'une nouvelle ligne
	logger := log.New(logFile, "[" + prefix + "] ", log.LstdFlags)
	logger.Println(textToAppend)
	// Nota : exemple de ligne écrite :
	//         [main] 2023/09/30 15:12:02 Script called
	
}