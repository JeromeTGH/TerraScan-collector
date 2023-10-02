package dboperations

import (
	"fmt"
	"strings"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/utils/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/utils/dboperations/db"
	_ "github.com/go-sql-driver/mysql"
)

func WriteTotalSuppliesInDb(dataFromLcd lcd.StructReponseTotalSupplies) {

	// Génération des valeurs à enregistrer
	currentTime := time.Now().UTC()
		nAnnee := currentTime.Year()
		nMois := currentTime.Month()
		nJours := currentTime.Day()
		nHeures := currentTime.Hour()
		nMinutes := currentTime.Minute()
		nSecondes := currentTime.Second()
		dayOfTheWeek := currentTime.Weekday()		// Dimanche = 0
	code := fmt.Sprintf("%d%02d%02d%02d%02d", nAnnee, nMois, nJours, nHeures, nMinutes)
	dateUTCpourMysql := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ", nAnnee, nMois, nJours, nHeures, nMinutes, nSecondes)
	bH1 := true;
	bH4 := (nHeures == 0 || nHeures == 4 || nHeures == 8 || nHeures == 12 || nHeures == 16 || nHeures == 20);
	bD1 := nHeures == 0;
	bW1 := (dayOfTheWeek == 1 && nHeures == 0);       // Lundi à Oh...
	bM1 := (nJours == 1 && nHeures == 0);
	bY1 := (nMois == 1 && nJours == 1 && nHeures == 0);


	// Construction de la requête
	rqt := "INSERT INTO tblTotalSupplies2 VALUES (null, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	// Exécution de la requête
	var bCreateTableNeeded = false
	_, errInsert := db.ExecInsert(rqt, code, dateUTCpourMysql, bH1, bH4, bD1, bW1, bM1, bY1, dataFromLcd.LuncTotalSupply, dataFromLcd.UstcTotalSupply)	
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "doesn't exist") {
			// Si c'est une erreur du type "table inexistante", alors on va essayer de la créer, et de refaire l'insertion
			bCreateTableNeeded = true
		} else {
			// Autre erreur, on quitte cette fonction
			return
		}
	}

	// Check s'il y a eu une erreur, de type "table inexistante"
	if bCreateTableNeeded {
		// Création de la table, car inexistante
		errCreation := CreateTotalSuppliesTable()
		if errCreation != nil {
			fmt.Println(errCreation)
			return
		}

		// Et re-tentative d'insertion
		_, errInsert2 := db.ExecInsert(rqt, code, dateUTCpourMysql, bH1, bH4, bD1, bW1, bM1, bY1, dataFromLcd.LuncTotalSupply, dataFromLcd.UstcTotalSupply)	

		if errInsert2 != nil {
			fmt.Println(errInsert2)
			return
		}
	}


}