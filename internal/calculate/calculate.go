package calculate

import (
	"fmt"
	"math"
)

// Общая проверка входных данных для всех функций
func validateInput(amount float64, interestRate float64, years int) (string, error) {
	if amount <= 0 {
		return "", fmt.Errorf("amount must be greater than zero")
	}

	if interestRate < 0 {
		return "", fmt.Errorf("interest rate cannot be negative")
	}

	if years <= 0 {
		return "", fmt.Errorf("years must be a positive integer")
	}

	return "", nil
}

// Функция для расчета ежемесячного платежа
func calculateMonthlyPayment(amount float64, interestRate float64, years int) (float64, error) {
	_, err := validateInput(amount, interestRate, years)

	if err != nil {
		return 0, err
	}

	monthlyRate := interestRate / 12 / 100
	totalMonths := years * 12

	if monthlyRate == 0 {
		return amount / float64(totalMonths), nil
	}

	monthlyPayment := (amount * monthlyRate) / (1 - math.Pow(1+monthlyRate, -float64(totalMonths)))

	return monthlyPayment, nil
}

// Функция для запуска симуляции кредита и формирования отчета
func runCreditSimulation(amount float64, interestRate float64, years int) (string, string, error) {
	_, err := validateInput(amount, interestRate, years)

	if err != nil {
		return "", "", err
	}

	monthlyPayment, err := calculateMonthlyPayment(amount, interestRate, years)

	if err != nil {
		return "", "", err
	}

	totalMonths := years * 12
	monthlyRate := interestRate / 12 / 100
	currentBalance := amount
	yearlyOverpayment := 0.0
	totalOverpayment := 0.0

	yearlyReport := fmt.Sprintf("изначальная сумма: %.2f процент: %.2f срок выплаты: %d лет (%d месяцев)\n", amount, interestRate, years, totalMonths)
	detailedReport := yearlyReport

	yearlyReport += "+-------+---------------+-------------------+-----------------+\n"
	yearlyReport += "|  Год  | Остаток долга | Переплата за год  | Общая переплата |\n"
	yearlyReport += "+-------+---------------+-------------------+-----------------+\n"

	detailedReport += "+--------+------------+------------+------------+---------------+\n"
	detailedReport += "| Месяц  |   Платеж   |  Проценты  | Тело долга | Остаток долга |\n"
	detailedReport += "+--------+------------+------------+------------+---------------+\n"

	for month := 1; month <= totalMonths; month++ {
		interestPayment := currentBalance * monthlyRate
		principalPayment := monthlyPayment - interestPayment
		currentBalance -= principalPayment
		yearlyOverpayment += interestPayment
		totalOverpayment += interestPayment

		if currentBalance < 0 {
			currentBalance = 0
		}

		detailedReport += fmt.Sprintf("| %-6d | %-10.2f | %-10.2f | %-10.2f | %-13.2f |\n",
			month, monthlyPayment, interestPayment, principalPayment, currentBalance)

		if month%12 == 0 {
			detailedReport += "+--------+------------+------------+------------+---------------+\n"
			yearlyReport += fmt.Sprintf("| %-5d | %-13.2f | %-17.2f | %-15.2f |\n",
				month/12, currentBalance, yearlyOverpayment, totalOverpayment)

			yearlyOverpayment = 0.0
		}
	}

	yearlyReport += "+-------+---------------+-------------------+-----------------+\n"

	totalCost := amount + totalOverpayment

	yearlyReport += fmt.Sprintf("Общая переплата: %.2f, Общая стоимость кредита: %.2f\n", totalOverpayment, totalCost)
	detailedReport += fmt.Sprintf("Общая переплата: %.2f, Общая стоимость кредита: %.2f\n", totalOverpayment, totalCost)

	return yearlyReport, detailedReport, nil
}

// Функция для расчета деталей кредита
func CalculateCreditDetails(amount float64, interestRate float64, years int) (string, error) {
	_, detailed, err := runCreditSimulation(amount, interestRate, years)

	if err != nil {
		return "", err
	}

	return detailed, nil
}

// Функция для расчета кредита и формирования отчета
func CalculateCredit(amount float64, interestRate float64, years int) (string, error) {
	yearly, _, err := runCreditSimulation(amount, interestRate, years)

	if err != nil {
		return "", err
	}

	return yearly, nil
}
