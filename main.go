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
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Сумма кредита должна быть больше нуля.")
		return
	}

	fmt.Print("Введите процентную ставку: ")
	fmt.Scanln(&interestRate)

	if interestRate < 0 {
		fmt.Println("Процентная ставка не может быть отрицательной.")
		return
	}

	fmt.Print("Введите срок выплаты (лет): ")
	fmt.Scanln(&years)

	if years <= 0 {
		fmt.Println("Срок выплаты должен быть положительным числом.")
		return
	}

	report, err := calculate.CalculateCredit(amount, interestRate, years)

	if err != nil {
		fmt.Println("Ошибка при расчете кредита:", err)
		return
	}

	fmt.Println("Результат расчета:")
	fmt.Println(report)
}
