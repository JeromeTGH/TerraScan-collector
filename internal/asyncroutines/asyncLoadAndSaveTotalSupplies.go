package asyncroutines

import (
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations"
)


func AsyncLoadAndSaveTotalSupplies() chan int {

	// Création d'un channel de retour, pour cette fonction asynchrone
	r := make(chan int)

	go func() {
		// Chargement des données, en faisant appel au LCD
		dataFromLcd := dataloader.LoadTotalSupplies()

		// Écriture en base de données
		dboperations.WriteTotalSuppliesInDb(dataFromLcd)

		// Et signalement de fin, via le channel de cette fonction
		r <- 1
	}()

	return r

}