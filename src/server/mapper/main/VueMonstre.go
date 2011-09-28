package main

import (
	"fmt"
	"os"
	"strconv"
)

type VueMonstre struct {
	X      int16
	Y      int16
	Id     uint
	Nom    string
	Taille string
	Niveau uint
	IdType uint
}

func (o *VueMonstre) readCsv(cells []string) (err os.Error) {
	if len(cells) < 9 {
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
	o.Nom = cells[5]
	o.Taille = cells[6]
	if o.Niveau, err = strconv.Atoui(cells[7]); err != nil {
		return
	}
	if o.IdType, err = strconv.Atoui(cells[8]); err != nil {
		return
	}
	return
}
