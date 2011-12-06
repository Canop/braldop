package main

import (
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
	versionActuelleExtension = Version{[]uint{2, 2}}
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
		fmt.Println("Erreur encodage réponse : ", err)
		return
	}
	fmt.Fprint(w, "receiveFromMapServer(")
	w.Write(bout)
	fmt.Fprint(w, ")")
}

func vérifieVersion(vs string) string {
	if version, err := ParseVersion(vs); err != nil {
		fmt.Println(" version utilisateur incomprise : " + vs)
	} else if CompareVersions(version, &versionActuelleExtension) == -1 {
		fmt.Println(" version utilisateur obsolète : " + vs)
		return "L'extension Braldop n'est pas à jour.<br>Vous devriez installer <a href=http://canop.org/braldop/carte_et_extension.html>la nouvelle version</a>."
	}
	return ""
}

func (ms *MapServer) ServeHTTP(w http.ResponseWriter, hr *http.Request) {
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
		fmt.Println("Erreur décodage : ", err.Error())
		return
	}
	//fmt.Printf("Message reçu : %+v\n", in)
	out.Text = vérifieVersion(in.Version)
	if in.IdBraldun == 0 || len(in.Mdpr) != 64 {
		fmt.Println("IdBraldun ou Mot de passe restreint invalide")
		return
	}
	fmt.Println("IdBraldun : ", in.IdBraldun, "   Mdpr : ", in.Mdpr)
	if in.Vue == nil || len(in.Vue.Couches) == 0 {
		fmt.Println("Pas de données de vue")
		return
	}
	hasher := sha1.New()
	hasher.Write(bin)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum())
	fmt.Print("Clef SHA1 : ", sha)
	dirBase := fmt.Sprintf("%s/%d-%s", *ms.répertoireCartes, in.IdBraldun, in.Mdpr)
	dir := dirBase + "/" + time.LocalTime().Format("2006/01/02")
	path := dir + "/carte-" + sha + ".json"
	if _, err = os.Stat(path); err != nil { // le fichier n'existe pas, ce sont des données intéressantes
		fmt.Println(" -> Carte à modifier")
		//> on sauvegarde le fichier json
		os.MkdirAll(dir, 0777)
		f, _ := os.Create(path)
		defer f.Close()
		f.Write(bin)
		//> on crée ou enrichit l'image png correspondant à la couche
		in.Vue.Couches[0].ConstruitPNG(dirBase, true)
	} else {
		fmt.Println(" -> Carte inchangée")
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
	fmt.Printf("mapserver démarre sur le port %d\n", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("Erreur au lancement : ", err)
	}
}

func main() {
	ms := new(MapServer)
	ms.répertoireCartes = flag.String("cartes", "", "répertoire des cartes")
	cpuprofile := flag.String("cpuprofile", "", "fichier dans lequel écrire un bilan de profiling cpu")
	flag.Parse()
	if *ms.répertoireCartes == "" {
		fmt.Println("Chemin des cartes non fourni")
	} else {
		fmt.Println("Répertoire des cartes : " + *ms.répertoireCartes)
	}
	if *cpuprofile != "" {
		fmt.Println("Profiling actif, résultats dans le fichier ", *cpuprofile)
		fp, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(fp)
		go func() {
			for {
				sig := <-signal.Incoming
				fmt.Printf("Signal : %+v", sig)
				if usig, ok := sig.(os.UnixSignal); ok {
					if usig==syscall.SIGTERM || usig==syscall.SIGINT {
						fmt.Printf("Mapserver tué ! (", sig, ")")
						pprof.StopCPUProfile()
						os.Exit(0)
					}
				}
			}
		}()
	}
	ms.Start()
}
