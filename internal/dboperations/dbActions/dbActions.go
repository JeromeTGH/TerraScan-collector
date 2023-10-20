package dbActions

import (
	"database/sql"
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	_ "github.com/go-sql-driver/mysql"
)


func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.AppConfig.Bdd.UserName,
		config.AppConfig.Bdd.Password,
		config.AppConfig.Bdd.HostName,
		config.AppConfig.Bdd.Port,
		config.AppConfig.Bdd.DbName)
}


func ExecCreateOrDrop (rqt string) error {

	db, errOpen := sql.Open("mysql", dsn())
	if errOpen != nil {
		// fmt.Printf("failed to create connection with mysql server : %s\n", errOpen)
		return errOpen
    }
	defer db.Close()

	_, errExec := db.Exec(rqt)

	if errExec != nil {
		// fmt.Printf("failed to execute rqt : %s\n", errExec)
		return errExec
	}

	return nil
}



func ExecInsert (rqt string, args ...interface{}) (int, error) {

	db, errOpen := sql.Open("mysql", dsn())
	if errOpen != nil {
		// fmt.Printf("failed to create connection with mysql server : %s\n", errOpen)
		return -1, errOpen
    }
	defer db.Close()

	insert, errPrepare := db.Prepare(rqt)
	if errPrepare != nil {
		// fmt.Printf("failed to prepare rqt : %s\n", errPrepare)
		return -1, errPrepare
	}

	resp, errExec := insert.Exec(args...)
	insert.Close()

	if errExec != nil {
		// fmt.Printf("failed to execute rqt : %s\n", errExec)
		return -1, errExec
	}

	lastInsertId, errLastInsertId := resp.LastInsertId()
	if errLastInsertId != nil {
		// fmt.Printf("failed to fetch LastInsertId : %s\n", errLastInsertId)
		return -1, errLastInsertId
	}

	// fmt.Printf("Values added ! LastInsertId = %d\n", lastInsertId)
	return int(lastInsertId), nil
}
