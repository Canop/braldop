package main

/*
couche stockée en mémoire
*/

import (
	"fmt"
)

type MemCouche struct {
	Z                   int16
	BosquetsParXY       map[int32]*VueBosquet
	ChampsParXY         map[int32]*VueChamp
	CrevassesParXY      map[int32]*VueCrevasse
	EchoppesParXY       map[int32]*VueEchoppe
	PalissadesParXY     map[int32]*VuePalissade
	EnvironnementsParXY map[int32]*VueEnvironnement
	LieuxParXY          map[int32]*VueLieu
	RoutesParXY         map[int32]*VueRoute
}

func NewMemCouche() (mc *MemCouche) {
	mc = new(MemCouche)
	mc.BosquetsParXY = make(map[int32]*VueBosquet)
	mc.ChampsParXY = make(map[int32]*VueChamp)
	mc.CrevassesParXY = make(map[int32]*VueCrevasse)
	mc.EchoppesParXY = make(map[int32]*VueEchoppe)
	mc.EnvironnementsParXY = make(map[int32]*VueEnvironnement)
	mc.PalissadesParXY = make(map[int32]*VuePalissade)
	mc.LieuxParXY = make(map[int32]*VueLieu)
	mc.RoutesParXY = make(map[int32]*VueRoute)
	return mc
}

func (mc *MemCouche) Compile(mm *MemMap) (m *Couche) {
	m = NewCouche()
	m.Z = mc.Z
	cases := make(map[int32]*Case) // map suivant PosKey(x,y)
	for _, e := range mc.EnvironnementsParXY {
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
	for _, b := range mc.BosquetsParXY {
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
	for _, r := range mc.RoutesParXY {
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
	for _, r := range mc.CrevassesParXY {
		key := PosKey(r.X, r.Y)
		c, ok := cases[key]
		if !ok {
			fmt.Println("Crevasse sans case en ", r.X, ", ", r.Y, ", ", r.Z)
		} else {
			fmt.Println("Crevasse OK")
			c.Fond = c.Fond + "-crevasse"
		}
	}
	for _, b := range mc.PalissadesParXY {
		m.Palissades = append(m.Palissades, b)
	}
	for _, e := range mc.EchoppesParXY {
		// on renseigne si possible le nom du braldun
		if mmb, ok := mm.Bralduns[e.IdBraldun]; ok {
			e.NomCompletBraldun = mmb.Prénom + " " + mmb.Nom
		} else {
			fmt.Printf("Braldun introuvable : %d\n", e.IdBraldun)
		}
		m.Echoppes = append(m.Echoppes, e)
	}
	for _, l := range mc.LieuxParXY {
		m.Lieux = append(m.Lieux, l)
	}
	for _, e := range mc.ChampsParXY {
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
	return m
}
