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


type nbStakedLuncLcdStructure struct {
	Pool struct {
		Not_bonded_tokens string `json:"not_bonded_tokens"`
		Bonded_tokens 	 string `json:"bonded_tokens"`
	}
	Error struct {
		Code 	int 		`json:"code"`
		Message string		`json:"message"`
	}
}

type StructReponseNbStakedLunc struct {
	NbStakedLunc int
	// Autres infos inutilisées, pour l'instant
}


func GetNbStakedLunc(channelForErrors chan<- string) (StructReponseNbStakedLunc, string) {

	// Initialisation de la struct qui sera renvoyée en retour
	var reponseEnRetour StructReponseNbStakedLunc

	// Path, pour accéder à ce qui nous intéresse
	var path = "/cosmos/staking/v1beta1/pool"

	// Récupération de l'url du LCD
	var LCDurl = config.AppConfig.Lcd.Url

	// Création d'un client HTTP (avec timeout fixé à 30 secondes)
	client := &http.Client{
        Timeout: time.Duration(config.AppConfig.Lcd.GetTimeoutInSeconds) * time.Second,
    }

	// Lancement du GET
	reponse, errGet := client.Get(LCDurl + path)
	if errGet != nil {
		return reponseEnRetour, "failed to fetch 'nb staked lunc' from LCD"
	}

	// Lecture de la réponse du GET
	body, errReadAll := io.ReadAll(reponse.Body)
	if errReadAll != nil {
		return reponseEnRetour, "failed to read 'nb staked lunc' answer from LCD"
	}

	// Transformation byte[] -> string pour avoir du contenu JSON "en clair"
	reponseJSON := string(body)

	// Stockage des données dans une struct
	dataStruct := nbStakedLuncLcdStructure{}
	json.Unmarshal([]byte(reponseJSON), &dataStruct)

	// Sortie si erreur retournée
	if dataStruct.Error.Message != "" {
		return reponseEnRetour, "an error was returned while fetching 'nb staked lunc' from LCD"
	}

	// Récupération du nombre de Lunc stakés ("bonded tokens")
	var nbStakedLunc int = -1
	bonded_tokens, errBondedTokens := strconv.ParseFloat(dataStruct.Pool.Bonded_tokens, 64)
	if errBondedTokens != nil {
		return reponseEnRetour, "failed to convert 'uluna' amount in 'lunc' from LCD"
	}
	nbStakedLunc = int(bonded_tokens / 1000000)

	// Si jamais ce n'est pas supérieur à zéro, alors on remonte une erreur
	if(nbStakedLunc <= 0) {
		stringToReturn := fmt.Sprintf("[dataloader] GetNbStakedLunc : -1 or 0 returned by function.\nError = %s", reponseJSON)
		channelForErrors <- stringToReturn
		return reponseEnRetour, "failed to get 'bonded tokens' from LCD (-1 or 0 returned)"
	}

	// Et renvoi à l'appeleur
	reponseEnRetour.NbStakedLunc = nbStakedLunc
	return reponseEnRetour, ""

}