package bra

/*
	Les fonctions ici définies permettent la construction ou l'enrichissement d'images contenant
	 les terrains d'une couche.	
*/

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const SEMI_LARGEUR = 800
const SEMI_HAUTEUR = 500

var palette color.Palette                  //  palette standard, utilisée pour les nouvelles images (peut être différent des images modifiées)
var indexes = make(map[string]uint8)       // donne les index des couleurs par type d'environnement
var couleurs = make(map[string]color.RGBA) // donne les couleurs par type d'environnement

// le cache est utilisé exclusivement par la fonction EnrichitPNG
var mutexCachePng = new(sync.Mutex)             // utilisé pour verrouiller le cache des png
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
		indexes[env+".p"] = uint8(index) // ceci ne sert qu'en raison du remappage possible de palette
		palette[index] = couleur254
		index++
	}
}

func couleursEgales(c1 color.RGBA, c2 color.RGBA) bool {
	return c1.R == c2.R && c1.G == c2.G && c1.B == c2.B && c1.A == c2.A
}

// dessine la couche sur l'image qui peut avoir une palette différente de la palette standard.
func dessine(img *image.Paletted, couche *Couche) {
	imgIndexes := make(map[string]uint8)   // similaire à indexes (fond->index) mais relatif à la palette de l'image et non à la palette standard	
	caseAPalissade := make(map[int32]bool) // map suivant PosKey(x,y) : true ssi une palissade est en x,y
	for _, p := range couche.Palissades {
		caseAPalissade[PosKey(p.X, p.Y)] = true
	}
	imgPalette := img.Palette
	déplacementsDansPalette := 0
	ajoutsPalette := 0
	n := len(imgPalette)
	nbAbsences := make(map[string]uint) // je note les fonds manquants dans ma palette, ils peuvent correspondre à des évolutions du jeu Braldahim
	for _, c := range couche.Cases {
		x, y := int(c.X)+SEMI_LARGEUR, SEMI_HAUTEUR-int(c.Y)
		key := c.Fond
		if caseAPalissade[PosKey(c.X, c.Y)] {
			key += ".p"
		}
		imgIndex, ok := imgIndexes[key] // index de la couleur du fond dans la palette de l'image
		if !ok {
			index, ok := indexes[key]
			if ok {
				c := palette[index].(color.RGBA)
				if index < uint8(n) && couleursEgales(c, imgPalette[index].(color.RGBA)) { // test rapide : si la couleur est au même index dans imgPalette que dans la palette standard
					imgIndex = index
					imgIndexes[key] = imgIndex
				} else {
					found := false
					for i := 0; i < n; i++ {
						if couleursEgales(c, imgPalette[i].(color.RGBA)) {
							found = true
							imgIndex = uint8(i)
							imgIndexes[key] = imgIndex
							break
						}
					}
					if found {
						déplacementsDansPalette++
					} else {
						log.Printf(" couleur \"%s\" absente de la palette de l'image\n", key)
						imgIndex = uint8(len(imgPalette))
						imgIndexes[key] = imgIndex
						img.Palette = append(img.Palette, c)
						imgPalette = img.Palette
						ajoutsPalette++
					}
				}
			} else { // fond inconnu y compris pour la palette standard
				nbAbsences[c.Fond] = nbAbsences[c.Fond] + 1
			}
		}
		img.SetColorIndex(x, y, imgIndex) // si pas ok, ça doit passer transparent (imgIndex=0)
	}
	if ajoutsPalette+déplacementsDansPalette != 0 {
		log.Println(" Transformations palette : ", déplacementsDansPalette, " déplacements et ", ajoutsPalette, "ajouts")
	}
	if len(nbAbsences) != 0 {
		log.Println(" Fonds manquants :")
		for fond, nb := range nbAbsences {
			log.Println(" ", fond, " : ", nb)
		}
	}
}

