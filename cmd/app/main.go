package main

import (
	"strconv"
	"strings"
)

func calc(input string) int {
	if input == "" {
		return 0
	}
	var delimeter string
	if strings.HasPrefix(input, "//") {
		input, _ = strings.CutPrefix(input, "//")
		if strings.HasPrefix(input, "[") {
			input, _ := strings.CutPrefix(input, "[")
			delimeter += string(input[0])
			for string(input[0]) != "]" {
				delimeter += string(input[0])
				input = input[1:]
			}
		}

	}
	number, err := strconv.ParseInt(input, 10, 32)
	if err == nil {
		return int(number)
	}

	var numbers []string
	if delimeter == "" {
		if strings.Contains(input, ",") {
			nums := strings.Split(input, ",")
			for _, num := range nums {
				numbers = append(numbers, strings.Split(num, "\n")...)
			}
		} else if strings.Contains(input, "\n") {
			nums := strings.Split(input, "\n")
			for _, num := range nums {
				numbers = append(numbers, strings.Split(num, ",")...)
			}
		}
	} else {
		numbers = strings.Split(input, delimeter)
	}
	var sum int = 0
	for _, num := range numbers {
		n, err := strconv.ParseInt(num, 10, 32)
		if err != nil {
			panic("invalid number")
		}
		if n < 0 {
			panic("negative number")
		}
		if n > 1000 {
			continue
		}
		sum += int(n)
	}
	return sum
}

func main() {
	calc("")
}
