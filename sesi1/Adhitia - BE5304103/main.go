package main

import (
	"fmt"
)

func main() {
	grading := map[string]int{
		"Test":  100,
		"Test2": 90,
		"Test3": 70,
		"Test4": 50,
	}

	hasil := GroupingStudentByPoint(grading)

	fmt.Println(hasil)
}

func GroupingStundentByPoint(grading map[string]int) map[string][]string {

	var lulus, tidakLulus []string

	for name, score := range grading {
		if score >= 60 {
			lulus = append(lulus, name)
		} else {
			tidakLulus = append(tidakLulus, name)
		}
	}

	result := make(map[string][]string)
	result["lulus"] = lulus
	result["tidak_lulus"] = tidakLulus

	return result
}
