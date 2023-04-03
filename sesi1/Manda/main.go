package main

import (
	"fmt"
)

func main() {

	students := make(map[string]int)
	var lulus []string
	var tidak_lulus []string
	var name string
	var score int

	for {
		fmt.Printf("Masukkan nama siswa : ")
		fmt.Scanln(&name)

		fmt.Printf("Masukkan nilai siswa : ")
		fmt.Scanln(&score)

		students[name] = score

		fmt.Println(GroupingStundentByPoint(students, lulus, tidak_lulus))

		var pilihan string
		fmt.Print("Apakah Anda ingin mengkonversi kembali? (y/n): ")
		fmt.Scanln(&pilihan)

		if pilihan == "n" {
			break
		}
	}

}

func GroupingStundentByPoint(students map[string]int, lulus []string, tidak_lulus []string) map[string][]string {
	// lengkapi logicnya

	for name, score := range students {
		if score > 60 {
			lulus = (append(lulus, name))
		} else {
			tidak_lulus = (append(tidak_lulus, name))
		}
	}
	hasil := map[string][]string{
		"Lulus ":      lulus,
		"tidak_lulus": tidak_lulus,
	}

	return hasil

}
