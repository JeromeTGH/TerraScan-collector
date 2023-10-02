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
	errDropTable := db.Exec(rqt)	
	if errDropTable != nil {
		fmt.Println(errDropTable)
	}

	fmt.Println("Table effacée")


}