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

		helpers.Print(biodata, absen)
	}
}
