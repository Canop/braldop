package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (ls *LecteurScripts) parseLigneFichierStatique(line string, alloue func() Visible) {
	cells := strings.Split(line, ";")
	o := alloue()
	if err := o.readCsv(cells); err != nil {
		fmt.Printf(" Erreur lecture : %+v \n cellules : %+v", err, cells)
	} else {
		o.store(ls.MemMap)
	}
}
func (ls *LecteurScripts) parseFichierStatique(file *os.File, alloue func() Visible) os.Error {
	r := bufio.NewReader(file)
	line, err := readLine(r)
	if err != nil {
		fmt.Println("Fichier vide")
		return err
	}
	line, err = readLine(r) // on saute une ligne car la première contient les en-têtes
	for err == nil {
		ls.parseLigneFichierStatique(line, alloue)
		line, err = readLine(r)
	}
	ls.NbReadFiles++
	if err != os.EOF {
		fmt.Println("Error in parseFichierStatique :")
		fmt.Println(err)
		return err
	}
	return nil
}