// Merge les deux images de telle sorte qu'après l'opération les pixels de img1 :
//   - soient inchangés là où l'index dans img2 est 0
//   - prennent la valeur de img2 ailleurs
// Les deux images doivent être de mêmes dimensions exactement
// La palette de img1 est enrichie si nécessaire
func Fusionne(img1 *image.Paletted, img2 *image.Paletted) error {
	r1 := img1.Bounds()
	r2 := img2.Bounds()
	if r1.Min.X != r2.Min.X || r1.Max.X != r2.Max.X || r1.Min.Y != r2.Min.Y || r1.Max.Y != r2.Max.Y {
		return errors.New("image dimensions not compatible")
	}
	p1 := img1.Palette
	p2 := img2.Palette
	n1 := len(p1)
	n2 := len(p2)
	op := make([]uint8, n2) // tableau donnant l'index dans la palette de img1 d'une couleur de la palette de img2
	for i2 := 0; i2 < n2; i2++ {
		c := p2[i2].(color.RGBA)
		if i2 < n1 && couleursEgales(c, p1[i2].(color.RGBA)) {
			op[i2] = uint8(i2)
		} else {
			found := false
			for i1 := 0; i1 < n1; i1++ {
				if couleursEgales(c, p1[i1].(color.RGBA)) {
					op[i2] = uint8(i1)
					found = true
					break
				}
			}
			if !found {
				op[i2] = uint8(len(p1))
				p1 = append(p1, c)
				img1.Palette = p1
			}
		}
	}
	l2 := len(img2.Pix)
	for x := 0; x < l2; x++ {
		if img2.Pix[x] != 0 {
			img1.Pix[x] = op[img2.Pix[x]]
		}
	}
	return nil
}

// utilise les images couchexxx.png du répertoire cheminRépertoireSource pour construire
//  ou enrichir les images couchexxx.png du répertoire cheminRépertoireEcriture
func EnrichitCouchesPNG(cheminRépertoireEcriture string, cheminRépertoireSource string) error {
	srcdir, err := os.Open(cheminRépertoireSource)
	if err!=nil {
		log.Println("Erreur ouverture répertoire source dans EnrichitCouchesPNG : ", err)
		return err
	}
	defer srcdir.Close()
	filenames, err := srcdir.Readdirnames(-1)
	if err!=nil {
		log.Println("Erreur listage fichiers sources dans EnrichitCouchesPNG : ", err)
		return err
	}
	for _, filename := range(filenames) {
		if filepath.Ext(filename)!=".png" {
			continue
		}
		bn := filepath.Base(filename)
		if !strings.HasPrefix(bn, "couche") {
			continue
		}
		var z int
		z, err = strconv.Atoi(bn[len("couche"):len(filename)-len(".png")])
		if err!=nil {
			continue 
		}
		log.Println("Fichier ok : ", bn, " z :", z)
		cheminFichierSource := filepath.Join(cheminRépertoireSource, bn)
		log.Println("cheminFichierSource : ", cheminFichierSource)
		cheminFichierDestination := filepath.Join(cheminRépertoireEcriture, bn)
		log.Println("cheminFichierDestination : ", cheminFichierDestination)
		sf, err := os.Open(cheminFichierSource)
		if err!=nil {
			return err 
		}		
		defer sf.Close()
		if df, err := os.Open(cheminFichierDestination); err == nil { // le fichier existe, on le charge
			destimg, _, err := image.Decode(df)
			df.Close()
			if err != nil {
				return err
			}
			srcimg, _, err := image.Decode(sf)
			if err != nil {
				return err
			}			
			Fusionne(destimg.(*image.Paletted), srcimg.(*image.Paletted))
			
			df, err = os.Create(cheminFichierDestination) // ouverture en écriture
			if err != nil {
				return err
			}
			defer df.Close()
			png.Encode(df, destimg)
			
		} else { // le fichier n'existe pas, il suffit de copier la source
			df, err := os.Create(cheminFichierDestination)
			if err != nil {
				return err
			}
			defer df.Close()
			io.Copy(df, sf)
		}
	}
	return nil
}


