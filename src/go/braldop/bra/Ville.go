package bra

import (
	"errors"
	"fmt"
	"strconv"
)

type Ville struct {
	Id          uint
	Nom         string
	EstCapitale bool
	XMin        int16
	XMax        int16
	YMin        int16
	YMax        int16
}

func (o *Ville) ReadCsv(cells []string) (err error) {
	if len(cells) < 8 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	if o.Id, err = strconv.Atoui(cells[0]); err != nil {
		return
	}
	o.Nom = cells[1]
	o.EstCapitale = "oui" == cells[2]
	if o.XMin, err = Atoi16(cells[3]); err != nil {
		return
	}
	if o.YMin, err = Atoi16(cells[4]); err != nil {
		return
	}
	if o.XMax, err = Atoi16(cells[5]); err != nil {
		return
	}
	if o.YMax, err = Atoi16(cells[6]); err != nil {
		return
	}
	return
}

func (o *Ville) Store(mm *MemMap) {
	mm.StoreVille(o)
}
