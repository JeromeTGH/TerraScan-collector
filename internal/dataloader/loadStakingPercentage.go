package dataloader

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"github.com/JeromeTGH/TerraScan-collector/config"
// 	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
// 	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
// 	"github.com/JeromeTGH/TerraScan-collector/internal/mailsender"
// )

// func LoadStakingPercentage() (float64) {

// 	// config.AppConfig.Lcd.NbOfAttemptsIfFailure :   				par défaut = 5 		(donc on essaye cinq fois maximum, puis on arrête)
// 	// config.AppConfig.Lcd.NbMinutesOfBreakBetweenAttempts :		par défaut = 5 		(donc 5 minutes de pause, entre chaque tentative infructueuse)

// 	var idxTentatitves int

// 	for idxTentatitves = 1 ; idxTentatitves <= config.AppConfig.Lcd.NbOfAttemptsIfFailure ; idxTentatitves++ {
// 		stakingPercentage, errGetStakingPercentage := lcd.GetStakingPercentage()
// 		if errGetStakingPercentage == "" {
// 			if idxTentatitves > 1 {
// 				stringToReturn4 := fmt.Sprintf("LoadStakingPercentage : success of attempt %d/%d", idxTentatitves, config.AppConfig.Lcd.NbOfAttemptsIfFailure)
// 				logger.WriteLog("dataloader", stringToReturn4)
// 				mailsender.Sendmail("[TerraScan-collector] " + strconv.Itoa(idxTentatitves) + " attempts to get data from LCD successfully", "<p><strong>LoadStakingPercentage</strong></p><p>Script did " + strconv.Itoa(idxTentatitves) + " attempts to get data from LCD successfully</p>")
// 			}
// 			return stakingPercentage
// 		} else {
// 			stringToReturn1 := fmt.Sprintf("LoadStakingPercentage : failed attempt %d/%d", idxTentatitves, config.AppConfig.Lcd.NbOfAttemptsIfFailure)
// 			logger.WriteLog("dataloader", stringToReturn1)
// 			stringToReturn2 := fmt.Sprintf("LoadStakingPercentage : %s", errGetStakingPercentage)
// 			logger.WriteLog("dataloader", stringToReturn2)
// 			// Pause de X minutes, avant de retenter, s'il reste des tentatives à faire
// 			if idxTentatitves != config.AppConfig.Lcd.NbOfAttemptsIfFailure {
// 				time.Sleep(time.Duration(config.AppConfig.Lcd.NbMinutesOfBreakBetweenAttempts) * time.Second)
// 			}
// 		}
// 	}

// 	mailsender.Sendmail("[TerraScan-collector] impossible to load datas from LCD", "<p><strong>Impossible to load datas from LCD</strong></p><p>" + strconv.Itoa(config.AppConfig.Lcd.NbOfAttemptsIfFailure) + " attempts, and no success</p>")

// 	stringToReturn3 := fmt.Sprintf("LoadStakingPercentage : impossible to load datas from LCD, even after %d attempts", config.AppConfig.Lcd.NbOfAttemptsIfFailure)
// 	logger.WriteLog("dataloader", stringToReturn3)

// 	return -1
// }