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


type oraclePoolContentStructure struct {
	Balances []struct {
		Denom string `json:"denom"`
		Amount 	 string `json:"amount"`
	}
	Error struct {
		Code 	int 		`json:"code"`
		Message string		`json:"message"`
	}
}

type StructReponseOraclePoolContent struct {
	NbLuncInOraclePool int
	NbUstcInOraclePool int
}


func GetOraclePoolContent(channelForLogsMsgs chan<- string) (StructReponseOraclePoolContent, string) {

	// Initialisation de la struct qui sera renvoyée en retour
	var reponseEnRetour StructReponseOraclePoolContent

	// Path, pour accéder à ce qui nous intéresse
	var path = config.AppConfig.Lcd.PathForOraclePoolContent

	// Récupération de l'url du LCD
	var LCDurl = config.AppConfig.Lcd.Url

	// Création d'un client HTTP (avec timeout fixé à 30 secondes)
	client := &http.Client{
        Timeout: time.Duration(config.AppConfig.Lcd.GetTimeoutInSeconds) * time.Second,
    }

	// Lancement du GET
	reponse, errGet := client.Get(LCDurl + path)
	if errGet != nil {
		return reponseEnRetour, "failed to fetch 'oracle pool content' from LCD"
	}

	// Lecture de la réponse du GET
	body, errReadAll := io.ReadAll(reponse.Body)
	if errReadAll != nil {
		return reponseEnRetour, "failed to read 'oracle pool content' answer from LCD"
	}

	// Transformation byte[] -> string pour avoir du contenu JSON "en clair"
	reponseJSON := string(body)

	// Stockage des données dans une struct
	dataStruct := oraclePoolContentStructure{}
	json.Unmarshal([]byte(reponseJSON), &dataStruct)

	// Sortie si erreur retournée
	if dataStruct.Error.Message != "" {
		return reponseEnRetour, "failed to unmarshal 'oracle pool content' from LCD"
	}

	// Récupération du nombre de Lunc et Ustc contenus dans l'Oracle Pool
	var nbLuncInOraclePool int = -1
	var nbUstcInOraclePool int = -1
	for i:=0 ; i<len(dataStruct.Balances) ; i++ {
		if dataStruct.Balances[i].Denom == "uluna" {
			uluna, errUluna := strconv.ParseFloat(dataStruct.Balances[i].Amount, 64)
			if errUluna != nil {
				return reponseEnRetour, "failed to convert 'uluna' amount in 'lunc' from LCD"
			}
			nbLuncInOraclePool = int(uluna / 1000000)
		}
		if dataStruct.Balances[i].Denom == "uusd" {
			uusd, errUusd := strconv.ParseFloat(dataStruct.Balances[i].Amount, 64)
			if errUusd != nil {
				return reponseEnRetour, "failed to convert 'uusd' amount in 'ustc' from LCD"
			}
			nbUstcInOraclePool = int(uusd / 1000000)
		}
	}

	// Si les valeurs à retourner ne sont pas supérieures à zéro, alors on remonte une erreur
	if(nbLuncInOraclePool <= 0 || nbUstcInOraclePool <= 0) {
		stringToReturn := fmt.Sprintf("[dataloader] GetOraclePoolContent : -1 or 0 returned by function.\nError = %s", reponseJSON)
		channelForLogsMsgs <- stringToReturn
		return reponseEnRetour, "failed to get 'oracle pool content' from LCD (-1 or 0 returned)"
	}

	// Et renvoi à l'appeleur
	reponseEnRetour.NbLuncInOraclePool = nbLuncInOraclePool
	reponseEnRetour.NbUstcInOraclePool = nbUstcInOraclePool
	return reponseEnRetour, ""

}