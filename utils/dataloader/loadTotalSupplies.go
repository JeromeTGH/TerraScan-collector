package dataloader

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/lcd"
)


func LoadTotalSupplies(appConfig *config.Config) (lcd.StructReponseTotalSupplies, error) {

	// appConfig.Lcd.NbOfAttemptsIfFailure :   				par défaut = 5 		(donc on essaye cinq fois maximum, puis on arrête)
	// appConfig.Lcd.NbMinutesOfBreakBetweenAttempts :		par défaut = 5 		(donc 5 minutes de pause, entre chaque tentative infructueuse)

	var idxTentatitves int

	for idxTentatitves = 1 ; idxTentatitves <= appConfig.Lcd.NbOfAttemptsIfFailure ; idxTentatitves++ {
		totalSupplies, errGetTotalSupplies := lcd.GetTotalSupplies(appConfig)
		if errGetTotalSupplies == nil {
			return totalSupplies, nil
		} else {
			fmt.Println("[dataloader] Failed attempt", idxTentatitves, "/", appConfig.Lcd.NbOfAttemptsIfFailure)
			fmt.Println(errGetTotalSupplies.Error())
			// Pause de X minutes, avant de retenter
			time.Sleep(time.Duration(appConfig.Lcd.NbMinutesOfBreakBetweenAttempts) * time.Minute)
		}
	}

	return lcd.StructReponseTotalSupplies{}, errors.New("[dataloader] impossible to load datas from LCD, even after " + strconv.Itoa(appConfig.Lcd.NbOfAttemptsIfFailure) + " attempts")
}