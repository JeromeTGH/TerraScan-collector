package dataloader

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
	"github.com/JeromeTGH/TerraScan-collector/internal/mailsender"
)


func LoadTotalSupplies() (lcd.StructReponseTotalSupplies) {

	// config.AppConfig.Lcd.NbOfAttemptsIfFailure :   				par défaut = 5 		(donc on essaye cinq fois maximum, puis on arrête)
	// config.AppConfig.Lcd.NbMinutesOfBreakBetweenAttempts :		par défaut = 5 		(donc 5 minutes de pause, entre chaque tentative infructueuse)

	var idxTentatitves int

	for idxTentatitves = 1 ; idxTentatitves <= config.AppConfig.Lcd.NbOfAttemptsIfFailure ; idxTentatitves++ {
		totalSupplies, errGetTotalSupplies := lcd.GetTotalSupplies()
		if errGetTotalSupplies == "" {
			if idxTentatitves > 1 {
				mailsender.Sendmail("[TerraScan-collector] " + strconv.Itoa(idxTentatitves) + " attemps to get data from LCD", "<p><strong>LoadTotalSupplies</strong></p><p>" + strconv.Itoa(idxTentatitves) + " attemps to get data from LCD</p>")
			}
			return totalSupplies
		} else {
			stringToReturn1 := fmt.Sprintf("LoadTotalSupplies : failed attempt %d/%d", idxTentatitves, config.AppConfig.Lcd.NbOfAttemptsIfFailure)
			logger.WriteLog("dataloader", stringToReturn1)
			stringToReturn2 := fmt.Sprintf("LoadTotalSupplies : %s", errGetTotalSupplies)
			logger.WriteLog("dataloader", stringToReturn2)
			// Pause de X minutes, avant de retenter, s'il reste des tentatives à faire
			if idxTentatitves != config.AppConfig.Lcd.NbOfAttemptsIfFailure {
				time.Sleep(time.Duration(config.AppConfig.Lcd.NbMinutesOfBreakBetweenAttempts) * time.Second)
			}
		}
	}

	mailsender.Sendmail("[TerraScan-collector] impossible to load datas from LCD", "<p><strong>Impossible to load datas from LCD</strong></p><p>" + strconv.Itoa(config.AppConfig.Lcd.NbOfAttemptsIfFailure) + " attemps, and no success</p>")

	stringToReturn3 := fmt.Sprintf("LoadTotalSupplies : impossible to load datas from LCD, even after %d attempts", config.AppConfig.Lcd.NbOfAttemptsIfFailure)
	logger.WriteLog("dataloader", stringToReturn3)
	os.Exit(500)

	return lcd.StructReponseTotalSupplies{}		// Ne sera jamais exécuté, car os.Exit juste avant
}