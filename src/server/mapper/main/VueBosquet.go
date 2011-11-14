package main

import (
	"errors"
	"fmt"
)

type VueBosquet struct {
	X       int16
	Y       int16
	Z       int16
	NomType string
}

func (o *VueBosquet) readCsv(cells []string) (err error) {
	if len(cells) < 6 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.X, _ = Atoi16(cells[1])
	o.Y, _ = Atoi16(cells[2])
	o.Z, _ = Atoi16(cells[3])
	o.NomType = cells[5]
	return
}
