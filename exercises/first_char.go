package main

import "fmt"

// example :supercalifragilisticexpialidocious -> f

func findFirstCharWithCount(word string, pos int, count int) rune {
	chars := make(map[rune]int)
	posCount := 0

	for _, char := range word {
		chars[char] = chars[char] + 1
	}
	for _, char := range word {
		if chars[char] == count {
			if posCount == pos {
				return char
			}
			posCount++
		}
	}
	return ' '
}

func main() {
	word := "supercalifragilisticexpialidocious"
	result := findFirstCharWithCount(word, 2, 3)
	fmt.Println(string(result))
}

// nth char that occurs x amount of times
