package main

import (
	"fmt"
)

func main() {
	//const
	const fullName string = "alwan"

	//varibale
	var (
		fname string
		lname string
	)
	fname = "Idelia"
	lname = "Muthia"

	//conversion
	var nilaiKecil int32 = 1000000001
	var nilaiBesar int64 = int64(nilaiKecil)

	var name = "Ley"
	var e byte = name[0]
	var testAll = string(e)

	//Type Declaratons (alias tipe data)
	type NoKTP string
	type Married bool

	var ktp NoKTP
	var status Married
	ktp = "3170207040003"
	status = true

	fmt.Println(fname, lname, fullName, nilaiBesar, testAll, ktp, status)

}
