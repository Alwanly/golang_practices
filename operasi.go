package main

import "fmt"

func main() {
	matematika()
	perbandingan()
	logika()
}

func matematika() {
	var result = 10 + 10
	fmt.Println(-result)
}

func perbandingan() {
	var (
		a = 100
		b = 200
	)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a != b)
}

func logika() {
	var (
		nilai = 70
		kkm   = 75
	)
	var lulus = nilai >= kkm

	fmt.Print("lulus = ", lulus)
}
