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
	"path/filepath"
	"flag"
	"fmt"
	"json"
	"log"
	"os"
	"runtime/pprof"
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

func readLine(r *bufio.Reader) (cells []string, err os.Error) {
	var line string
	if line, err = r.ReadString('\n'); err == nil {
		line = line[0 : len(line)-1]
		cells = strings.Split(line, ";")
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
		s += " CEST"
		t, err := time.Parse("2006-1-2-15-04 MST", s)
		if err != nil {
			//~ fmt.Printf("Erreur parsing date \"%s\" : %+v\n", s, err)
			return 0
		}
		//~ println(s, " -> ", t.Seconds())
		return t.Seconds()
	} else {
		//~ fmt.Printf("readTimeFromFilePath : indices non trouvés (i1=%d, i2=%d)\n", i1, i2)
	}
	return 0
}

func IndexOfStringIn(s string, a []string) int {
	for i, t := range a {
		if s == t {
			return i
		}
	}
	return -1
}

// implémentation de filepath.Visitor
func (ls *LecteurScripts) VisitFile(path string, f *os.FileInfo) {
	pathToken := strings.Split(path, "/")
	if strings.HasSuffix(path, ".csv") {
		f, err := os.Open(path)
		if err != nil {
			return
		}
		defer f.Close()
		r := bufio.NewReader(f)
		switch pathToken[len(pathToken)-1] {
		case "bralduns.csv":
			ls.parseFichierStatique(r, func() Visible { return new(Braldun) })
		case "communautes.csv":
			ls.parseFichierStatique(r, func() Visible { return new(Communauté) })
		case "lieux_villes.csv":
			ls.parseFichierStatique(r, func() Visible { return new(LieuVille) })
		case "regions.csv":
			ls.parseFichierStatique(r, func() Visible { return new(Région) })
		case "villes.csv":
			ls.parseFichierStatique(r, func() Visible { return new(Ville) })
		default:
			vue := ls.parseFichierDynamique(r, ls.readTimeFromFilePath(pathToken))
			if vue != nil {
				//~ fmt.Printf("    -> vue : %+v\n", vue)
				if vue.Voyeur > 0 && vue.Time > 0 {
					if ls.MemMap.DernièresVues[vue.Voyeur] == nil || vue.Time > ls.MemMap.DernièresVues[vue.Voyeur].Time {
						ls.MemMap.DernièresVues[vue.Voyeur] = vue
					}
				}
			}
		}
	}
}

// implémentation de filepath.Visitor
func (ls *LecteurScripts) VisitDir(path string, f *os.FileInfo) bool {
	pathToken := strings.Split(path, "/")
	indexTokenPrivate := IndexOfStringIn("private", pathToken)
	if indexTokenPrivate == -1 || indexTokenPrivate+1 == len(pathToken) {
		return true
	}
	if IndexOfStringIn(pathToken[indexTokenPrivate+1], ls.IdBralduns) == -1 {
		if ls.verbose > 0 {
			fmt.Printf("Fichier non autorisé : %s\n", path)
		}
		return false
	}
	return true
}

/*
 * Paramètres :
 *  - chemin d'un répertoire contenant dans des sous répertoires les fichiers csv
 *  - liste des id des bralduns séparés par des virgules
 *  - chemin du répertoire dans lequel écrire les fichiers de sortie
 */
func main() {
	cheminFichiersCsv := flag.String("in", "", "répertoire des fichiers csv")
	cheminRepertoireExport := flag.String("out", "", "répertoire d'export")
	idBraldunsBruts := flag.String("bralduns", "", "ids des bralduns, séparés par des virgules")
	cpuprofile := flag.String("cpuprofile", "", "fichier dans lequel écrire un bilan de profiling cpu")
	flag.Parse()
	
	if *cpuprofile != "" {
		fmt.Println("Profiling actif, résultats dans le fichier ", *cpuprofile)
        fp, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(fp)
        defer pprof.StopCPUProfile()
    }
	
	ls := NewLecteurScripts()

	startTime := time.Seconds()
	ls.IdBralduns = strings.Split(*idBraldunsBruts, ",")

	//> lecture des fichiers csv
	filepath.Walk(*cheminFichiersCsv, func(path string, info *os.FileInfo, err os.Error) os.Error {
		if info.IsDirectory() {
			if !ls.VisitDir(path, info) {
				return filepath.SkipDir
			}
		} else {
			ls.VisitFile(path, info)
		}
		return nil
	})

	//> compilation de la carte
	carte := ls.MemMap.Compile()

	//> export de la carte compilée
	cheminFichierJson := *cheminRepertoireExport + "/carte.json"
	f, err := os.Create(cheminFichierJson)
	if err != nil {
		fmt.Printf("Erreur à la création du fichier : %s", cheminFichierJson)
		return
	}
	defer f.Close()
	bout, _ := json.Marshal(carte)
	f.Write(bout)

	//> export des images PNG
	for _, couche := range carte.Couches {
		couche.ConstruitPNG(*cheminRepertoireExport)
	}

	//> affichage d'un petit bilan
	fmt.Printf("Fichiers lus : %d\nCarte compilée dans %s en %d secondes\n\n", ls.NbReadFiles, *cheminRepertoireExport, time.Seconds()-startTime)

}
