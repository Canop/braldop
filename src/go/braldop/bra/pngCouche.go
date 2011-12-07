package bra

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"sync"
	"time"
)

const SEMI_LARGEUR = 800
const SEMI_HAUTEUR = 500

var palette color.Palette
var indexes = make(map[string]uint8)       // donne les index des couleurs par type d'environnement
var couleurs = make(map[string]color.RGBA) // donne les couleurs par type d'environnement

// le cache est utilisé exclusivement par la fonction EnrichitPNG
var mutexCachePng = new(sync.Mutex) // utilisé pour verrouiller le cache des png
var cachePng = make(map[string]élémentCachePng) // cache des images, la clef est le chemin du fichier
type élémentCachePng struct {
	img *image.Paletted
	vue int64 // date dernier accès (Nanoseconds)
}


// initialise la palette de couleurs des cartes PNG
func init() {
	// attention : Les couleurs suivantes doivent impérativement être toutes différentes.
	//             Elles seront utilisées par le client pour connaitre le terrain.
	// pour les palissades, on joue sur l'alpha
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
	couleurs["peupliers-gr"] = color.RGBA{176, 217, 92, 255}    // #b0d95c
	couleurs["caverne-crevasse"] = color.RGBA{78, 65, 100, 255} // #4e4164
	couleurs["caverne"] = color.RGBA{163, 145, 159, 255}        // #a3919f

	palette = make(color.Palette, 2*len(couleurs)+1)
	palette[0] = color.RGBA{0, 0, 0, 0}
	index := 1
	for env, couleur := range couleurs {
		// les couleurs de base, correspondant aux environnements de fond : alpha = 255
		indexes[env] = uint8(index)
		palette[index] = couleur
		index++
		// la couleur traduisant la présence d'une palissade : alpha = 254
		couleur254 := color.RGBA{couleur.R, couleur.G, couleur.B, 254}
		palette[index] = couleur254
		index++
	}
}


func (couche *Couche) dessine(img *image.Paletted) {
	caseAPalissade := make(map[int32]bool) // map suivant PosKey(x,y) : true ssi une palissade est en x,y
	for _, p := range couche.Palissades {
		caseAPalissade[PosKey(p.X, p.Y)] = true
	}

	nbAbsences := make(map[string]uint) // je note les fonds manquants dans ma palette, ils peuvent correspondre à des évolutions du jeu Braldahim
	for _, c := range couche.Cases {
		if colorIndex, ok := indexes[c.Fond]; ok {
			x, y := int(c.X)+SEMI_LARGEUR, SEMI_HAUTEUR-int(c.Y)
			if caseAPalissade[PosKey(c.X, c.Y)] {
				colorIndex++
			}
			img.SetColorIndex(x, y, colorIndex)
		} else {
			nbAbsences[c.Fond] = nbAbsences[c.Fond] + 1
		}
	}
	if len(nbAbsences) > 0 {
		fmt.Println("Fonds manquants :")
		for fond, nb := range nbAbsences {
			fmt.Println(" ", fond, " : ", nb)
		}
	}

}

