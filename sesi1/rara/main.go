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

func GroupingStundentByPoint(data map[string]int) map[string][]string {
	result := map[string][]string{}

	for key, val := range data {
		if val > 60 {
			result["lulus"] = append(result["lulus"], key)
		} else {
			result["tidak lulus"] = append(result["tidak lulus"], key)
		}
	}
	return result
}
