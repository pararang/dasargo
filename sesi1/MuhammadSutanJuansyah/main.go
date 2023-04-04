package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	input := map[string]int{
		"A": 59,
		"B": 60,
		"C": 61,
		"D": 0,
		"E": 86,
	}

	expectedResult := map[string][]string{
		"lulus": {
			"C", "E",
		},
		"tidak_lulus": {
			"A", "B", "D",
		},
	}

	result := GroupingStudentByPoint(input)

	sort.Strings(result["lulus"])
	sort.Strings(result["tidak_lulus"])
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("want: %v\ngot : %v\n", expectedResult, result)
	} else {
		fmt.Println("Well Done!")
	}
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
