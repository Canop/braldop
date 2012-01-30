package main

import (
	"braldop/bra"
	"encoding/json"
	"fmt"
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
	return fusion.Vues[i].Time>fusion.Vues[j].Time
}
func (fusion FusionVues) Swap(i, j int) {
	fusion.Vues[i], fusion.Vues[j] = fusion.Vues[j], fusion.Vues[i]
}

type FusionneurVue struct {
	vues map[uint]*bra.Vue // clef : id braldun
}

func (fv *FusionneurVue) Reçoit(vue *bra.Vue) {
	vue.Time = time.Seconds()
	fv.vues[vue.Voyeur] = vue
}



// renvoie les vues
func (fv *FusionneurVue) Complète(vue *bra.Vue, amis []*bra.CompteBraldop) []*bra.Vue {
	fusion := new(FusionVues)
	fusion.Vues = make([]*bra.Vue,1,len(amis))
	fusion.Vues[0] = vue
	for _, ami := range(amis) {
		if dvb, ok := fv.vues[ami.IdBraldun];ok {
			fusion.Vues = append(fusion.Vues, dvb)
		}
	}
	for _, v := range(fusion.Vues) {
		fmt.Println("fusion vue ", v.Voyeur, " time : ", v.Time)
	}
	sort.Sort(fusion) // index faible : plus récent
	// on construit un nouveau tableau des vues, épurées des doublons et tenant compte des dates de MAJ
	bralduns := make(map[uint]*bra.Braldun)
	monstres := make(map[uint]*bra.VueMonstre)
	l := len(fusion.Vues)
	vues := make([]*bra.Vue, l, l)
	for i, vi := range(fusion.Vues) {
		v := bra.NewVue()
		v.Z = fusion.Vues[i].Z
		v.Time = fusion.Vues[i].Time
		v.Voyeur = fusion.Vues[i].Voyeur
		v.PrénomVoyeur = fusion.Vues[i].PrénomVoyeur
		v.XMin = fusion.Vues[i].XMin
		v.XMax = fusion.Vues[i].XMax
		v.YMin = fusion.Vues[i].YMin
		v.YMax = fusion.Vues[i].YMax
		vues[i]=v
		intersectvues := make([]*bra.Vue, 0, i)
		for j:=0; j<i; j++ {
			vj := fusion.Vues[j]
			if vi.Z==vj.Z && vi.XMin<=vj.XMax && vj.XMin<=vi.XMax && vi.YMin<=vj.YMax && vj.YMin<=vi.YMax {
				intersectvues = append(intersectvues, vj)
			}
		}
		for _, b := range(vi.Bralduns) {
			if b.Id==v.Voyeur {
				v.PrénomVoyeur = b.Prénom
			}
			if _, ok := bralduns[b.Id]; !ok { // si ce braldun n'est pas dans une vue plus récente
				bralduns[b.Id] = b
				if bra.PointEnDehors(b.X, b.Y, intersectvues) {
					v.Bralduns = append(v.Bralduns, b)
				}
			}
		}
		for _, m := range(vi.Monstres) {
			if _, ok := monstres[m.Id]; !ok {
				monstres[m.Id] = m
				if bra.PointEnDehors(m.X, m.Y, intersectvues) {
					v.Monstres = append(v.Monstres, m)
				}
			}
		}
		for _, c := range(vi.Cadavres) {
			if bra.PointEnDehors(c.X, c.Y, intersectvues) {
				v.Cadavres = append(v.Cadavres, c)
			}
		}
		for _, o := range(vi.Objets) {
			if bra.PointEnDehors(o.X, o.Y, intersectvues) {
				v.Objets = append(v.Objets, o)
			}
		}
	}
	return vues
}

// appelée au lancement cette fonction charge les donnéesVueBraldun
//  à partir des fichiers json (le dernier de chaque utilisateur).
// Notons qu'on ne vérifie pas ici les mdpr
func (fv *FusionneurVue) Charge(répertoireCartes string) error {
	fv.vues = make(map[uint]*bra.Vue)
	rep, err := os.Open(répertoireCartes)
	if err!=nil {
		return err
	}
	defer rep.Close()
	cfis, _ := rep.Readdir(-1)
	for _, cfi := range(cfis) {
		t := strings.Split(cfi.Name, "-")
		if len(t)!=2 {
			continue
		}
		idBraldun, _ := strconv.Atoui(t[0])
		if idBraldun==0 {
			continue
		}
		plusRécentChemin := ""
		plusRécentDate := int64(0)
		// on parcoure toute l'arborescence, pas le plus efficace mais le plus simple à coder
		filepath.Walk(filepath.Join(répertoireCartes, cfi.Name), func(path string, info *os.FileInfo, err error) error {
			if info.IsRegular() && strings.HasPrefix(info.Name, "carte-") && strings.HasSuffix(info.Name, ".json") && info.Mtime_ns>plusRécentDate {
				plusRécentChemin = path
				plusRécentDate = info.Mtime_ns
			}
			return nil
		});
		if plusRécentChemin=="" {
			log.Println("  Pas de fichier json pour", idBraldun)
			continue
		}
		fjson, err := os.Open(plusRécentChemin)
		if err!=nil {
			log.Println(err)
			continue
		}
		in := new(MessageIn)
		bin, err := ioutil.ReadAll(fjson)
		defer fjson.Close()
		if err!=nil {
			log.Println("  Erreur à la lecture du dernier fichier json du braldun", idBraldun)
			continue
		}
		err = json.Unmarshal(bin, in)
		if err!=nil {
			log.Println("  Erreur à la lecture du dernier fichier json du braldun", idBraldun)
			continue
		}
		if in.Vue==nil || len(in.Vue.Vues)==0 {
			log.Println("  Pas de vue dans le dernier fichier json du braldun", idBraldun)
			continue
		}
		in.Vue.Vues[0].Time = plusRécentDate/1e9
		fv.vues[idBraldun] = in.Vue.Vues[0]
		log.Println("  Données vue chargées pour", idBraldun)
	}
	return nil
}
