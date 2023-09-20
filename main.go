package main

import (
	"fmt"
	"os"
)

type Teman struct {
	id         uint
	nama       string
	alamat    string
	pekerjaan string
	alasan     string
}

var teman []Teman = []Teman{
		{
            id:        0,
            nama:      "Januwar",
            alamat:    "Purwokerto",
            pekerjaan: "Mahasiswa",
            alasan:    "Karena Golang Populer",
        },
        {
            id:        1,
            nama:      "Yayan",
            alamat:    "Karanglewas",
            pekerjaan: "Software Engineer",
            alasan:    "Karena golang bahasa pemrograman yang keren",
        },
        {
            id:        2,
            nama:      "Aldi",
            alamat:    "Banyumas",
            pekerjaan: "Programmer",
            alasan:    "Karena saya sedang belajar Golang",
        },
        {
            id:        3,
            nama:      "Budi",
            alamat:    "Jakarta",
            pekerjaan: "Data Analyst",
            alasan:    "Golang cocok untuk analisis data",
        },
        {
            id:        4,
            nama:      "Siti",
            alamat:    "Surabaya",
            pekerjaan: "Web Developer",
            alasan:    "Golang memungkinkan pengembangan web yang cepat",
        },
        {
            id:        5,
            nama:      "Rina",
            alamat:    "Semarang",
            pekerjaan: "Sistem Administrator",
            alasan:    "Golang mudah untuk dikelola dalam sistem",
        },
        {
            id:        6,
            nama:      "Diana",
            alamat:    "Bandung",
            pekerjaan: "Front-end Developer",
            alasan:    "Golang memiliki dukungan komunitas yang kuat",
        },
        {
            id:        7,
            nama:      "Ivan",
            alamat:    "Makassar",
            pekerjaan: "UI/UX Designer",
            alasan:    "Golang membantu dalam desain antarmuka yang efisien",
        },
        {
            id:        8,
            nama:      "Eva",
            alamat:    "Medan",
            pekerjaan: "Mobile App Developer",
            alasan:    "Golang cocok untuk pengembangan aplikasi mobile",
        },
        {
            id:        9,
            nama:      "Faisal",
            alamat:    "Aceh",
            pekerjaan: "QA Engineer",
            alasan:    "Golang mempermudah dalam pengujian perangkat lunak",
        },
}
func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Gunakan: go run biodata.go [nama Teman]")
		return
	}

	nama := args[1]
	CariTeman(nama)
}

func CariTeman(nama string) {
	found := false
	for _, t := range teman {
		if nama == t.nama {
			fmt.Println("Informasi Teman:")
			fmt.Println("ID:", t.id)
			fmt.Println("Nama:", t.nama)
			fmt.Println("Alamat:", t.alamat)
			fmt.Println("Pekerjaan:", t.pekerjaan)
			fmt.Println("Alasan:", t.alasan)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Teman dengan nama", nama, "tidak ditemukan")
	}
}
