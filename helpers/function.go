package helpers

import (
	"fmt"
)

func Print(biodata []Biodata, absen int) {
	fmt.Println("Nama\t\t\t\t: ", biodata[absen].Nama)
	fmt.Println("Alamat\t\t\t\t: ", biodata[absen].Alamat)
	fmt.Println("Pekerjaan\t\t\t: ", biodata[absen].Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang\t: ", biodata[absen].Alasan)
}
