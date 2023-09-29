package main

import (
	"fmt"
	"os"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/utils/dataloader"
)

// Variables globales
var appConfig config.Config

func main() {

	// Chargement des données de configuration, dans la variable "appConfig"
	config.LoadConfig(&appConfig)

	// Chargement des données, en faisant appel au LCD (on passe la config dans la foulée, pour transmettre l'URL du LCD, stocké dedans)
	dataFromLcd, errFromLcd := dataloader.LoadTotalSupplies(&appConfig)
	if errFromLcd != nil {
		fmt.Println(errFromLcd)
		os.Exit(500)
	}

	// Afichage dans la console (debug)
	fmt.Printf("LUNCtotalSupply = %f\n", dataFromLcd.LuncTotalSupply)
	fmt.Printf("USTCtotalSupply = %f\n", dataFromLcd.UstcTotalSupply)

}