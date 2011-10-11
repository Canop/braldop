package main

import (
	"fmt"
	"os"
	"strconv"
)

type VueLieu struct {
	X                 int16
	Y                 int16
	Z                 int16
	Nom string
	IdTypeLieu uint
}

func (o *VueLieu) readCsv(cells []string) (err os.Error) {
	if len(cells) < 9 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.X, _ = Atoi16(cells[1])
	o.Y, _ = Atoi16(cells[2])
	o.Z, _ = Atoi16(cells[3])
	o.Nom = cells[5]
	if o.IdTypeLieu, err = strconv.Atoui(cells[8]); err != nil {
		return
	}
	return
}

