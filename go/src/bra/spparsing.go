package bra

// parsing de fichier tel que produit dynamiquement par un script public de Braldahim
// Voir sp.braldahim.com

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func readLine(r *bufio.Reader) (cells []string, err error) {
	var line string
	if line, err = r.ReadString('\n'); err == nil {
		line = line[0 : len(line)-1]
		cells = strings.Split(line, ";")
	}
	return
}

// remplit l'objet Vue et optionnellement (si elle est fournie) la MemMap
//  à partir d'une ligne d'un flux csv
func ParseLigneFichierDynamique(cells []string, vue *Vue, memmap *MemMap, verbose bool) {
	if len(cells) < 3 {
		fmt.Println("  Ligne trop courte : ", cells)
		return
	}
	var err error
	displayErrors := true
	switch cells[0] {
	case "ALIMENT":
		o := new(VueObjet)
		if err = o.ReadCsvAliment(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "BALLON_SOULE":
		o := new(VueObjet)
		if err = o.ReadCsvSimple(cells, "ballon", "Ballon de soule"); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "BOSQUET":
		o := new(VueBosquet)
		if err = o.ReadCsv(cells); err == nil {
			if memmap != nil {
				memmap.StoreBosquet(o)
			}
		}
	case "BRALDUN":
		o := new(Braldun)
		if err = o.ReadCsvDynamique(cells); err == nil {
			vue.Bralduns = append(vue.Bralduns, o)
		}
	case "BUISSON":
		o := new(VueObjet)
		if err = o.ReadCsvSimpleLabel(cells, "buisson"); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "CADAVRE":
		o := new(VueCadavre)
		if err = o.ReadCsv(cells); err == nil {
			vue.Cadavres = append(vue.Cadavres, o)
		}
	case "CHAMP":
		o := new(VueChamp)
		if err = o.ReadCsv(cells); err == nil {
			if memmap != nil {
				memmap.StoreChamp(o)
			}
		}
	case "CHARRETTE":
		o := new(VueObjet)
		if err = o.ReadCsvSimpleLabel(cells, "charrette"); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "CREVASSE":
		o := new(VueCrevasse)
		if err = o.ReadCsv(cells); err == nil {
			if memmap != nil {
				memmap.StoreCrevasse(o)
			}
		}
	case "ECHOPPE":
		o := new(VueEchoppe)
		if err = o.ReadCsv(cells); err == nil {
			if memmap != nil {
				memmap.StoreEchoppe(o)
			}
		}
	case "ELEMENT":
		o := new(VueObjet)
		if err = o.ReadCsvElement(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "ENVIRONNEMENT":
		o := new(VueEnvironnement)
		if err = o.ReadCsv(cells); err == nil {
			if memmap != nil {
				memmap.StoreEnvironnement(o)
			}
		}
	case "GRAINE":
		o := new(VueObjet)
		if err = o.ReadCsvGraine(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "INGREDIENT":
		o := new(VueObjet)
		if err = o.ReadCsvQLB(cells, "ingrédient"); err == nil {
			vue.Objets = append(vue.Objets, o)
		} else {
			displayErrors = verbose
		}
	case "LIEU":
		o := new(VueLieu)
		if err = o.ReadCsv(cells); err == nil {
			if memmap != nil {
				memmap.StoreLieu(o)
			}
		} else {
			displayErrors = verbose
		}
	case "LINGOT":
		o := new(VueObjet)
		if err = o.ReadCsvLingot(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "MINERAI_BRUT":
		o := new(VueObjet)
		if err = o.ReadCsvMinerai(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "MONSTRE":
		o := new(VueMonstre)
		if err = o.ReadCsv(cells); err == nil {
			vue.Monstres = append(vue.Monstres, o)
		} else {
			displayErrors = verbose
		}
	case "MUNITION":
		o := new(VueObjet)
		if err = o.ReadCsvMunition(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "NID":
		o := new(VueObjet)
		if err = o.ReadCsvSimpleLabel(cells, "nid"); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "PALISSADE":
		o := new(VuePalissade)
		if err = o.ReadCsv(cells, false); err == nil {
			if memmap != nil {
				memmap.StorePalissade(o)
			}
		}
	case "PLANTE_BRUTE":
		o := new(VueObjet)
		if err = o.ReadCsvPlante(cells, true); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "PLANTE_PREPAREE":
		o := new(VueObjet)
		if err = o.ReadCsvPlante(cells, false); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "POTION":
		o := new(VueObjet)
		if err = o.ReadCsvPotion(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "PORTAIL":
		o := new(VuePalissade)
		if err = o.ReadCsv(cells, true); err == nil {
			if memmap != nil {
				memmap.StorePalissade(o)
			}
		}
	case "POSITION":
		o := new(VuePosition)
		if err = o.ReadCsv(cells); err == nil {
			vue.Z = o.Z
			vue.Voyeur = o.IdBraldun
			vue.XMin = o.XMin
			vue.XMax = o.XMax
			vue.YMin = o.YMin
			vue.YMax = o.YMax
		}
	case "ROUTE":
		o := new(VueRoute)
		if err = o.ReadCsv(cells); err == nil {
			if memmap != nil {
				memmap.StoreRoute(o)
			}
		}
	case "RUNE":
		o := new(VueObjet)
		if err = o.ReadCsvRune(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	case "TABAC":
		o := new(VueObjet)
		if err = o.ReadCsvTabac(cells); err == nil {
			vue.Objets = append(vue.Objets, o)
		}
	}
	if displayErrors && err != nil {
		fmt.Printf("Erreur lecture : %+v \n cellules : %+v\n", err, cells)
	}
}

// remplit l'objet Vue et optionnellement (si elle est fournie) la MemMap
//  à partir d'un flux csv.
func ParseFichierDynamique(r *bufio.Reader, time int64, memmap *MemMap, verbose bool) (vue *Vue) {
	vue = NewVue()
	vue.Time = time
	for {
		line, err := readLine(r)
		if err != nil {
			if err != io.EOF {
				log.Println("Error in ParseFichierDynamique :", err)
			}
			return
		}
		ParseLigneFichierDynamique(line, vue, memmap, verbose)
	}
	return
}
