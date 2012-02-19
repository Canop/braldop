package bra

import (
	"errors"
	"fmt"
)

type VueCadavre struct {
	X      int16
	Y      int16
	Id     uint
	Nom    string
	Taille string
}

func (o *VueCadavre) ReadCsv(cells []string) (err error) {
	if len(cells) < 7 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	if o.X, err = Atoi16(cells[1]); err != nil {
		return
	}
	if o.Y, err = Atoi16(cells[2]); err != nil {
		return
	}
	if o.Id, err = Atoui(cells[4]); err != nil {
		return
	}
	o.Nom = cells[5]
	o.Taille = cells[6]
	return
}
