package dataloader

import (
	"errors"
	"fmt"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/application"
	"github.com/JeromeTGH/TerraScan-collector/lcd"
)


func LoadTotalSupplies(appConfig *application.Config) (lcd.StructReponseTotalSupplies, error) {

	// appConfig.Lcd.NbTentativesSiEchec :   				par défaut = 5 		(donc on essaye cinq fois maximum, puis on arrête)
	// appConfig.Lcd.NbMinutesDePauseEntreTentatives :		par défaut = 5 		(donc 5 minutes de pause, entre chaque tentative infructueuse)

	var idxTentatitves uint8

	for idxTentatitves = 1 ; idxTentatitves <= appConfig.Lcd.NbTentativesSiEchec ; idxTentatitves++ {
		totalSupplies, errGetTotalSupplies := lcd.GetTotalSupplies(appConfig)
		if errGetTotalSupplies == nil {
			return totalSupplies, nil
		} else {
			fmt.Println("[LoadTotalSupplies] Échec tentative", idxTentatitves, "/", appConfig.Lcd.NbTentativesSiEchec)
			fmt.Println(errGetTotalSupplies.Error())
			// Pause de X minutes, avant de retenter
			time.Sleep(time.Duration(appConfig.Lcd.NbMinutesDePauseEntreTentatives) * time.Minute)
		}
	}

	return lcd.StructReponseTotalSupplies{}, errors.New("impossible to load datas from LCD")
}