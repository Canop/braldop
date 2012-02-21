package main

// test de Version. C'est pas que ce soit bien compliqué mais je voulais voir comment fonctionne
//  le framework de test...

import (
	"testing"
)

func testParse(t *testing.T, s string, ok bool, parts ...uint64) {
	v, e := ParseVersion(s)
	for i, p := range(parts) {
		if i>=len(v.Parts) || p!=v.Parts[i] {
			t.Errorf("Version \"%s\" should be parsed as %s", s, Version{parts})
		}
	}
	if (e==nil) != ok {
		if ok {
			t.Errorf("Unexpected error while parsing \"%s\"", s)
		} else {
			t.Errorf("Unexpected lack of error while parsing \"%s\". Gave  %s", s, v)
		}
	}
}

func testParseAndCompare(t *testing.T, s1 string, s2 string, c int) {
	v1, _ := ParseVersion(s1)
	v2, _ := ParseVersion(s2)
	if CompareVersions(&v1, &v2)!=c {
		t.Errorf("CompareVersions(\"%s\",\"%s\") should return %d", s1, s2, c)
	}
}

func TestVersion(t *testing.T) {
	testParse(t, "version absente", false)
	testParse(t, "", false)
	testParse(t, ".1", false)
	testParse(t, "0", true, 0)
	testParse(t, "1.2.3.5", true, 1, 2, 3, 5)
	testParse(t, "1.2.3.b", false, 1, 2, 3)
	testParse(t, "1.2.3b", false, 1, 2)
	testParse(t, "1.2.02", true, 1, 2, 2)
	testParse(t, "1.2.02.", false, 1, 2, 2)
	testParse(t, "1.-2.02", false, 1)
	testParseAndCompare(t, "0.0.0", "", 1)
	testParseAndCompare(t, "", "0.0.0", -1)
	testParseAndCompare(t, "", ".0.0.0", 0)
	testParseAndCompare(t, "0.2.0.0", "0.2.0", 1)
	testParseAndCompare(t, "0.2.a.0", "0.2", 0)
	testParseAndCompare(t, "1.2.3..5", " 1.2.3.4a ", 0)
}
