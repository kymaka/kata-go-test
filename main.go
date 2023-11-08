package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type NumberType int

const (
	Arabic NumberType = iota
	Roman
	Wrong
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fields, err := getInputAsTokens(reader)
		if err != nil {
			log.Fatal(err)
		}

		numType, err := getNumbersType(fields[0], fields[2])
		if err != nil {
			log.Fatal(err)
		}

		operator := fields[1]
		switch numType {
		case Arabic:
			a, b, err := parseNumbers(fields[0], fields[2])
			if err != nil {
				log.Fatal(err)
			}

			result, err := calculate(operator, a, b)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
		case Roman:
			a, err := parseRomanNumber(fields[0])
			if err != nil {
				log.Fatal(err)
			}

			b, err := parseRomanNumber(fields[2])
			if err != nil {
				log.Fatal(err)
			}

			result, err := calculate(operator, a, b)
			if err != nil {
				log.Fatal(err)
			}

			if result <= 0 {
				fmt.Println("roman number can't be zero or negative")
			} else {
				fmt.Println(integerToRoman(result))
			}
		case Wrong:
			log.Fatal("unexpected error")
		}
	}
}

func getNumbersType(a string, b string) (NumberType, error) {
	_, errRFirst := parseRomanNumber(a)
	_, errRSecond := parseRomanNumber(b)

	_, _, err := parseNumbers(a, b)
	if err == nil {
		return Arabic, nil
	}

	if errRFirst == nil && errRSecond == nil {
		return Roman, nil
	}

	return Wrong, errors.New("bad input number")
}

func getInputAsTokens(reader *bufio.Reader) ([]string, error) {
	expressionError := `Error! Please enter valid expression.
	Patterns:
		a + b
		a * b
		a - b
		a / b`

	fmt.Println("Please enter an expression to calculate.")
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("can't read line from input")
	}

	fields := strings.Fields(line)
	if len(fields) != 3 {
		return nil, errors.New(expressionError)
	}

	return fields, nil
}

func parseNumbers(field1 string, field2 string) (int, int, error) {
	a, errA := strconv.Atoi(field1)
	b, errB := strconv.Atoi(field2)

	if errA != nil || errB != nil {
		return 0, 0, errors.New("invalid number format")
	}
	if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
		return a, b, nil
	} else {
		return a, b, errors.New("number is not correct")
	}

}

func parseRomanNumber(num string) (int, error) {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	if val, ok := romanNumerals[num]; ok {
		return val, nil
	}

	return 0, errors.New("no allowed roman number is found")
}

func integerToRoman(number int) string {
	maxRomanNumber := 100
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func calculate(operator string, a int, b int) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("bad operator")
	}
}
