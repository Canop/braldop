package main

/*******************************************************************
 * 
 * 
 * 
 * TODO
 * 	- utiliser la position pour effacer les objets qui ont disparu
 *    (échopes, etc.)
 * 
 * 
 * 
 *******************************************************************/

import (
	"bufio"
	"fmt"
	"json"
	"os"
	"strings"
	"time"
)

type LecteurScripts struct {
	NbReadFiles uint
	MemMap      *MemMap
	IdBralduns  []string // la liste des bralduns dont on peut regarder la vue
	verbose     uint     // 0 : peu, 1 : un peu, 2 : beaucoup
}

type Visible interface {
	readCsv(cells []string) (err os.Error)
	store(mm *MemMap)
}

func NewLecteurScripts() (ls *LecteurScripts) {
	ls = new(LecteurScripts)
	ls.MemMap = NewMemMap()
	return ls
}

func readLine(r *bufio.Reader) (line string, err os.Error) {
	if line, err = r.ReadString('\n'); err == nil {
		line = line[0 : len(line)-1]
	}
	return
}

// renvoie le nombre de secondes depuis 1970 caché dans le chemin vers le fichier : année/mois/jour/truc-heurehminutes.csv
// Le parsage des dates est pour moi le gros WTF du go... si quelqu'un arrive à faire plus propre...
func (ls *LecteurScripts) readTimeFromFilePath(path []string) int64 {
	//~ fmt.Printf("readTimeFromFilePath : %+v\n", path)
	l := len(path)
	if l < 4 {
		//~ fmt.Printf("readTimeFromFilePath : chemin trop court (l=%d)\n", l)
		return 0
	}
	s := path[l-4] + "-" + path[l-3] + "-" + path[l-2] // année-mois-jour
	name := path[l-1]
	i1 := strings.LastIndex(name, "h")
	i2 := strings.LastIndex(name, "-")
	if i2 > i1 && i1 > 0 {
		s += "-" + name[0:i1] + "-" + name[i1+1:i2]
		if i2 == i1+1 {
			// cas où on n'a pas les minutes
			s += "00"
		}
		//~ fmt.Println("  date formatée : ", s)
		t, err := time.Parse("2006-1-2-15-04", s)
		if err != nil {
			//~ fmt.Printf("Erreur parsing date \"%s\" : %+v\n", s, err)
			return 0
		}
		return t.Seconds()
	} else {
		//~ fmt.Printf("readTimeFromFilePath : indices non trouvés (i1=%d, i2=%d)\n", i1, i2)
	}
	return 0
}

// fichier ou répertoire
func (ls *LecteurScripts) traiteFichier(f *os.File) os.Error {
	childs, err := f.Readdir(-1)
	if err == nil {
		//fmt.Println("Entering directory " + f.Name())
		for _, fi := range childs {
			err := ls.traiteNomFichier(f.Name() + "/" + fi.Name)
			if err != nil {
				return err
			}
		}
		//fmt.Println("Leaving directory " + f.Name())
	} else {
		path := strings.Split(f.Name(), "/")
		filename := path[len(path)-1]
		indexTokenPrivate := -1
		for i, t := range path {
			if t == "private" {
				indexTokenPrivate = i
				break
			}
		}
		if indexTokenPrivate >= 0 {
			ok := false
			for _, id := range ls.IdBralduns {
				if id == path[indexTokenPrivate+1] {
					ok = true
					break
				}
			}
			if !ok {
				if ls.verbose > 0 {
					fmt.Printf("Fichier non autorisé : %s\n", f.Name())
				}
				return nil
			}
		}
		if strings.HasSuffix(filename, ".csv") {
			//~ fmt.Printf("   parsed file : %s\n", f.Name())
			switch filename {
			case "bralduns.csv":
				return ls.parseFichierStatique(f, func() Visible { return new(Braldun) })
			case "communautes.csv":
				return ls.parseFichierStatique(f, func() Visible { return new(Communauté) })
			case "lieux_villes.csv":
				return ls.parseFichierStatique(f, func() Visible { return new(LieuVille) })
			case "regions.csv":
				return ls.parseFichierStatique(f, func() Visible { return new(Région) })
			case "villes.csv":
				return ls.parseFichierStatique(f, func() Visible { return new(Ville) })
			default:
				vue, err := ls.parseFichierDynamique(f, ls.readTimeFromFilePath(path))
				if err != nil {
					fmt.Printf("erreur parsing fichier dynamique : %+v\n", err)
				} else {
					//~ fmt.Printf("    -> vue : %+v\n", vue)
					if vue.Voyeur > 0 && vue.Time > 0 {
						if ls.MemMap.DernièresVues[vue.Voyeur] == nil || vue.Time > ls.MemMap.DernièresVues[vue.Voyeur].Time {
							ls.MemMap.DernièresVues[vue.Voyeur] = vue
						}
					}
				}
				return err
			}
		} else {
			//~ fmt.Println("   ignored file : " + f.Name())
		}
	}
	return nil
}

// fichier ou répertoire
func (ls *LecteurScripts) traiteNomFichier(filename string) os.Error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return ls.traiteFichier(f)
}

/*
 * Paramètres :
 *  - chemin d'un répertoire contenant dans des sous répertoires les fichiers csv
 *  - liste des id des bralduns séparés par des virgules
 *  - chemin du répertoire dans lequel écrire les fichiers de sortie
 */
func main() {
	if len(os.Args) < 4 {
		fmt.Println(os.EINVAL)
		return
	}
	ls := NewLecteurScripts()

	fmt.Printf("Export : %s\n", os.Args[3])

	ls.IdBralduns = strings.Split(os.Args[2], ",")
	fmt.Printf("Bralduns du groupe : %+v\n", ls.IdBralduns)

	//> lecture des fichiers csv
	cheminRacine := os.Args[1]
	startTime := time.Seconds()
	err := ls.traiteNomFichier(cheminRacine)
	if err != nil {
		fmt.Printf("Erreur à la lecture des fichiers : %v", err)
		return
	}

	//> compilation de la carte
	carte := ls.MemMap.Compile()

	//> export de la carte compilée
	cheminFichierJson := os.Args[3] + "/carte.json"
	f, err := os.Create(cheminFichierJson)
	if err != nil {
		fmt.Printf("Erreur à la création du fichier : %s", cheminFichierJson)
		return
	}
	fmt.Printf("Carte compilée : %s\n", f.Name())
	defer f.Close()
	bout, _ := json.Marshal(carte)
	f.Write(bout)

	//> affichage d'un petit bilan
	fmt.Printf("Fini en %d secondes\n Fichiers lus : %d\n", time.Seconds()-startTime, ls.NbReadFiles)
	fmt.Println()
}
