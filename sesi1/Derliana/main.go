package main

import "fmt"

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
	// input value kedalam slice
	input := map[string]int{
		"ikhsan": 30,
		"sans":	  70,
		"ikh":	  80,
		"albi":   50,
		"dela":   70,
		"ryzal":  90,
		"tuni":   100,
		"yola":	  80,
	}

	// variabel student untuk menampung function GroupingstudentByPoint dengan parameter input
	student := GroupingStudentByPoint(input)

	fmt.Println(student)
}

