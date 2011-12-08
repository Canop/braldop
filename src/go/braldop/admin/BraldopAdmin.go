package main

import (
	"braldop/bra"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	exportePalette := flag.Bool("palette", false, "si oui alors la palette des environnements est exportée (exemple : \"bradmin -palette\")")
	in := flag.String("in", "", "source des données")
	cmd := flag.String("cmd", "", "commande")
	flag.Parse()

	if *exportePalette {
		bra.ExportePalettePng(os.Stdout)
	} else if *cmd == "png" {
		if *in == "" {
			log.Println("Source de données non précisée (in devrait être le chemin d'un fichier json)")
		} else {
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
			dir, _ := filepath.Split(*in)
			messin.Vue.Couches[0].EnrichitPNG(dir, 0)
		}
	} else {
		log.Println("Usage :")
		flag.PrintDefaults()
	}
}
