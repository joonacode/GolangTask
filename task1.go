package main

// Cetak gambar

import (
	"fmt"
)

func cetakGambar(param int) {
	if param % 2 == 0 {
		fmt.Println("Parameter harus bilangan ganjil")
	} else{
		penengah := (param / 2) + 1
		fmt.Println("--- panjang ---")
		for i := 1; i <= param; i++ {
			tampung := ""
			for x := 1; x <= param; x++{
				if x == 1 || x == param ||  i == penengah {
					tampung += "* "
				}else{
					tampung += "= "
				}
			}
			fmt.Println(tampung)
		}	
	}
	fmt.Println("")
}

func main() {
	cetakGambar(3)
	cetakGambar(5)
	cetakGambar(9)
	cetakGambar(15)
	cetakGambar(10)
}