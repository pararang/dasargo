package main

import "fmt"

func main() {
	m := map[string]int{
		"ikhsan": 30,
		"sans":   70,
		"ikh":    80,
		"john":   50,
	}

	result := GroupingStudentByPoint(m)
	fmt.Println(result)
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
