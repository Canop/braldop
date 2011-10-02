package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (ls *LecteurScripts) parseLigneFichierDynamique(line string, vue *Vue) {
	cells := strings.Split(line, ";")
	//fmt.Println(" cells : " + strings.Join(cells, "#"))
	if len(cells) < 3 {
		fmt.Printf("  Ligne trop courte : %s\n", line)
		return
	}
	switch cells[0] {
	case "ALIMENT":
		o := new(VueObjet)
		if err := o.readCsvAliment(cells); err != nil {
			fmt.Printf(" Erreur lecture ALIMENT : %+v \n cellules : %+v", err, cells)
		} else {
			fmt.Printf(" ALIMENT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "BALLON_SOULE":
		o := new(VueObjet)
		if err := o.readCsvSimple(cells, "ballon", "Ballon de soule"); err != nil {
			if ls.verbose > 0 {
				fmt.Printf(" Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
			}
		} else {
			//~ fmt.Printf(" BALLON_SOULE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "BOSQUET":
		o := new(VueBosquet)
		if err := o.readCsv(cells); err != nil {
			if ls.verbose > 0 {
				fmt.Printf(" Erreur lecture VueBosquet : %+v \n cellules : %+v\n", err, cells)
			}
		} else {
			//~ fmt.Printf(" VueBosquet : %+v\n", o)
			ls.MemMap.StoreBosquet(o)
		}
	case "BRALDUN":
		o := new(Braldun)
		if err := o.readCsvDynamique(cells); err != nil {
			fmt.Printf(" Erreur lecture Braldun : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" Braldun : %+v\n", o)
			vue.Bralduns = append(vue.Bralduns, o)
		}
	case "BUISSON":
		o := new(VueObjet)
		if err := o.readCsvSimpleLabel(cells, "buisson"); err != nil {
			fmt.Printf(" Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" BUISSON : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "CADAVRE":
		o := new(VueCadavre)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueCadavre : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" VueCadavre : %+v\n", o)
			vue.Cadavres = append(vue.Cadavres, o)
		}
	case "CHAMP":
		o := new(VueChamp)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueChamp : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" VueChamp : %+v\n", o)
			ls.MemMap.StoreChamp(o)
		}
	case "CHARRETTE":
		o := new(VueObjet)
		if err := o.readCsvSimpleLabel(cells, "charrette"); err != nil {
			fmt.Printf(" Erreur lecture VueObjet : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" CHARRETTE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "ECHOPPE":
		o := new(VueEchoppe)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueEchoppe : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" VueEchoppe : %+v\n", o)
			ls.MemMap.StoreEchoppe(o)
		}
	case "ELEMENT":
		o := new(VueObjet)
		if err := o.readCsvElement(cells); err != nil {
			fmt.Printf(" Erreur lecture VueObjet : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" ELEMENT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "ENVIRONNEMENT":
		o := new(VueEnvironnement)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueEnvironnement : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VueEnvironnement : %+v\n", o)
			ls.MemMap.StoreEnvironnement(o)
		}
	case "GRAINE":
		o := new(VueObjet)
		if err := o.readCsvGraine(cells); err != nil {
			fmt.Printf(" Erreur lecture GRAINE : %+v \n cellules : %+v", err, cells)
		} else {
			fmt.Printf(" GRAINE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "INGREDIENT":
		o := new(VueObjet)
		if err := o.readCsvQLB(cells, "ingrédient"); err != nil {
			if ls.verbose > 0 {
				fmt.Printf(" Erreur lecture INGREDIENT : %+v \n cellules : %+v", err, cells)
			}
		} else {
			//~ fmt.Printf(" INGREDIENT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "LINGOT":
		o := new(VueObjet)
		if err := o.readCsvLingot(cells); err != nil {
			fmt.Printf(" Erreur lecture LINGOT : %+v \n cellules : %+v", err, cells)
		} else {
			fmt.Printf(" LINGOT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "MINERAI_BRUT":
		o := new(VueObjet)
		if err := o.readCsvMinerai(cells); err != nil {
			fmt.Printf(" Erreur lecture MINERAI_BRUT : %+v \n cellules : %+v", err, cells)
		} else {
			fmt.Printf(" MINERAI_BRUT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "MONSTRE":
		o := new(VueMonstre)
		if err := o.readCsv(cells); err != nil {
			if ls.verbose > 0 {
				fmt.Printf(" Erreur lecture VueMonstre : %+v \n cellules : %+v", err, cells)
			}
		} else {
			//~ fmt.Printf(" VueMonstre : %+v\n", o)
			vue.Monstres = append(vue.Monstres, o)
		}
	case "MUNITION":
		o := new(VueObjet)
		if err := o.readCsvMunition(cells); err != nil {
			fmt.Printf(" Erreur lecture VueObjet : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" MUNITION : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "PALISSADE":
		o := new(VuePalissade)
		if err := o.readCsv(cells, false); err != nil {
			fmt.Printf(" Erreur lecture VuePalissade : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VuePalissade : %+v\n", o)
			ls.MemMap.StorePalissade(o)
		}
	case "PLANTE_BRUTE":
		o := new(VueObjet)
		if err := o.readCsvPlante(cells, true); err != nil {
			fmt.Printf(" Erreur lecture PLANTE_BRUTE : %+v \n cellules : %+v", err, cells)
		} else {
			fmt.Printf(" PLANTE_BRUTE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "PLANTE_PREPAREE":
		o := new(VueObjet)
		if err := o.readCsvPlante(cells, false); err != nil {
			fmt.Printf(" Erreur lecture PLANTE_PREPAREE : %+v \n cellules : %+v", err, cells)
		} else {
			fmt.Printf(" PLANTE_PREPAREE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "POTION":
		o := new(VueObjet)
		if err := o.readCsvPotion(cells); err != nil {
			fmt.Printf(" Erreur lecture POTION : %+v \n cellules : %+v", err, cells)
		} else {
			fmt.Printf(" POTION : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "PORTAIL":
		o := new(VuePalissade)
		if err := o.readCsv(cells, true); err != nil {
			fmt.Printf(" Erreur lecture VuePalissade : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VuePalissade : %+v\n", o)
			ls.MemMap.StorePalissade(o)
		}
	case "POSITION":
		o := new(VuePosition)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture position : %+v \n cellules : %+v", err, cells)
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
		o := new(VueRoute)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture Route : %+v \n cellules : %+v", err, cells)
		} else {
			//fmt.Printf(" Route : %+v\n", o)
			ls.MemMap.StoreRoute(o)
		}
	case "RUNE":
		o := new(VueObjet)
		if err := o.readCsvRune(cells); err != nil {
			fmt.Printf(" Erreur lecture VueObjet : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" RUNE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "TABAC":
		o := new(VueObjet)
		if err := o.readCsvTabac(cells); err != nil {
			fmt.Printf(" Erreur lecture TABAC : %+v \n cellules : %+v", err, cells)
		} else {
			fmt.Printf(" TABAC : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	}
}

func (ls *LecteurScripts) parseFichierDynamique(file *os.File, time int64) (vue *Vue, err os.Error) {
	r := bufio.NewReader(file)
	vue = NewVue()
	vue.Time = time
	line, err := readLine(r)
	for err == nil {
		ls.parseLigneFichierDynamique(line, vue)
		line, err = readLine(r)
	}
	ls.NbReadFiles++
	if err != os.EOF {
		fmt.Println("Error in parsing (parseFichierDynamique) :")
		fmt.Println(err)
		return
	}
	err = nil
	return
}
