package main

import (
	"fmt"
	"os"
)

type VueEnvironnement struct {
	Point
	NomSystemeEnvironnement string
	NomEnvironnement        string
}

func (o *VueEnvironnement) readCsv(cells []string) (err os.Error) {
	if len(cells) < 6 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.readCsvPoint(cells)
	o.NomSystemeEnvironnement = cells[4]
	o.NomEnvironnement = cells[5]
	return
}
