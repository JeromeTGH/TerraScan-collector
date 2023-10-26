package dbOraclePoolContent

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbActions"
	_ "github.com/go-sql-driver/mysql"
)

func DropOraclePoolContentTable(channelForLogsMsgs chan<- string) error {

	// Construction de la requête
	rqt := "DROP TABLE IF EXISTS " + config.AppConfig.Bdd.TblOraclePoolContent

	// Exécution de la requête
	errExec := dbActions.ExecCreateOrDrop(rqt)	
	if errExec != nil {
		stringToReturn := fmt.Sprintf("[dboperations] DropOraclePoolContentTable : failed (%s)", errExec.Error())
		channelForLogsMsgs <- stringToReturn
		return errExec
	}

	return nil

}