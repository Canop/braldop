package main

import (
	"errors"
	"fmt"
)

type VueEnvironnement struct {
	Point
	NomSystemeEnvironnement string
	NomEnvironnement        string
}

func (o *VueEnvironnement) readCsv(cells []string) (err error) {
	if len(cells) < 6 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.readCsvPoint(cells)
	o.NomSystemeEnvironnement = cells[4]
	o.NomEnvironnement = cells[5]
	return
}
