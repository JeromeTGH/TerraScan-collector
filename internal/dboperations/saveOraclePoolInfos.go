package dboperations

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbOraclePoolContent"
)


func SaveOraclePoolInfos(oraclePoolContentChannel <-chan lcd.StructReponseOraclePoolContent, exitChannel chan<- bool, channelForErrors chan<- string) () {

	oraclePoolContent := <- oraclePoolContentChannel
	
	if(oraclePoolContent != lcd.StructReponseOraclePoolContent{}) {

		// Enregistrement du nombre de LUNC et USTC contenus dans l'Oracle Pool
		dbOraclePoolContent.WriteOraclePoolContentInDb(oraclePoolContent, channelForErrors)

	} else {
		fmt.Println("SaveCommunityPoolInfos error")
	}

	exitChannel <- true
}