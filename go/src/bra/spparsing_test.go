package bra

// test des fonctions de parsing des scripts publics définies dans spparsing.go

import (
	"bufio"
	"bytes"
	"testing"
	"time"
)

func assertint(t *testing.T, label string, a, b int) {
	if a != b {
		t.Errorf("Unexpected value for %s : %d instead of %d", label, a, b)
	}
}

func testLitBraldun(t *testing.T) {
	csv := []byte(
		`idBraldun;prenom;nom;x;y;z;paRestant;DLA;DureeProchainTour;dateDebutTour;dateFinTour;dateFinLatence;dateDebutCumul;dureeCourantTour;dureeBmTour;PvRestant;bmPVmax;bbdf;nivAgilite;nivForce;nivVigueur;nivSagesse;bmAgilite;bmForce;bmVigueur;bmSagesse;bmBddfAgilite;bmBddfForce;bmBddfVigueur;bmBddfSagesse;bmVue;regeneration;bmRegeneration;pxPerso;pxCommun;pi;niveau;poidsTransportable;poidsTransporte;armureNaturelle;armureEquipement;bmAttaque;bmDegat;bmDefense;nbKo;nbKill;nbKoBraldun;estEngage;estEngageProchainTour;estIntangible;nbPlaquagesSubis;nbPlaquagesEffectues
754;Canopée;du Haut-Rac;5;6;0;4;2011-09-12 02:51:54;24:00:00;2011-09-11 00:18:54;2011-09-12 02:51:54;2011-09-11 06:57:09;2011-09-11 13:35:24;26:33:00;0;26;0;38;1;0;0;0;0;0;0;0;0;0;0;0;0;1;0;6;0;0;1;3;1.151;2;0;0;0;0;0;0;0;oui;non;non;0;0
`)
	csvreader := bufio.NewReader(bytes.NewReader(csv))
	eb, err := LitEtatBraldunDansCsvProfil(csvreader)
	if err != nil {
		t.Errorf("Unexpected error while parsing \"%s\"", err.Error())
	}
	if eb == nil {
		t.Error("LitEtatBraldunDansCsvProfil didn't return an EtatBraldun")
		return
	}
	assertint(t, "IdBraldun", int(eb.IdBraldun), 754)
	assertint(t, "PA", eb.PA, 4)
	duréeTour, _ := time.ParseDuration("24h")
	assertint(t, "DuréeTour", eb.DuréeTour, int(duréeTour.Seconds()))
	assertint(t, "DLA", eb.DLA, 1315693134) // calculé indépendemment en javascript : (new Date("2011-09-11 00:18:54")).getTime()/1000
	assertint(t, "PV", eb.PV, 26)
	assertint(t, "PVMax", eb.PVMax, 40)
	assertint(t, "Faim", eb.Faim, 38)
}

func TestSpParsing(t *testing.T) {
	testLitBraldun(t)
}
