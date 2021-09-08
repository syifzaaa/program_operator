package main

import "fmt"

func main() {
	fmt.Print("Masukan Jumlah Rupiah \t= ")
	var rupiah int
	fmt.Scanln(&rupiah)

	dollar := float64(rupiah) * 0.00007 // rate google
	fmt.Printf("Jumlah Dollar yang Didapat \t= %6.2f\n", dollar)

	euro := float64(rupiah) * 0.000059 // rate google
	fmt.Printf("Jumlah Euro yang Didapat \t= %6.2f\n", euro)

	fmt.Print("Masukan Jumlah GB Pounds \t= ")
	var gbp int
	fmt.Scanln(&gbp)

	knut := gbp * 100
	fmt.Printf("Jumlah Knut yang Didapat \t= %d\n", knut)
	sickle := knut / 29
	knut = knut % 29

	galleon := sickle / 17
	sickle = sickle % 17
	fmt.Printf("Hasil Penukaran Mendapatkan \t= %d Galleon(s)\n", galleon)
	fmt.Printf("Sisa Ditukar Menjadi \t= %d Sickle(s)\n", sickle)
	fmt.Printf("Keping Knut yang Tersisa \t= %d Knut(s)\n", knut)

}