// ajoute les fonds de la couche passée à un PNG existant (ou un nouveau s'il n'existait pas)
//  nommé couchexxx.png dans le répertoire cheminRépertoire.
// Il est recommandé de faire tous les appels avec la même valeur de cacheSize.
// On backupe l'ancien fichier avant l'écriture pour qu'en
//  cas de crash durant l'écriture on puisse disposer de l'ancien fichier.
func EnrichitCouchePNG(cheminRépertoire string, couche *Couche, cacheSize int) {
	startTime := time.Nanoseconds()
	cheminFichierImage :=  filepath.Join(cheminRépertoire, fmt.Sprintf("couche%d.png", couche.Z))
	var img *image.Paletted
	cheminFichierBackup := ""
	mutexCachePng.Lock()
	defer mutexCachePng.Unlock()
	if élémentCache, ok := cachePng[cheminFichierImage]; !ok {
		if f, err := os.Open(cheminFichierImage); err == nil { // le fichier existe, on le charge
			ancienneImage, _, err := image.Decode(f)
			f.Close()
			if err == nil { // image décodée
				img = ancienneImage.(*image.Paletted)
				cheminFichierBackup = fmt.Sprintf("%s/couche%d-%s.png", cheminRépertoire, couche.Z, time.LocalTime().Format("20060102_1504_05.000"))
				os.Rename(cheminFichierImage, cheminFichierBackup)
			} else {
				log.Println("ERREUR DECODAGE IMAGE ", cheminFichierImage)
				img = image.NewPaletted(image.Rect(0, 0, SEMI_LARGEUR*2, SEMI_HAUTEUR*2), palette)
			}
		} else {
			log.Println(" Pas de fichier image existant")
			img = image.NewPaletted(image.Rect(0, 0, SEMI_LARGEUR*2, SEMI_HAUTEUR*2), palette)
		}
		if cacheSize > 0 { // mise en cache, et réduction du cache si nécessaire
			log.Println(" Image pas en cache")
			now := time.Nanoseconds()
			for len(cachePng) > cacheSize {
				oldestImagePath := ""
				oldestAge := now
				for path, élément := range cachePng {
					if élément.vue <= oldestAge {
						oldestAge = élément.vue
						oldestImagePath = path
					}
				}
				delete(cachePng, oldestImagePath)
				log.Println(" Image supprimée du cache : " + oldestImagePath)
			}
			log.Println(" Ajout au cache : " + cheminFichierImage)
			cachePng[cheminFichierImage] = élémentCachePng{img, now}
		}
	} else {
		log.Println(" Image trouvée en cache")
		img = élémentCache.img
	}

	dessine(img, couche)

	f, err := os.Create(cheminFichierImage)
	if err != nil {
		log.Println(" Erreur à la création du fichier ", cheminFichierImage)
		return
	}
	//~ log.Println(" Ecriture du fichier ", cheminFichierImage)
	defer f.Close()
	png.Encode(f, img)

	if cheminFichierBackup != "" {
		os.Remove(cheminFichierBackup)
	}

	log.Printf(" Enrichissement carte PNG en %d ms\n", (time.Nanoseconds()-startTime)/1e6)
}

// construit l'image PNG d'une couche
func ConstruitNouveauPNG(cheminRépertoire string, couche *Couche) {
	startTime := time.Nanoseconds()
	img := image.NewPaletted(image.Rect(0, 0, SEMI_LARGEUR*2, SEMI_HAUTEUR*2), palette)
	dessine(img, couche)
	cheminFichierImage := fmt.Sprintf("%s/couche%d.png", cheminRépertoire, couche.Z)
	f, err := os.Create(cheminFichierImage)
	if err != nil {
		log.Println("Erreur à la création du fichier ", cheminFichierImage)
		return
	}
	defer f.Close()
	png.Encode(f, img)
	log.Printf("Construction carte PNG en %d ms\n", (time.Nanoseconds()-startTime)/1e6)
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
//  sont celles qui sont potentiellement destructives). Celles en cours s'achêvent.
func BloqueEcrituresPNG() {
	mutexCachePng.Lock()
}
