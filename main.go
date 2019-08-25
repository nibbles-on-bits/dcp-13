package main

import "fmt"

func main() {

	s := "abcbcbajkjk"
	k := 2
	fmt.Println("SOLUTION: " + doit(s, k))

}

func doit(s string, k int) string {

	innerLoopReps := 0
	for x := len(s); x > 1; x-- {
		innerLoopReps++

		for y := 0; y < innerLoopReps; y++ {
			if countUniqueChars(s[y:y+x]) == k {
				return s[y : y+x]
			}
		}
	}
	return ""
}

func countUniqueChars(s string) int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}
	return len(m)
}
