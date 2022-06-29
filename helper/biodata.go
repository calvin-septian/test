package helper

import (
	"fmt"
)

type Data struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func GetBiodata() {
	var listData = []Data{
		{Nama: "andi", Alamat: "jalan 1", Pekerjaan: "pegawai", Alasan: "belajar"},
		{Nama: "budi", Alamat: "jalan 2", Pekerjaan: "pegawai", Alasan: "belajar"},
		{Nama: "bono", Alamat: "jalan 3", Pekerjaan: "pegawai", Alasan: "belajar"},
		{Nama: "eko", Alamat: "jalan 4", Pekerjaan: "pegawai", Alasan: "belajar"},
	}

	for _, v := range listData {
		fmt.Printf("%s %s %s %s\n", v.Nama, v.Alamat, v.Pekerjaan, v.Alasan)
	}

	// index, _ := strconv.Atoi(os.Args[1])
	// fmt.Println(listData[index])
}