// ajoute les fonds à un PNG existant (ou un nouveau s'il n'existait pas).
// Il est recommandé de faire tous les appels avec la même valeur de cacheSize.
// On backupe l'ancien fichier avant l'écriture pour qu'en
//  cas de crash durant l'écriture on puisse disposer de l'ancien fichier.
func (couche *Couche) EnrichitPNG(cheminRépertoire string, cacheSize int) {
	startTime := time.Nanoseconds()
	cheminFichierImage := fmt.Sprintf("%s/couche%d.png", cheminRépertoire, couche.Z)
	var img *image.Paletted
	cheminFichierBackup := ""
	mutexCachePng.Lock()
	defer mutexCachePng.Unlock()
	if élémentCache, ok := cachePng[cheminFichierImage]; !ok {
		img = image.NewPaletted(image.Rect(0, 0, SEMI_LARGEUR*2, SEMI_HAUTEUR*2), palette)
		if cacheSize>0 { // mise en cache, et réduction du cache si nécessaire
			fmt.Println(" Image pas en cache")
			now := time.Nanoseconds()
			for len(cachePng)>cacheSize {
				oldestImagePath := ""
				oldestAge := now
				for path, élément := range(cachePng) {
					if élément.vue <= oldestAge {
						oldestAge = élément.vue
						oldestImagePath = path
					}
				}
				delete(cachePng, oldestImagePath)
				fmt.Println(" Image supprimée du cache : " + oldestImagePath)
			}
			fmt.Println(" Ajout au cache : " + cheminFichierImage)
			cachePng[cheminFichierImage] = élémentCachePng{img, now}
		}
		if f, err := os.Open(cheminFichierImage); err == nil { // le fichier existe, on le charge
			ancienneImage, _, err := image.Decode(f)
			f.Close()
			if err==nil { // image décodée
				cheminFichierBackup = fmt.Sprintf("%s/couche%d-%s.png", cheminRépertoire, couche.Z, time.LocalTime().Format("20060102_1504_05.000"))
				os.Rename(cheminFichierImage, cheminFichierBackup)
				rectangleAncienneImage := ancienneImage.Bounds()
				if rectangleAncienneImage.Min.X != 0 || rectangleAncienneImage.Min.Y != 0 || rectangleAncienneImage.Max.X != SEMI_LARGEUR*2 || rectangleAncienneImage.Max.Y != SEMI_HAUTEUR*2 {
					fmt.Printf("Dimensions ancienne image incorrectes : %+v\n", rectangleAncienneImage)
				} else {
					draw.Draw(img, ancienneImage.Bounds(), ancienneImage, ancienneImage.Bounds().Min, draw.Src)
				}
			}
		} else {
			fmt.Println("Pas de fichier image existant")
		}
	} else {
		fmt.Println(" Image trouvée en cache")
		img = élémentCache.img
	}
	
	couche.dessine(img)

	f, err := os.Create(cheminFichierImage)
	if err != nil {
		fmt.Println("Erreur à la création du fichier ", cheminFichierImage)
		return
	}
	fmt.Println("Création du fichier ", cheminFichierImage)
	defer f.Close()
	png.Encode(f, img)

	if cheminFichierBackup!="" {
		os.Remove(cheminFichierBackup)
	}

	fmt.Printf("Enrichissement carte PNG en %d ms\n", (time.Nanoseconds()-startTime)/1e6)
}


// construit l'image PNG d'une couche
func (couche *Couche) ConstruitNouveauPNG(cheminRépertoire string) {
	startTime := time.Nanoseconds()
	img := image.NewPaletted(image.Rect(0, 0, SEMI_LARGEUR*2, SEMI_HAUTEUR*2), palette)
	couche.dessine(img)
	cheminFichierImage := fmt.Sprintf("%s/couche%d.png", cheminRépertoire, couche.Z)
	f, err := os.Create(cheminFichierImage)
	if err != nil {
		fmt.Println("Erreur à la création du fichier ", cheminFichierImage)
		return
	}
	defer f.Close()
	png.Encode(f, img)
	fmt.Printf("Construction carte PNG en %d ms\n", (time.Nanoseconds()-startTime)/1e6)
}

// décrit la palette pour une inclusion plus aisée dans le javascript
// 
func ExportePalettePng(w io.Writer) {
	fmt.Fprintln(w, "Palette des environnements")
	for nom, c := range couleurs {
		v := (uint32(c.R) << 16) + (uint32(c.G) << 8) + uint32(c.B)
		vp := (uint32(c.R-1) << 16) + (uint32(c.G-1) << 8) + uint32(c.B-1) // couleur visible dans getImageData pour un alpha de 254
		fmt.Fprintf(w, "\t%d: \"%s\", %d: \"%s\",\n", v, nom, vp, nom)
	}
}

// arrête (pour un arrêt du logiciel) les écritures faites via EnrichitPNG (ce
//  sont celles qui. Celles en cours s'achêvent.
func BloqueEcrituresPNG() {
	mutexCachePng.Lock()
}
