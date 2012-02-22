package bra

// Interrogations des Scripts Publics de Braldahim

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
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

// récupère par script public la vue du braldun
func VueParScriptPublic(idBraldun uint, mdpr string) (*DonnéesVue, error) {
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
	return ParseFichierDynamiqueDonnéesVue(r, time.Now().Unix(), true), nil
}
