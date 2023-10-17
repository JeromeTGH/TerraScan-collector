package dataloader

import (
	"fmt"
	"strconv"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
	"github.com/JeromeTGH/TerraScan-collector/internal/mailsender"
)


func LoadNbStakedLunc(c chan<- lcd.StructReponseNbStakedLunc) () {

	var idxTentatives int

	for idxTentatives = 1 ; idxTentatives <= config.AppConfig.Lcd.NbOfAttemptsIfFailure ; idxTentatives++ {
		nbStakedLunc, errGetNbStakedLunc := lcd.GetNbStakedLunc()
		if errGetNbStakedLunc == "" {
			if idxTentatives > 1 {
				stringToReturn4 := fmt.Sprintf("LoadNbStakedLunc : success of attempt %d/%d", idxTentatives, config.AppConfig.Lcd.NbOfAttemptsIfFailure)
				logger.WriteLog("dataloader", stringToReturn4)
				mailsender.Sendmail("[TerraScan-collector] " + strconv.Itoa(idxTentatives) + " attempts to get data from LCD successfully", "<p><strong>LoadNbStakedLunc</strong></p><p>Script did " + strconv.Itoa(idxTentatives) + " attempts to get data from LCD successfully</p>")
			} else {
				logger.WriteLog("dataloader", "LoadNbStakedLunc : success")
			}
			
			c <- nbStakedLunc
			return
		} else {
			fmt.Println("Erreur")
			stringToReturn1 := fmt.Sprintf("LoadNbStakedLunc : failed attempt %d/%d", idxTentatives, config.AppConfig.Lcd.NbOfAttemptsIfFailure)
			logger.WriteLog("dataloader", stringToReturn1)
			stringToReturn2 := fmt.Sprintf("LoadNbStakedLunc : %s", errGetNbStakedLunc)
			logger.WriteLog("dataloader", stringToReturn2)
			// Pause de X minutes, avant de retenter, s'il reste des tentatives à faire
			if idxTentatives != config.AppConfig.Lcd.NbOfAttemptsIfFailure {
				time.Sleep(time.Duration(config.AppConfig.Lcd.NbMinutesOfBreakBetweenAttempts) * time.Second)
			}
		}
	}

	mailsender.Sendmail("[TerraScan-collector] impossible to load datas from LCD", "<p><strong>Impossible to load datas from LCD</strong></p><p>" + strconv.Itoa(config.AppConfig.Lcd.NbOfAttemptsIfFailure) + " attempts, and no success</p>")

	stringToReturn3 := fmt.Sprintf("LoadNbStakedLunc : impossible to load datas from LCD, even after %d attempts", config.AppConfig.Lcd.NbOfAttemptsIfFailure)
	logger.WriteLog("dataloader", stringToReturn3)

	c <- lcd.StructReponseNbStakedLunc{}

}
