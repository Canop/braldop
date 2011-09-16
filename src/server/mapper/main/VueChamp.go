package main

import (
	"fmt"
	"os"
	"strconv"
)

type VueChamp struct {
	X         int16
	Y         int16
	IdBraldun uint
}

func (o *VueChamp) readCsv(cells []string) (err os.Error) {
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
	if o.IdBraldun, err = strconv.Atoui(cells[5]); err != nil {
		return
	}
	return
}

func (o *VueChamp) store(mm *MemMap) {
	mm.StoreChamp(o)
}
