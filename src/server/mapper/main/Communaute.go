package main

/*
Objet exportable en json.

*/

import (
	"fmt"
	"os"
	"strconv"
)

type Communauté struct {
	Id         uint
	Nom        string
}

// cette méthode est appelée pour le décodage du fichier statique
func (o *Communauté) readCsv(cells []string) (err os.Error) {
	if len(cells) < 3 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
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