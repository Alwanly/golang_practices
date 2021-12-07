package main

import "fmt"

func main() {
	funArray()
	nextArray()
	funSlice()
}

func funArray() {
	var names [3]string

	names[0] = "Alwan"
	names[1] = "Idel"
	names[2] = "Ley"
	fmt.Println(names)
}

func nextArray() {
	var nilais = [3]int{
		90,
		80,
		30,
	}
	fmt.Println(nilais, len(nilais))
}

func funSlice() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var triwulanPertama = months[0:3]
	var triwulanKedua = months[3:6]
	var triwulanTiga = months[6:9]

	fmt.Println(triwulanPertama)
	fmt.Println(triwulanTiga)
	fmt.Println(triwulanKedua)
}
