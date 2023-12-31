package lcd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/JeromeTGH/TerraScan-collector/config"
)


type totalSuppliesLcdStructure struct {
	Supply [] struct {
		Denom 	string      `json:"denom"`
		Amount 	string     	`json:"amount"`
	}
	Pagination struct {
		Next_key string     `json:"next_key"`
		Total 	 string    	`json:"total"`
	}
	Error struct {
		Code 	int 		`json:"code"`
		Message string		`json:"message"`
	}
}

type StructReponseTotalSupplies struct {
	LuncTotalSupply int
	UstcTotalSupply int
}


func GetTotalSupplies(channelForLogsMsgs chan<- string) (StructReponseTotalSupplies, string) {

	// Initialisation de la struct qui sera renvoyée en retour
	var reponseEnRetour StructReponseTotalSupplies

	// Path, pour accéder à ce qui nous intéresse
	var path = config.AppConfig.Lcd.PathForTotalSupplies

	// Récupération de l'url du LCD
	var LCDurl = config.AppConfig.Lcd.Url

	// Création d'un client HTTP (avec timeout fixé à 30 secondes)
	client := &http.Client{
        Timeout: time.Duration(config.AppConfig.Lcd.GetTimeoutInSeconds) * time.Second,
    }

	// Lancement du GET
	reponse, errGet := client.Get(LCDurl + path)
	if errGet != nil {
		return reponseEnRetour, "failed to fetch 'total supplies' from LCD"
	}

	// Lecture de la réponse du GET
	body, errReadAll := io.ReadAll(reponse.Body)
	if errReadAll != nil {
		return reponseEnRetour, "failed to read 'total supplies' answer from LCD"
	}

	// Transformation byte[] -> string pour avoir du contenu JSON "en clair"
	reponseJSON := string(body)

	// Stockage des données dans une struct
	dataStruct := totalSuppliesLcdStructure{}
	json.Unmarshal([]byte(reponseJSON), &dataStruct)

	// Sortie si erreur retournée
	if dataStruct.Error.Message != "" {
		return reponseEnRetour, "failed to unmarshal 'total supplies' from LCD"
	}

	// Récupération des total supplies du LUNC (uluna) et de l'USTC (uusd)
	var LUNCtotalSupply int = -1
	var USTCtotalSupply int = -1
	for i:=0 ; i<len(dataStruct.Supply) ; i++ {
		if dataStruct.Supply[i].Denom == "uluna" {
			uluna, errUluna := strconv.ParseFloat(dataStruct.Supply[i].Amount, 64)
			if errUluna != nil {
				return reponseEnRetour, "failed to convert 'uluna' amount in 'lunc' from LCD"
			}
			LUNCtotalSupply = int(uluna / 1000000)
		}
		if dataStruct.Supply[i].Denom == "uusd" {
			uusd, errUusd := strconv.ParseFloat(dataStruct.Supply[i].Amount, 64)
			if errUusd != nil {
				return reponseEnRetour, "failed to convert 'uusd' amount in 'ustc' from LCD"
			}
			USTCtotalSupply = int(uusd / 1000000)
		}
	}

	// Si jamais les variables "LUNCtotalSupply" et "USTCtotalSupply" n'ont pas été impactées, alors on remonte une erreur
	if(LUNCtotalSupply == -1 || USTCtotalSupply == -1) {
		stringToReturn := fmt.Sprintf("[dataloader] GetTotalSupplies : -1 returned by function.\nError = %s", reponseJSON)
		channelForLogsMsgs <- stringToReturn
		return reponseEnRetour, "failed to get 'uusd' or 'uluna' amount from LCD (-1 returned)"
	}

	// Et renvoi à l'appeleur
	reponseEnRetour.LuncTotalSupply = LUNCtotalSupply
	reponseEnRetour.UstcTotalSupply = USTCtotalSupply
	return reponseEnRetour, ""

}