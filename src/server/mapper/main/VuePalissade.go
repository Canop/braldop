package main
/*
palissade ou portail
*/

import (
	"fmt"
	"os"
)

type VuePalissade struct {
	X       int16
	Y       int16
	Z int16
	Portail bool
}

func (o *VuePalissade) readCsv(cells []string, estPortail bool) (err os.Error) {
	if len(cells) < 6 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.X, _ = Atoi16(cells[1])
	o.Y, _ = Atoi16(cells[2])
	o.Z, _ = Atoi16(cells[3])
	o.Portail = estPortail
	return
}
