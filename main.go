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
	arabic = iota
	roman
	wrong
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
		if numType == arabic {
			a, errA := parseNumber(fields[0])
			b, errB := parseNumber(fields[2])
			if errA != nil || errB != nil {
				log.Fatal(err)
			}

			result, err := calculate(operator, a, b)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
		} else if numType == roman {
			a, _ := parseRomanNumber(fields[0])
			b, _ := parseRomanNumber(fields[2])
			result, err := calculate(operator, a, b)
			if err != nil {
				log.Fatal(err)
			}
			if result <= 0 {
				log.Fatal("roman number can't be zero or negative")
			}

			fmt.Println(integerToRoman(result))
		}
	}
}

func getNumbersType(a string, b string) (NumberType, error) {
	first, errFirst := parseRomanNumber(a)
	second, errSecond := parseRomanNumber(b)
	if errFirst != nil || errSecond != nil {
		return wrong, errors.New("illegal roman number is supplied")
	}

	if first == 0 && second == 0 {
		return arabic, nil
	} else if first != 0 && second != 0 {
		return roman, nil
	} else {
		return wrong, errors.New("mismatched types")
	}
}

func getInputAsTokens(reader *bufio.Reader) ([]string, error) {
	expressionError := `Error! Please enter valid expression.
	Patterns:
		a + b
		a * b
		a - b
		a / b`

	fmt.Println("Please enter an expression to calculate.")
	line, err := reader.ReadString('\n') // Read string until enter is pressed.
	if err != nil {
		return []string{""}, errors.New("can't read line from input")
	}

	fields := strings.Fields(line) // Get items separated by spaces.
	if len(fields) != 3 {
		return nil, errors.New(expressionError)
	}

	return fields, nil
}

func parseNumber(field string) (int, error) {
	num, err := strconv.Atoi(field)
	if err != nil {
		return num, err
	}

	return num, nil
}

func parseRomanNumber(num string) (int, error) {
	switch num {
	case "I":
		return 1, nil
	case "II":
		return 2, nil
	case "III":
		return 3, nil
	case "IV":
		return 4, nil
	case "V":
		return 5, nil
	case "VI":
		return 6, nil
	case "VII":
		return 7, nil
	case "VIII":
		return 8, nil
	case "IX":
		return 9, nil
	case "X":
		return 10, nil
	default:
		return 0, errors.New("no allowed roman number is found")
	}
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
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return result, errors.New("bad operator")
	}
	return result, nil
}
