package bra

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

const SEMI_LARGEUR = 800
const SEMI_HAUTEUR = 500

var palette color.Palette
var indexes = make(map[string]uint8) // donne les index des couleurs par type d'environnement

// initialise la palette de couleurs des cartes PNG
func init() {
	// attention : Les couleurs suivantes doivent impérativement être toutes différentes.
	//             Elles seront utilisées par le client pour connaitre le terrain.
	couleurs := make(map[string]color.RGBA)                     // donne les couleurs par type d'environnement
	couleurs["plaine"] = color.RGBA{183, 221, 129, 255}         // #b7dd81
	couleurs["plaine-gr"] = color.RGBA{145, 192, 117, 255}      // #91c075
	couleurs["peuprofonde"] = color.RGBA{100, 140, 195, 255}    // #648cc3
	couleurs["pave"] = color.RGBA{211, 203, 202, 255}           // #d3cbca
	couleurs["route"] = color.RGBA{198, 195, 187, 255}          // #c6c3bb
	couleurs["montagne"] = color.RGBA{210, 185, 170, 255}       // #d2b9aa
	couleurs["montagne-gr"] = color.RGBA{172, 148, 139, 255}    // #ac948b
	couleurs["gazon"] = color.RGBA{120, 202, 74, 255}           // #78ca4a
	couleurs["gazon-gr"] = color.RGBA{188, 237, 166, 255}       // #bceda6
	couleurs["marais"] = color.RGBA{184, 227, 200, 255}         // #b8e3c8
	couleurs["marais-gr"] = color.RGBA{132, 187, 149, 255}      // #84bb95
	couleurs["profonde"] = color.RGBA{74, 110, 153, 255}        // #4a6e99
	couleurs["tunnel"] = color.RGBA{184, 152, 97, 255}          // #b89861
	couleurs["mine"] = color.RGBA{156, 125, 123, 255}           // #9c7d7b
	couleurs["hetres"] = color.RGBA{170, 174, 91, 255}          // #aaae5b
	couleurs["hetres-gr"] = color.RGBA{122, 134, 73, 255}       // #7a8649
	couleurs["erables"] = color.RGBA{180, 134, 91, 255}         // #b4865b
	couleurs["erables-gr"] = color.RGBA{182, 120, 79, 255}      // #b6784f
	couleurs["chenes"] = color.RGBA{108, 156, 94, 255}          // #6c9c5e
	couleurs["lac"] = color.RGBA{94, 143, 195, 255}             // #5e8fc3
	couleurs["peupliers"] = color.RGBA{151, 217, 92, 255}       // #97d95c
	couleurs["caverne-crevasse"] = color.RGBA{78, 65, 100, 255} // #4e4164
	couleurs["caverne"] = color.RGBA{163, 145, 159, 255}        // #a3919f

	palette = make(color.Palette, len(couleurs)+1)
	palette[0] = color.RGBA{0, 0, 0, 0}
	index := 1
	for env, couleur := range couleurs {
		indexes[env] = uint8(index)
		palette[index] = couleur
		index++
	}
}

// construit l'image PNG d'une couche
func (couche *Couche) ConstruitPNG(cheminRépertoire string) {
	startTime := time.Nanoseconds()

	img := image.NewPaletted(image.Rect(0, 0, SEMI_LARGEUR*2, SEMI_HAUTEUR*2), palette)

	nbAbsences := make(map[string]uint) // je note les fonds manquants dans ma palette, ils peuvent correspondre à des évolutions du jeu Braldahim
	for _, c := range couche.Cases {
		if colorIndex, ok := indexes[c.Fond]; ok {
			x, y := int(c.X)+SEMI_LARGEUR, SEMI_HAUTEUR-int(c.Y)
			img.SetColorIndex(x, y, colorIndex)
		} else {
			nbAbsences[c.Fond] = nbAbsences[c.Fond] + 1
		}
	}

	cheminFichierImage := fmt.Sprintf("%s/couche%d.png", cheminRépertoire, couche.Z)
	f, err := os.Create(cheminFichierImage)
	if err != nil {
		fmt.Println("Erreur à la création du fichier ", cheminFichierImage)
		return
	}
	defer f.Close()
	png.Encode(f, img)

	if len(nbAbsences) > 0 {
		fmt.Println("Fonds manquants :")
		for fond, nb := range nbAbsences {
			fmt.Println(" ", fond, " : ", nb)
		}
	}

	fmt.Printf("Construction carte PNG en %d ms\n", (time.Nanoseconds()-startTime)/1e6)
}
