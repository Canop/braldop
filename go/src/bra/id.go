package bra

import (
	"strconv"
)

// renvoie 0 si la chaine a ne contient pas un id non signé. Utiliser strconv.ParseUint si l'erreur est nécessaire.
func AtoId(a string) uint {
	id64, _ := strconv.ParseUint(a, 10, 0)
	return uint(id64)
}

