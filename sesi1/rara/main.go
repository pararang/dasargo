package main

import "fmt"

func main() {
	fmt.Println(GroupingStundentByPoint(map[string]int{
		"ikhsan":  30,
		"sans":    70,
		"ikh":     80,
		"lintang": 55,
		"indri":   79,
		"rara":    80,
		"hilmi":   60,
	}))
}

func GroupingStundentByPoint(data map[string]int) map[string][]string {
	result := map[string][]string{}

	for key, val := range data {
		if val > 60 {
			result["lulus"] = append(result["lulus"], key)
		} else {
			result["tidak lulus"] = append(result["tidak lulus"], key)
		}
	}
	return result
}
