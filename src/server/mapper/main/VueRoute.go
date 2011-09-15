package main

import (
	"fmt"
	"os"
	"strconv"
)

type VueRoute struct {
	Point
	IdRoute   uint
	TypeRoute string
}

func (o *VueRoute) readCsv(cells []string) (err os.Error) {
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
	if o.Z, err = Atoi16(cells[3]); err != nil {
		return
	}
	if o.IdRoute, err = strconv.Atoui(cells[4]); err != nil {
		return
	}
	o.TypeRoute = cells[5]
	return
}

func (o *VueRoute) store(mm *MemMap) {
	mm.StoreRoute(o)
}
