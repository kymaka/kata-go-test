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
			a := parseNumber(fields[0])
			b := parseNumber(fields[2])
			result, err := calculate(operator, a, b)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
		} else if numType == roman {
			a := parseRomanNumber(fields[0])
			b := parseRomanNumber(fields[2])
			result, err := calculate(operator, a, b)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
		}
	}
}

func getNumbersType(a string, b string) (NumberType, error) {
	first := parseRomanNumber(a)
	second := parseRomanNumber(b)

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
		log.Fatal(err)
	}

	fields := strings.Fields(line) // Get items separated by spaces.
	if len(fields) != 3 {
		return nil, errors.New(expressionError)
	}
	return fields, nil

}

func parseNumber(field string) int {

	num, err := strconv.Atoi(field)
	if err != nil {
		log.Fatal("Please enter valid number.", err)
	}
	return num
}

func parseRomanNumber(num string) int {
	switch num {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	}
	return 0
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
		return result, errors.New("can't calculate")
	}
	return result, nil
}
