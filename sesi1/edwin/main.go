package main

import (
	"fmt"
	"reflect"
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
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("want: %v\ngot : %v\n", expectedResult, result)
	} else {
		fmt.Println("Well Done!")
	}
}

func GroupingStundentByPoint(studentGrade map[string]int) map[string][]string {
	var lulus []string
	var tidak_lulus []string

	for key, value := range studentGrade {
		if value >= 60 {
			lulus = append(lulus, key)
		} else {
			tidak_lulus = append(tidak_lulus, key)
		}
	}

	result := map[string][]string{
		"lulus":       lulus,
		"tidak_lulus": tidak_lulus,
	}

	return result
}
