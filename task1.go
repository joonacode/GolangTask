package main

// Cetak gambar

import (
	"fmt"
)

func cetakGambar(param int) {
	if param % 2 == 0 {
		fmt.Println("Parameter harus bilangan ganjil")
	} else{
		penengah := param - ((param - 1) / 2)
		for i := 1; i <= param; i++ {
			tampung := ""
			for x := 1; x <= param; x++{
				if i == penengah {
					tampung += "* "
				}else{
					if x == 1 || x == param{
						tampung += "* "
					}else{
						tampung += "= "
					}
				}
			}
			fmt.Println(tampung)
		}	
	}
	fmt.Println("")
}

func main() {
	cetakGambar(5)
	cetakGambar(10)
	cetakGambar(3)
	cetakGambar(9)
}