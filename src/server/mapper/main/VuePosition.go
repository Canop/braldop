package main

import (
	"fmt"
	"os"
	"strconv"
)

type VuePosition struct {
	Point
	XMin       int16
	XMax       int16
	YMin       int16
	YMax       int16
	IdBraldun  uint
	VueNbCases int
	VueBm      int
}

func (o *VuePosition) readCsv(cells []string) (err os.Error) {
	if len(cells) < 11 {
		err = os.NewError(fmt.Sprintf("pas assez de champs (%d)", len(cells)))
		return
	}
	o.readCsvPoint(cells)
	if o.XMin, err = Atoi16(cells[4]); err != nil {
		return
	}
	if o.XMax, err = Atoi16(cells[5]); err != nil {
		return
	}
	if o.YMin, err = Atoi16(cells[6]); err != nil {
		return
	}
	if o.YMax, err = Atoi16(cells[7]); err != nil {
		return
	}
	if o.IdBraldun, err = strconv.Atoui(cells[8]); err != nil {
		return
	}
	if o.VueNbCases, err = strconv.Atoi(cells[9]); err != nil {
		return
	}
	if o.VueBm, err = strconv.Atoi(cells[10]); err != nil {
		return
	}
	return
}
