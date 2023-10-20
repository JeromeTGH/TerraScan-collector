package dbLuncStaking

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbActions"
	_ "github.com/go-sql-driver/mysql"
)

func DropNbStakedLuncTable(channelForErrors chan<- string) error {

	// Construction de la requête
	rqt := "DROP TABLE IF EXISTS " + config.AppConfig.Bdd.TblLuncStaking

	// Exécution de la requête
	errExec := dbActions.ExecCreateOrDrop(rqt)	
	if errExec != nil {
		stringToReturn := fmt.Sprintf("[dboperations] DropNbStakedLuncTable : failed (%s)", errExec.Error())
		channelForErrors <- stringToReturn
		return errExec
	}

	return nil

}