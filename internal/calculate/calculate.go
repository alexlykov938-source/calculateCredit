package calculate

import (
	"fmt"
	"math"
)

func calculateMonthlyPayment(amount float64, interestRate float64, years int) (float64, error) {
	if amount <= 0 {
		return 0, fmt.Errorf("amount must be greater than zero")
	}

	if interestRate < 0 {
		return 0, fmt.Errorf("interest rate cannot be negative")
	}

	if years <= 0 {
		return 0, fmt.Errorf("years must be a positive integer")
	}

	monthlyInterestRate := interestRate / 100 / 12
	numberOfPayments := float64(years * 12)

	// Calculate the monthly payment
	monthlyPayment := (amount * monthlyInterestRate) / (1 - math.Pow(1+monthlyInterestRate, -numberOfPayments))

	return monthlyPayment, nil
}

func CalculateCredit(amount float64, interestRate float64, years int) (string, error) {
	monthlyPayment, err := calculateMonthlyPayment(amount, interestRate, years)
	if err != nil {
		return "", err
	}

	currentBalance := amount
	yearlyOverpayment := 0.0
	totalOverpayment := 0.0
	totalMonths := years * 12
	reportResult := fmt.Sprintf("изначальная сумма: %.2f процент: %.2f срок выплаты: %d лет (%d месяцев)\n", amount, interestRate, years, totalMonths)

	for months := 1; months <= years*12; months++ {
		interestPayment := currentBalance * (interestRate / 100 / 12)
		principalPayment := monthlyPayment - interestPayment
		currentBalance -= principalPayment
		totalOverpayment += interestPayment
		yearlyOverpayment += interestPayment

		if currentBalance < 0 {
			currentBalance = 0
		}

		if months%12 == 0 {
			reportResult += fmt.Sprintf("Год %d: Остаток долга: %.2f, Переплата за год: %.2f, Общая переплата: %.2f\n", months/12, currentBalance, yearlyOverpayment, totalOverpayment)
			yearlyOverpayment = 0
		}
	}

	totalCost := amount + totalOverpayment
	reportResult += fmt.Sprintf("Общая стоимость кредита: %.2f\n", totalCost)

	return reportResult, nil
}
