package main

/*
définit les structures, mappées avec du json, en entrée et sortie du mapserver.
*/

import (
	"braldop/bra"
)

type MessageIn struct {
	IdBraldun uint
	Mdpr      string // mot de passe restreint
	Vue       *DonnéesVue
}

type DonnéesVue struct {
	Couches []bra.Couche
	Vues []bra.Vue
	Position bra.VuePosition
}

type MessageOut struct {
	Erreur string
	Text   string
	UrlPngCouche string
	PngCoucheBase64 string
}
