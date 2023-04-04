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

func GroupingStudentByPoint(m map[string]int) map[string][]string {
	lulus := []string{}
	tidakLulus := []string{}
	for k, v := range m {
		if v >= 60 {
			lulus = append(lulus, k)
		} else {
			tidakLulus = append(tidakLulus, k)
		}
	}
	result := map[string][]string{
		"lulus":       lulus,
		"tidak_lulus": tidakLulus,
	}
	return result
}
