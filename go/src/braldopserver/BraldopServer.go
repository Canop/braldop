package main

import (
	"bra"
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
	"path/filepath"
	"runtime/pprof"
	"strconv"
	"time"
)

const (
	port = 8001
)

var versionActuelleExtension Version

func init() {
	versionActuelleExtension = MakeVersion(3, 0)
}

type MapServer struct {
	répertoireDonnées string // répertoire racine dans lequel on trouve les répertoires des utilisateurs, les fichiers csv publics, etc.
	répertoireCartes  string // répertoire racine dans lequel on trouve les répertoires des utilisateurs, les fichiers csv publics, etc.
	bd                *bra.BaseMysql
	fv                FusionneurVue
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
		log.Println("Version utilisateur incomprise : " + vs)
	} else if CompareVersions(&version, &versionActuelleExtension) == -1 {
		log.Println("Version utilisateur obsolète : " + vs)
		html = "L'extension Braldop n'est pas à jour.<br>Vous devriez installer <a href=http://canop.org/braldop/index.html>la nouvelle version</a>."
	}
	return
}

func (ms *MapServer) répertoireCartesBraldun(idBraldun uint, mdpr string) string {
	return fmt.Sprintf("%s/%d-%s", ms.répertoireCartes, idBraldun, mdpr)
}

// bv est du binaire correspondant à l'encodage json d'un MessageIn
func (ms *MapServer) stockeVue(idVoyeur uint, mdprVoyeur string, bv []byte, couche *bra.Couche) (nouveau bool) {
	dirBase := ms.répertoireCartesBraldun(idVoyeur, mdprVoyeur)
	hasher := sha1.New()
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	dir := dirBase + "/" + time.Now().Format("2006/01/02")
	path := dir + "/carte-" + sha + ".json"
	if _, err := os.Stat(path); err != nil { // nouveau fichier, donc nouvelle vue
		//> on sauvegarde le fichier json
		os.MkdirAll(dir, 0777)
		f, _ := os.Create(path)
		defer f.Close()
		f.Write(bv)
		//> on crée ou enrichit l'image png correspondant à la couche
		bra.EnrichitCouchePNG(dirBase, couche, 20)
		return true
	}
	return false
}

