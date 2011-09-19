package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (ls *LecteurScripts) parseLigneFichierDynamique(line string, vue *Vue) {
	cells := strings.Split(line, ";")
	//fmt.Println(" cells : " + strings.Join(cells, "#"))
	if len(cells) < 3 {
		fmt.Printf("  Ligne trop courte : %s\n", line)
		return
	}
	switch cells[0] {
	case "CHAMP":
		o := new(VueChamp)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueChamp : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VueChamp : %+v\n", o)
			ls.MemMap.StoreChamp(o)
		}
	case "ECHOPPE":
		o := new(VueEchoppe)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueEchoppe : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VueEchoppe : %+v\n", o)
			ls.MemMap.StoreEchoppe(o)
		}
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
			vue.Voyeur = o.IdBraldun
			vue.XMin = o.XMin
			vue.XMax = o.XMax
			vue.YMin = o.YMin
			vue.YMax = o.YMax
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

func (ls *LecteurScripts) parseFichierDynamique(file *os.File, time int64) (vue *Vue, err os.Error) {
	r := bufio.NewReader(file)
	vue = NewVue()
	line, err := readLine(r)
	for err == nil {
		ls.parseLigneFichierDynamique(line, vue)
		line, err = readLine(r)
	}
	ls.NbReadFiles++
	if err != os.EOF {
		fmt.Println("Error in parsing (parseFichierDynamique) :")
		fmt.Println(err)
		return
	}
	err = nil
	return
}
