package dboperations

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
)


func SaveCommunityPoolInfos(communityPoolContentChannel <-chan lcd.StructReponseCommunityPoolContent, exitChannel chan<- bool, channelForErrors chan<- string) () {

	communityPoolContent := <- communityPoolContentChannel

	fmt.Println("LUNCs in community pool :", communityPoolContent.NbLuncInCommunityPool)
	fmt.Println("USTCs in community pool :", communityPoolContent.NbUstcInCommunityPool)
	
	// if(totalSuppliesStruct != lcd.StructReponseTotalSupplies{}) {

	// 	// Enregistrement des total supplies
	// 	dbTotalSupplies.WriteTotalSuppliesInDb(totalSuppliesStruct, channelForErrors)

	// 	// Enregistrement du nombre de LUNC stakés, et du taux de staking
	// 	dbNbStakedLunc.WriteNbStakedLuncInDb(nbStakedLuncStruct, stakingPercentage, channelForErrors)

	// } else {
	// 	fmt.Println("Total supplies error")
	// }

	exitChannel <- true
}