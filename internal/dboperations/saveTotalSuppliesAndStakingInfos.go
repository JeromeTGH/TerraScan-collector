package dboperations

import (
	"fmt"
	"math"

	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbLuncStaking"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbTotalSupplies"
)


func SaveTotalSuppliesAndStakingInfos(totalSuppliesChannel <-chan lcd.StructReponseTotalSupplies, nbStakedLuncChannel <-chan lcd.StructReponseNbStakedLunc, exitChannel chan<- bool, channelForErrors chan<- string) () {

	totalSuppliesStruct := <- totalSuppliesChannel
	nbStakedLuncStruct := <- nbStakedLuncChannel

	stakingPercentage := 100 * float64(nbStakedLuncStruct.NbStakedLunc) / float64(totalSuppliesStruct.LuncTotalSupply)
	stakingPercentage = math.Round(stakingPercentage*100)/100		// Arrondi à 2 chiffres après la virgule
	
	if(totalSuppliesStruct != lcd.StructReponseTotalSupplies{}) {

		// Enregistrement des total supplies
		dbTotalSupplies.WriteTotalSuppliesInDb(totalSuppliesStruct, channelForErrors)

		// Enregistrement du nombre de LUNC stakés, et du taux de staking
		dbLuncStaking.WriteNbStakedLuncInDb(nbStakedLuncStruct, stakingPercentage, channelForErrors)

	} else {
		fmt.Println("Total supplies error")
	}

	exitChannel <- true
}