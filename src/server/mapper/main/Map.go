package main

/*
carte compilée, exportable en json par exemple.
*/

import (
//~ "fmt"
)

type Case struct {
	X    int16
	Y    int16
	Fond string
}

type Map struct {
	Cases       []*Case
	Villes      []*Ville
	LieuxVilles []*LieuVille
	Régions     []*Région
}

func NewMap() (m *Map) {
	m = new(Map)
	m.Cases = make([]*Case, 0, 40)
	return m
}
