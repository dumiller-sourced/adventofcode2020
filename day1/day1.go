package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	expenseData, fileError := os.Open("expenses.txt")

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

	expenseData.Close()

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

	fmt.Printf("Expense #1: %d \nExpense #2: %d\n", expense1, expense2)
	totalProduct := expense1 * expense2
	fmt.Printf("Expense #1 multiplied by Expense #2 = %d\n", totalProduct)

}
