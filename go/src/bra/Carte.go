package bra

/*
carte compilée, exportable en json par exemple.
*/

type Carte struct {
	Couches     []*Couche // les couches correspondant chacune à une profondeur
	Villes      []*Ville
	LieuxVilles []*LieuVille
	Régions     []*Région
	Vues        []*Vue
	Communautés []*Communauté // on exporte l'ensemble des communautés, il n'y en a quasiment pas
}

func NewCarte() (m *Carte) {
	m = new(Carte)
	m.Couches = make([]*Couche, 0, 3)
	m.Vues = make([]*Vue, 0, 1)
	return m
}
