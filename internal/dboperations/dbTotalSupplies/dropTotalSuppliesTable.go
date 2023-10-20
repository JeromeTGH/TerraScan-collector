package dbTotalSupplies

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/dbActions"
	_ "github.com/go-sql-driver/mysql"
)

func DropTotalSuppliesTable(channelForErrors chan<- string) error {

	// Construction de la requête
	rqt := "DROP TABLE IF EXISTS tblTotalSupplies3"

	// Exécution de la requête
	errExec := dbActions.ExecCreateOrDrop(rqt)	
	if errExec != nil {
		stringToReturn := fmt.Sprintf("[dboperations] DropTotalSuppliesTable : failed (%s)", errExec.Error())
		channelForErrors <- stringToReturn
		return errExec
	}

	return nil

}