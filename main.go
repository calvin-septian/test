package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello Word")

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("ganjil")
		} else {
			fmt.Println("genap")
		}
	}

	name := []string{"andi", "budi", "cacing"}
	for _, v := range name {
		fmt.Println(v)
	}

	var listBiodata = []Data{
		{Nama: "andi", Alamat: "jalan 1", Pekerjaan: "pegawai", Alasan: "belajar"},
		{Nama: "budi", Alamat: "jalan 2", Pekerjaan: "pegawai", Alasan: "belajar"},
		{Nama: "bono", Alamat: "jalan 3", Pekerjaan: "pegawai", Alasan: "belajar"},
		{Nama: "eko", Alamat: "jalan 4", Pekerjaan: "pegawai", Alasan: "belajar"},
	}

	index, _ := strconv.Atoi(os.Args[1])
	fmt.Println(listBiodata[index])

	biodata()
}
