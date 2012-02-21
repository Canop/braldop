package main

// représente une version, sous forme d'une séquence d'entiers positifs ou nuls.
// Initialiser une version :
//     v := MakeVersion(4, 3, 1) 
// ou bien
//     v := ParseVersion("4.3.1")
// Calcul d'antériorité :
//     CompareVersions(&MakeVersion(2, 7, 7), &MakeVersion(3,0)) // == -1
// La version vide (Version{}) est inférieure à toutes les autres versions et
//  cette logique s'applique aux parties ("1.2.0" > "1.2")

import (
	"strconv"
	"strings"
)

type Version struct {
	Parts []uint64
}

func MakeVersion(parts ...uint64) Version {
	return Version{parts}
}

func (v Version) String() string {
	sp := make([]string, len(v.Parts))
	for i, p := range v.Parts {
		sp[i] = strconv.FormatUint(uint64(p), 10)
	}
	return strings.Join(sp, ".")
}

// construit la version correspondant à la chaine passée.
// En cas de composante incomprise, renvoie la version comprise jusqu'à cette composante ainsi qu'une erreur
// Si l'on ignore les erreurs (elles en provoquent toutes), les versions suivantes sont équivalentes :
//    "1.2.3."    " 1.2.3.4a "   "1.2.3.b"   "1.2.3..5"
func ParseVersion(s string) (Version, error) {
	tokens := strings.Split(strings.Trim(s, " "), ".")
	parts := make([]uint64, len(tokens))
	for i, t := range tokens {
		if n, pe := strconv.ParseUint(t, 10, 0); pe != nil {
			return Version{parts[:i]}, pe
		} else {
			parts[i] = n
		}
	}
	return Version{parts}, nil
}

// a == b : returns 0
// a > b  : returns 1
// a < b  : return -1
func CompareVersions(a *Version, b *Version) int {
	for i := 0; ; i++ {
		if i >= len(a.Parts) {
			if i >= len(b.Parts) {
				return 0
			} else {
				return -1
			}
		}
		if i >= len(b.Parts) {
			return 1
		}
		if a.Parts[i] > b.Parts[i] {
			return 1
		} else if a.Parts[i] < b.Parts[i] {
			return -1
		}
	}
	return 0
}
