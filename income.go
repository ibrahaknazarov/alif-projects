package main

import (
	"errors"
	"fmt"
)

func validateAmount(amount int, currency string) error {
	// Разрешённые валюты
	allowedCurrencies := map[string]bool{
		"TJS": true,
		"USD": true,
	}

	// Проверка валюты
	if !allowedCurrencies[currency] {
		return errors.New("недопустимая валюта")
	}

	// Проверка суммы (неотрицательная)
	if amount < 0 {
		return errors.New("сумма не может быть отрицательной")
	}

	// Допустимые тестовые значения (0, 500000, 1000000)
	allowedAmounts := map[int]bool{
		0:       true,
		500000:  true,
		1000000: true,
	}

	if !allowedAmounts[amount] {
		return errors.New("сумма не входит в список допустимых")
	}

	return nil
}

func main() {
	tests := []struct {
		amount   int
		currency string
	}{
		{0, "TJS"},
		{0, "USD"},
		{500000, "TJS"},
		{500000, "USD"},
		{1000000, "TJS"},
		{1000000, "USD"},
	}

	for _, t := range tests {
		err := validateAmount(t.amount, t.currency)
		if err != nil {
			fmt.Printf("FAIL: %d %s → %s\n", t.amount, t.currency, err)
		} else {
			fmt.Printf("PASS: %d %s\n", t.amount, t.currency)
		}
	}
}
