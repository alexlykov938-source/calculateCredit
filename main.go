package main

import (
	"creditCalculate/internal/calculate"
	"fmt"
	"strings"
)

func main() {
	var (
		amount       float64
		interestRate float64
		years        int
		reportMode   int
	)

	for {
		fmt.Print("Введите сумму кредита: ")
		fmt.Scanln(&amount)

		if amount <= 0 {
			fmt.Println("Сумма кредита должна быть больше нуля.")
			continue
		}

		fmt.Print("Введите процентную ставку: ")
		fmt.Scanln(&interestRate)

		if interestRate < 0 {
			fmt.Println("Процентная ставка не может быть отрицательной.")
			continue
		}

		fmt.Print("Введите срок выплаты (лет): ")
		fmt.Scanln(&years)

		if years <= 0 {
			fmt.Println("Срок выплаты должен быть положительным числом.")
			continue
		}

		fmt.Println("\nВыберите тип отчета:")
		fmt.Println("1 — Краткий отчет (переплата по годам)")
		fmt.Println("2 — Детальный отчет (подробно по месяцам)")
		fmt.Print("Ваш выбор: ")
		fmt.Scanln(&reportMode)

		// Если пользователь ввел не 1 и не 2, пускаем цикл заново
		if reportMode != 1 && reportMode != 2 {
			fmt.Println("Неверный выбор режима отчета. Начните сначала.")
			continue
		}

		var report string
		var err error

		if reportMode == 1 {
			// Вызываем краткий отчет по годам
			report, err = calculate.CalculateCredit(amount, interestRate, years)
		} else {
			// Вызываем детальный отчет по месяцам
			report, err = calculate.CalculateCreditDetails(amount, interestRate, years)
		}

		// Твоя стандартная проверка ошибки от калькулятора (тоже меняем return на continue!)
		if err != nil {
			fmt.Println("Ошибка при расчете кредита:", err)
			continue
		}

		fmt.Println("Результат расчета:")
		fmt.Println(report)

		fmt.Print("\nХотите сделать еще один расчет? Нажмите любую клавишу для нового расчёта или (N) для выхода: ")

		var input string

		fmt.Scanln(&input)
		if strings.ToUpper(input) == "N" {
			break
		}
	}
}
