package dbLuncStaking

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbActions"
	_ "github.com/go-sql-driver/mysql"
)

func CreateNbStakedLuncTable(channelForErrors chan<- string) error {

	// Construction de la requête
	rqt := "CREATE TABLE IF NOT EXISTS " + config.AppConfig.Bdd.TblLuncStaking + " ("
	rqt += "enregNumber INT AUTO_INCREMENT PRIMARY KEY,"
	rqt += "code VARCHAR(12) UNIQUE,"
	rqt += "datetimeUTC DATETIME,"
	rqt += "bH1 BOOLEAN NOT NULL DEFAULT TRUE,"
	rqt += "bH4 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bD1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bW1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bM1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "bY1 BOOLEAN NOT NULL DEFAULT FALSE,"
	rqt += "nbStakedLunc BIGINT,"
	rqt += "stakingPercentage FLOAT"
	rqt += ");"

	// Exécution de la requête
	errExec := dbActions.ExecCreateOrDrop(rqt)
	if errExec != nil {
		stringToReturn := fmt.Sprintf("[dboperations] CreateNbStakedLuncTable : failed (%s)", errExec.Error())
		channelForErrors <- stringToReturn
		return errExec
	}

	return nil

}