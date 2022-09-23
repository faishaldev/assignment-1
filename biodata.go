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

		helpers.Print(biodata, absen)

		fmt.Println()
	}
}
