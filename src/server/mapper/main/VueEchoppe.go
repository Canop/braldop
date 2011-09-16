package main

import (
	"fmt"
	"os"
	"strconv"
)

type VueEchoppe struct {
	X         int16
	Y         int16
	Id        uint
	Nom       string
	Métier    string
	IdBraldun uint
}

func (o *VueEchoppe) readCsv(cells []string) (err os.Error) {
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
	o.Métier = cells[6]
	if o.IdBraldun, err = strconv.Atoui(cells[8]); err != nil {
		return
	}
	return
}

func (o *VueEchoppe) store(mm *MemMap) {
	mm.StoreEchoppe(o)
}
