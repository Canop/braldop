package main

import (
	"fmt"
	"os"
)

type VueObjet struct {
	X        int16
	Y        int16
	Type     string
	Quantité uint
}


func (o *VueObjet) readCsv(Type string, cells []string) (err os.Error) {
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
	o.Type = Type
	return
}

