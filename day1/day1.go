package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(expenseFile string) (expenses []int) {
	expenseData, fileError := os.Open(expenseFile)

	if fileError != nil {
		fmt.Println("Error encountered when opening file:", fileError)
		return
	}

	scanFile := bufio.NewScanner(expenseData)
	scanFile.Split(bufio.ScanLines)

	for scanFile.Scan() {
		convertInt, _ := strconv.Atoi(scanFile.Text())
		expenses = append(expenses, convertInt)
	}

	expenseData.Close()

	return
}

func twoExpenses(expenses []int) (expense1 int, expense2 int) {
	var expenseTotal int

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

func threeExpenses(expenses []int) (expense1 int, expense2 int, expense3 int) {

	var expenseTotal int

Calculator:
	for i := 0; i < len(expenses); i++ {
		for x := i + 1; x < len(expenses); x++ {
			for y := i + 2; y < len(expenses); y++ {
				expenseTotal = expenses[i] + expenses[x] + expenses[y]
				if expenseTotal == 2020 {
					expense1, expense2, expense3 = expenses[i], expenses[x], expenses[y]
					break Calculator
				}
			}
		}
	}

	return
}

func main() {

	expenses := readFile("expenses.txt")

	expense1, expense2 := twoExpenses(expenses)

	fmt.Printf("Expense #1: %d \nExpense #2: %d\n", expense1, expense2)
	twoProduct := expense1 * expense2
	fmt.Printf("Expense #1 multiplied by Expense #2 = %d\n", twoProduct)

	expense1, expense2, expense3 := threeExpenses(expenses)

	fmt.Printf("Expense #1: %d \nExpense #2: %d\nExpense #3: %d\n", expense1, expense2, expense3)
	threeProduct := expense1 * expense2 * expense3
	fmt.Printf("Expense #1, #2 and #3 mulitplied together = %d\n", threeProduct)

}
