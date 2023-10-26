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


type communityPoolContentStructure struct {
	Pool []struct {
		Denom string `json:"denom"`
		Amount 	 string `json:"amount"`
	}
	Error struct {
		Code 	int 		`json:"code"`
		Message string		`json:"message"`
	}
}

type StructReponseCommunityPoolContent struct {
	NbLuncInCommunityPool int
	NbUstcInCommunityPool int
}


func GetCommunityPoolContent(channelForLogsMsgs chan<- string) (StructReponseCommunityPoolContent, string) {

	// Initialisation de la struct qui sera renvoyée en retour
	var reponseEnRetour StructReponseCommunityPoolContent

	// Path, pour accéder à ce qui nous intéresse
	var path = config.AppConfig.Lcd.PathForCommunityPoolContent

	// Récupération de l'url du LCD
	var LCDurl = config.AppConfig.Lcd.Url

	// Création d'un client HTTP (avec timeout fixé à 30 secondes)
	client := &http.Client{
        Timeout: time.Duration(config.AppConfig.Lcd.GetTimeoutInSeconds) * time.Second,
    }

	// Lancement du GET
	reponse, errGet := client.Get(LCDurl + path)
	if errGet != nil {
		return reponseEnRetour, "failed to fetch 'community pool content' from LCD"
	}

	// Lecture de la réponse du GET
	body, errReadAll := io.ReadAll(reponse.Body)
	if errReadAll != nil {
		return reponseEnRetour, "failed to read 'community pool content' answer from LCD"
	}

	// Transformation byte[] -> string pour avoir du contenu JSON "en clair"
	reponseJSON := string(body)

	// Stockage des données dans une struct
	dataStruct := communityPoolContentStructure{}
	json.Unmarshal([]byte(reponseJSON), &dataStruct)

	// Sortie si erreur retournée
	if dataStruct.Error.Message != "" {
		return reponseEnRetour, "failed to unmarshal 'community pool content' from LCD"
	}

	// Récupération du nombre de Lunc et Ustc contenus dans le Community Pool
	var nbLuncInCommunityPool int = -1
	var nbUstcInCommunityPool int = -1
	for i:=0 ; i<len(dataStruct.Pool) ; i++ {
		if dataStruct.Pool[i].Denom == "uluna" {
			uluna, errUluna := strconv.ParseFloat(dataStruct.Pool[i].Amount, 64)
			if errUluna != nil {
				return reponseEnRetour, "failed to convert 'uluna' amount in 'lunc' from LCD"
			}
			nbLuncInCommunityPool = int(uluna / 1000000)
		}
		if dataStruct.Pool[i].Denom == "uusd" {
			uusd, errUusd := strconv.ParseFloat(dataStruct.Pool[i].Amount, 64)
			if errUusd != nil {
				return reponseEnRetour, "failed to convert 'uusd' amount in 'ustc' from LCD"
			}
			nbUstcInCommunityPool = int(uusd / 1000000)
		}
	}

	// Si les valeurs à retourner ne sont pas supérieures à zéro, alors on remonte une erreur
	if(nbLuncInCommunityPool <= 0 || nbUstcInCommunityPool <= 0) {
		stringToReturn := fmt.Sprintf("[dataloader] GetCommunityPoolContent : -1 or 0 returned by function.\nError = %s", reponseJSON)
		channelForLogsMsgs <- stringToReturn
		return reponseEnRetour, "failed to get 'community pool content' from LCD (-1 or 0 returned)"
	}

	// Et renvoi à l'appeleur
	reponseEnRetour.NbLuncInCommunityPool = nbLuncInCommunityPool
	reponseEnRetour.NbUstcInCommunityPool = nbUstcInCommunityPool
	return reponseEnRetour, ""

}