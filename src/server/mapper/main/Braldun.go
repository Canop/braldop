package main

/*
Objet exportable en json.

Certains champs peuvent provenir du fichier statique, d'autres du dynamique (de la vue)

*/

import (
	"fmt"
	"os"
	"strconv"
)

type Braldun struct {
	Id                uint
	X                 int16
	Y                 int16
	Prénom            string
	Nom               string
	Niveau            uint
	Sexe              string // "f" ou "m"
	KO                bool
	Intangible        bool
	Camp              string // équipe de soule : "a" ou "b" ou ""
	IdCommunauté      uint
	PointsGredin      int16 // cellule 17
	PointsRedresseur  int16 // cellule 18
	PointsDistinction int16 // cellule 19
}

// cette méthode est appelée pour le décodage du fichier statique
func (o *Braldun) readCsv(cells []string) (err os.Error) {
	if len(cells) < 20 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	if o.Id, err = strconv.Atoui(cells[0]); err != nil {
		return
	}
	o.Prénom = cells[1]
	o.Nom = cells[2]
	o.Niveau, _ = strconv.Atoui(cells[3])
	if len(cells[4]) < 1 {
		return
	}
	o.Sexe = string(cells[4][0])
	o.IdCommunauté, _ = strconv.Atoui(cells[12])
	o.PointsGredin, _ = Atoi16(cells[17])
	o.PointsRedresseur, _ = Atoi16(cells[18])
	o.PointsDistinction, _ = Atoi16(cells[19])
	return
}

// cette méthode est appelée pour le décodage du fichier dynamique
func (o *Braldun) readCsvDynamique(cells []string) (err os.Error) {
	if len(cells) < 6 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	if o.X, err = Atoi16(cells[1]); err != nil {
		return
	}
	if o.Y, err = Atoi16(cells[2]); err != nil {
		return
	}
	if o.Id, err = strconv.Atoui(cells[4]); err != nil {
		return
	}
	o.KO = cells[5] == "oui"
	o.Intangible = cells[6] == "oui"
	o.Camp = cells[8]
	return
}

func (o *Braldun) store(mm *MemMap) {
	mm.Bralduns[o.Id] = o
}
