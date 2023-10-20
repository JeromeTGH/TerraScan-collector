package dboperations

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
)


func SaveOraclePoolInfos(oraclePoolContentChannel <-chan lcd.StructReponseOraclePoolContent, exitChannel chan<- bool, channelForErrors chan<- string) () {

	oraclePoolContent := <- oraclePoolContentChannel

	fmt.Println("LUNCs in oracle pool :", oraclePoolContent.NbLuncInOraclePool)
	fmt.Println("USTCs in oracle pool :", oraclePoolContent.NbUstcInOraclePool)
	
	// if(oraclePoolContent != lcd.StructReponseOraclePoolContent{}) {

	// 	// Enregistrement du nombre de LUNC et USTC contenus dans l'Oracle Pool
	// 	dbOraclePoolContent.WriteCommunityPoolContentInDb(oraclePoolContent, channelForErrors)

	// } else {
	// 	fmt.Println("SaveCommunityPoolInfos error")
	// }

	exitChannel <- true
}