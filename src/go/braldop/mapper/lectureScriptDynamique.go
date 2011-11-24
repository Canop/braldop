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
	switch cells[0] {
	case "ALIMENT":
		o := new(bra.VueObjet)
		if err := o.ReadCsvAliment(cells); err != nil {
			fmt.Printf("Erreur lecture ALIMENT : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" ALIMENT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "BALLON_SOULE":
		o := new(bra.VueObjet)
		if err := o.ReadCsvSimple(cells, "ballon", "Ballon de soule"); err != nil {
			if ls.verbose > 0 {
				fmt.Printf("Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
			}
		} else {
			//~ fmt.Printf(" BALLON_SOULE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "BOSQUET":
		o := new(bra.VueBosquet)
		if err := o.ReadCsv(cells); err != nil {
			if ls.verbose > 0 {
				fmt.Printf("Erreur lecture VueBosquet : %+v \n cellules : %+v\n", err, cells)
			}
		} else {
			//~ fmt.Printf(" VueBosquet : %+v\n", o)
			ls.MemMap.StoreBosquet(o)
		}
	case "BRALDUN":
		o := new(bra.Braldun)
		if err := o.ReadCsvDynamique(cells); err != nil {
			fmt.Printf("Erreur lecture Braldun : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" Braldun : %+v\n", o)
			vue.Bralduns = append(vue.Bralduns, o)
		}
	case "BUISSON":
		o := new(bra.VueObjet)
		if err := o.ReadCsvSimpleLabel(cells, "buisson"); err != nil {
			fmt.Printf("Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" BUISSON : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "CADAVRE":
		o := new(bra.VueCadavre)
		if err := o.ReadCsv(cells); err != nil {
			fmt.Printf("Erreur lecture VueCadavre : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" VueCadavre : %+v\n", o)
			vue.Cadavres = append(vue.Cadavres, o)
		}
	case "CHAMP":
		o := new(bra.VueChamp)
		if err := o.ReadCsv(cells); err != nil {
			fmt.Printf("Erreur lecture VueChamp : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" VueChamp : %+v\n", o)
			ls.MemMap.StoreChamp(o)
		}
	case "CHARRETTE":
		o := new(bra.VueObjet)
		if err := o.ReadCsvSimpleLabel(cells, "charrette"); err != nil {
			fmt.Printf("Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" CHARRETTE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "CREVASSE":
		o := new(bra.VueCrevasse)
		if err := o.ReadCsv(cells); err != nil {
			fmt.Printf("Erreur lecture VueCrevasse : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" VueCrevasse : %+v\n", o)
			ls.MemMap.StoreCrevasse(o)
		}
	case "ECHOPPE":
		o := new(bra.VueEchoppe)
		if err := o.ReadCsv(cells); err != nil {
			fmt.Printf("Erreur lecture VueEchoppe : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" VueEchoppe : %+v\n", o)
			ls.MemMap.StoreEchoppe(o)
		}
	case "ELEMENT":
		o := new(bra.VueObjet)
		if err := o.ReadCsvElement(cells); err != nil {
			fmt.Printf("Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" ELEMENT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "ENVIRONNEMENT":
		o := new(bra.VueEnvironnement)
		if err := o.ReadCsv(cells); err != nil {
			fmt.Printf("Erreur lecture VueEnvironnement : %+v\n", err)
		} else {
			//~ fmt.Printf(" VueEnvironnement : %+v\n", o)
			ls.MemMap.StoreEnvironnement(o)
		}
	case "GRAINE":
		o := new(bra.VueObjet)
		if err := o.ReadCsvGraine(cells); err != nil {
			fmt.Printf("Erreur lecture GRAINE : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" GRAINE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "INGREDIENT":
		o := new(bra.VueObjet)
		if err := o.ReadCsvQLB(cells, "ingrédient"); err != nil {
			if ls.verbose > 0 {
				fmt.Printf("Erreur lecture INGREDIENT : %+v \n cellules : %+v\n", err, cells)
			}
		} else {
			//~ fmt.Printf(" INGREDIENT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "LIEU":
		o := new(bra.VueLieu)
		if err := o.ReadCsv(cells); err != nil {
			if ls.verbose > 0 {
				fmt.Printf("Erreur lecture LIEU : %+v \n cellules : %+v\n", err, cells)
			}
		} else {
			//~ fmt.Printf(" LIEU : %+v\n", o)
			ls.MemMap.StoreLieu(o)
		}
	case "LINGOT":
		o := new(bra.VueObjet)
		if err := o.ReadCsvLingot(cells); err != nil {
			fmt.Printf("Erreur lecture LINGOT : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" LINGOT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "MINERAI_BRUT":
		o := new(bra.VueObjet)
		if err := o.ReadCsvMinerai(cells); err != nil {
			fmt.Printf("Erreur lecture MINERAI_BRUT : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" MINERAI_BRUT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "MONSTRE":
		o := new(bra.VueMonstre)
		if err := o.ReadCsv(cells); err != nil {
			if ls.verbose > 0 {
				fmt.Printf("Erreur lecture VueMonstre : %+v \n cellules : %+v\n", err, cells)
			}
		} else {
			//~ fmt.Printf(" VueMonstre : %+v\n", o)
			vue.Monstres = append(vue.Monstres, o)
		}
	case "MUNITION":
		o := new(bra.VueObjet)
		if err := o.ReadCsvMunition(cells); err != nil {
			fmt.Printf("Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" MUNITION : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "NID":
		o := new(bra.VueObjet)
		if err := o.ReadCsvSimpleLabel(cells, "nid"); err != nil {
			fmt.Printf("Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" NID : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "PALISSADE":
		o := new(bra.VuePalissade)
		if err := o.ReadCsv(cells, false); err != nil {
			fmt.Printf("Erreur lecture VuePalissade : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" VuePalissade : %+v\n", o)
			ls.MemMap.StorePalissade(o)
		}
	case "PLANTE_BRUTE":
		o := new(bra.VueObjet)
		if err := o.ReadCsvPlante(cells, true); err != nil {
			fmt.Printf("Erreur lecture PLANTE_BRUTE : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" PLANTE_BRUTE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "PLANTE_PREPAREE":
		o := new(bra.VueObjet)
		if err := o.ReadCsvPlante(cells, false); err != nil {
			fmt.Printf("Erreur lecture PLANTE_PREPAREE : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" PLANTE_PREPAREE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "POTION":
		o := new(bra.VueObjet)
		if err := o.ReadCsvPotion(cells); err != nil {
			fmt.Printf("Erreur lecture POTION : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" POTION : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "PORTAIL":
		o := new(bra.VuePalissade)
		if err := o.ReadCsv(cells, true); err != nil {
			fmt.Printf("Erreur lecture VuePalissade : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" VuePalissade : %+v\n", o)
			ls.MemMap.StorePalissade(o)
		}
	case "POSITION":
		o := new(bra.VuePosition)
		if err := o.ReadCsv(cells); err != nil {
			fmt.Printf("Erreur lecture position : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" Position : %+v\n", o)
			vue.Z = o.Z
			vue.Voyeur = o.IdBraldun
			vue.XMin = o.XMin
			vue.XMax = o.XMax
			vue.YMin = o.YMin
			vue.YMax = o.YMax
		}
	case "ROUTE":
		o := new(bra.VueRoute)
		if err := o.ReadCsv(cells); err != nil {
			fmt.Printf("Erreur lecture Route : %+v \n cellules : %+v\n", err, cells)
		} else {
			//fmt.Printf(" Route : %+v\n", o)
			ls.MemMap.StoreRoute(o)
		}
	case "RUNE":
		o := new(bra.VueObjet)
		if err := o.ReadCsvRune(cells); err != nil {
			fmt.Printf("Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" RUNE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "TABAC":
		o := new(bra.VueObjet)
		if err := o.ReadCsvTabac(cells); err != nil {
			fmt.Printf("Erreur lecture TABAC : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" TABAC : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
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
