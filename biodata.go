package main

import (
	"fmt"
)

type data struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func biodata() {

	// listData := []data{
	// 	{Nama: "andi", Alamat: "jalan 1", Pekerjaan: "pegawai", Alasan: "belajar"},
	// 	{Nama: "budi", Alamat: "jalan 2", Pekerjaan: "pegawai", Alasan: "belajar"},
	// 	{Nama: "bono", Alamat: "jalan 3", Pekerjaan: "pegawai", Alasan: "belajar"},
	// 	{Nama: "eko", Alamat: "jalan 4", Pekerjaan: "pegawai", Alasan: "belajar"},
	// }

	listData := []data{}

	Data1 := data{}
	Data1.Nama = "andi"
	Data1.Alamat = "jalan 1"
	Data1.Pekerjaan = "pegawai"
	Data1.Alasan = "belajar lebih dalam"
	listData = append(listData, Data1)

	Data2 := data{}
	Data2.Nama = "budi"
	Data2.Alamat = "jalan 2"
	Data2.Pekerjaan = "pegawai"
	Data2.Alasan = "belajar lebih dalam"
	listData = append(listData, Data2)

	Data3 := data{}
	Data3.Nama = "bono"
	Data3.Alamat = "jalan 3"
	Data3.Pekerjaan = "pegawai"
	Data3.Alasan = "belajar lebih dalam"
	listData = append(listData, Data3)

	Data4 := data{}
	Data4.Nama = "eko"
	Data4.Alamat = "jalan 4"
	Data4.Pekerjaan = "pegawai"
	Data4.Alasan = "belajar lebih dalam"
	listData = append(listData, Data4)

	for _, v := range listData {
		fmt.Printf("%s\n%s\n%s\n%s\n", v.Nama, v.Alamat, v.Pekerjaan, v.Alasan)
	}
}
