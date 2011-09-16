package main

/*
carte compilée, exportable en json par exemple.
*/

type Case struct {
	X    int16
	Y    int16
	Fond string
}

type Map struct {
	Cases       []*Case
	Champs      []*VueChamp
	Echoppes    []*VueEchoppe
	Villes      []*Ville
	LieuxVilles []*LieuVille
	Régions     []*Région
}

func NewMap() (m *Map) {
	m = new(Map)
	m.Cases = make([]*Case, 0, 40)
	m.Echoppes = make([]*VueEchoppe, 0, 40)
	return m
}
