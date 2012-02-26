package bra

import (
	"log"
)

type MemMap struct {
	Bralduns      map[uint]*Braldun // il s'agit des bralduns récupérés depuis le fichier statique donc certaines informations sont absentes
	Communautés   map[uint]*Communauté
	Couches       map[int16]*MemCouche // les couches, indexées par profondeur
	Villes        []*Ville
	LieuxVilles   []*LieuVille
	Régions       []*Région
	DernièresVues map[uint]*Vue // les dernières vues avec pour clé l'id du Braldun voyeur
}

func NewMemMap() (mm *MemMap) {
	mm = new(MemMap)
	mm.Couches = make(map[int16]*MemCouche)
	mm.Bralduns = make(map[uint]*Braldun)
	mm.Communautés = make(map[uint]*Communauté)
	mm.Villes = make([]*Ville, 0, 10)
	mm.LieuxVilles = make([]*LieuVille, 0, 10)
	mm.Régions = make([]*Région, 0, 10)
	mm.DernièresVues = make(map[uint]*Vue)
	return mm
}

// renvoie la couche de profondeur z, la créant si nécessaire
func (mm *MemMap) GetCouche(z int16) *MemCouche {
	c, ok := mm.Couches[z]
	if !ok {
		c = NewMemCouche()
		c.Z = z
		mm.Couches[z] = c
	}
	return c
}

func (mm *MemMap) StoreBosquet(o *VueBosquet) {
	mm.GetCouche(o.Z).BosquetsParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreChamp(o *VueChamp) {
	mm.GetCouche(o.Z).ChampsParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreCrevasse(o *VueCrevasse) {
	mm.GetCouche(o.Z).CrevassesParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreEchoppe(o *VueEchoppe) {
	mm.GetCouche(o.Z).EchoppesParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreEnvironnement(o *VueEnvironnement) {
	mm.GetCouche(o.Z).EnvironnementsParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StorePalissade(o *VuePalissade) {
	mm.GetCouche(o.Z).PalissadesParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreRoute(o *VueRoute) {
	mm.GetCouche(o.Z).RoutesParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreVille(o *Ville) { // on va supposer qu'on ne lit pas deux fois la même ville
	mm.Villes = append(mm.Villes, o)
}
func (mm *MemMap) StoreLieuVille(o *LieuVille) { // on va supposer qu'on ne lit pas deux fois le même lieu
	mm.LieuxVilles = append(mm.LieuxVilles, o)
}
func (mm *MemMap) StoreLieu(o *VueLieu) {
	mm.GetCouche(o.Z).LieuxParXY[PosKey(o.X, o.Y)] = o
}
func (mm *MemMap) StoreRégion(o *Région) {
	mm.Régions = append(mm.Régions, o)
}

func (mm *MemMap) CompleteBralduns(v *Vue) {
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
			b.IdCommunauté = mmb.IdCommunauté
			b.PointsGredin = mmb.PointsGredin
			b.PointsRedresseur = mmb.PointsRedresseur
			b.PointsDistinction = mmb.PointsDistinction
		} else {
			log.Printf("Braldun introuvable : %d\n", b.Id)
		}
	}
}

func (mm *MemMap) Compile() (carte *Carte) {
	carte = NewCarte()
	for _, mc := range mm.Couches {
		carte.Couches = append(carte.Couches, mc.Compile(mm))
	}
	carte.Villes = mm.Villes
	carte.LieuxVilles = mm.LieuxVilles
	carte.Régions = mm.Régions
	for _, v := range mm.DernièresVues {
		mm.CompleteBralduns(v)
	}
	// pour les communautés, j'ai des soucis avec l'export json des maps donc je fais un gros tableau
	maxId := 0
	for i, _ := range mm.Communautés {
		if int(i) > maxId {
			maxId = int(i)
		}
	}
	carte.Communautés = make([]*Communauté, maxId+1, maxId+1)
	for i, c := range mm.Communautés {
		carte.Communautés[i] = c
	}
	return
}
