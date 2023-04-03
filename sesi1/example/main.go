package main

import "fmt"

func main() {
	pariable := map[string]int{
		"ikhsan": 30, "sans": 70, "ikh": 80,
	}

	yey := GroupingStundentByPoint(pariable)

	fmt.Println(yey)

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
