package main

import (
	"fmt"
	"os"
	"strconv"
)

type Région struct {
	Id     uint
	Nom    string
	XMin   int16
	XMax   int16
	YMin   int16
	YMax   int16
	EstPvp bool
}

func (o *Région) readCsv(cells []string) (err os.Error) {
	if len(cells) < 7 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	if o.Id, err = strconv.Atoui(cells[0]); err != nil {
		return
	}
	o.Nom = cells[1]
	if o.XMin, err = Atoi16(cells[2]); err != nil {
		return
	}
	if o.XMax, err = Atoi16(cells[3]); err != nil {
		return
	}
	if o.YMin, err = Atoi16(cells[4]); err != nil {
		return
	}
	if o.YMax, err = Atoi16(cells[5]); err != nil {
		return
	}
	o.EstPvp = "oui" == cells[6]
	return
}

func (o *Région) store(mm *MemMap) {
	mm.StoreRégion(o)
}
