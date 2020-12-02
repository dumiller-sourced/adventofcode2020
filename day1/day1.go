package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(expenseFile string) []int {
	expenseData, fileError := os.Open(expenseFile)

	if fileError != nil {
		fmt.Println("Error encountered when opening file:", fileError)
		return
	}

	scanFile := bufio.NewScanner(expenseData)
	scanFile.Split(bufio.ScanLines)

	var expenses []int

	for scanFile.Scan() {
		convertInt, _ := strconv.Atoi(scanFile.Text())
		expenses = append(expenses, convertInt)
	}

	expenseData.Close()

	return expenses
}

func twoExpenses(expenses []int) (int, int) {
	var expenseTotal int
	var expense1, expense2 int

Calculator:
	for i := 0; i < len(expenses); i++ {
		for x := i + 1; x < len(expenses); x++ {
			expenseTotal = expenses[i] + expenses[x]
			if expenseTotal == 2020 {
				expense1, expense2 = expenses[i], expenses[x]
				break Calculator
			}
		}
	}

	return expense1, expense2
}

func main() {

	expenses := readFile("expenses.txt")

	expense1, expense2 := twoExpenses(expenses)

	fmt.Printf("Expense #1: %d \nExpense #2: %d\n", expense1, expense2)
	twoProduct := expense1 * expense2
	fmt.Printf("Expense #1 multiplied by Expense #2 = %d\n", twoProduct)

	/* expenseData, fileError := os.Open("expenses.txt")

	if fileError != nil {
		fmt.Println("Error encountered when opening file:", fileError)
		return
	}

	readFile := bufio.NewScanner(expenseData)
	readFile.Split(bufio.ScanLines)

	var expenses []int

	for readFile.Scan() {
		convertInt, _ := strconv.Atoi(readFile.Text())
		expenses = append(expenses, convertInt)
	}

	expenseData.Close() */

	/* var expenseTotal int
		var expense1, expense2 int

	Calculator:
		for i := 0; i < len(expenses); i++ {
			for x := i + 1; x < len(expenses); x++ {
				expenseTotal = expenses[i] + expenses[x]
				if expenseTotal == 2020 {
					expense1, expense2 = expenses[i], expenses[x]
					break Calculator
				}
			}
		} */
}
