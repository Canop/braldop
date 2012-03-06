package bra

// parsing de fichier tel que produit dynamiquement par un script public de Braldahim
// Voir sp.braldahim.com

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Visible interface {
	ReadCsv(cells []string) (err error)
	Store(mm *MemMap)
}

func readLine(r *bufio.Reader) (cells []string, err error) {
	var line string
	if line, err = r.ReadString('\n'); err == nil {
		line = line[0 : len(line)-1]
		cells = strings.Split(line, ";")
	}
	return
}

func skipLines(r *bufio.Reader, nblines int) error {
	for i:=0; i<nblines; i++ {
		_, err := readLine(r)
		if err!=nil {
			return err
		}
	}
	return nil
}

// remplit l'objet Vue et optionnellement (si elle est fournie) la MemMap
//  à partir d'une ligne d'un flux csv
// Renvoie éventuellement un objet VuePosition (oui, ça devient le bordel...)
func ParseLigneFichierDynamique(cells []string, vue *Vue, memmap *MemMap, verbose bool) *VuePosition {
	if len(cells) < 3 {
		fmt.Println("  Ligne trop courte : ", cells)
		return nil
	}
	var pos *VuePosition
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
			pos = o
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
	return pos
}

// remplit l'objet Vue et optionnellement (si elle est fournie) la MemMap
//  à partir d'un flux csv.
func ParseFichierCsvDynamique(r *bufio.Reader, time int64, memmap *MemMap, verbose bool) (vue *Vue, pos *VuePosition) {
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
		pos = ParseLigneFichierDynamique(line, vue, memmap, verbose)
	}
	return
}

func ParseLigneFichierCsvStatique(cells []string, memmap *MemMap, alloue func() Visible) {
	o := alloue()
	if err := o.ReadCsv(cells); err != nil {
		log.Printf(" Erreur lecture : %+v \n cellules : %+v", err, cells)
	} else {
		o.Store(memmap)
	}
}

func ParseFichierCsvStatique(r *bufio.Reader, memmap *MemMap, alloue func() Visible) {
	_, _ = readLine(r) // on saute une ligne car la première contient les en-têtes
	for {
		line, err := readLine(r)
		if err != nil {
			if err != io.EOF {
				log.Println("Error in parsing (parseFichierDynamique) :", err)
			}
			return
		}
		ParseLigneFichierCsvStatique(line, memmap, alloue)
	}
}


// inputs :
//   - le fichier bralduns.csv obtenu par script public
//   - le fichier communautes.csv obtenu par script public
//   - un fichier de vue de script public
//   
func ConstruitDonnéesVue(spvue *bufio.Reader, spbralduns *bufio.Reader, spcommunautes *bufio.Reader, verbose bool) *DonnéesVue {
	time := time.Now().Unix()
	memmap := NewMemMap()
	ParseFichierCsvStatique(spbralduns, memmap, func() Visible { return new(Braldun) })
	ParseFichierCsvStatique(spcommunautes, memmap, func() Visible { return new(Communauté) })
	dv := new(DonnéesVue)
	var vue *Vue
	vue, dv.Position = ParseFichierCsvDynamique(spvue, time, memmap, verbose)
	carte := memmap.Compile()
	dv.Couches = carte.Couches
	memmap.CompleteBralduns(vue)
	dv.Vues = []*Vue{vue}
	return dv
}

func ChargeDonnéesStatiquesPubliques(cheminRépertoireCsvPublic string, verbose bool) (*MemMap, error) {
	memmap := NewMemMap()
	f, err := os.Open(filepath.Join(cheminRépertoireCsvPublic, "bralduns.csv"))
	if err != nil {
		log.Println(" Erreur à l'ouverture du fichier bralduns.csv :", err)
		return nil, err
	}
	ParseFichierCsvStatique(bufio.NewReader(f), memmap, func() Visible { return new(Braldun) })
	f.Close()
	f, err = os.Open(filepath.Join(cheminRépertoireCsvPublic, "communautes.csv"))
	if err != nil {
		log.Println(" Erreur à l'ouverture du fichier communautes.csv :", err)
		return nil, err
	}
	ParseFichierCsvStatique(bufio.NewReader(f), memmap, func() Visible { return new(Communauté) })
	f.Close()
	f, err = os.Open(filepath.Join(cheminRépertoireCsvPublic, "lieux_villes.csv"))
	if err != nil {
		log.Println(" Erreur à l'ouverture du fichier lieux_villes.csv :", err)
		return nil, err
	}
	ParseFichierCsvStatique(bufio.NewReader(f), memmap, func() Visible { return new(LieuVille) })
	f.Close()
	return memmap, nil
}

// le reader doit correspondre à un fichier csv de type profil
func LitEtatBraldunDansCsvProfil(r *bufio.Reader) (*EtatBraldun, error) {
	err := skipLines(r, 2) // on saute les deux lignes d'en-têtes
	if err!=nil {
		return nil, err
	}
	line, err := readLine(r)
	if err!=nil {
		return nil, err
	}
	eb := new(EtatBraldun)
	eb.IdBraldun = AtoId(line[0])
	eb.PA, _ = strconv.Atoi(line[6])
	cd := strings.Split(line[8], ":")
	duréeTour, err := time.ParseDuration(cd[0]+"h"+cd[1]+"m"+cd[2]+"s")
	if err!=nil {
		return nil, err
	}
	eb.DuréeTour = int(duréeTour.Seconds())
	dla, err := time.Parse("2006-1-2 15:04:05 MST", line[9] + " CEST")
	if err!=nil {
		return nil, err
	}
	eb.DLA = int(dla.Unix())
	eb.PV, _ = strconv.Atoi(line[15])
	bmPVmax, _  := strconv.Atoi(line[16])
	niveauVigueur, _ := strconv.Atoi(line[20])
	eb.PVMax = niveauVigueur*10 + 40 + bmPVmax
	return eb, nil
}
