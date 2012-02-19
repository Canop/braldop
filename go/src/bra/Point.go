package bra

import "strconv"

type Point struct {
	X int16
	Y int16
	Z int16
}

// pas le plus efficace mais j'ai pas envie de refaire le boulot fait dans strconv.Btoui64...
func Atoi16(s string) (int16, error) {
	i, err := strconv.Atoi(s)
	return int16(i), err
}

func Atoui(a string) (uint, error) {
	i, err := strconv.ParseUint(a, 10, 0)
	return uint(i), err
}
func Atoi(a string) (int, error) {
	i, err := strconv.ParseInt(a, 10, 0)
	return int(i), err
}


// fournit une clé standard Int32 à partir de deux nombres Int16
func PosKey(x int16, y int16) int32 {
	return (int32(x) << 16) + (int32(y))
}

func (o *Point) readCsvPoint(cells []string) {
	o.X, _ = Atoi16(cells[1])
	o.Y, _ = Atoi16(cells[2])
	o.Z, _ = Atoi16(cells[3])
}
