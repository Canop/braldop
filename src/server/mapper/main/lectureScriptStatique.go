package main

import (
	"bufio"
	"fmt"
	"os"
)

func (ls *LecteurScripts) parseLigneFichierStatique(cells []string, alloue func() Visible) {
	o := alloue()
	if err := o.readCsv(cells); err != nil {
		fmt.Printf(" Erreur lecture : %+v \n cellules : %+v", err, cells)
	} else {
		o.store(ls.MemMap)
	}
}
func (ls *LecteurScripts) parseFichierStatique(r *bufio.Reader, alloue func() Visible) {
	ls.NbReadFiles++
	_, _ = readLine(r) // on saute une ligne car la première contient les en-têtes
	for {
		line, err := readLine(r)
		if err != nil {
			if err != os.EOF {
				fmt.Println("Error in parsing (parseFichierDynamique) :")
				fmt.Println(err)
			}
			return
		}
		ls.parseLigneFichierStatique(line, alloue)
	}
}
