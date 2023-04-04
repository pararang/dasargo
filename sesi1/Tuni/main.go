package main

import "fmt"

func GroupingStundentByPoint(studentGrade map[string]int) map[string][]string {
	// lengkapi logicnya
	var lulus []string
	var tidaklulus []string

	//pengecekan nilai studentGrade
	for key, value := range studentGrade {
		if value >= 60 {
			lulus = append(lulus, key)
		} else {
			tidaklulus = append(tidaklulus, key)
		}
	}

	//map untuk terurn value
	result := map[string][]string{
		"lulus":      lulus,
		"tidaklulus": tidaklulus,
	}
	//return value
	return result
}

func main() {
	//input value kedalam slice
	input := map[string]int{
		"ikhsan": 30,
		"sans":   70,
		"ikh":    80,
		"albi":   50,
		"dela":   70,
		"ryzal":  90,
		"tuni":   100,
		"yola":   80,
	}
	student := GroupingStundentByPoint(input)

	fmt.Println(student)
}
