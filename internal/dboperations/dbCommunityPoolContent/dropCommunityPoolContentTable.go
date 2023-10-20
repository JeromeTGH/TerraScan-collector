package dbCommunityPoolContent

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbActions"
	_ "github.com/go-sql-driver/mysql"
)

func DropCommunityPoolContentTable(channelForErrors chan<- string) error {

	// Construction de la requête
	rqt := "DROP TABLE IF EXISTS " + config.AppConfig.Bdd.TblCommunityPoolContent

	// Exécution de la requête
	errExec := dbActions.ExecCreateOrDrop(rqt)	
	if errExec != nil {
		stringToReturn := fmt.Sprintf("[dboperations] DropCommunityPoolContentTable : failed (%s)", errExec.Error())
		channelForErrors <- stringToReturn
		return errExec
	}

	return nil

}