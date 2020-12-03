package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(passwordFile string) (passwords []string) {
	passwordData, fileError := os.Open(passwordFile)

	if fileError != nil {
		fmt.Println("Error encountered when opening file:", fileError)
		return
	}

	scanFile := bufio.NewScanner(passwordData)
	scanFile.Split(bufio.ScanLines)

	for scanFile.Scan() {
		passwords = append(passwords, scanFile.Text())
	}

	passwordData.Close()

	return
}

func parseString(stringList []string) (min int, max int, target string, pw string) {
	min, _ = strconv.Atoi(string(stringList[0][0]))
	max, _ = strconv.Atoi(string(stringList[0][2]))
	target = string(stringList[1][0])
	pw = stringList[2]

	return
}

func validatePassword(min int, max int, target string, pw string) (validatedPW bool, countChar int) {
	countChar = 0
	for x := 0; x < len(pw); x++ {
		if string(pw[x]) == target {
			countChar++
		}
	}
	if countChar >= min {
		if countChar <= max {
			validatedPW = true
			return
		}
	}
	validatedPW = false
	return
}

func main() {
	passwords := readFile("password.txt")

	var numValid int

	for i := 0; i < len(passwords); i++ {
		stringList := strings.Fields(passwords[i])
		min, max, target, pw := parseString(stringList)
		fmt.Printf("Validating password %s: The character %s must be present between %d and %d times.\n", pw, target, min, max)

		validatedPW, countChar := validatePassword(min, max, target, pw)

		fmt.Printf("Character %s found %d times. Password Valid?: %t\n", target, countChar, validatedPW)

		if validatedPW == true {
			numValid++
		}
	}

	fmt.Printf("Number of Valid passwords: %d\n", numValid)
}
