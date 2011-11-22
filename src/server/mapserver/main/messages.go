package main

/*
définit les structures, mappées avec du json, en entrée et sortie du mapserver.
*/

type MessageIn struct {
	IdBraldun uint
	Mdpr string // mot de passe restreint
	Vue *DonnéesVue
}

type DonnéesVue struct {
}


type MessageOut struct {
	Erreur string
	Text string
}


