package main

import "fmt"

func main() {
	input := make(map[string]int)
	input["ikhsan"] = 30
	input["sans"] = 70
	input["ikh"] = 80
	fmt.Println(GroupingStundentByPoint(input))
}

func GroupingStundentByPoint(values map[string]int) map[string][]string {
	result := make(map[string][]string)
	result["lulus"] = []string{}
	result["tidak_lulus"] = []string{}

	for student, score := range values {
		if score > 60 {
			result["lulus"] = append(result["lulus"], student)
		} else {
			result["tidak_lulus"] = append(result["tidak_lulus"], student)
		}
	}

	return result
}
