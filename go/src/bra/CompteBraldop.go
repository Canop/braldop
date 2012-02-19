package bra

import ()

// cette structure est persistée dans la table mysql compte (voir store_comptes.go)
type CompteBraldop struct {
	IdBraldun         uint
	Mdpr              string
	Authentifié       bool
	X, Y, Z           int16
	DernièreMiseAJour int64 // secondes
}

func (cb *CompteBraldop) Clone() *CompteBraldop {
	return &CompteBraldop{cb.IdBraldun, cb.Mdpr, cb.Authentifié, cb.X, cb.Y, cb.Z, cb.DernièreMiseAJour}
}
