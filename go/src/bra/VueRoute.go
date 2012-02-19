package bra

import (
	"errors"
	"fmt"
)

type VueRoute struct {
	Point
	IdRoute   uint
	TypeRoute string
}

func (o *VueRoute) ReadCsv(cells []string) (err error) {
	if len(cells) < 6 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.readCsvPoint(cells)
	o.IdRoute, _ = Atoui(cells[4])
	o.TypeRoute = cells[5]
	return
}

func (o *VueRoute) Store(mm *MemMap) {
	mm.StoreRoute(o)
}
