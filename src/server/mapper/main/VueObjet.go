package main

/*
 * correspond aux
 *  éléments
 *  charettes
 *  runes
 *  buissons
 */

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type VueObjet struct {
	X        int16
	Y        int16
	Type     string
	Quantité uint
	Label    string
}

func (o *VueObjet) readCsvBuisson(cells []string) (err os.Error) {
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
	o.Type = "buisson"
	o.Label = cells[5]
	return
}
func (o *VueObjet) readCsvCharrette(cells []string) (err os.Error) {
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
	o.Type = "charrette"
	o.Label = cells[5]
	return
}

func (o *VueObjet) readCsvElement(cells []string) (err os.Error) {
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
	o.Type = strings.ToLower(cells[4])
	o.Quantité, _ = strconv.Atoui(cells[5])

	o.Label = fmt.Sprintf("%d %s", o.Quantité, o.Type)
	// lignes suivantes provisoires
	if o.Quantité > 1 {
		if len(cells) > 6 {
			o.Label += cells[6]
		} else {
			o.Label += "s"
		}
	}
	return
}

func (o *VueObjet) readCsvRune(cells []string) (err os.Error) {
	if len(cells) < 3 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	if o.X, err = Atoi16(cells[1]); err != nil {
		return
	}
	if o.Y, err = Atoi16(cells[2]); err != nil {
		return
	}
	o.Type = "rune"
	o.Label = "rune"
	return
}
