package dbOraclePoolContent

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbActions"
	_ "github.com/go-sql-driver/mysql"
)

func CreateOraclePoolContentTable(channelForLogsMsgs chan<- string) error {

	// Construction de la requête
	rqt := "CREATE TABLE IF NOT EXISTS " + config.AppConfig.Bdd.TblOraclePoolContent + " ("
	rqt += "enregNumber INT AUTO_INCREMENT PRIMARY KEY,"
	rqt += "code VARCHAR(12) UNIQUE,"
	rqt += "datetimeUTC DATETIME,"
	rqt += "bH1 BOOLEAN NOT NULL DEFAULT TRUE,"
	rqt += "bH4 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bD1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bW1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bM1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bY1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "nbLuncInOP BIGINT,"
	rqt += "nbUstcInOP BIGINT"
	rqt += ");"

	// Exécution de la requête
	errExec := dbActions.ExecCreateOrDrop(rqt)
	if errExec != nil {
		stringToReturn := fmt.Sprintf("[dboperations] CreateOraclePoolContentTable : failed (%s)", errExec.Error())
		channelForLogsMsgs <- stringToReturn
		return errExec
	}

	return nil

}