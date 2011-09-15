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
			fmt.Printf("   parsed file : %s\n", f.Name())
			switch filename {
			case "lieux_villes.csv":
				return ls.parseFichierStatique(f, func() Visible { return new(LieuVille) })
			case "regions.csv":
				return ls.parseFichierStatique(f, func() Visible { return new(Région) })
			case "villes.csv":
				return ls.parseFichierStatique(f, func() Visible { return new(Ville) })
			default:
				return ls.parseFichierDynamique(f)
			}
		} else {
			fmt.Println("   ignored file : " + f.Name())
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
		fmt.Printf("Erreur à la lecture des événements : %v", err)
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
