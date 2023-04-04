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

	result := GroupingStundentByPoint(input)

	sort.Strings(result["lulus"])
	sort.Strings(result["tidak_lulus"])
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("want: %v\ngot : %v\n", expectedResult, result)
	} else {
		fmt.Println("Well Done!")
	}
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
