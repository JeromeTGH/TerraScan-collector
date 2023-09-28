package dataloader

import (
	"errors"
	"fmt"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/application"
	"github.com/JeromeTGH/TerraScan-collector/lcd"
)


func LoadTotalSupplies(appConfig *application.Config) (lcd.StructReponseTotalSupplies, error) {

	const nbMaxTentatives = 5						// Par défaut, on essaye 5 fois, puis on arrête
	const nbMinutesDePauseEntreChaqueEssai = 5		// avec une pause de 5 minutes par défaut, entre chaque essai

	var idxTentatitves uint8

	for idxTentatitves = 1 ; idxTentatitves <= nbMaxTentatives ; idxTentatitves++ {
		totalSupplies, errGetTotalSupplies := lcd.GetTotalSupplies(appConfig)
		if errGetTotalSupplies == nil {
			return totalSupplies, nil
		} else {
			fmt.Println("[LoadTotalSupplies] Échec tentative", idxTentatitves, "/", nbMaxTentatives)
			fmt.Println(errGetTotalSupplies.Error())
			// Pause de X minutes, avant de retenter
			time.Sleep(nbMinutesDePauseEntreChaqueEssai * time.Minute)
		}
	}

	return lcd.StructReponseTotalSupplies{}, errors.New("impossible to load datas from LCD")
}