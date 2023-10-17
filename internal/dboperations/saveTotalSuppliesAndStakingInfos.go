package dboperations

import (
	"fmt"
	"math"

	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbTotalSupplies"
)


func SaveTotalSuppliesAndStakingInfos(totalSuppliesChannel <-chan lcd.StructReponseTotalSupplies, nbStakedLuncChannel <-chan lcd.StructReponseNbStakedLunc, exitChannel chan<- bool, channelForErrors chan<- string) () {

	totalSuppliesStruct := <- totalSuppliesChannel
	nbStakedLuncStruct := <- nbStakedLuncChannel

	stakingPercentage := 100 * float64(nbStakedLuncStruct.NbStakedLunc) / float64(totalSuppliesStruct.LuncTotalSupply)
	stakingPercentage = math.Round(stakingPercentage*100)/100		// Arrondi à 2 chiffres après la virgule

	fmt.Println("LuncTotalSupply :", totalSuppliesStruct.LuncTotalSupply)
	fmt.Println("nbStakedLunc :", nbStakedLuncStruct.NbStakedLunc)
	fmt.Println("stakingPercentage :", stakingPercentage, "%")
	
	if(totalSuppliesStruct != lcd.StructReponseTotalSupplies{}) {

		// Enregistrement des total supplies
		dbTotalSupplies.WriteTotalSuppliesInDb(totalSuppliesStruct, channelForErrors)

		// // Enregistrement du taux de staking
		// totalSupplyOfLunc := totalSupplies.LuncTotalSupply
		// nbStakedLunc := <- channelForNbStakedLunc
		// dboperations.WriteStakingPercentageInDb(nbStakedLunc, totalSupplyOfLunc)

	} else {
		fmt.Println("Total supplies error")
	}

	exitChannel <- true
}