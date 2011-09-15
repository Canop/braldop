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
	if o.X, err = Atoi16(cells[1]); err != nil {
		return
	}
	if o.Y, err = Atoi16(cells[2]); err != nil {
		return
	}
	if o.Z, err = Atoi16(cells[3]); err != nil {
		return
	}
	o.NomSystemeEnvironnement = cells[4]
	o.NomEnvironnement = cells[5]
	return
}
