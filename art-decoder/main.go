package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Define flags for multi-line and encode modes
	multiLine := flag.Bool("m", false, "Enable multi-line mode")
	encodeMode := flag.Bool("e", false, "Enable encode mode")
	flag.Parse()

	// when the user selects multiline
	if *multiLine {
		fmt.Println("Enter multi-line input (Press enter on an empty line to finish):")
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			lines = append(lines, line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error")
			return
		}
		for _, line := range lines {
			ProcessLine(line, encodeMode)
		}

		// when the user does not select multiline
	} else if len(flag.Args()) > 0 {
		input := flag.Args()[0]
		ProcessLine(input, encodeMode)
	} else {
		fmt.Println("Error")
	}
}

func ProcessLine(line string, encodeMode *bool) {
	if *encodeMode {
		fmt.Println(encodeInput(line))
	} else {
		if !isBalancedBrackets(line) {
			fmt.Println("Error")
			return
		}
		decoded, err := decodeInput(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(strings.Join(decoded, ""))
	}
}

func decodeInput(input string) ([]string, error) {
	re := regexp.MustCompile(`\[([^\]]*)\]`)
	matches := re.FindAllStringSubmatchIndex(input, -1)

	// If no matches are found, return the original input
	if len(matches) == 0 {
		return []string{input}, nil
	}

	var results []string
	lastIndex := 0

	for _, match := range matches {
		// Append the part of the string before the pattern
		results = append(results, input[lastIndex:match[0]])

		// Extract number of repetitions and text to repeat
		part := input[match[2]:match[3]]
		parts := strings.SplitN(part, " ", 2)

		if len(parts) != 2 || parts[1] == "" {
			return nil, fmt.Errorf("Error\n")
		}

		n, err := strconv.Atoi(parts[0])
		x := parts[1]

		if err != nil {
			return nil, fmt.Errorf("Error\n")
		}

		// Check for incorrect input inside brackets
		if strings.Contains(x, "[") || strings.Contains(x, "]") {
			return nil, fmt.Errorf("Error\n")
		}

		// Repeat the string x, n times and add to results
		result := strings.Repeat(x, n)
		results = append(results, result)

		// Update the last index
		lastIndex = match[1]
	}

	// Append the remaining part of the string after the last pattern
	results = append(results, input[lastIndex:])

	return results, nil
}

func encodeInput(input string) string {
	var encodedBuilder strings.Builder
	i := 0

	for i < len(input) {
		for unitSize := 1; unitSize <= min(len(input)/2, 3); unitSize++ {
			unit := input[i:min(i+unitSize, len(input))]
			count := countRepeatingUnits(input, i, unit)

			if count > 1 && len(unit) <= 3 {
				// Encode the repeating unit
				encodedBuilder.WriteString(fmt.Sprintf("[%d %s]", count, unit))
				i += unitSize * count
				break
			} else if unitSize == min(len(input)/2, 3) {
				// If no repeating unit is found, or it's too long, append the current character
				encodedBuilder.WriteString(string(input[i]))
				i++
			}
		}
	}

	return encodedBuilder.String()
}

func isBalancedBrackets(s string) bool {

	balance := 0
	for _, char := range s {
		if char == '[' {
			balance++
		} else if char == ']' {
			balance--
			if balance < 0 {
				return false
			}
		}
	}
	return balance == 0
}

func countRepeatingUnits(str string, start int, unit string) int {
	count := 0
	for i := start; i+len(unit) <= len(str); i += len(unit) {
		if str[i:i+len(unit)] == unit {
			count++
		} else {
			break
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
