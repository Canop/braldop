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
	Portail bool
}

func (o *VuePalissade) readCsv(cells []string, estPortail bool) (err os.Error) {
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
	o.Portail = estPortail
	return
}
