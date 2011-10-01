package main

import (
	"fmt"
	"os"
	"strconv"
)

type VueEchoppe struct {
	X                 int16
	Y                 int16
	Z int16
	Id                uint
	Nom               string
	Métier            string
	IdBraldun         uint
	NomCompletBraldun string
}

func (o *VueEchoppe) readCsv(cells []string) (err os.Error) {
	if len(cells) < 9 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.X, _ = Atoi16(cells[1])
	o.Y, _ = Atoi16(cells[2])
	o.Z, _ = Atoi16(cells[3])
	if o.Id, err = strconv.Atoui(cells[4]); err != nil {
		return
	}
	o.Nom = cells[5]
	o.Métier = cells[6]
	if o.IdBraldun, err = strconv.Atoui(cells[8]); err != nil {
		return
	}
	return
}

func (o *VueEchoppe) store(mm *MemMap) {
	mm.StoreEchoppe(o)
}
