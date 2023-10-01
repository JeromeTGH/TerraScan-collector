package dbwriter

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/lcd"
)

func WriteTotalSuppliesInDb(dataFromLcd lcd.StructReponseTotalSupplies) {

	// Afichage dans la console de ces données (debug)
	fmt.Printf("LUNCtotalSupply = %d\n", dataFromLcd.LuncTotalSupply)
	fmt.Printf("USTCtotalSupply = %d\n", dataFromLcd.UstcTotalSupply)

}