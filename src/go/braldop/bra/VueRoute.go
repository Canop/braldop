package bra

import (
	"errors"
	"fmt"
	"strconv"
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
	o.IdRoute, _ = strconv.Atoui(cells[4])
	o.TypeRoute = cells[5]
	return
}

func (o *VueRoute) Store(mm *MemMap) {
	mm.StoreRoute(o)
}
