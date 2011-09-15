package main

import (
	"fmt"
	"os"
	"strconv"
)

type LieuVille struct {
	Id         uint
	Nom        string
	IdTypeLieu int16
	X          int16
	Y          int16
}

func (o *LieuVille) readCsv(cells []string) (err os.Error) {
	if len(cells) < 8 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	if o.Id, err = strconv.Atoui(cells[0]); err != nil {
		return
	}
	o.Nom = cells[1]
	if o.IdTypeLieu, err = Atoi16(cells[4]); err != nil {
		return
	}
	if o.X, err = Atoi16(cells[5]); err != nil {
		return
	}
	if o.Y, err = Atoi16(cells[6]); err != nil {
		return
	}
	return
}

func (o *LieuVille) store(mm *MemMap) {
	mm.StoreLieuVille(o)
}
