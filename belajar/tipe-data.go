package main

import (
	"fmt"
	"reflect"
)

func main(){

// numerik non desimal
	var positifNumber uint8 = 89
	var negatifNumber = -1243423644

	fmt.Printf("bilangan positif: %d (tipe: %s) \n", positifNumber, reflect.TypeOf(positifNumber))
	fmt.Printf("bilangan negatif: %d (tipe: %s) \n", negatifNumber, reflect.TypeOf(negatifNumber))


// numerik desimal
	var decimalNumber = 2.62

	fmt.Println("bilangan desimal: ", decimalNumber)
	fmt.Println("bilangan desimal: ", decimalNumber)

// boolean
	var exist bool = true

	fmt.Printf("exist? %t \n \n", exist)

// string
	kata := "apapun"
	fmt.Printf("message: %s \n", kata)

	kalimat := `halo
	aku
	programmer "golang".`
	fmt.Printf("message: %s \n", kalimat)

// nil atau zero value
	/*
	zero value otomatis ada pada tipe data yang tidak ada isinya
	nil hanya bisa digunakan tipe data non primitive =
		pointer
		tipe data fungsi
		slice
		map
		channel
		interface kogong atau any
	*/
}