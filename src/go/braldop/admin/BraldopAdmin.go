package main

import (
	"braldop/bra"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Exemples de commandes :
//  > bradmin -palette
//    exporte sur stdout la palette des couleurs des fonds à utiliser dans le javascript pour décoder les png
//  > bradmin -cmd png -in tata/toto.json 
//    construit un png (tata/couchexx.png) dont le contenu correspond à la Couche encodée dans le fichier tata/toto.json
//  > bradmin -cmd png -in toto;tutu -out titi
//    enrichit (ou crée) les fichiers titi/couchexx.png à partir des données couchexx.png des répertoires toto et tutu
func main() {
	exportePalette := flag.Bool("palette", false, "si oui alors la palette des environnements est exportée (exemple : \"bradmin -palette\")")
	in := flag.String("in", "", "source des données")
	out := flag.String("out", "", "répertoire de sortie")
	cmd := flag.String("cmd", "", "commande")
	flag.Parse()
	var err error
	if *exportePalette {
		bra.ExportePalettePng(os.Stdout)
	} else if *cmd == "png" {
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
					messin := new(MessageIn)
					bin, _ := ioutil.ReadAll(filein)
					err = json.Unmarshal(bin, messin)
					if err != nil {
						log.Fatal(err)
					}
					if messin.Vue == nil || len(messin.Vue.Couches) == 0 {
						log.Fatal(" Pas de données de vue")
					}
					bra.EnrichitCouchePNG(dir, &messin.Vue.Couches[0], 0)
				} else { // si pas json on suppose pour l'instant qu'il s'agit d'un répertoire de fichiers png
					log.Println("********\nEnrichitCouchesPNG", dir, cheminIn)
					err = bra.EnrichitCouchesPNG(dir, cheminIn)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	} else {
		log.Println("Usage :")
		flag.PrintDefaults()
	}
}
