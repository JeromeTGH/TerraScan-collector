package dboperations

import (
	"fmt"
	"strings"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/db"
	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
	"github.com/JeromeTGH/TerraScan-collector/internal/mailsender"

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
	rqt := "INSERT INTO tblTotalSupplies3 VALUES (null, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	// Exécution de la requête
	var bCreateTableNeeded = false
	lastInsertId, errInsert := db.ExecInsert(rqt, code, dateUTCpourMysql, bH1, bH4, bD1, bW1, bM1, bY1, dataFromLcd.LuncTotalSupply, dataFromLcd.UstcTotalSupply)	
	if errInsert != nil {
		stringToReturn := fmt.Sprintf("WriteTotalSuppliesInDb : failed (%s)", errInsert.Error())
		logger.WriteLog("dboperations", stringToReturn)	
		if strings.Contains(errInsert.Error(), "doesn't exist") {
			// Si c'est une erreur du type "table inexistante", alors on va essayer de la créer, et de refaire l'insertion
			bCreateTableNeeded = true
		} else {
			// Autre erreur, on quitte cette fonction
			mailsender.Sendmail("[TerraScan-collector] failed to insert data in DB", "<p><strong>Failed to insert data in DB, on first try</strong></p><p>Error : " +  errInsert.Error() + "</p>")
			return
		}
	}

	// Check s'il y a eu une erreur, de type "table inexistante"
	if bCreateTableNeeded {
		// Création de la table, car inexistante
		errCreation := CreateTotalSuppliesTable()
		if errCreation != nil {
			stringToReturn2 := fmt.Sprintf("WriteTotalSuppliesInDb : failed (%s)", errCreation.Error())
			logger.WriteLog("dboperations", stringToReturn2)
			mailsender.Sendmail("[TerraScan-collector] failed to create table in DB", "<p><strong>Failed to create table in DB</strong></p><p>Error : " +  errCreation.Error() + "</p>")
			return
		}
		stringToReturn3 := "WriteTotalSuppliesInDb : new table created successfully"
		logger.WriteLog("dboperations", stringToReturn3)

		// Et re-tentative d'insertion
		lastInsertId2, errInsert2 := db.ExecInsert(rqt, code, dateUTCpourMysql, bH1, bH4, bD1, bW1, bM1, bY1, dataFromLcd.LuncTotalSupply, dataFromLcd.UstcTotalSupply)	

		if errInsert2 != nil {
			stringToReturn4 := fmt.Sprintf("WriteTotalSuppliesInDb : failed (%s)", errInsert2.Error())
			logger.WriteLog("dboperations", stringToReturn4)
			mailsender.Sendmail("[TerraScan-collector] failed to insert data in DB", "<p><strong>Failed to insert data in DB, on second try</strong></p><p>Error : " +  errInsert2.Error() + "</p>")
			return
		}

		stringToReturn5 := fmt.Sprintf("WriteTotalSuppliesInDb : insert completed successfully (lastInsertId = %d)", lastInsertId2)
		logger.WriteLog("dboperations", stringToReturn5)
	} else {
		stringToReturn6 := fmt.Sprintf("WriteTotalSuppliesInDb : insert completed successfully (lastInsertId = %d)", lastInsertId)
		logger.WriteLog("dboperations", stringToReturn6)
	}


}