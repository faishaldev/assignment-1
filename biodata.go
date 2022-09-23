package main

import (
	"assignment-1/helpers"
	"os"
	"strconv"
)

func main() {
	biodata := helpers.Init()

	absen, _ := strconv.Atoi(os.Args[1])

	helpers.Print(biodata, absen)
}
