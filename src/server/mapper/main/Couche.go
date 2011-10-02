package main

/*
couche, exportable en json par exemple.
*/

type Case struct {
	X    int16
	Y    int16
	Fond string
}

type Couche struct {
	Z        int16
	Cases    []*Case
	Champs   []*VueChamp
	Echoppes []*VueEchoppe
}

func NewCouche() (c *Couche) {
	c = new(Couche)
	c.Cases = make([]*Case, 0, 40)
	c.Echoppes = make([]*VueEchoppe, 0, 40)
	return
}
