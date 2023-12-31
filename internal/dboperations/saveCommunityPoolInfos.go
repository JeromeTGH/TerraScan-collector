package dboperations

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbCommunityPoolContent"
)


func SaveCommunityPoolInfos(communityPoolContentChannel <-chan lcd.StructReponseCommunityPoolContent, exitChannel chan<- bool, channelForLogsMsgs chan<- string) () {

	communityPoolContent := <- communityPoolContentChannel
	
	if(communityPoolContent != lcd.StructReponseCommunityPoolContent{}) {

		// Enregistrement du nombre de LUNC et USTC contenus dans la Community Pool
		dbCommunityPoolContent.WriteCommunityPoolContentInDb(communityPoolContent, channelForLogsMsgs)

	} else {
		fmt.Println("SaveCommunityPoolInfos error")
	}

	exitChannel <- true
}