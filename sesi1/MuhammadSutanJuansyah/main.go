package main

import (
	"fmt"
	"strings"
)

func main() {
	nilaiSiswa := map[string]int{
		"Lorem":   90,
		"Ipsum":   80,
		"Dolor":   60,
		"Sitamet": 50,
	}

	hasil := GroupingStudentByPoint(nilaiSiswa)
	fmt.Println("Siswa yang lulus:", strings.Join(hasil["lulus"], ", "))
	fmt.Println("Siswa yang tidak lulus:", strings.Join(hasil["tidak_lulus"], ", "))
}

func GroupingStudentByPoint(nilaiSiswa map[string]int) map[string][]string {
	var lulus, tidakLulus []string

	for namaSiswa, nilai := range nilaiSiswa {
		if nilai >= 70 {
			lulus = append(lulus, namaSiswa)
		} else {
			tidakLulus = append(tidakLulus, namaSiswa)
		}
	}

	hasil := make(map[string][]string)
	hasil["lulus"] = lulus
	hasil["tidak_lulus"] = tidakLulus

	return hasil
}
