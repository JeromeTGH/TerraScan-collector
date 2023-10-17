package dbTotalSupplies

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/db"
	_ "github.com/go-sql-driver/mysql"
)

func CreateTotalSuppliesTable(channelForErrors chan<- string) error {

	// Construction de la requête
	rqt := "CREATE TABLE IF NOT EXISTS tblTotalSupplies3 ("
	rqt += "enregNumber INT AUTO_INCREMENT PRIMARY KEY,"
	rqt += "code VARCHAR(12) UNIQUE,"
	rqt += "datetimeUTC DATETIME,"
	rqt += "bH1 BOOLEAN NOT NULL DEFAULT TRUE,"
	rqt += "bH4 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bD1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bW1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bM1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bY1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "luncAmount BIGINT,"
	rqt += "ustcAmount BIGINT"
	rqt += ");"

	// Exécution de la requête
	errExec := db.ExecCreateOrDrop(rqt)	
	if errExec != nil {
		stringToReturn := fmt.Sprintf("[dboperations] CreateTotalSuppliesTable : failed (%s)", errExec.Error())
		channelForErrors <- stringToReturn
		return errExec
	}

	return nil

}