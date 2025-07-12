package main

import (
	"fmt"
)

func main() {
	var annualSpending float64 // в сомони

	fmt.Print("Введите вашу сумму годовых трат (в сомони): ")
	fmt.Scan(&annualSpending)

	// перевод в дирамы (1 сомони = 100 дирамов)
	spendingDirams := int(annualSpending * 100)

	// процент кэшбэка в минимуме и максимуме
	minPercent := 0.005  // 0.5%
	maxPercent := 0.015  // 1.5%

	// расчет дохода в дирамах
	minIncomeDirams := int(float64(spendingDirams) * minPercent)
	maxIncomeDirams := int(float64(spendingDirams) * maxPercent)

	// перевод обратно в сомони (делим на 100), округляем вниз
	minIncomeSomoni := minIncomeDirams / 100
	maxIncomeSomoni := maxIncomeDirams / 100

	fmt.Println("Ожидаемый доход от Корти Милли в год:")
	fmt.Println("Минимальный:", minIncomeSomoni, "TJS")
	fmt.Println("Максимальный:", maxIncomeSomoni, "TJS")
}
