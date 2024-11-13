package main

import "fmt"

func main()  {
	fmt.Println("halo dunia")

	// tipe data (case sensitive)
// integer = 8, 16, 32, 64
	// int = -2,147,483,648 to 2,147,483,647
	// unint = 0 to 255

// floating point
	// desimal (,)
	// complex number (statistik dll)

// alias
	// byte = uint8
	// rune = int32
	// u/int = sesuai OSnya 32 / 64 bit

	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 2)
	fmt.Println("Tiga koma Lima = ", 3.5) 

// boolean 
	fmt.Println("benar = ", true)
	fmt.Println("salah = ", false)

// string
	fmt.Println("Oktavari Budi Arief")

	// function untuk string
		fmt.Println("panjang string = ", len("Oktavari Budi Arief"))
		fmt.Println("index ke 3 = ", "oktavari budi arief"[0]) // mengambalikan byte (bukan karakternya)

// variabel
	var name string // bisa untuk strict type
	// var age int8 = 21	// bisa untuk stict type
	var age = 21	// bisa langsung inisialisasi agar fleksibel
	alamat := "jakarta" // deklarasi awal agar tidak menggunakan var 


	name = "Oktavari Budi Arief"
	fmt.Println(name)

	name = "Budi Arie"
	fmt.Println("nama = ", name, "umur = ", age, "alamat =", alamat)

	// membuat variabel sekaligus
		var (
			firstName = "Oktavari"
			lastName = "Budi Arief"
		)
		fmt.Println(firstName, lastName)

	// variabel constant
		const namaAwal string = "budi gemink"
		const namaAkhir = "arief"

		// error 
		// namaAwal = "pipin"

		// bisa multiple seperti variabel

// konversi tipe data ( tipeData() ) = lihat lagi dokumentasi untuk besaran yang diampu oleh tipe data
		var nilai int32 = 2000000
		var nilai64 int64 = int64(nilai)

		// melebihi batas
		var nilai8 int8 = int8(nilai64)

		fmt.Println(nilai)
		fmt.Println(nilai64)
		fmt.Println(nilai8)

		// string (e = uint8)
		var nama string = "Oktavari"
		var e byte = nama[0]
		var eString string = string(e)

		fmt.Println(nama)
		fmt.Println(eString)

// type declaration membuat alias dari tipe data
		
	type Nama string
	type Married bool

	var namaUser1 Nama = "Oktavari"
	var namaUser2 Nama = "Budi Arief"
	var married1 Married = true

	fmt.Println(namaUser1)
	fmt.Println(namaUser2)
	fmt.Println(married1)


// operasi matematika
		// + - * / %
			var result = 1 + 1
			fmt.Println("1 + 1 = ", result)

			var a = 2
			var b = 3
			var c = a * b
			fmt.Println(a, " * ", b, " = ", c)
		
		// a += b
			var i = 10
			i += 10
			fmt.Println(i)

		// unary operator
			// ++ = +1
			// -- = -1
			// ! = not
			i++
			fmt.Println(i)

// operasi perbandingan
		// <
		// >
		// <=
		// >=
		// ==
		// !=

			// coba
				var name1 string = "Oktavari"
				var name2 = "Budi Arief"
				name3 := "jujung"

				var result1 = name1 == name2 || name1 != name3

				fmt.Println(result1)

//  operasi boolean
		// && = and
		// || = or
		// ! = not

			// coba
				var nilaiAkhir = 90
				var nilaiAbsen = 80

				var lulusAkhir bool = nilaiAkhir > 80
				var lulusAbsen bool = nilaiAbsen > 80

				var lulus = lulusAkhir && lulusAbsen

				fmt.Println(lulus)

// array
	// coba
		var hewan [3]string
		
		hewan[0] = "Kucing"
		hewan[1] = "Anjing"
		hewan[2] = "Kelinci"

		fmt.Println(hewan[0])
		fmt.Println(hewan[1])
		fmt.Println(hewan[2])

			// array langsung
				var nomor = [3]int{
					1,
					2,
					3,
				}

				fmt.Println(nomor)

					// function array
						// len(array) = panjang array
						// array[index] = ambil data array
						// array[index] = value = ubah data array

							// coba
							var value = [...]int{ // [...] =  artinya tidak menentukan secara eksplisit panjang data arraynya
								1,
								2,
								3,
							}
			
							fmt.Println(len(value))
							fmt.Println(value[0])
							value[0] = 50
							fmt.Println(value[0])

// tipe data slice (sering dipakai)
	// potongan dari array 
			// pointer (penunjuk data pertama)
			// length (panjang data)
			// capasity (dari pointer sampai akhir data)
							// contoh 
								// array[low:high] index awal sampai index akhir

									// coba
										mhs := [...]string{"okta", "budi", "arief", "jujung"}
										slice := mhs[1:4]

										fmt.Println(slice[0]) // ketika di print slicenya maka akan menampilkan data awal dari si slice bukan dari array keseluruhannya 

											// penjelasan
												// var array [3]string = ditentukan panjang datanya
												// var slice []string = tidak ditentukan panjang datanya

													// function dalam slice
														// len(slice) = panjang slice
															fmt.Println(len(slice))

														// cap(slice) = kapasitas slice
															fmt.Println(cap(slice))

														// append(slice, data) = menambah data ke slice di akhir
															slice = append(slice, "pipin")
															fmt.Println(slice)
															fmt.Println(mhs)
												// catatan: jika menggunakan append dan bila kapasitas array awal sudah habis maka akan dibuat array baru, namun bisa juga inisialisasi nama slice baru jika tidak ingin slice lama berubah

														// make([]string, 2, 5)
															// penjelasan
																// membuat slice namun arraynya otomatis dibuatkan
																	// panjang 2 = datanya ada 2
																	// kapasitasnya 5 = maka slicenya jika di append tidak membuat array baru

																		// coba
																			var newSlice = make([]string, 2, 5)
																			newSlice[0] = "icip"
																			newSlice[1] = "ucup"

																			fmt.Println(newSlice)
																			fmt.Println(len(newSlice))
																			fmt.Println(cap(newSlice))

																			newSlice = append(newSlice, "acap", "ocop", "ecep", "scsp", "lclp")
																			fmt.Println(newSlice)
}															

