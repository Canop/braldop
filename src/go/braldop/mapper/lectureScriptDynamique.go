package main

import (
	"braldop/bra"
	"bufio"
	"fmt"
	"io"
)

func (ls *LecteurScripts) parseLigneFichierDynamique(cells []string, vue *bra.Vue) {
	if len(cells) < 3 {
		fmt.Println("  Ligne trop courte : ", cells)
		return
	}
	var err error
	displayErrors := true
	switch cells[0] {
	case "ALIMENT":
		o := new(bra.VueObjet)
		if err = o.ReadCsvAliment(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "BALLON_SOULE":
		o := new(bra.VueObjet)
		if err = o.ReadCsvSimple(cells, "ballon", "Ballon de soule"); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "BOSQUET":
		o := new(bra.VueBosquet)
		if err = o.ReadCsv(cells); err == nil {
			ls.MemMap.StoreBosquet(o)
		}
	case "BRALDUN":
		o := new(bra.Braldun)
		if err = o.ReadCsvDynamique(cells); err == nil {
			vue.Bralduns = append(vue.Bralduns, o)
		}
	case "BUISSON":
		o := new(bra.VueObjet)
		if err = o.ReadCsvSimpleLabel(cells, "buisson"); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "CADAVRE":
		o := new(bra.VueCadavre)
		if err = o.ReadCsv(cells); err == nil {
			vue.Cadavres = append(vue.Cadavres, o)
		}
	case "CHAMP":
		o := new(bra.VueChamp)
		if err = o.ReadCsv(cells); err == nil {
			ls.MemMap.StoreChamp(o)
		}
	case "CHARRETTE":
		o := new(bra.VueObjet)
		if err = o.ReadCsvSimpleLabel(cells, "charrette"); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "CREVASSE":
		o := new(bra.VueCrevasse)
		if err = o.ReadCsv(cells); err == nil {
			ls.MemMap.StoreCrevasse(o)
		}
	case "ECHOPPE":
		o := new(bra.VueEchoppe)
		if err = o.ReadCsv(cells); err == nil {
			ls.MemMap.StoreEchoppe(o)
		}
	case "ELEMENT":
		o := new(bra.VueObjet)
		if err = o.ReadCsvElement(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "ENVIRONNEMENT":
		o := new(bra.VueEnvironnement)
		if err = o.ReadCsv(cells); err == nil {
			ls.MemMap.StoreEnvironnement(o)
		}
	case "GRAINE":
		o := new(bra.VueObjet)
		if err = o.ReadCsvGraine(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "INGREDIENT":
		o := new(bra.VueObjet)
		if err = o.ReadCsvQLB(cells, "ingrédient"); err == nil {
			vue.Objets = append(vue.Objets, o)
		} else {
			displayErrors = ls.verbose
		}
	case "LIEU":
		o := new(bra.VueLieu)
		if err = o.ReadCsv(cells); err == nil {
			ls.MemMap.StoreLieu(o)
		} else {
			displayErrors = ls.verbose
		}
	case "LINGOT":
		o := new(bra.VueObjet)
		if err = o.ReadCsvLingot(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "MINERAI_BRUT":
		o := new(bra.VueObjet)
		if err = o.ReadCsvMinerai(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "MONSTRE":
		o := new(bra.VueMonstre)
		if err = o.ReadCsv(cells); err == nil {
			vue.Monstres = append(vue.Monstres, o)
		} else {
			displayErrors = ls.verbose
		}
	case "MUNITION":
		o := new(bra.VueObjet)
		if err = o.ReadCsvMunition(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "NID":
		o := new(bra.VueObjet)
		if err = o.ReadCsvSimpleLabel(cells, "nid"); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "PALISSADE":
		o := new(bra.VuePalissade)
		if err = o.ReadCsv(cells, false); err == nil {
			ls.MemMap.StorePalissade(o)
		}
	case "PLANTE_BRUTE":
		o := new(bra.VueObjet)
		if err = o.ReadCsvPlante(cells, true); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "PLANTE_PREPAREE":
		o := new(bra.VueObjet)
		if err = o.ReadCsvPlante(cells, false); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "POTION":
		o := new(bra.VueObjet)
		if err = o.ReadCsvPotion(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "PORTAIL":
		o := new(bra.VuePalissade)
		if err = o.ReadCsv(cells, true); err == nil {
			ls.MemMap.StorePalissade(o)
		}
	case "POSITION":
		o := new(bra.VuePosition)
		if err = o.ReadCsv(cells); err == nil {
			vue.Z = o.Z
			vue.Voyeur = o.IdBraldun
			vue.XMin = o.XMin
			vue.XMax = o.XMax
			vue.YMin = o.YMin
			vue.YMax = o.YMax
		}
	case "ROUTE":
		o := new(bra.VueRoute)
		if err = o.ReadCsv(cells); err == nil {
			ls.MemMap.StoreRoute(o)
		}
	case "RUNE":
		o := new(bra.VueObjet)
		if err = o.ReadCsvRune(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "TABAC":
		o := new(bra.VueObjet)
		if err = o.ReadCsvTabac(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	}
	if displayErrors && err!=nil {
		fmt.Printf("Erreur lecture : %+v \n cellules : %+v\n", err, cells)
	}
}

func (ls *LecteurScripts) parseFichierDynamique(r *bufio.Reader, time int64) (vue *bra.Vue) {
	vue = bra.NewVue()
	vue.Time = time
	ls.NbReadFiles++
	for {
		line, err := readLine(r)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error in parsing (parseFichierDynamique) :")
				fmt.Println(err)
			}
			return
		}
		ls.parseLigneFichierDynamique(line, vue)
	}
	return
}
