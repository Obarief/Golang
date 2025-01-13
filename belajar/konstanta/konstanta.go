package main

import (
	"fmt"
	"reflect"
)

func main(){
	const nama = "okta"

	fmt.Print("halo", " ", nama, "!\n")

	const (
		hari = "senin"
		sekarang // nilainya sama yaitu senin
		angka = 21
		desimal = 2.21
	)

	fmt.Printf("hari ini: %s (tipe: %s) \n", hari, reflect.TypeOf(hari))
	fmt.Printf("sekarang: %s (tipe: %s) \n", sekarang, reflect.TypeOf(sekarang))
	fmt.Printf("sekarang: %.3f (tipe: %s) \n", desimal, reflect.TypeOf(desimal))
	
}