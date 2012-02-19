package bra

import (
	"errors"
	"fmt"
)

type VueEchoppe struct {
	X                 int16
	Y                 int16
	Z                 int16
	Id                uint
	Nom               string
	Métier            string
	IdBraldun         uint
	NomCompletBraldun string
}

func (o *VueEchoppe) ReadCsv(cells []string) (err error) {
	if len(cells) < 9 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.X, _ = Atoi16(cells[1])
	o.Y, _ = Atoi16(cells[2])
	o.Z, _ = Atoi16(cells[3])
	if o.Id, err = Atoui(cells[4]); err != nil {
		return
	}
	o.Nom = cells[5]
	o.Métier = cells[6]
	if o.IdBraldun, err = Atoui(cells[8]); err != nil {
		return
	}
	return
}

func (o *VueEchoppe) Store(mm *MemMap) {
	mm.StoreEchoppe(o)
}
