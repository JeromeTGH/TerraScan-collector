package dbwriter

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/lcd"
	_ "github.com/go-sql-driver/mysql"
)

// func connectDb() (sql.DB, error) {
// }

func WriteTotalSuppliesInDb(dataFromLcd lcd.StructReponseTotalSupplies) {

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
	rqt += "?)"		// uusdAmount


	// Exécution de la requête
	insertResult, err2 := db.ExecContext(context.Background(), rqt, "code", "datetimeUTC", true, false, false, false, false, false, 12, 13)
	if err2 != nil {
		fmt.Printf("insert failed : %s", err2)
		return
	}

	// Récupération du dernier ID inséré
	id, err3 := insertResult.LastInsertId()
	if err3 != nil {
		fmt.Printf("impossible to retrieve last inserted id: %s", err3)
		return
	}
	fmt.Printf("inserted id: %d", id)

}