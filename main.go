package main

import (
	"creditCalculate/internal/calculate"
	"fmt"
)

func main() {
	var amount float64
	var interestRate float64
	var years int

	fmt.Print("Введите сумму кредита: ")
	fmt.Scanf("%f", &amount)

	if amount <= 0 {
		fmt.Println("Сумма кредита должна быть больше нуля.")
		return
	}

	fmt.Print("Введите процентную ставку: ")
	fmt.Scanf("%f", &interestRate)

	if interestRate < 0 {
		fmt.Println("Процентная ставка не может быть отрицательной.")
		return
	}

	fmt.Print("Введите срок выплаты (лет): ")
	fmt.Scanf("%d", &years)

	if years <= 0 {
		fmt.Println("Срок выплаты должен быть положительным числом.")
		return
	}

	fmt.Println(calculate.CalculateCredit(amount, interestRate, years)) // Example usage
}
