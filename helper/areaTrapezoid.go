package helper

import "fmt"

func AreaTrapezoid(a, b, h float64) {
	// Masukkan nilai
	// panjang alas (a),
	// panjang atas (b),
	// dan tinggi (h)
	//trapesium


	// Hitung luas trapesium
	area := 0.5 * (a + b) * h

	fmt.Printf("Luas trapesium dengan alas %f, atas %f, dan tinggi %f adalah %f\n", a, b, h, area)
}