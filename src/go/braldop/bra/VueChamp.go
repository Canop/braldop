package bra

import (
	"errors"
	"fmt"
	"strconv"
)

type VueChamp struct {
	X                 int16
	Y                 int16
	Z                 int16
	IdBraldun         uint
	NomCompletBraldun string
}

func (o *VueChamp) ReadCsv(cells []string) (err error) {
	if len(cells) < 6 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.X, _ = Atoi16(cells[1])
	o.Y, _ = Atoi16(cells[2])
	o.Z, _ = Atoi16(cells[3])
	if o.IdBraldun, err = strconv.Atoui(cells[5]); err != nil {
		return
	}
	return
}

func (o *VueChamp) Store(mm *MemMap) {
	mm.StoreChamp(o)
}
