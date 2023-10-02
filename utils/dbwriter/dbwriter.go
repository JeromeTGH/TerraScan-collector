package dbwriter

import (
	"fmt"
	"strings"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/utils/dataloader/lcd"
	"github.com/JeromeTGH/TerraScan-collector/utils/dbwriter/db"
	_ "github.com/go-sql-driver/mysql"
)

func WriteTotalSuppliesInDb(dataFromLcd lcd.StructReponseTotalSupplies) {

	// Afichage dans la console de ces données (debug)
	fmt.Printf("LUNCtotalSupply = %d\n", dataFromLcd.LuncTotalSupply)
	fmt.Printf("USTCtotalSupply = %d\n", dataFromLcd.UstcTotalSupply)
	

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
	_, errInsert := db.InsertIntoDb(rqt, code, dateUTCpourMysql, bH1, bH4, bD1, bW1, bM1, bY1, dataFromLcd.LuncTotalSupply, dataFromLcd.UstcTotalSupply)	
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "doesn't exist") {
			bCreateTableNeeded = true
		} else {
			// Autre erreur, on quitte cette fonction
			return
		}
	}

	// Check s'il y a eu une erreur, de type "table inexistante"
	if bCreateTableNeeded {
		fmt.Println("Création de table nécessaire")

		// Construction de la requête
		rqt2 := "CREATE TABLE IF NOT EXISTS tblTotalSupplies2 ("
		rqt2 += "enregNumber INT AUTO_INCREMENT PRIMARY KEY,"
		rqt2 += "code VARCHAR(12) UNIQUE,"
		rqt2 += "datetimeUTC DATETIME,"
		rqt2 += "bH1 BOOLEAN NOT NULL DEFAULT TRUE,"
		rqt2 += "bH4 BOOLEAN NOT NULL DEFAULT FALSE,"
		rqt2 += "bD1 BOOLEAN NOT NULL DEFAULT FALSE,"
		rqt2 += "bW1 BOOLEAN NOT NULL DEFAULT FALSE,"
		rqt2 += "bM1 BOOLEAN NOT NULL DEFAULT FALSE,"
		rqt2 += "bY1 BOOLEAN NOT NULL DEFAULT FALSE,"
		rqt2 += "luncAmount BIGINT,"
		rqt2 += "ustcAmount BIGINT"
		rqt2 += ");"

		id, errCreation := db.CreateTableInDb(rqt2)
		if errCreation != nil {
			fmt.Println(errCreation)
		} else {
			fmt.Println("Id =", id)
		}
	} else {
		fmt.Println("Création de table inutile")
	}

	// // Si la table visée n'existe pas, on l'a créé
	// if bCreateTableNeeded {

	// 	// Construction de la requête
	// 	rqt2 := "CREATE TABLE IF NOT EXISTS tblTotalSupplies2 ("
	// 	rqt2 += "enregNumber INT AUTO_INCREMENT PRIMARY KEY,"
	// 	rqt2 += "code VARCHAR(12) UNIQUE,"
	// 	rqt2 += "datetimeUTC DATETIME,"
	// 	rqt2 += "bH1 BOOLEAN NOT NULL DEFAULT TRUE,"
	// 	rqt2 += "bH4 BOOLEAN NOT NULL DEFAULT FALSE,"
	// 	rqt2 += "bD1 BOOLEAN NOT NULL DEFAULT FALSE,"
	// 	rqt2 += "bW1 BOOLEAN NOT NULL DEFAULT FALSE,"
	// 	rqt2 += "bM1 BOOLEAN NOT NULL DEFAULT FALSE,"
	// 	rqt2 += "bY1 BOOLEAN NOT NULL DEFAULT FALSE,"
	// 	rqt2 += "luncAmount BIGINT,"
	// 	rqt2 += "ustcAmount BIGINT"
	// 	rqt2 += ");"

	// 	_, err2b := db.InsertIntoDb(rqt2)  
	// 	if err2b != nil {  
	// 		fmt.Printf("Error %s when creating table", err2b)
	// 		return
	// 	}

	// 	insertResult2, err2c := db.InsertIntoDb(rqt)
	// 	if err2c != nil {
	// 		fmt.Printf("insert failed : %s", err2c)
	// 		return
	// 	}

	// 	// // Récupération du dernier ID inséré
	// 	// id, err2d := insertResult2.LastInsertId()
	// 	// if err2d != nil {
	// 	// 	fmt.Printf("impossible to retrieve last inserted id: %s", err2d)
	// 	// 	return
	// 	// }
	// 	// fmt.Printf("inserted id: %d", id)
	// 	fmt.Println(insertResult2)


	// } else {

	// 	// Récupération du dernier ID inséré
	// 	// id, err3 := insertResult.LastInsertId()
	// 	// if err3 != nil {
	// 	// 	fmt.Printf("impossible to retrieve last inserted id: %s", err3)
	// 	// 	return
	// 	// }
	// 	// fmt.Printf("inserted id: %d", id)
	// 	fmt.Println(insertResult)

	// }



}