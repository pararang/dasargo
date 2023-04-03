package main

import "fmt"

func main() {
	a := map[string]int{
		"Alfa":  90,
		"Daffa": 20,
		"Aceng": 30,
	}

	fmt.Println(GroupingStundentByPoint(a))
}

func GroupingStundentByPoint(siswa map[string]int) map[string][]string {

	var tidak_lulus = []string{}
	var lulus = []string{}

	for key, s := range siswa {
		if s > 60 {
			lulus = append(lulus, key)
		} else if s <= 60 {
			tidak_lulus = append(tidak_lulus, key)
		}

	}

	result := map[string][]string{
		"lulus":       lulus,
		"tidak_lulus": tidak_lulus,
	}

	return result // lengkapi logicnya
}
