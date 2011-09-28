package main

/*
carte stockée en mémoire
  TODO : renommer en MemCouche
*/

import (
	"fmt"
)

type MemMap struct {
	Bralduns            map[uint]*Braldun // il s'agit des bralduns récupérés depuis le fichier statique donc certaines informations sont absentes
	BosquetsParXY       map[int32]*VueBosquet
	ChampsParXY         map[int32]*VueChamp
	EchoppesParXY       map[int32]*VueEchoppe
	EnvironnementsParXY map[int32]*VueEnvironnement
	RoutesParXY         map[int32]*VueRoute
	Villes              []*Ville
	LieuxVilles         []*LieuVille
	Régions             []*Région
	DernièresVues       map[uint]*Vue // les dernières vues avec pour clé l'id du Braldun voyeur
	Objets              []*VueObjet
}

func NewMemMap() (mm *MemMap) {
	mm = new(MemMap)
	mm.BosquetsParXY = make(map[int32]*VueBosquet)
	mm.Bralduns = make(map[uint]*Braldun)
	mm.ChampsParXY = make(map[int32]*VueChamp)
	mm.EchoppesParXY = make(map[int32]*VueEchoppe)
	mm.EnvironnementsParXY = make(map[int32]*VueEnvironnement)
	mm.RoutesParXY = make(map[int32]*VueRoute)
	mm.Villes = make([]*Ville, 0, 10)
	mm.LieuxVilles = make([]*LieuVille, 0, 10)
	mm.Régions = make([]*Région, 0, 10)
	mm.DernièresVues = make(map[uint]*Vue)
	mm.Objets = make([]*VueObjet, 0, 10)
	return mm
}

func (mm *MemMap) StoreBosquet(o *VueBosquet) {
	mm.BosquetsParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreChamp(o *VueChamp) {
	mm.ChampsParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreEchoppe(o *VueEchoppe) {
	mm.EchoppesParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreEnvironnement(o *VueEnvironnement) {
	mm.EnvironnementsParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreRoute(o *VueRoute) {
	mm.RoutesParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreVille(o *Ville) { // on va supposer qu'on ne lit pas deux fois la même ville
	mm.Villes = append(mm.Villes, o)
}
func (mm *MemMap) StoreLieuVille(o *LieuVille) { // on va supposer qu'on ne lit pas deux fois le même lieu
	mm.LieuxVilles = append(mm.LieuxVilles, o)
}
func (mm *MemMap) StoreRégion(o *Région) {
	mm.Régions = append(mm.Régions, o)
}
func (mm *MemMap) StoreObjet(o *VueObjet) {
	mm.Objets = append(mm.Objets, o)
}

func (mm *MemMap) Compile() (m *Map) {
	m = NewMap()
	cases := make(map[int32]*Case) // map suivant PosKey(x,y)
	for _, e := range mm.EnvironnementsParXY {
		key := PosKey(e.X, e.Y)
		c, ok := cases[key]
		if !ok {
			c = new(Case)
			c.X = e.X
			c.Y = e.Y
			cases[key] = c
		}
		c.Fond = e.NomSystemeEnvironnement
	}
	for _, b := range mm.BosquetsParXY {
		key := PosKey(b.X, b.Y)
		c, ok := cases[key]
		if !ok {
			c = new(Case)
			c.X = b.X
			c.Y = b.Y
			cases[key] = c
		}
		c.Fond = b.NomType
	}
	for _, r := range mm.RoutesParXY {
		key := PosKey(r.X, r.Y)
		c, ok := cases[key]
		if !ok {
			c = new(Case)
			c.X = r.X
			c.Y = r.Y
			cases[key] = c
		}
		switch r.TypeRoute {
		case "balise":
			c.Fond = c.Fond + "-gr"
		case "echoppe":
			c.Fond = "pave"
		case "route":
			c.Fond = "route"
		case "ville":
			c.Fond = "pave"
		case "ruine":
			c.Fond = "pave"
		}
	}
	for _, e := range mm.EchoppesParXY {
		// on renseigne si possible le nom du braldun
		if mmb, ok := mm.Bralduns[e.IdBraldun]; ok {
			e.NomCompletBraldun = mmb.Prénom + " " + mmb.Nom
		} else {
			fmt.Printf("Braldun introuvable : %d\n", e.IdBraldun)
		}
		m.Echoppes = append(m.Echoppes, e)
	}
	for _, e := range mm.ChampsParXY {
		// on renseigne si possible le nom du braldun
		if mmb, ok := mm.Bralduns[e.IdBraldun]; ok {
			e.NomCompletBraldun = mmb.Prénom + " " + mmb.Nom
		} else {
			fmt.Printf("Braldun introuvable : %d\n", e.IdBraldun)
		}
		m.Champs = append(m.Champs, e)
	}
	for _, c := range cases {
		m.Cases = append(m.Cases, c)
	}
	m.Villes = mm.Villes           // pour l'instant pour les villes c'est simple...
	m.LieuxVilles = mm.LieuxVilles // et pour leurs lieux aussi
	m.Régions = mm.Régions         // et les régions itou
	for _, v := range mm.DernièresVues {
		m.Vues = append(m.Vues, v)
		// on indique le prénom du voyeur
		if b, ok := mm.Bralduns[v.Voyeur]; ok {
			v.PrénomVoyeur = b.Prénom
		}
		// on remplit les champs manquant des bralduns
		for _, b := range v.Bralduns {
			if mmb, ok := mm.Bralduns[b.Id]; ok {
				b.Prénom = mmb.Prénom
				b.Nom = mmb.Nom
				b.Niveau = mmb.Niveau
				b.Sexe = mmb.Sexe
			} else {
				fmt.Printf("Braldun introuvable : %d\n", b.Id)
			}
		}
	}
	return m
}
