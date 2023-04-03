package main

import "fmt"

func main() {
	fmt.Println(GroupingStundentByPoint(map[string]int{
		"ikhsan":    30,
		"sans":      70,
		"ikh":       80,
		"lintang":   100,
		"indriyani": 80,
		"rara":      10,
		"hilmi":     30,
	}))
}

func GroupingStundentByPoint(data map[string]int) map[string][]string {
	hasil := map[string][]string{}
	for key, val := range data {
		if val > 60 {
			hasil["lulus"] = append(hasil["lulus"], key)
		} else {
			hasil["tidak lulus"] = append(hasil["tidak lulus"], key)
		}
	}
	return hasil
}
