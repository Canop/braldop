package main

/*
Objet exportable en json.

*/

import (
	"errors"
	"fmt"
	"strconv"
)

type Communauté struct {
	Id  uint
	Nom string
}

// cette méthode est appelée pour le décodage du fichier statique
func (o *Communauté) readCsv(cells []string) (err error) {
	if len(cells) < 3 {
		err = errors.New(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	if o.Id, err = strconv.Atoui(cells[0]); err != nil {
		return
	}
	o.Nom = cells[1]
	return
}

func (o *Communauté) store(mm *MemMap) {
	mm.Communautés[o.Id] = o
}
