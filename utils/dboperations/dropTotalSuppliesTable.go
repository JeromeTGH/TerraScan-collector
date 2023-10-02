package dboperations

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/utils/dboperations/db"
	_ "github.com/go-sql-driver/mysql"
)

func DropTotalSuppliesTable() {

	// Construction de la requête
	rqt := "DROP TABLE IF EXISTS tblTotalSupplies2"

	// Exécution de la requête
	errExec := db.Exec(rqt)	
	if errExec != nil {
		fmt.Println(errExec)
	}

	// fmt.Println("Table effacée")

}