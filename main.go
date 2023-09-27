package main

import (
	"github.com/JeromeTGH/TerraScan-collector/application"
	"github.com/JeromeTGH/TerraScan-collector/lcd"
)

// Variables globales
var appConfig application.Config

func main() {

	// Chargement des données de configuration, dans la variable "appConfig"
	application.LoadConfig(&appConfig)

	// Chargement des données, en faisant appel au LCD (on passe la config dans la foulée, pour transmettre l'URL du LCD, stocké dedans)
	lcd.GetTotalSupplies(&appConfig)

}