package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func calculate(eq string) int {
	eq = strings.Trim(eq, " ")
	fmt.Println(eq)
	var stack []int
	var nums []rune
	operator := 1
	var result int
	i := 0
	for i < len(eq) {
		if eq[i] == '-' {
			fmt.Println("operator -1")
			operator = -1
		} else if eq[i] == '+' {
			operator = 1
		} else if eq[i] == '(' {
			fmt.Println("pushing: ", result)
			stack = append(stack, result)
			fmt.Println("pushing: ", operator)
			stack = append(stack, operator)
			result = 0
			operator = 1
		} else if eq[i] == ')' {
			result *= pop(&stack)
			result += pop(&stack)
		} else if unicode.IsDigit(rune(eq[i])) {
			for i < len(eq) && unicode.IsDigit(rune(eq[i])) {
				nums = append(nums, rune(eq[i]))
				i++
			}
			num, _ := strconv.Atoi(string(nums))
			fmt.Printf("%d + (%d)%d\n", result, operator, num)
			result += operator * num
			// reset the nums for the next number
			nums = []rune{}
			continue
		}
		i++
	}
	return result
}

func pop(stack *[]int) int {
	l := len(*stack)
	if l == 0 {
		fmt.Errorf("popping from empty arr")
		return -1
	}
	res := (*stack)[l-1]
	*stack = (*stack)[:l-1]
	return res
}

//func main() {
//	b := "(1+(4+5+2)-3)+(6+8)"
//	result := calculate(b)
//	fmt.Println(result)
//}
