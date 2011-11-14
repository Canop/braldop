package main
/*
palissade ou portail
*/

import (
	"errors"
	"fmt"
	"time"
)

type VuePalissade struct {
	X            int16
	Y            int16
	Z            int16
	Portail      bool
	Destructible bool
	DateFin      int64 // secondes depuis 1970 (0 si pas de date)
}

func (o *VuePalissade) readCsv(cells []string, estPortail bool) (err error) {
	if len(cells) < 6 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.X, _ = Atoi16(cells[1])
	o.Y, _ = Atoi16(cells[2])
	o.Z, _ = Atoi16(cells[3])
	o.Portail = estPortail

	o.Destructible = cells[5] == "oui"
	if len(cells) >= 7 {
		t, terr := time.Parse("2006-01-02 15:04:05 MST", cells[6]+" CEST")
		if terr == nil {
			o.DateFin = t.Seconds()
		}
	}
	return
}
