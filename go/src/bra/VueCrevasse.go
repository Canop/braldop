package bra

import (
	"errors"
	"fmt"
)

type VueCrevasse struct {
	Point
	Id uint
}

func (o *VueCrevasse) ReadCsv(cells []string) (err error) {
	if len(cells) < 5 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.readCsvPoint(cells)
	if o.Id, err = Atoui(cells[4]); err != nil {
		return
	}
	return
}
