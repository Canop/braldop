// braldopadmin est utilisé pour effectuer certaines manipulations occasionnelles ou de test.
// Pour avoir la documentation, tapez simplement braldopadmin à la ligne de commande.

package main

import (
	"bra"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cmd := os.Args[1]
	flags := flag.NewFlagSet(cmd, flag.ExitOnError)
	in := flags.String("in", "", "source des données")
	out := flags.String("out", "", "répertoire de sortie (pour la commande png)")
	id := flags.Int("id", 0, "id braldun")
	mdpr := flags.String("mdpr", "", "mot de passe restreint")
	data := flags.String("data", "", "chemin de base du stockage (contient en principe des répertoires 'private', 'public' et 'cartes')")
	flags.Parse(os.Args[2:])
	var err error
	switch cmd {
	case "palette":
		bra.ExportePalettePng(os.Stdout)
	case "stats":
		if *data == "" {
			log.Println("Chemin de stockage non précisé (nécessaire pour lire bralduns.csv et communautes.csv)")
		} else {
			memmap, err := bra.ChargeDonnéesStatiquesPubliques(filepath.Join(*data, "public"), true)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(len(memmap.Bralduns), "bralduns")
		}
	case "vue":
		if *id == 0 {
			log.Println("Id du braldun non précisé")
		} else if *mdpr == "" {
			log.Println("Mot de passe restreint non précisé")
		} else if *data == "" {
			log.Println("Chemin de stockage non précisé (nécessaire pour lire bralduns.csv et communautes.csv)")
		} else {
			dv, err := bra.VueParScriptPublic(uint(*id), *mdpr, filepath.Join(*data, "public"))
			if err != nil {
				log.Fatal(err)
			}
			bout, err := json.Marshal(dv)
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write(bout)
		}
	case "etat":
		if *id == 0 {
			log.Println("Id du braldun non précisé")
		} else if *mdpr == "" {
			log.Println("Mot de passe restreint non précisé")
		} else {
			état, err := bra.EtatBraldunParScriptPublic(uint(*id), *mdpr)
			if err != nil {
				log.Fatal(err)
			}
			bout, err := json.Marshal(état)
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write(bout)
		}
	case "check":
		if *id == 0 {
			log.Println("Id du braldun non précisé")
		} else if *mdpr == "" {
			log.Println("Mot de passe restreint non précisé")
		} else {
			auth, err := bra.AuthentifieCompteParScriptPublic(uint(*id), *mdpr)
			if err != nil {
				log.Println("Unable to authenticate", err)
			} else {
				log.Println("Autentication : ", auth)
			}
		}
	case "png":
		if *in == "" {
			log.Println("Source de données non précisée (in devrait être le chemin d'un fichier json)")
		} else {
			dir := *out
			cheminsIn := strings.Split(*in, ";")
			if dir == "" {
				dir, _ = filepath.Split(cheminsIn[0])
				cheminsIn = cheminsIn[1:]
			}
			dir, err = filepath.Abs(dir)
			if err != nil {
				log.Fatal(err)
			}
			for _, cheminIn := range cheminsIn {
				cheminIn, err = filepath.Abs(cheminIn)
				if err != nil {
					log.Fatal(err)
				}
				if filepath.Ext(cheminIn) == ".json" {
					filein, err := os.Open(*in)
					if err != nil {
						log.Fatal(err)
					}
					defer filein.Close()
					messin := new(bra.MessageIn)
					bin, _ := ioutil.ReadAll(filein)
					err = json.Unmarshal(bin, messin)
					if err != nil {
						log.Fatal(err)
					}
					if messin.Vue == nil || len(messin.Vue.Couches) == 0 {
						log.Fatal(" Pas de données de vue")
					}
					bra.EnrichitCouchePNG(dir, messin.Vue.Couches[0], 0)
				} else { // si pas json on suppose pour l'instant qu'il s'agit d'un répertoire de fichiers png
					log.Println("********\nEnrichitCouchesPNG", dir, cheminIn)
					err = bra.EnrichitCouchesPNG(dir, cheminIn)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	default:
		fmt.Println("Syntaxe :")
		fmt.Println(" > braldopadmin [{check|palette|png|etat|vue}] [paramètres]")
		fmt.Println("Paramètres :")
		flags.PrintDefaults()
		fmt.Println("Exemples :")
		fmt.Println(" > braldopadmin palette")
		fmt.Println("   exporte sur stdout la palette des couleurs des fonds à utiliser dans le javascript pour décoder les png")
		fmt.Println(" > braldopadmin png -in tata/toto.json")
		fmt.Println("   construit un png (tata/couchexx.png) dont le contenu correspond à la Couche encodée dans le fichier tata/toto.json")
		fmt.Println(" > braldopadmin png -in toto;tutu -out titi")
		fmt.Println("   enrichit (ou crée) les fichiers titi/couchexx.png à partir des données couchexx.png des répertoires toto et tutu")
		fmt.Println(" > braldopadmin check -id 754 -mdpr XXXX")
		fmt.Println("   vérifie que le braldun 754 est connu de Braldahim et a bien le mot de passe restreint XXXX")
		fmt.Println(" > braldopadmin vue -id 754 -mdpr XXXX")
		fmt.Println("   renvoie en json la vue du braldun obtenue par script public braldahim")
		fmt.Println(" > braldopadmin etat -id 754 -mdpr XXXX")
		fmt.Println("   renvoie en json l'état du braldun obtenu par script public braldahim")
		fmt.Println("Attention : certaines de ces commandes font des requêtes au serveur sp.braldahim.com et décrémentent le nombre de requêtes possibles pour la journée pour le braldun en question.")
	}
	fmt.Println()
}
