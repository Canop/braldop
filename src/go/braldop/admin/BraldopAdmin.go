package main

import (
	"braldop/bra"
	"flag"
	"fmt"
	"os"
)

func main() {

	exportePalette := flag.Bool("palette", false, "si oui alors la palette des environnements est exportée (exemple : \"bradmin -palette\")")
	flag.Parse()

	if *exportePalette {
		bra.ExportePalettePng(os.Stdout)
	} else {
		fmt.Fprintln(os.Stderr, "Usage :")
		flag.PrintDefaults()
	}
}
