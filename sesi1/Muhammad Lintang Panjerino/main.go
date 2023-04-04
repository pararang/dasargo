package main

import "fmt"

func main() {
	fmt.Println(GroupingStundentByPoint(map[string]int{
		"ikhsan": 30,
		"sans": 70,
		"ikh": 80,
		"lintang": 100,
		"indriyani": 80,
		"rara": 10,
		"hilmi": 30,
	}))
}

func GroupingStundentByPoint( data map[string]int ) map[string][]string {
	// lengkapi logicnya

	//define map untuk result
	result := map[string][]string{}

	for key, nilai := range data {
		if nilai > 60 {
			result["lulus"] = append(result["lulus"], key)
		} else {
			result["tidak lulus"] = append(result["tidak lulus"], key)
		}
	}

	return result
}
