package main

import (
	"bufio"
	"fmt"
	"io"
)

func (ls *LecteurScripts) parseLigneFichierStatique(cells []string, alloue func() Visible) {
	o := alloue()
	if err := o.ReadCsv(cells); err != nil {
		fmt.Printf(" Erreur lecture : %+v \n cellules : %+v", err, cells)
	} else {
		o.Store(ls.MemMap)
	}
}
func (ls *LecteurScripts) parseFichierStatique(r *bufio.Reader, alloue func() Visible) {
	ls.NbReadFiles++
	_, _ = readLine(r) // on saute une ligne car la première contient les en-têtes
	for {
		line, err := readLine(r)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error in parsing (parseFichierDynamique) :")
				fmt.Println(err)
			}
			return
		}
		ls.parseLigneFichierStatique(line, alloue)
	}
}
