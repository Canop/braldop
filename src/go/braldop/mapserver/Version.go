package main

import (
	"errors"
	"strconv"
	"strings"
)

type Version struct {
	Parts []uint
}

func (v *Version) String() string {
	s := ""
	for i, p := range v.Parts {
		if i > 0 {
			s += "."
		}
		s += strconv.Uitoa(p)
	}
	return s
}

func ParseVersion(s string) (v *Version, err error) {
	if len(s) == 0 {
		return nil, errors.New("empty string")
	}
	tokens := strings.Split(strings.Trim(s, " "), ".")
	v = new(Version)
	v.Parts = make([]uint, len(tokens))
	for i, t := range tokens {
		if n, pe := strconv.Atoui(t); pe != nil {
			return nil, pe
		} else {
			v.Parts[i] = n
		}
	}
	return v, nil
}

/**
 * a == b : returns 0
 * a > b  : returns 1
 * a < b  : return -1
 */
func CompareVersions(a *Version, b *Version) int {
	for i := 0; i < 100; i++ {
		if i >= len(a.Parts) {
			if i >= len(b.Parts) {
				return 0
			} else {
				return -1
			}
		}
		if i > len(b.Parts) {
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
