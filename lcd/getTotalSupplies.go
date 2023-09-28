package lcd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/application"
)


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


type structReponseTotalSupplies struct {
	luncTotalSupply float64
	ustcTotalSupply float64
}


func GetTotalSupplies(appConfig *application.Config) (structReponseTotalSupplies, error) {

	// Initialisation de la struct qui sera renvoyée en retour
	var reponseEnRetour structReponseTotalSupplies

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
		return reponseEnRetour, errors.New("failed go fetch 'total supplies' on LCD")
	}

	// Lecture de la réponse du GET
	body, errReadAll := io.ReadAll(reponse.Body)
	if errReadAll != nil {
		return reponseEnRetour, errors.New("failed go read 'total supplies' answer, from LCD")
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
				return reponseEnRetour, errors.New("failed go convert 'uluna' amount in 'lunc'")
			}
			LUNCtotalSupply = uluna / 1000000
		}
		if dataStruct.Supply[i].Denom == "uusd" {
			uusd, errUusd := strconv.ParseFloat(dataStruct.Supply[i].Amount, 64)
			if errUusd != nil {
				return reponseEnRetour, errors.New("failed go convert 'uusd' amount in 'ustc'")
			}
			USTCtotalSupply = uusd / 1000000
		}
	}

	// Afichage dans la console (debug)
	fmt.Printf("LUNCtotalSupply = %f\n", LUNCtotalSupply)
	fmt.Printf("USTCtotalSupply = %f\n", USTCtotalSupply)

	// Et renvoi à l'appeleur
	reponseEnRetour.luncTotalSupply = LUNCtotalSupply
	reponseEnRetour.ustcTotalSupply = USTCtotalSupply
	return reponseEnRetour, nil

}