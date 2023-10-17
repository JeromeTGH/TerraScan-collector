package dbTotalSupplies

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/internal/dboperations/db"
	"github.com/JeromeTGH/TerraScan-collector/internal/logger"
	_ "github.com/go-sql-driver/mysql"
)

func DropTotalSuppliesTable() error {

	// Construction de la requête
	rqt := "DROP TABLE IF EXISTS tblTotalSupplies3"

	// Exécution de la requête
	errExec := db.ExecCreateOrDrop(rqt)	
	if errExec != nil {
		stringToReturn := fmt.Sprintf("DropTotalSuppliesTable : failed (%s)", errExec.Error())
		logger.WriteLog("dboperations", stringToReturn)
		return errExec
	}

	return nil

}