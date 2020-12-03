package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

func parseString(stringList []string) (num1 int, num2 int, target string, pw string) {

	intList := strings.Split(stringList[0], "-")
	num1, _ = strconv.Atoi(intList[0])
	num2, _ = strconv.Atoi(intList[1])
	target = string(stringList[1][0])
	pw = stringList[2]

	return
}

func validateOldPassword(min int, max int, target string, pw string) (validatedPW bool) {

	//fmt.Printf("Validating password %s: The character %s must be present between %d and %d times.\n", pw, target, min, max)

	var countChar int = 0

	for x := 0; x < len(pw); x++ {
		if string(pw[x]) == target {
			countChar++
		}
	}

	//fmt.Printf("Character %s found %d times. ", target, countChar)

	if countChar >= min {
		if countChar <= max {
			validatedPW = true
			return
		}
	}

	validatedPW = false
	return
}

func validateNewPassword(pos1 int, pos2 int, target string, pw string) (validatedPW bool) {

	//fmt.Printf("Validating password %s: The character %s must be present at positions %d and %d\n", pw, target, pos1, pos2)

	var checkPos1 bool = false
	var checkPos2 bool = false

	for x := 0; x < len(pw); x++ {
		if string(pw[x]) == target {
			if pos1 == x+1 {
				checkPos1 = true
			} else if pos2 == x+1 {
				checkPos2 = true
			}
		}
	}

	//fmt.Printf("Position 1 match: %t. Position 2 match: %t. ", checkPos1, checkPos2)

	if checkPos1 != checkPos2 {
		validatedPW = true
		return
	}

	return
}

func passwordPolicy(passwords []string, policyVersion string) (numValid int) {

	var validatedPW bool = false

	for i := 0; i < len(passwords); i++ {
		stringList := strings.Fields(passwords[i])
		num1, num2, target, pw := parseString(stringList)

		if policyVersion == "old" {
			validatedPW = validateOldPassword(num1, num2, target, pw)
		} else if policyVersion == "new" {
			validatedPW = validateNewPassword(num1, num2, target, pw)
		}

		//fmt.Printf("Password Valid?: %t\n", validatedPW)

		if validatedPW == true {
			numValid++
		}
	}

	return
}

func main() {

	start := time.Now()

	passwords := readFile("password.txt")

	oldPolicyValid := passwordPolicy(passwords, "old")
	newPolicyValid := passwordPolicy(passwords, "new")

	fmt.Printf("Number of Valid passwords using OLD policy: %d\n", oldPolicyValid)
	fmt.Printf("Number of Valid passwords using NEW policy: %d\n", newPolicyValid)

	elapsed := time.Since(start)
	log.Printf("Program execution took %s", elapsed)
}
