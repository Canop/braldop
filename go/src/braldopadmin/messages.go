package main

// définit les structures, mappées avec du json, en entrée et sortie du mapserver.

import (
	"bra"
)

type MessageIn struct {
	IdBraldun uint
	Mdpr      string // mot de passe restreint
	Vue       *DonnéesVue
	Version   string // version du client
}

type DonnéesVue struct {
	Couches  []bra.Couche
	Vues     []bra.Vue
	Position bra.VuePosition
}

type MessageOut struct {
	Erreur    string
	PngCouche string
	Text      string // message du serveur à l'utilisateur
}
