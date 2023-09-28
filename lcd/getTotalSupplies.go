package lcd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/application"
)

func GetTotalSupplies(appConfig *application.Config) {

	type totalSuppliesLcdStructure struct {
		Supply [] struct {
			Denom string      `json:"denom"`
			Amount string     `json:"amount"`
		}
		Pagination struct {
			Next_key string     `json:"next_key"`
			Total string     	`json:"total"`
		}
	}

	// Path, pour accéder à ce qui nous intéresse
	var path = "/cosmos/bank/v1beta1/supply?pagination.limit=9999"

	// Récupération de l'url du LCD
	var LCDurl = appConfig.Lcd.Url

	// Création d'un client HTTP (avec timeout fixé à 30 secondes)
	client := &http.Client{
        Timeout: 30 * time.Second,
    }

	// Lancement du GET
	reponse, errGet := client.Get(LCDurl + path)
	if errGet != nil {
		panic(errGet)
	}

	// Lecture de la réponse du GET
	body, errReadAll := io.ReadAll(reponse.Body)
	if errReadAll != nil {
		panic(errReadAll)
	}

	// Transformation byte[] -> string pour avoir du contenu JSON "en clair"
	reponseJSON := string(body)

	// Stockage des données dans une struct
	dataStruct := totalSuppliesLcdStructure{}
	json.Unmarshal([]byte(reponseJSON), &dataStruct)

	// Récupération des total supplies du LUNC (uluna) et de l'USTC (uusd)
	var LUNCtotalSupply float64 = -1
	var USTCtotalSupply float64 = -1
	for i:=0 ; i<len(dataStruct.Supply) ; i++ {
		if dataStruct.Supply[i].Denom == "uluna" {
			uluna, errUluna := strconv.ParseFloat(dataStruct.Supply[i].Amount, 64)
			if errUluna != nil {
				panic(errUluna)
			}
			LUNCtotalSupply = uluna / 1000000
		}
		if dataStruct.Supply[i].Denom == "uusd" {
			uusd, errUusd := strconv.ParseFloat(dataStruct.Supply[i].Amount, 64)
			if errUusd != nil {
				panic(errUusd)
			}
			USTCtotalSupply = uusd / 1000000
		}
	}


	// Conversion de ces valeurs en string, et renvoi à l'appeleur








	// Afichage dans la console, pour commencer
	fmt.Printf("LUNCtotalSupply = %f\n", LUNCtotalSupply)
	fmt.Printf("USTCtotalSupply = %f\n", USTCtotalSupply)

}