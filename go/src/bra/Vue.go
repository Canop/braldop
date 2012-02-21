package bra

// Objet exportable en json correspondant à la vue d'un Braldun (en fait pas tous les objets pour l'instant
//  certains, les plus statiques (crevasses, champs, échoppes) étant stockés dans la memmap)

type Vue struct {
	Z            int16 // la profondeur de la couche, 0 pour la surface
	Time         int64 // secondes depuis 1970. Une date à 0 signifie généralement que l'objet est vide ou invalide
	Voyeur       uint  // id du braldun. Un id à 0 signifie que l'objet est vide ou invalide
	PrénomVoyeur string
	XMin         int16
	XMax         int16
	YMin         int16
	YMax         int16
	Bralduns     []*Braldun
	Cadavres     []*VueCadavre
	Monstres     []*VueMonstre
	Objets       []*VueObjet
}

func NewVue() *Vue {
	vue := new(Vue)
	vue.Bralduns = make([]*Braldun, 0, 5)
	vue.Cadavres = make([]*VueCadavre, 0, 5)
	vue.Monstres = make([]*VueMonstre, 0, 0)
	vue.Objets = make([]*VueObjet, 0, 0)
	return vue
}

// effectue un clonage à un niveau : les objets (Braldun, etc.) ne
//  sont pas eux même clonés
func (vue *Vue) clone() (c *Vue) {
	c = new(Vue)
	c.Z = vue.Z
	c.Time = vue.Time
	c.Voyeur = vue.Voyeur
	c.PrénomVoyeur = vue.PrénomVoyeur
	c.XMin = vue.XMin
	c.XMax = vue.XMax
	c.YMin = vue.YMin
	c.YMax = vue.YMax
	c.Bralduns = make([]*Braldun, 0, len(vue.Bralduns))
	copy(c.Bralduns, vue.Bralduns[0:])
	c.Cadavres = make([]*VueCadavre, 0, len(vue.Cadavres))
	copy(c.Cadavres, vue.Cadavres[0:])
	c.Monstres = make([]*VueMonstre, 0, len(vue.Monstres))
	copy(c.Monstres, vue.Monstres[0:])
	c.Objets = make([]*VueObjet, 0, len(vue.Objets))
	copy(c.Objets, vue.Objets[0:])
	return c
}

// non testé
func (vue *Vue) cloneCoupé(xmin, xmax, ymin, ymax int16) (c *Vue) {
	c = new(Vue)
	c.Z = vue.Z
	c.Time = vue.Time
	c.Voyeur = vue.Voyeur
	c.PrénomVoyeur = vue.PrénomVoyeur
	c.XMin = vue.XMin
	c.XMax = vue.XMax
	c.YMin = vue.YMin
	c.YMax = vue.YMax
	if xmin > vue.XMax || vue.XMin < xmax || ymin > vue.YMax || vue.YMin < ymax {
		c.Bralduns = make([]*Braldun, len(vue.Bralduns))
		c.Cadavres = make([]*VueCadavre, len(vue.Cadavres))
		c.Monstres = make([]*VueMonstre, len(vue.Monstres))
		c.Objets = make([]*VueObjet, len(vue.Objets))
		copy(c.Bralduns, vue.Bralduns[0:])
		copy(c.Monstres, vue.Monstres[0:])
		copy(c.Cadavres, vue.Cadavres[0:])
		copy(c.Objets, vue.Objets[0:])
	} else {
		c.Bralduns = make([]*Braldun, 0, len(vue.Bralduns))
		c.Cadavres = make([]*VueCadavre, 0, len(vue.Cadavres))
		c.Monstres = make([]*VueMonstre, 0, len(vue.Monstres))
		c.Objets = make([]*VueObjet, 0, len(vue.Objets))
		for _, o := range vue.Bralduns {
			if o.X < xmin || o.X > xmax || o.Y < ymin || o.Y > ymax {
				c.Bralduns = append(c.Bralduns, o)
			}
		}
	}
	return c
}

// indique si (x,y) est en dehors des zones de vue
func PointEnDehors(x, y int16, vues []*Vue) bool {
	for _, vj := range vues {
		if x >= vj.XMin && x <= vj.XMax && y >= vj.YMin && y <= vj.YMax {
			return false
		}
	}
	return true
}
