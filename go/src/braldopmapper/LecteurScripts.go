package main

import (
	"bra"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strings"
	"time"
)

type LecteurScripts struct {
	NbReadFiles uint
	MemMap      *bra.MemMap
	IdBralduns  []string // la liste des bralduns dont on peut regarder la vue
	verbose     bool
}



func NewLecteurScripts() (ls *LecteurScripts) {
	ls = new(LecteurScripts)
	ls.MemMap = bra.NewMemMap()
	return ls
}

func readLine(r *bufio.Reader) (cells []string, err error) {
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
	l := len(path)
	if l < 4 {
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
		s += " CEST"
		t, err := time.Parse("2006-1-2-15-04 MST", s)
		if err != nil {
			//~ fmt.Printf("Erreur parsing date \"%s\" : %+v\n", s, err)
			return 0
		}
		return t.Unix()
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
func (ls *LecteurScripts) VisitFile(path string, f os.FileInfo) {
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
			bra.ParseFichierCsvStatique(r, ls.MemMap, func() bra.Visible { return new(bra.Braldun) })
		case "communautes.csv":
			bra.ParseFichierCsvStatique(r, ls.MemMap, func() bra.Visible { return new(bra.Communauté) })
		case "lieux_villes.csv":
			bra.ParseFichierCsvStatique(r, ls.MemMap, func() bra.Visible { return new(bra.LieuVille) })
		case "regions.csv":
			bra.ParseFichierCsvStatique(r, ls.MemMap, func() bra.Visible { return new(bra.Région) })
		case "villes.csv":
			bra.ParseFichierCsvStatique(r, ls.MemMap, func() bra.Visible { return new(bra.Ville) })
		default:
			vue, _ := bra.ParseFichierCsvDynamique(r, ls.readTimeFromFilePath(pathToken), ls.MemMap, ls.verbose)
			ls.NbReadFiles++
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
func (ls *LecteurScripts) VisitDir(path string, f os.FileInfo) bool {
	pathToken := strings.Split(path, "/")
	indexTokenPrivate := IndexOfStringIn("private", pathToken)
	if indexTokenPrivate == -1 || indexTokenPrivate+1 == len(pathToken) {
		return true
	}
	if IndexOfStringIn(pathToken[indexTokenPrivate+1], ls.IdBralduns) == -1 {
		if ls.verbose {
			fmt.Printf("Fichier non autorisé : %s\n", path)
		}
		return false
	}
	return true
}

func main() {
	ls := NewLecteurScripts()

	cheminFichiersCsv := flag.String("in", "", "répertoire des fichiers csv")
	cheminRepertoireExport := flag.String("out", "", "répertoire d'export")
	idBraldunsBruts := flag.String("bralduns", "", "ids des bralduns, séparés par des virgules")
	cpuprofile := flag.String("cpuprofile", "", "fichier dans lequel écrire un bilan de profiling cpu")
	exportEnv := flag.Bool("exportenv", false, "exporte les environnements dans le json (fichier plus gros, false par défaut)")
	ls.verbose = *flag.Bool("verbose", false, "active le mode verbeux (faux par défaut)")
	flag.Parse()

	if *cpuprofile != "" {
		fmt.Println("Profiling actif, résultats dans le fichier -> abandon", *cpuprofile)
		fp, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(fp)
		defer pprof.StopCPUProfile()
	}

	startTime := time.Now().Unix()
	ls.IdBralduns = strings.Split(*idBraldunsBruts, ",")

	//> lecture des fichiers csv
	filepath.Walk(*cheminFichiersCsv, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
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

	//> export des images PNG
	//  et on supprime les cases de fond (maintenant inutile : on encode dans le png)
	for _, couche := range carte.Couches {
		bra.ConstruitNouveauPNG(*cheminRepertoireExport, couche)
		if !*exportEnv {
			couche.Cases = nil
		}
	}

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

	//> affichage d'un petit bilan
	fmt.Printf("Fichiers lus : %d\nCarte compilée dans %s en %d secondes\n\n", ls.NbReadFiles, *cheminRepertoireExport, time.Now().Unix()-startTime)

}
