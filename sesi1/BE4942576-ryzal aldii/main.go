package main

import "fmt"

func GroupingStundentByPoint(studentGrade map[string]int) map[string][]string {

	var lulus []string
	var tidaklulus []string

	for key, value := range studentGrade {
		if value >= 60 {
			lulus = append(lulus, key)
		} else {
			tidaklulus = append(tidaklulus, key)
		}
	}

	result := map[string][]string{
		"lulus":     lulus,
		"tidakulus": tidaklulus,
	}

	return result
}

func main() {
	input := map[string]int{
		"ikhsan": 30,
		"sans":   70,
		"ikh":    80,
		"albi":   50,
		"della":  70,
		"ryzal":  90,
		"tuni":   100,
		"yola":   80,
	}

	student := GroupingStundentByPoint(input)

	fmt.Println(student)
}
