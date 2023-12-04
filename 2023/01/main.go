package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isNum(in byte) (int, bool) {
	numMap := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	if num, ok := numMap[in]; ok {
		return num, true
	}

	return 0, false
}

func lastTextDigit(in string) (int, bool) {
	if len(in) < 3 {
		return 0, false
	}
	stringDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	strIndex := -1
	strDigit := 0
	for i, v := range stringDigits {
		index := strings.LastIndex(in, v)
		if index != -1 {
			if strIndex == -1 || index > strIndex {
				strIndex = index
				strDigit = i + 1
			}
		}
	}
	if strIndex != -1 {
		return strDigit, true
	}

	return 0, false
}

func firstTextDigit(in string) (int, bool) {
	if len(in) < 3 {
		return 0, false
	}
	stringDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	strIndex := -1
	strDigit := 0
	for i, v := range stringDigits {
		index := strings.Index(in, v)
		if index != -1 {
			if strIndex == -1 || index < strIndex {
				strIndex = index
				strDigit = i + 1
			}
		}
	}
	if strIndex != -1 {
		return strDigit, true
	}

	return 0, false
}

func firstDigit(in string) int {
	digitIndex, digitNum := 0, 0
	for i := 0; i < len(in); i++ {
		if num, ok := isNum(in[i]); ok {
			digitNum = num
			digitIndex = i
			break
		}
	}
	if textNum, ok := firstTextDigit(in[:digitIndex]); ok {
		return textNum
	}

	return digitNum
}

func lastDigit(in string) int {
	digitIndex, digitNum := 0, 0
	for i := len(in) - 1; i >= 0; i-- {
		if num, ok := isNum(in[i]); ok {
			digitNum = num
			digitIndex = i
			break
		}
	}

	if textNum, ok := lastTextDigit(in[digitIndex+1:]); ok {
		return textNum
	}
	return digitNum
}

func processTextToNum(in string) int {
	f := firstDigit(in)
	l := lastDigit(in)

	return f*10 + l
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to open input.txt")
		os.Exit(1)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += processTextToNum(scanner.Text())
	}

	fmt.Println("result is ", sum)
}
