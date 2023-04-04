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

	result := GroupingStundentByPoint(input, []string{}, []string{})

	sort.Strings(result["lulus"])
	sort.Strings(result["tidak_lulus"])
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("want: %v\ngot : %v\n", expectedResult, result)
	} else {
		fmt.Println("Well Done!")
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
