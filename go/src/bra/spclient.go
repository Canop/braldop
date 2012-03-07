package bra

// Interrogations des Scripts Publics de Braldahim

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// vérifie auprès du serveur de scripts publics que le couple login/password est correct
func AuthentifieCompteParScriptPublic(idBraldun uint, mdpr string) (bool, error) {
	httpClient := new(http.Client)
	request := fmt.Sprintf("http://sp.braldahim.com/scripts/profil/?idBraldun=%d&mdpRestreint=%s&version=2", idBraldun, mdpr)
	resp, err := httpClient.Get(request)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	r := bufio.NewReader(resp.Body)
	bline, _, _ := r.ReadLine()
	line := string(bline)
	log.Printf(" Réponse authentification %d : %s\n", idBraldun, line)
	if strings.Contains(line, "ERREUR") {
		return false, nil
	}
	return true, nil
}

// récupère par script public la vue du braldun, et la construit en exploitant les données 
//  des fichiers bralduns.csv et communautes.csv trouvés dans le répertoire fourni (on mettra ça en cache plus tard
//   mais ça impliquera de faire en sorte que les données ne soient pas obsolètes)
func VueParScriptPublic(idBraldun uint, mdpr string, cheminRépertoireCsvPublic string) (*DonnéesVue, error) {
	spbralduns, err := os.Open(filepath.Join(cheminRépertoireCsvPublic, "bralduns.csv"))
	if err != nil {
		log.Println(" Erreur à l'ouverture du fichier bralduns.csv :", err)
		return nil, err
	}
	defer spbralduns.Close()
	spcommunautes, err := os.Open(filepath.Join(cheminRépertoireCsvPublic, "communautes.csv"))
	if err != nil {
		log.Println(" Erreur à l'ouverture du fichier communautes.csv :", err)
		return nil, err
	}
	defer spcommunautes.Close()
	httpClient := new(http.Client)
	request := fmt.Sprintf("http://sp.braldahim.com/scripts/vue/?idBraldun=%d&mdpRestreint=%s&version=5", idBraldun, mdpr)
	resp, err := httpClient.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r := bufio.NewReader(resp.Body)
	bline, _, _ := r.ReadLine()
	line := string(bline)
	if strings.Contains(line, "ERREUR") {
		return nil, errors.New("Erreur script public : " + line)
	}
	return ConstruitDonnéesVue(r, bufio.NewReader(spbralduns), bufio.NewReader(spcommunautes), true), nil
}

func EtatBraldunParScriptPublic(idBraldun uint, mdpr string) (*EtatBraldun, error) {
	httpClient := new(http.Client)
	request := fmt.Sprintf("http://sp.braldahim.com/scripts/profil/?idBraldun=%d&mdpRestreint=%s&version=2", idBraldun, mdpr)
	resp, err := httpClient.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r := bufio.NewReader(resp.Body)
	bline, _, _ := r.ReadLine()
	line := string(bline)
	if strings.Contains(line, "ERREUR") {
		return nil, errors.New("Erreur script public : " + line)
	}
	return LitEtatBraldunDansCsvProfil(r)
}
