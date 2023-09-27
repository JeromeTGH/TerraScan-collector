package lcd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/application"
)

func GetTotalSupplies(appConfig *application.Config) {

	// Structure des données, qui seront récupérées auprès du LCD (au format JSON, nativement)
	type supplies struct {
		Denom string      `json:"denom"`
		Amount string     `json:"amount"`
	}

	type totalSupplies struct {
		Supply []supplies      `json:"supply"`
		Pagination struct {
			Next_key string      `json:"next_key"`
			Total string     `json:"total"`
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

	// Lecture de la réponse
	body, errReadAll := io.ReadAll(reponse.Body)
	if errReadAll != nil {
		panic(errReadAll)
	}

	// Transformation byte[] -> string avec du contenu JSON
	reponseJSON := string(body)

	// Stockage des données dans une struct
	dataStruct := totalSupplies{}
	json.Unmarshal([]byte(reponseJSON), &dataStruct)



	// Afichage dans la console, pour commencer
	fmt.Println(dataStruct.Supply[1].Denom)
	fmt.Println(dataStruct.Supply[1].Amount)

}