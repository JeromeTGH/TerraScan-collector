package db

import (
	"database/sql"
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	_ "github.com/go-sql-driver/mysql"
)

type InsertParams struct {
	Code string
	DateUTCpourMysql string
	BH1 bool
	BH4 bool
	BD1 bool
	BW1 bool
	BM1 bool
	BY1 bool
	LuncTotalSupply int
	UstcTotalSupply int
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.AppConfig.Bdd.UserName,
		config.AppConfig.Bdd.Password,
		config.AppConfig.Bdd.HostName,
		config.AppConfig.Bdd.Port,
		config.AppConfig.Bdd.DbName)
}

// // Effacement de la table
// rqtZero := "DROP TABLE IF EXISTS tblTotalSupplies2"
// _, errZero := db.ExecContext(context.Background(), rqtZero)
// if errZero != nil {  
// 	fmt.Printf("Error %s when dropping table", errZero)
// 	return
// }

func InsertIntoDb (rqt string, params InsertParams) {

	db, errOpen := sql.Open("mysql", dsn())
	if errOpen != nil {
		fmt.Printf("failed to create connection with mysql server : %s", errOpen)
		return
    }
	defer db.Close()

	insert, errPrepare := db.Prepare(rqt)
	if errPrepare != nil {
		fmt.Printf("failed to prepare rqt : %s", errPrepare)
		return
	}

	resp, errExec := insert.Exec(params.Code, params.DateUTCpourMysql, params.BH1, params.BH4, params.BD1, params.BW1, params.BM1, params.BY1, params.LuncTotalSupply, params.UstcTotalSupply)
	insert.Close()

	if errExec != nil {
		fmt.Printf("failed to execute rqt : %s", errExec)
		return
	}

	lastInsertId, errLastInsertId := resp.LastInsertId()
	if errLastInsertId != nil {
		fmt.Printf("failed to fetch LastInsertId : %s", errLastInsertId)
		return
	}

	fmt.Printf("Values added ! LastInsertId = %d", lastInsertId)
}

