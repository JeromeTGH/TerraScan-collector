package dbwriter

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/lcd"
	_ "github.com/go-sql-driver/mysql"
)

// func connectDb() (sql.DB, error) {
// }

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



	// Afichage dans la console de ces données (debug)
	fmt.Printf("LUNCtotalSupply = %d\n", dataFromLcd.LuncTotalSupply)
	fmt.Printf("USTCtotalSupply = %d\n", dataFromLcd.UstcTotalSupply)

	// Connexion à la base de données
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.AppConfig.Bdd.UserName, config.AppConfig.Bdd.Password, config.AppConfig.Bdd.HostName, config.AppConfig.Bdd.Port, config.AppConfig.Bdd.DbName)
	db, err := sql.Open("mysql", dsn)
    if err != nil {
		fmt.Printf("impossible to create the connection: %s", err)
		return
    }
	defer db.Close()

	// Paramètres de connexion
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)


	// // Effacement de la table
	// rqtZero := "DROP TABLE IF EXISTS tblTotalSupplies2"
	// _, errZero := db.ExecContext(context.Background(), rqtZero)
	// if errZero != nil {  
	// 	fmt.Printf("Error %s when dropping table", errZero)
	// 	return
	// }

	// Construction de la requête
	rqt := "INSERT INTO tblTotalSupplies2 VALUES ("
	rqt += "null,"
	rqt += "?,"		// code
	rqt += "?,"		// datetimeUTC
	rqt += "?,"		// bH1
	rqt += "?,"		// bH4
	rqt += "?,"		// bD1
	rqt += "?,"		// bW1
	rqt += "?,"		// bM1
	rqt += "?,"		// bY1
	rqt += "?,"		// luncAmount
	rqt += "?)"		// ustcAmount


	// Exécution de la requête
	var bCreateTableNeeded = false
	insertResult, err2 := db.ExecContext(context.Background(), rqt, code, dateUTCpourMysql, bH1, bH4, bD1, bW1, bM1, bY1, dataFromLcd.LuncTotalSupply, dataFromLcd.UstcTotalSupply)
	if err2 != nil {
		if strings.Contains(err2.Error(), "doesn't exist") {
			bCreateTableNeeded = true
		} else {
			fmt.Printf("insert failed : %s", err2)
			return
		}
	}

	// Si la table visée n'existe pas, on l'a créé
	if bCreateTableNeeded {

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

		_, err2b := db.ExecContext(context.Background(), rqt2)  
		if err2b != nil {  
			fmt.Printf("Error %s when creating table", err2b)
			return
		}

		insertResult2, err2c := db.ExecContext(context.Background(), rqt, code, dateUTCpourMysql, bH1, bH4, bD1, bW1, bM1, bY1, dataFromLcd.LuncTotalSupply, dataFromLcd.UstcTotalSupply)
		if err2c != nil {
			fmt.Printf("insert failed : %s", err2c)
			return
		}

		// Récupération du dernier ID inséré
		id, err2d := insertResult2.LastInsertId()
		if err2d != nil {
			fmt.Printf("impossible to retrieve last inserted id: %s", err2d)
			return
		}
		fmt.Printf("inserted id: %d", id)


	} else {

		// Récupération du dernier ID inséré
		id, err3 := insertResult.LastInsertId()
		if err3 != nil {
			fmt.Printf("impossible to retrieve last inserted id: %s", err3)
			return
		}
		fmt.Printf("inserted id: %d", id)

	}



}