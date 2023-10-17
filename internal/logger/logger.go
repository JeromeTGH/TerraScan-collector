package logger

import (
	"fmt"
	"log"
	"os"
)

func WriteLog(textToAppend string) {

	// Écriture dans le fichier log
	WriteLogWithoutPrinting(textToAppend)

	// Et écriture dans la console aussi
	fmt.Println(textToAppend)
	
}

func WriteLogWithoutPrinting(textToAppend string) {

	// Ouverture du fichier log
	logFile, errOpenLogFile := os.OpenFile("./logs/activity.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)	// 6 = rw pour le créateur, 4=read only pour les autres
	if errOpenLogFile != nil {
		panic(errOpenLogFile)
	}
	defer logFile.Close()

	// Écriture d'une nouvelle ligne
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(textToAppend)
	// Nota : exemple de ligne écrite :
	//              2023/10/17 21:46:18 [main] script called
	
}