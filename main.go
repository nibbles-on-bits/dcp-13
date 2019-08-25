package main

import "fmt"

func main() {

	s := "abcba"
	//k := 2

	innerLoopReps := 0
	for x := len(s); x > 1; x-- {
		innerLoopReps++

		fmt.Printf("x=%d\n", x)
		for y := 0; y < innerLoopReps; y++ {
			fmt.Printf("y=%d\n", y)
			fmt.Printf("s[y:y+x-1] = %v\n", s[y:y+x])
			// to count unique characters, build a map?
			// count unique chars

		}

		/*
			5 - 1
			4 - 2
			3 - 3
			2 - 4
			1 - 5
		*/

		fmt.Println(countUniqueChars("abcab"))

	}

}

func countUniqueChars(s string) int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		..fmt.Println(string(s[i]))

		//fmt.Println(m[string(s[i])])
		m[string(s[i])]++
	}
	//fmt.Println(m)
	//fmt.Println(len(m))
	return len(m)
}
