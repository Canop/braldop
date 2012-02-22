package main

// définit les structures, mappées avec du json, en entrée et sortie du mapserver.

import (
	"bra"
)

type MessageIn struct {
	IdBraldun uint
	Mdpr      string // mot de passe restreint
	Vue       *bra.DonnéesVue
	Version   string // version du client
	ZRequis   int    // la profondeur pour laquelle on veut en retour des données (png+compléments)
	Cmd       string
	Action    string
	Cible     uint
}



type MessageOut struct {
	Erreur    string
	PngCouche string
	Text      string         // message du serveur à l'utilisateur
	Z         int            // la profondeur correspondant aux données envoyées (en particulier le png)
	ZConnus   []int          // les profondeurs pour lesquelles on peut proposer des données
	DV        *bra.DonnéesVue    // des compléments
	Partages  []*bra.Partage // uniquement envoyé si demandé
}
