package main

/*
 * correspond aux
 *  éléments
 *  charettes
 *  runes
 *  buissons
 *  munition
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

func (o *VueObjet) readCsvSimple(cells []string, Type string, Label string) (err os.Error) {
	if len(cells) < 5 {
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
	o.Label = Label
	return
}

func (o *VueObjet) readCsvSimpleLabel(cells []string, Type string) (err os.Error) {
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
func (o *VueObjet) readCsvMunition(cells []string) (err os.Error) {
	if len(cells) < 8 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	if o.X, err = Atoi16(cells[1]); err != nil {
		return
	}
	if o.Y, err = Atoi16(cells[2]); err != nil {
		return
	}
	o.Type = "munition"
	o.Quantité, _ = strconv.Atoui(cells[6])
	o.Label = fmt.Sprintf("%d %s", o.Quantité, cells[5])
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
