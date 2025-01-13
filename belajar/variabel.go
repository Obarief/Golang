package main

import "fmt"

func main(){

	// inisialisasi variabel
	var firstName string = "Okta"

	var lastName string = "Arief"

	fmt.Printf("halo %s %s!\n", firstName, lastName)


	// tanpa inisialisasi variabel
	nama := "Oktavari budi Arief"
	
	fmt.Printf("halo %s\n", nama)

	// multi variabel
	var melon, semangka, jeruk string
	melon, semangka, jeruk = "melon", "watermelon", "orange"
	fmt.Println(melon, semangka, jeruk)

	satu, dua, tiga := "first", "second", "third"
	fmt.Println(satu,dua,tiga)

	one, isMarried, twoPointOne, say := 1, false, 2.1, "hello"
	fmt.Println(one,isMarried, twoPointOne, say)

	// balckhole dengan _ (digunakan sebagai variabel penampung) tidak dapat ditampilkan
	name, _ := "okta", "arief"
	fmt.Println(name)

	// pointer dengan new
	// untuk melihat isinya maka dengan * di awal nama variabel
	barang := new(string)
	*barang = "okta"

	fmt.Println(barang) // alamat memory dalam bentuk hexadesimal
	fmt.Println(*barang) // value variabel

}