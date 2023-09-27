package main

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/application"
)

// Variables globales
var appConfig application.Config

func main() {

	// TEST : Récupération de la config
	application.LoadConfig(&appConfig)
	fmt.Println(appConfig.Lcd.Url)

}