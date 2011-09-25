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
		if strings.HasSuffix(filename, ".csv") {
			//~ fmt.Printf("   parsed file : %s\n", f.Name())
			switch filename {
			case "bralduns.csv":
				return ls.parseFichierStatique(f, func() Visible { return new(Braldun) })
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
 *  - chemin d'un répertoire contenant dans des sous répertoires
 */
func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.EINVAL)
		return
	}

	//> lecture des fichiers csv
	cheminRacine := os.Args[1]
	startTime := time.Seconds()
	ls := NewLecteurScripts()
	err := ls.traiteNomFichier(cheminRacine)
	if err != nil {
		fmt.Printf("Erreur à la lecture des fichiers : %v", err)
		return
	}

	//> compilation de la carte
	m := ls.MemMap.Compile()

	//> export de la carte compilée
	cheminFichierJson := cheminRacine + "/carte.json"
	f, err := os.Create(cheminFichierJson)
	if err != nil {
		fmt.Printf("Erreur à la création du fichier : %s", cheminFichierJson)
		return
	}
	defer f.Close()
	bout, _ := json.Marshal(m)
	f.Write(bout)

	//> affichage d'un petit bilan
	fmt.Printf("Fini en %d secondes\n Fichiers lus : %d\n", time.Seconds()-startTime, ls.NbReadFiles)
	fmt.Println()
}
