package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (ls *LecteurScripts) parseLigneFichierDynamique(line string, date uint) {
	cells := strings.Split(line, ";")
	//fmt.Println(" cells : " + strings.Join(cells, "#"))
	if len(cells) < 3 {
		fmt.Printf("  Ligne trop courte : %s\n", line)
		return
	}
	switch cells[0] {
	case "ENVIRONNEMENT":
		o := new(VueEnvironnement)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueEnvironnement : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VueEnvironnement : %+v\n", o)
			ls.MemMap.StoreEnvironnement(o)
		}
	case "POSITION":
		o := new(VuePosition)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture position : %+v \n cellules : %+v", err, cells)
		} else {
			//fmt.Printf(" Position : %+v\n", o)
		}
	case "ROUTE":
		o := new(VueRoute)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture Route : %+v \n cellules : %+v", err, cells)
		} else {
			//fmt.Printf(" Route : %+v\n", o)
			ls.MemMap.StoreRoute(o)
		}
	}
	//fmt.Printf("  %s\n", line)
}

func (ls *LecteurScripts) parseFichierDynamique(file *os.File) os.Error {
	r := bufio.NewReader(file)
	line, err := readLine(r)
	for err == nil {
		ls.parseLigneFichierDynamique(line, 0) // TODO date à construire
		line, err = readLine(r)
	}
	ls.NbReadFiles++
	if err != os.EOF {
		fmt.Println("Error in parsing (parseFichierDynamique) :")
		fmt.Println(err)
		return err
	}
	return nil
}
