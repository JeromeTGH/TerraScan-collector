package dataloader

import (
	"fmt"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/lcd"
	"github.com/JeromeTGH/TerraScan-collector/utils/logger"
)


func LoadTotalSupplies(appConfig *config.Config) (lcd.StructReponseTotalSupplies, string) {

	// appConfig.Lcd.NbOfAttemptsIfFailure :   				par défaut = 5 		(donc on essaye cinq fois maximum, puis on arrête)
	// appConfig.Lcd.NbMinutesOfBreakBetweenAttempts :		par défaut = 5 		(donc 5 minutes de pause, entre chaque tentative infructueuse)

	var idxTentatitves int

	for idxTentatitves = 1 ; idxTentatitves <= appConfig.Lcd.NbOfAttemptsIfFailure ; idxTentatitves++ {
		totalSupplies, errGetTotalSupplies := lcd.GetTotalSupplies(appConfig)
		if errGetTotalSupplies == "" {
			return totalSupplies, ""
		} else {
			stringToReturn1 := fmt.Sprintf("LoadTotalSupplies : failed attempt %d/%d", idxTentatitves, appConfig.Lcd.NbOfAttemptsIfFailure)
			logger.WriteLog("dataloader", stringToReturn1)
			stringToReturn2 := fmt.Sprintf("LoadTotalSupplies : %s", errGetTotalSupplies)
			logger.WriteLog("dataloader", stringToReturn2)
			// Pause de X minutes, avant de retenter, s'il reste des tentatives à faire
			if idxTentatitves != appConfig.Lcd.NbOfAttemptsIfFailure {
				time.Sleep(time.Duration(appConfig.Lcd.NbMinutesOfBreakBetweenAttempts) * time.Second)
			}
		}
	}

	stringToReturn3 := fmt.Sprintf("LoadTotalSupplies : impossible to load datas from LCD, even after %d attempts", appConfig.Lcd.NbOfAttemptsIfFailure)
	return lcd.StructReponseTotalSupplies{}, stringToReturn3
}