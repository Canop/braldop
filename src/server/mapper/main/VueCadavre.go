package main

import (
	"fmt"
	"os"
	"strconv"
)

type VueCadavre struct {
	X      int16
	Y      int16
	Id     uint
	Nom    string
	Taille string
	Niveau uint
}

func (o *VueCadavre) readCsv(cells []string) (err os.Error) {
	if len(cells) < 8 {
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
	return
}
