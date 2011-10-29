package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const SEMI_LARGEUR = 800
const SEMI_HAUTEUR = 500

func (couche *Couche) ConstruitPNG(cheminRépertoire string) {
	img := image.NewRGBA(image.Rect(0, 0, SEMI_LARGEUR*2, SEMI_HAUTEUR*2))

	couleurs := make(map[string]color.RGBA)
	nbAbsences := make(map[string]uint)

	couleurs["plaine"] = color.RGBA{183, 221, 129, 255}         // #b7dd81
	couleurs["plaine-gr"] = color.RGBA{145, 192, 117, 255}      // #91c075
	couleurs["peuprofonde"] = color.RGBA{100, 140, 195, 255}    // #648cc3
	couleurs["pave"] = color.RGBA{211, 203, 202, 255}           // #d3cbca
	couleurs["route"] = color.RGBA{198, 195, 187, 255}          // #c6c3bb
	couleurs["montagne"] = color.RGBA{210, 185, 170, 255}       // #d2b9aa
	couleurs["montagne-gr"] = color.RGBA{172, 148, 139, 255}    // #ac948b
	couleurs["gazon"] = color.RGBA{120, 202, 74, 255}           // #78ca4a
	couleurs["marais"] = color.RGBA{184, 227, 200, 255}         // #b8e3c8
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

	for _, c := range couche.Cases {
		if couleur, ok := couleurs[c.Fond]; ok {
			x, y := int(c.X)+SEMI_LARGEUR, SEMI_HAUTEUR-int(c.Y)
			img.SetRGBA(x, y, couleur)
		} else {
			nbAbsences[c.Fond] = nbAbsences[c.Fond] + 1
		}
	}

	cheminFichierImage := fmt.Sprintf("%s/couche%d.png", cheminRépertoire, couche.Z)
	f, err := os.Create(cheminFichierImage)
	if err != nil {
		fmt.Printf("Erreur à la création du fichier : %s", cheminFichierImage)
		return
	}
	png.Encode(f, img)

	fmt.Println("Fonds manquants :")
	for fond, nb := range nbAbsences {
		fmt.Println(fond, " : ", nb)
	}
}
