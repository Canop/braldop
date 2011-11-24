package bra

import (
	"errors"
	"fmt"
	"strconv"
)

type VueMonstre struct {
	X      int16
	Y      int16
	Id     uint
	Nom    string
	Taille string
	IdType uint
	Gibier bool
}

func (o *VueMonstre) ReadCsv(cells []string) (err error) {
	if len(cells) < 9 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
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
	if o.IdType, err = strconv.Atoui(cells[7]); err != nil {
		return
	}
	if cells[8] == "oui" {
		o.Gibier = true
	}
	return
}
