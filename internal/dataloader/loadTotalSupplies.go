package dataloader

import (
	"fmt"
	"strconv"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/mailsender"
)


func LoadTotalSupplies(channelForTotalSupplies chan<- lcd.StructReponseTotalSupplies, channelForErrors chan<- string) {

	var idxTentatives int

	for idxTentatives = 1 ; idxTentatives <= config.AppConfig.Lcd.NbOfAttemptsIfFailure ; idxTentatives++ {
		totalSupplies, errGetTotalSupplies := lcd.GetTotalSupplies(channelForErrors)
		if errGetTotalSupplies == "" {
			if idxTentatives > 1 {
				stringToReturn4 := fmt.Sprintf("[dataloader] LoadTotalSupplies : success of attempt %d/%d", idxTentatives, config.AppConfig.Lcd.NbOfAttemptsIfFailure)
				channelForErrors <- stringToReturn4
				mailsender.Sendmail("[TerraScan-collector] " + strconv.Itoa(idxTentatives) + " attempts to get data from LCD successfully", "<p><strong>LoadTotalSupplies</strong></p><p>Script did " + strconv.Itoa(idxTentatives) + " attempts to get data from LCD successfully</p>", channelForErrors)
			} else {
				channelForErrors <- "[dataloader] LoadTotalSupplies : success"
			}
			
			channelForTotalSupplies <- totalSupplies
			return
		} else {
			fmt.Println("Erreur")
			stringToReturn1 := fmt.Sprintf("[dataloader] LoadTotalSupplies : failed attempt %d/%d", idxTentatives, config.AppConfig.Lcd.NbOfAttemptsIfFailure)
			channelForErrors <- stringToReturn1
			stringToReturn2 := fmt.Sprintf("[dataloader] LoadTotalSupplies : %s", errGetTotalSupplies)
			channelForErrors <- stringToReturn2
			// Pause de X minutes, avant de retenter, s'il reste des tentatives à faire
			if idxTentatives != config.AppConfig.Lcd.NbOfAttemptsIfFailure {
				time.Sleep(time.Duration(config.AppConfig.Lcd.NbMinutesOfBreakBetweenAttempts) * time.Minute)
			}
		}
	}

	mailsender.Sendmail("[TerraScan-collector] impossible to load datas from LCD", "<p><strong>Impossible to load datas from LCD</strong></p><p>" + strconv.Itoa(config.AppConfig.Lcd.NbOfAttemptsIfFailure) + " attempts, and no success</p>", channelForErrors)

	stringToReturn3 := fmt.Sprintf("[dataloader] LoadTotalSupplies : impossible to load datas from LCD, even after %d attempts", config.AppConfig.Lcd.NbOfAttemptsIfFailure)
	channelForErrors <- stringToReturn3

	channelForTotalSupplies <- lcd.StructReponseTotalSupplies{}

}
