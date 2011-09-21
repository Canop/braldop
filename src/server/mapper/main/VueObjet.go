package main

import ()

type VueObjet struct {
	X        int16
	Y        int16
	Type     string
	Quantité uint
}

func (o *VueObjet) store(mm *MemMap) {
	mm.StoreObjet(o)
}
