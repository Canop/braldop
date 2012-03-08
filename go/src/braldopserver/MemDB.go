package main

import (
	"bra"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// pour l'instant la "fusion" n'est qu'un objet interne utilisé par la fonction Complète
type FusionVues struct {
	Vues []*bra.Vue
}

func (fusion FusionVues) Len() int {
	return len(fusion.Vues)
}
func (fusion FusionVues) Less(i, j int) bool {
	return fusion.Vues[i].Time < fusion.Vues[j].Time
}
func (fusion FusionVues) Swap(i, j int) {
	fusion.Vues[i], fusion.Vues[j] = fusion.Vues[j], fusion.Vues[i]
}

type MemDB struct {
	vues map[uint]*bra.Vue // clef : id braldun
}

func (mdb *MemDB) Reçoit(vue *bra.Vue) {
	vue.Time = time.Now().Unix()
	log.Println("  Réception vue", vue.Voyeur, "v.Time :", vue.Time)
	mdb.vues[vue.Voyeur] = vue
}

// renvoie true si la dernière maj de la vue de ce braldun est suffisamment ancienne
//  pour qu'on en redemande une
func (mdb *MemDB) MajPossible(idBraldun uint) bool {
	if v, ok := mdb.vues[idBraldun]; ok {
		âge := time.Now().Unix() - v.Time
		log.Println("  Maintenant :", time.Now().Unix())
		log.Println("  v.Time :", v.Time)
		log.Println("  Age vue en secondes :", âge)
		return âge > 3*60*60
	}
	return true
}

func (mdb *MemDB) Fusionne(idBraldun uint, amis []*bra.CompteBraldop) []*bra.Vue {
	fusion := new(FusionVues)
	fusion.Vues = make([]*bra.Vue, 1, len(amis)+1)
	var ok bool
	if fusion.Vues[0], ok = mdb.vues[idBraldun]; !ok {
		return nil // il faut au moins la vue du braldun principal
	}
	for _, ami := range amis {
		if dvb, ok := mdb.vues[ami.IdBraldun]; ok {
			fusion.Vues = append(fusion.Vues, dvb)
		}
	}
	sort.Sort(fusion) // index faible : plus ancien
	//> on ajoute le prénom du Braldun voyeur à chaque vue (faudrait faire ça ailleurs)
	for _, vi := range fusion.Vues {
		for _, b := range vi.Bralduns { // pas super joli...
			if b.Id == vi.Voyeur {
				vi.PrénomVoyeur = b.Prénom
				break
			}
		}
	}
	return fusion.Vues
}

// renvoie les vues
func (mdb *MemDB) Complète(vue *bra.Vue, amis []*bra.CompteBraldop) []*bra.Vue {
	fusion := new(FusionVues)
	fusion.Vues = make([]*bra.Vue, 1, len(amis)+1)
	fusion.Vues[0] = vue
	for _, ami := range amis {
		if dvb, ok := mdb.vues[ami.IdBraldun]; ok {
			fusion.Vues = append(fusion.Vues, dvb)
		}
	}
	sort.Sort(fusion) // index faible : plus ancien

	//> on ajoute le prénom du Braldun voyeur à chaque vue (faudrait faire ça ailleurs)
	for _, vi := range fusion.Vues {
		for _, b := range vi.Bralduns {
			if b.Id == vi.Voyeur {
				vi.PrénomVoyeur = b.Prénom
				break
			}
		}
	}
	return fusion.Vues
}

// appelée au lancement cette fonction charge les donnéesVueBraldun
//  à partir des fichiers json (le dernier de chaque utilisateur).
// Notons qu'on ne vérifie pas ici les mdpr
func (mdb *MemDB) Charge(répertoireCartes string) error {
	mdb.vues = make(map[uint]*bra.Vue)
	rep, err := os.Open(répertoireCartes)
	if err != nil {
		return err
	}
	defer rep.Close()
	cfis, _ := rep.Readdir(-1)
	for _, cfi := range cfis {
		t := strings.Split(cfi.Name(), "-")
		if len(t) != 2 {
			continue
		}
		idBraldun, _ := strconv.ParseUint(t[0], 10, 0)
		if idBraldun == 0 {
			continue
		}
		plusRécentChemin := ""
		plusRécentDate := int64(0)
		// on parcoure toute l'arborescence, pas le plus efficace mais le plus simple à coder
		filepath.Walk(filepath.Join(répertoireCartes, cfi.Name()), func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && strings.HasPrefix(info.Name(), "carte-") && strings.HasSuffix(info.Name(), ".json") && info.ModTime().Unix() > plusRécentDate {
				plusRécentChemin = path
				plusRécentDate = info.ModTime().Unix()
			}
			return nil
		})
		if plusRécentChemin == "" {
			log.Println("  Pas de fichier json pour", idBraldun)
			continue
		}
		fjson, err := os.Open(plusRécentChemin)
		if err != nil {
			log.Println(err)
			continue
		}
		in := new(bra.MessageIn)
		bin, err := ioutil.ReadAll(fjson)
		defer fjson.Close()
		if err != nil {
			log.Println("  Erreur à la lecture du dernier fichier json du braldun", idBraldun)
			continue
		}
		err = json.Unmarshal(bin, in)
		if err != nil {
			log.Println("  Erreur à la lecture du dernier fichier json du braldun", idBraldun)
			continue
		}
		if in.Vue == nil || len(in.Vue.Vues) == 0 {
			log.Println("  Pas de vue dans le dernier fichier json du braldun", idBraldun)
			continue
		}
		in.Vue.Vues[0].Time = plusRécentDate
		mdb.vues[uint(idBraldun)] = in.Vue.Vues[0]
		log.Printf("  Données vue chargées pour %d\n", idBraldun)
	}
	return nil
}
