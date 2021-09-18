package main

import (
	"fmt"
	"math"
)

type Hitung interface {
	keliling() float64
	luas() float64
}

type Segitiga struct {
	sisi   float64
	alas   float64
	tinggi float64
}

func (segitiga Segitiga) keliling() float64 {
	return segitiga.sisi * 3
}

func (segitiga Segitiga) luas() float64 {
	return (segitiga.alas * segitiga.tinggi) / 2
}

type Persegi struct {
	sisi float64
}

func (persegi Persegi) keliling() float64 {
	return persegi.sisi * 4
}

func (persegi Persegi) luas() float64 {
	return math.Pow(persegi.sisi, 2)
}

func menu() {
	var pilihan int
	var next string
	var bangunDatar Hitung
	var sisiSegitiga, alas, tinggi, sisiPersegi float64

	fmt.Println("===========================================")
	fmt.Println("==========  DAFTAR BANGUN DATAR  ==========")
	fmt.Println("===========================================")
	fmt.Println("1. Segitiga")
	fmt.Println("2. Persegi")
	fmt.Println("===========================================")
pilihan:
	fmt.Print("Masukan pilihan menu [1/2] : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukan sisi segitiga : ")
		fmt.Scan(&sisiSegitiga)
		fmt.Print("Masukan alas segitiga : ")
		fmt.Scan(&alas)
		fmt.Print("Masukan tinggi segitiga : ")
		fmt.Scan(&tinggi)

		bangunDatar = Segitiga{alas: alas, tinggi: tinggi, sisi: sisiSegitiga}
		fmt.Println("====segitiga======")
		fmt.Println("keliling : ", bangunDatar.keliling())
		fmt.Println("luas : ", bangunDatar.luas())
	} else if pilihan == 2 {
		fmt.Print("Masukan sisi persegi : ")
		fmt.Scan(&sisiPersegi)
		bangunDatar = Persegi{sisiPersegi}
		fmt.Println("====== persegi =======")
		fmt.Println("keliling : ", bangunDatar.keliling())
		fmt.Println("luas: ", bangunDatar.luas())
	} else {
		fmt.Print("Silahkan masukan pilihan kembali!\n")
		goto pilihan
	}

	fmt.Print("Hitung kembali? [y/t] : ")

	if pilihan != 0 {
		fmt.Scan(&next)
		if next == "y" || next == "Y" {
			main()
		} else {
			fmt.Println("================= END =====================")
		}
	}
}

func main() {
	menu()

}
