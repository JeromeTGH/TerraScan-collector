package dataloader

import (
	"fmt"
	"strconv"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/mailsender"
)


func LoadCommunityPoolContent(channelForCommunityPoolContent chan<- lcd.StructReponseCommunityPoolContent, channelForLogsMsgs chan<- string) {

	var idxTentatives int

	for idxTentatives = 1 ; idxTentatives <= config.AppConfig.Lcd.NbOfAttemptsIfFailure ; idxTentatives++ {
		communityPoolContent, errCommunityPoolContent := lcd.GetCommunityPoolContent(channelForLogsMsgs)
		if errCommunityPoolContent == "" {
			if idxTentatives > 1 {
				stringToReturn4 := fmt.Sprintf("[dataloader] LoadCommunityPoolContent : success of attempt %d/%d", idxTentatives, config.AppConfig.Lcd.NbOfAttemptsIfFailure)
				channelForLogsMsgs <- stringToReturn4
				mailsender.Sendmail("[TerraScan-collector] " + strconv.Itoa(idxTentatives) + " attempts to get data from LCD successfully", "<p><strong>LoadCommunityPoolContent</strong></p><p>Script did " + strconv.Itoa(idxTentatives) + " attempts to get data from LCD successfully</p>", channelForLogsMsgs)
			} else {
				channelForLogsMsgs <- "[dataloader] LoadCommunityPoolContent : success"
			}
			
			channelForCommunityPoolContent <- communityPoolContent
			return
		} else {
			stringToReturn1 := fmt.Sprintf("[dataloader] LoadCommunityPoolContent : failed attempt %d/%d", idxTentatives, config.AppConfig.Lcd.NbOfAttemptsIfFailure)
			channelForLogsMsgs <- stringToReturn1
			stringToReturn2 := fmt.Sprintf("[dataloader] LoadCommunityPoolContent : %s", errCommunityPoolContent)
			channelForLogsMsgs <- stringToReturn2
			// Pause de X minutes, avant de retenter, s'il reste des tentatives à faire
			if idxTentatives != config.AppConfig.Lcd.NbOfAttemptsIfFailure {
				time.Sleep(time.Duration(config.AppConfig.Lcd.NbMinutesOfBreakBetweenAttempts) * time.Minute)
			}
		}
	}

	mailsender.Sendmail("[TerraScan-collector] impossible to load datas from LCD", "<p><strong>Impossible to load datas from LCD</strong></p><p>" + strconv.Itoa(config.AppConfig.Lcd.NbOfAttemptsIfFailure) + " attempts, and no success</p>", channelForLogsMsgs)

	stringToReturn3 := fmt.Sprintf("[dataloader] LoadCommunityPoolContent : impossible to load datas from LCD, even after %d attempts", config.AppConfig.Lcd.NbOfAttemptsIfFailure)
	channelForLogsMsgs <- stringToReturn3

	channelForCommunityPoolContent <- lcd.StructReponseCommunityPoolContent{}

}
