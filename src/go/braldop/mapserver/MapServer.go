package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"flag"
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
	répertoireCartes *string
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
	//~ fmt.Println(getFormValue(hr, "in"))
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
	if in.Vue==nil || len(in.Vue.Couches)==0 {
		fmt.Println("Pas de données de vue")
		return
	}
	hasher := sha1.New()
	hasher.Write(bin)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum())
	fmt.Println("Clef SHA1 : ", sha)
	dirBase := fmt.Sprintf("%s/%d-%s", *ms.répertoireCartes, in.IdBraldun, in.Mdpr)
	dir := dirBase + "/" + time.LocalTime().Format("2006/01/02")
	path := dir + "/carte-"+sha+".json"
	if _,err=os.Stat(path); err!=nil { // le fichier n'existe pas, ce sont des données intéressantes
		fmt.Println(" Erreur : ", err)
		//> on sauvegarde le fichier json
		os.MkdirAll(dir, 0777)
		f, _ := os.Create(path)
		defer f.Close()
		f.Write(bin)
		//> on crée ou enrichit l'image png correspondant à la couche
		in.Vue.Couches[0].ConstruitPNG(dirBase, true)
	} else {
		fmt.Println(" Carte inchangée")
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
	ms.répertoireCartes = flag.String("cartes", "", "répertoire des cartes")
	flag.Parse()
	if *ms.répertoireCartes=="" {
		fmt.Println("Chemin des cartes non fourni")
	} else {
		fmt.Println("Répertoire des cartes : " + *ms.répertoireCartes)
	}
	ms.Start()
}
