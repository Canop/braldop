package main

/*
carte stockée en mémoire
*/

type MemMap struct {
	EchoppesParXY       map[int32]*VueEchoppe
	EnvironnementsParXY map[int32]*VueEnvironnement
	RoutesParXY         map[int32]*VueRoute
	Villes              []*Ville
	LieuxVilles         []*LieuVille
	Régions             []*Région
}

func NewMemMap() (mm *MemMap) {
	mm = new(MemMap)
	mm.EchoppesParXY = make(map[int32]*VueEchoppe)
	mm.EnvironnementsParXY = make(map[int32]*VueEnvironnement)
	mm.RoutesParXY = make(map[int32]*VueRoute)
	mm.Villes = make([]*Ville, 0, 10)
	mm.LieuxVilles = make([]*LieuVille, 0, 10)
	mm.Régions = make([]*Région, 0, 10)
	return mm
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
			c.Fond = "plaine-gr"
		case "echoppe":
			c.Fond = "pavé"
		case "route":
			c.Fond = "route"
		case "ville":
			c.Fond = "pavé"
		}
	}
	for _, e := range mm.EchoppesParXY {
		m.Echoppes = append(m.Echoppes, e)
	}
	for _, c := range cases {
		m.Cases = append(m.Cases, c)
	}
	m.Villes = mm.Villes           // pour l'instant pour les villes c'est simple...
	m.LieuxVilles = mm.LieuxVilles // et pour leurs lieux aussi
	m.Régions = mm.Régions         // et les régions itou

	return m
}
