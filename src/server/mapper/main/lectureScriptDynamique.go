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
	case "BRALDUN":
		o := new(Braldun)
		if err := o.readCsvDynamique(cells); err != nil {
			fmt.Printf(" Erreur lecture Braldun : %+v \n cellules : %+v\n", err, cells)
		} else {
			//~ fmt.Printf(" Braldun : %+v\n", o)
			vue.Bralduns = append(vue.Bralduns, o)
		}
	case "CADAVRE":
		o := new(VueCadavre)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueCadavre : %+v \n cellules : %+v", err, cells)
		} else {
			fmt.Printf(" VueCadavre : %+v\n", o)
			vue.Cadavres = append(vue.Cadavres, o)
		}
	case "CHAMP":
		o := new(VueChamp)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueChamp : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VueChamp : %+v\n", o)
			ls.MemMap.StoreChamp(o)
		}
	case "CHARRETTE":
		o := new(VueObjet)
		if err := o.readCsvCharrette(cells); err != nil {
			fmt.Printf(" Erreur lecture VueObjet : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" CHARRETTE : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "ECHOPPE":
		o := new(VueEchoppe)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueEchoppe : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VueEchoppe : %+v\n", o)
			ls.MemMap.StoreEchoppe(o)
		}
	case "ELEMENT":
		o := new(VueObjet)
		if err := o.readCsvElement(cells); err != nil {
			fmt.Printf(" Erreur lecture VueObjet : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" ELEMENT : %+v\n", o)
			vue.Objets = append(vue.Objets, o)
		}
	case "ENVIRONNEMENT":
		o := new(VueEnvironnement)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueEnvironnement : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VueEnvironnement : %+v\n", o)
			ls.MemMap.StoreEnvironnement(o)
		}
	case "MONSTRE":
		o := new(VueMonstre)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture VueMonstre : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" VueMonstre : %+v\n", o)
			vue.Monstres = append(vue.Monstres, o)
		}
	case "POSITION":
		o := new(VuePosition)
		if err := o.readCsv(cells); err != nil {
			fmt.Printf(" Erreur lecture position : %+v \n cellules : %+v", err, cells)
		} else {
			//~ fmt.Printf(" Position : %+v\n", o)
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
}

func (ls *LecteurScripts) parseFichierDynamique(file *os.File, time int64) (vue *Vue, err os.Error) {
	r := bufio.NewReader(file)
	vue = NewVue()
	vue.Time = time
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
