package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	port = 8001
)

type MapServer struct {
	répertoireCartes string
}

func getFormValue(hr *http.Request, name string) string {
	values := hr.Form[name]
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func envoieRéponse(w http.ResponseWriter, out *MessageOut) {
	bout, err := json.Marshal(out)
	if err != nil {
		fmt.Println("Erreur encodage réponse : ", err)
		return
	}
	fmt.Fprint(w, "receiveFromMapServer(")
	w.Write(bout)
	fmt.Fprint(w, ")")
}


func (ms *MapServer) ServeHTTP(w http.ResponseWriter, hr *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "GET")
	hr.ParseForm()

	in := new(MessageIn)
	out := new(MessageOut)
	defer envoieRéponse(w, out)
	fmt.Println(getFormValue(hr, "in"))
	bin := ([]byte)(getFormValue(hr, "in"))
	err := json.Unmarshal(bin, in)
	if err != nil {
		out.Erreur = "Erreur décodage : " + err.Error()
		fmt.Println("Erreur décodage : ", err.Error())
		return
	}
	//fmt.Printf("Message reçu : %+v\n", in)
	if in.IdBraldun==0 || len(in.Mdpr)!=64 {
		fmt.Println("IdBraldun ou Mot de passe restreint invalide")
		return
	}
	fmt.Println("IdBraldun : ", in.IdBraldun, "   Mdpr : ", in.Mdpr)
	hasher := sha1.New()
	hasher.Write(bin)
	sha := base64.StdEncoding.EncodeToString(hasher.Sum())
	fmt.Println("Clef SHA1 : ", sha)
	dir := fmt.Sprintf("%s/%d-%s/%s", ms.répertoireCartes, in.IdBraldun, sha, time.LocalTime().Format("2006/01/02"))
	path := dir + "/carte-"+sha+".json"
	if _,err=os.Stat(path); err!=nil { // le fichier n'existe pas, on le crée
		os.MkdirAll(dir, 0777)
		f, _ := os.Create(path)
		defer f.Close()
		f.Write(bin)
	}
}

func (server *MapServer) Start() {
	http.Handle("/", server)
	fmt.Printf("mapserver démarre sur le port %d\n", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("Erreur au lancement : ", err)
	}
}

func main() {
	ms := new(MapServer)
	ms.répertoireCartes = "/home/dys/braldop/cartes"
	ms.Start()
}
