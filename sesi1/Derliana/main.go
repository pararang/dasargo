package main

import (
	"fmt"
	"reflect"
	"sort"
)

func GroupingStudentByPoint(studentGrade map[string]int) map[string][]string {
	// membuat slice kosong 
	var lulus []string
	var tidakLulus[]string

	// pengecekan nilai studentGrade
	for key, value := range studentGrade {
		if value >= 60 {
			lulus = append(lulus, key)
		} else {
			tidakLulus =append(tidakLulus, key)
		}
	}

	// map untuk return value
	result := map[string][]string{
		"lulus":	lulus,
		"tidakLulus": tidakLulus,
	}
	// return value
	return result
}

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

