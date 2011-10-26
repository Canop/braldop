package main

import (
	"fmt"
	"os"
	"strconv"
)

type VueCrevasse struct {
	Point
	Id uint
}

func (o *VueCrevasse) readCsv(cells []string) (err os.Error) {
	if len(cells) < 5 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.readCsvPoint(cells)
	if o.Id, err = strconv.Atoui(cells[4]); err != nil {
		return
	}
	return
}
