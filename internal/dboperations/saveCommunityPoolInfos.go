package dboperations

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbCommunityPoolContent"
)


func SaveCommunityPoolInfos(communityPoolContentChannel <-chan lcd.StructReponseCommunityPoolContent, exitChannel chan<- bool, channelForErrors chan<- string) () {

	communityPoolContent := <- communityPoolContentChannel

	fmt.Println("LUNCs in community pool :", communityPoolContent.NbLuncInCommunityPool)
	fmt.Println("USTCs in community pool :", communityPoolContent.NbUstcInCommunityPool)
	
	if(communityPoolContent != lcd.StructReponseCommunityPoolContent{}) {

		// Enregistrement du nombre de LUNC et USTC contenus dans la Community Pool
		dbCommunityPoolContent.WriteCommunityPoolContentInDb(communityPoolContent, channelForErrors)

	} else {
		fmt.Println("SaveCommunityPoolInfos error")
	}

	exitChannel <- true
}