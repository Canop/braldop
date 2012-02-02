package bra

// Interrogations des Scripts Publics 


import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

// vérifie auprès du serveur de scripts publics que le couple login/password est correct
func AuthentifieCompteParScriptPublic(idBraldun uint, mdpr string) (bool, errorDetails string) {
	httpClient := new(http.Client)
	request := fmt.Sprintf("http://sp.braldahim.com/scripts/profil/?idBraldun=%d&mdpRestreint=%s&version=5", idBraldun, mdpr)
	resp, err := httpClient.Get(request)
	if err != nil {
		errorDetails = err.Error()
		return
	}
	defer resp.Body.Close()
	r := bufio.NewReader(resp.Body)
	bline, _, _ := r.ReadLine()
	line := string(bline)
	for line != "" {
		fmt.Println(line)
		bline, _, _ = r.ReadLine()
		line = string(bline)
	}
	return true
}
