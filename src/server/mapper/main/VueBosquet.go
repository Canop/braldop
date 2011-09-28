package main

import (
	"fmt"
	"os"
)

type VueBosquet struct {
	X       int16
	Y       int16
	NomType string
}

func (o *VueBosquet) readCsv(cells []string) (err os.Error) {
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
	o.NomType = cells[5]
	return
}
