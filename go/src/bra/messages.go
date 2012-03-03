package bra

// définit les structures, mappées avec du json, en entrée et sortie de braldopserver

type MessageIn struct {
	IdBraldun uint
	Etat      *EtatBraldun
	Mdpr      string // mot de passe restreint
	Vue       *DonnéesVue
	Version   string // version du client
	ZRequis   int    // la profondeur pour laquelle on veut en retour des données (png+compléments)
	Cmd       string
	Action    string
	Cible     uint
}

type MessageOut struct {
	Erreur    string
	PngCouche string
	Text      string      // message du serveur à l'utilisateur
	Z         int         // la profondeur correspondant aux données envoyées (en particulier le png)
	ZConnus   []int       // les profondeurs pour lesquelles on peut proposer des données
	DV        *DonnéesVue // des compléments
	Partages  []*Partage  // uniquement envoyé si demandé
	Etats []*EtatBraldun
}

type EtatBraldun struct {
	IdBraldun uint
	PV        int // les pv actuels
	PVMax     int // les pv max
	PA        int // pa disponibles
	DLA       int // timestamp unix
	DuréeTour int // en secondes
	Faim int // entre 0 et 100 je pense
}
