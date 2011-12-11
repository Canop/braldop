package main

import (
	"braldop/bra"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"
	"syscall"
	"time"
)

const (
	port = 8001
)

var versionActuelleExtension Version

func init() {
	versionActuelleExtension = Version{[]uint{2, 4}}
}

type MapServer struct {
	répertoireCartes *string // répertoire racine dans lequel on trouve les répertoires des utilisateurs
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
		log.Println("Erreur encodage réponse : ", err)
		return
	}
	fmt.Fprint(w, "receiveFromMapServer(")
	w.Write(bout)
	fmt.Fprint(w, ")")
}

func vérifieVersion(vs string) (html string) {
	if version, err := ParseVersion(vs); err != nil {
		log.Println(" version utilisateur incomprise : " + vs)
	} else if CompareVersions(version, &versionActuelleExtension) == -1 {
		log.Println(" version utilisateur obsolète : " + vs)
		html = "L'extension Braldop n'est pas à jour.<br>Vous devriez installer <a href=http://canop.org/braldop/extension-braldop.html>la nouvelle version</a>."
		html += "<br>Il s'agit de la première béta publique."
	}
	return
}

func (ms *MapServer) ServeHTTP(w http.ResponseWriter, hr *http.Request) {
	startTime := time.Nanoseconds()
	defer func() { log.Printf(" durée totale traitement requête : %d ms\n", (time.Nanoseconds()-startTime)/1e6) }()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "GET")
	hr.ParseForm()
	in := new(MessageIn)
	out := new(MessageOut)
	defer envoieRéponse(w, out)
	bin := ([]byte)(getFormValue(hr, "in"))
	err := json.Unmarshal(bin, in)
	if err != nil {
		out.Erreur = "Erreur décodage : " + err.Error()
		log.Println("Erreur décodage : ", err.Error())
		return
	}
	//fmt.Printf("Message reçu : %+v\n", in)
	out.Text = vérifieVersion(in.Version)
	if in.IdBraldun == 0 || len(in.Mdpr) != 64 {
		log.Println("IdBraldun ou Mot de passe restreint invalide")
		return
	}
	log.Println("Requête Braldun ", in.IdBraldun)
	if in.Vue == nil || len(in.Vue.Couches) == 0 {
		log.Println(" Pas de données de vue")
		return
	}
	hasher := sha1.New()
	hasher.Write(bin)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum())
	dirBase := fmt.Sprintf("%s/%d-%s", *ms.répertoireCartes, in.IdBraldun, in.Mdpr)
	dir := dirBase + "/" + time.LocalTime().Format("2006/01/02")
	path := dir + "/carte-" + sha + ".json"
	if _, err = os.Stat(path); err != nil { // le fichier n'existe pas, ce sont des données intéressantes
		log.Println(" Carte à modifier")
		//> on sauvegarde le fichier json
		os.MkdirAll(dir, 0777)
		f, _ := os.Create(path)
		defer f.Close()
		f.Write(bin)
		//> on crée ou enrichit l'image png correspondant à la couche
		bra.EnrichitCouchePNG(dirBase, &in.Vue.Couches[0], 20)
	} else {
		log.Println(" Carte inchangée")
	}
	cheminLocalImage := fmt.Sprintf("%s/%d-%s/couche%d.png", *ms.répertoireCartes, in.IdBraldun, in.Mdpr, in.Vue.Couches[0].Z)
	if f, err := os.Open(cheminLocalImage); err == nil {
		defer f.Close()
		bytes, _ := ioutil.ReadAll(f)
		out.PngCouche = "data:image/png;base64," + base64.StdEncoding.EncodeToString(bytes)
	}
}

func (server *MapServer) Start() {
	http.Handle("/", server)
	log.Printf("mapserver démarre sur le port %d\n", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Println("Erreur au lancement : ", err)
	}
}

func main() {
	ms := new(MapServer)
	ms.répertoireCartes = flag.String("cartes", "", "répertoire des cartes")
	cpuprofile := flag.String("cpuprofile", "", "fichier dans lequel écrire un bilan de profiling cpu")
	memprofile := flag.String("memprofile", "", "fichier dans lequel écrire un bilan de profiling mémoire (lors de l'ordre d'arrêt)")
	flag.Parse()
	if *ms.répertoireCartes == "" {
		log.Println("Chemin des cartes non fourni")
	} else {
		log.Println("Répertoire des cartes : " + *ms.répertoireCartes)
	}
	if *cpuprofile != "" {
		log.Println("Profiling CPU actif, résultats dans le fichier ", *cpuprofile)
		fp, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(fp)
	}
	go func() {
		for {
			sig := <-signal.Incoming
			log.Printf("Signal : %+v", sig)
			if usig, ok := sig.(os.UnixSignal); ok {
				if usig == syscall.SIGTERM || usig == syscall.SIGINT {
					log.Printf("Mapserver tué ! (", sig, ")")
					if *memprofile != "" {
						log.Println("Ecriture heap dans le fichier ", *memprofile)
						fp, err := os.Create(*memprofile)
						if err != nil {
							log.Fatal(err)
						}
						pprof.WriteHeapProfile(fp)
					}
					bra.BloqueEcrituresPNG()
					if *cpuprofile != "" {
						pprof.StopCPUProfile()
					}
					os.Exit(0)
				}
			}
		}
	}()
	ms.Start()
}