func (ms *MapServer) ServeHTTP(w http.ResponseWriter, hr *http.Request) {
	startTime := time.Now().UnixNano()
	defer func() {
		log.Printf(" durée totale traitement requête : %d ms\n", (time.Now().UnixNano()-startTime)/1e6)
	}()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "GET")

	//> analyse et vérification de la requête
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
	dirBase := ms.répertoireCartesBraldun(in.IdBraldun, in.Mdpr)
	log.Println("Requête Braldun ", in.IdBraldun)
	var couche *bra.Couche
	if in.Vue == nil || len(in.Vue.Couches) == 0 {
		log.Println(" Pas de données de vue")
	} else {
		couche = in.Vue.Couches[0]
	}

	//> récupération du compte braldop authentifié et du tableau des amis
	var cb *bra.CompteBraldop
	var amis []*bra.CompteBraldop
	con, err := ms.bd.DB()
	if err != nil {
		log.Println("Erreur à la connexion bd", err)
		out.Text = "Erreur Braldop : connexion BD"
		return
	}
	cb, errmess := con.AuthentifieCompte(in.IdBraldun, in.Mdpr, true)
	if errmess != "" {
		out.Text = "Une erreur s'est produite durant l'authentification sur Braldop : <i>" + errmess + "</i><br>Contactez Canopée du Haut-Rac pour plus d'informations."
		return // on a besoin du compte
	}
	log.Println(" compte authentifié")
	amis, err = con.Amis(cb.IdBraldun)
	if err != nil {
		log.Println(" erreur durant récupération amis")
	}

	log.Println(" commande :", in.Cmd)

	//> stockage des données de vue
	if couche != nil {
		ms.fv.Reçoit(in.Vue.Vues[0])
		if ms.stockeVue(in.IdBraldun, in.Mdpr, bin, couche) {
			//> on s'occupe aussi des amis
			if amis != nil {
				for _, ami := range amis {
					log.Println(" enrichissement carte ami ", ami.IdBraldun)
					bra.EnrichitCouchePNG(ms.répertoireCartesBraldun(ami.IdBraldun, ami.Mdpr), couche, 20)
				}
			}
		} else {
			log.Println(" Carte inchangée")
		}
	}

	if in.Cmd == "carte" || in.Cmd == "" { // pour la compatibilité ascendante, la commande est provisoirement optionnelle

		//> récupération et stockage, si demandé, de la vue d'un autre braldun
		if in.Action == "maj" && in.Cible > 0 {
			log.Println(" Demande mise à jour de la vue de ", in.Cible)
			if !ms.fv.MajPossible(in.Cible) {
				log.Println(" Impossible de mettre à jour la vue de ", in.Cible)
			} else {
				cc, errstr := con.GetCompteExistant(in.Cible)
				if errstr != "" {
					log.Println(" erreur durant la récupération du compte de", in.Cible)
				} else {
					dv, err := bra.VueParScriptPublic(in.Cible, cc.Mdpr, filepath.Join(ms.répertoireDonnées, "public"))
					if err != nil {
						log.Println(" erreur durant la récupération par script public de la vue de", in.Cible)
					} else {
						ms.fv.Reçoit(dv.Vues[0])
						spin := new(MessageIn) // on crée un messageIn car c'est sous ce format que sont stockés les données des bralduns
						spin.IdBraldun = in.Cible
						spin.Mdpr = cc.Mdpr
						spin.Vue = dv
						spin.Action = in.Action
						spin.Cmd = "sp_vue"
						spin.Version = versionActuelleExtension.String()
						bspin, err := json.Marshal(spin)
						if err != nil {
							log.Println(" erreur durant l'encodage en json de ", in.Cible)
						} else {
							if ms.stockeVue(in.Cible, cc.Mdpr, bspin, dv.Couches[0]) {
								camis, err := con.Amis(in.Cible)
								if err != nil {
									log.Println(" erreur durant récupération des amis de", in.Cible)
								}
								for _, cami := range camis {
									log.Println(" enrichissement carte ", cami.IdBraldun)
									bra.EnrichitCouchePNG(ms.répertoireCartesBraldun(cami.IdBraldun, cami.Mdpr), dv.Couches[0], 20)
								}
							}
						}
					}
				}
			}
		}

		//> renseignements sur les couches disponibles
		out.ZConnus, err = bra.CouchesPNGDisponibles(dirBase)
		if err != nil {
			log.Println(" erreur durant la détermination des couches disponibles")
		}

		//> renvoi des données de vues provenant des amis
		log.Println(" préparation DV")
		log.Printf(" Amis : %+v\n", amis)
		if amis != nil {
			vues := ms.fv.Complète(in.Vue.Vues[0], amis)
			log.Printf(" Vues : %+v\n", vues)
			if len(vues) > 0 {
				out.DV = new(bra.DonnéesVue)
				out.DV.Vues = vues
			}
		}
		log.Printf(" out.DV : %+v\n", out.DV)

		//> renvoi de la carte en png
		log.Println(" ZRequis : ", in.ZRequis)
		out.Z = in.ZRequis
		cheminLocalImage := fmt.Sprintf("%s/%d-%s/couche%d.png", ms.répertoireCartes, in.IdBraldun, in.Mdpr, in.ZRequis)
		if f, err := os.Open(cheminLocalImage); err == nil {
			defer f.Close()
			bytes, _ := ioutil.ReadAll(f)
			out.PngCouche = "data:image/png;base64," + base64.StdEncoding.EncodeToString(bytes)
		}
	} else if in.Cmd == "partages" {
		if in.Action != "" && in.Cible > 0 {
			con.ModifiePartage(in.IdBraldun, in.Cible, in.Action)
		}
		out.Partages, err = con.Partages(in.IdBraldun)
		if err != nil {
			log.Println(" erreur à la récupération des partages :", err)
		}
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
	datadir := flag.String("datadir", "", "répertoire des données (contient 'cartes', 'public' et éventuellement 'private')")
	cpuprofile := flag.String("cpuprofile", "", "fichier dans lequel écrire un bilan de profiling cpu")
	memprofile := flag.String("memprofile", "", "fichier dans lequel écrire un bilan de profiling mémoire (lors de l'ordre d'arrêt)")
	mysqluser := flag.String("mysqluser", "", "user pour l'accès mysql")
	mysqlmdp := flag.String("mysqlmdp", "", "mdp pour l'accès mysql")
	mysqldb := flag.String("mysqldb", "braldop", "base mysql")
	flag.Parse()

	ms.bd = bra.NewBaseMysql(*mysqluser, *mysqlmdp, *mysqldb)

	if *datadir == "" {
		log.Fatal("Chemin des cartes non fourni")
	}
	ms.répertoireDonnées = *datadir
	ms.répertoireCartes = filepath.Join(*datadir, "cartes")
	log.Println("Répertoire des données : " + ms.répertoireDonnées)
	if *cpuprofile != "" {
		log.Println("Profiling CPU actif, résultats dans le fichier ", *cpuprofile)
		fp, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(fp)
	}
	ms.fv.Charge(ms.répertoireCartes)
	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt, os.Kill)
		for {
			sig := <-sigchan
			log.Println("Signal : %+v", sig)
			log.Println("Mapserver tué ! (", sig, ")")
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
	}()
	ms.Start()
}
