package dbNbStakedLunc

import (
	"fmt"
	"strings"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbActions"
	"github.com/JeromeTGH/TerraScan-collector/internal/mailsender"

	_ "github.com/go-sql-driver/mysql"
)

func WriteNbStakedLuncInDb(dataFromLcd lcd.StructReponseNbStakedLunc, stakingPercentage float64, channelForErrors chan<- string) {

	// Génération des valeurs à enregistrer
	nbStakedLunc := dataFromLcd.NbStakedLunc
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
	rqt := "INSERT INTO " + config.AppConfig.Bdd.TblNbStakedLunc + " VALUES (null, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	// Exécution de la requête
	var bCreateTableNeeded = false
	lastInsertId, errInsert := dbActions.ExecInsert(rqt, code, dateUTCpourMysql, bH1, bH4, bD1, bW1, bM1, bY1, nbStakedLunc, stakingPercentage)
	if errInsert != nil {
		stringToReturn := fmt.Sprintf("[dboperations] WriteNbStakedLuncInDb : failed (%s)", errInsert.Error())
		channelForErrors <- stringToReturn
		if strings.Contains(errInsert.Error(), "doesn't exist") {
			// Si c'est une erreur du type "table inexistante", alors on va essayer de la créer, et de refaire l'insertion
			bCreateTableNeeded = true
		} else {
			// Autre erreur, on quitte cette fonction
			mailsender.Sendmail("[TerraScan-collector] failed to insert data in DB", "<p><strong>Failed to insert data in DB, on first try</strong></p><p>Error : " +  errInsert.Error() + "</p>", channelForErrors)
			return
		}
	}

	// Check s'il y a eu une erreur, de type "table inexistante"
	if bCreateTableNeeded {
		// Création de la table, car inexistante
		errCreation := CreateNbStakedLuncTable(channelForErrors)
		if errCreation != nil {
			stringToReturn2 := fmt.Sprintf("[dboperations] WriteNbStakedLuncInDb : failed (%s)", errCreation.Error())
			channelForErrors <- stringToReturn2
			mailsender.Sendmail("[TerraScan-collector] failed to create table in DB", "<p><strong>Failed to create table in DB</strong></p><p>Error : " +  errCreation.Error() + "</p>", channelForErrors)
			return
		}
		stringToReturn3 := "[dboperations] WriteNbStakedLuncInDb : new table created successfully"
		channelForErrors <- stringToReturn3

		// Et re-tentative d'insertion
		lastInsertId2, errInsert2 := dbActions.ExecInsert(rqt, code, dateUTCpourMysql, bH1, bH4, bD1, bW1, bM1, bY1, nbStakedLunc, stakingPercentage)

		if errInsert2 != nil {
			stringToReturn4 := fmt.Sprintf("[dboperations] WriteNbStakedLuncInDb : failed (%s)", errInsert2.Error())
			channelForErrors <- stringToReturn4
			mailsender.Sendmail("[TerraScan-collector] failed to insert data in DB", "<p><strong>Failed to insert data in DB, on second try</strong></p><p>Error : " +  errInsert2.Error() + "</p>", channelForErrors)
			return
		}

		stringToReturn5 := fmt.Sprintf("[dboperations] WriteNbStakedLuncInDb : insert completed successfully (lastInsertId = %d)", lastInsertId2)
		channelForErrors <- stringToReturn5
	} else {
		stringToReturn6 := fmt.Sprintf("[dboperations] WriteNbStakedLuncInDb : insert completed successfully (lastInsertId = %d)", lastInsertId)
		channelForErrors <- stringToReturn6
	}


}