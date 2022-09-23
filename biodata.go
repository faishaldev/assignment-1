package main

import (
	"assignment-1/helpers"
	"fmt"
	"os"
	"strconv"
)

func main() {
	biodata := helpers.Init()

	for i := 1; i < len(os.Args); i++ {
		absen, _ := strconv.Atoi(os.Args[i])

		fmt.Println()

		if absen == 0 || absen >= len(biodata) {
			fmt.Printf("Biodata nomor absen %v tidak ditemukan!\n", absen)
			continue
		}

		helpers.Print(biodata, absen)
	}
}
