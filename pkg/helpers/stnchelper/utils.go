package stnchelper

import (
	"math"
)

// PercentageCalculation yuzde hesaplama
func PercentageCalculation(targetMoney, totalPrice float64) (float64, float64, int) {
	var temp1, percentage float64
	var percentageRemaining int
	remainingMoney := targetMoney - totalPrice
	temp1 = remainingMoney / targetMoney
	percentage = math.Round(temp1 * 100)
	percentageRemaining = 100 - int(percentage)
	return remainingMoney, percentage, percentageRemaining
}

type (
	Utils struct{ Debug bool }
)

func (UtilModul Utils) PercentageCalculationForPercentageRemaining(targetMoney, totalPrice float64) int {
	return PercentageCalculationForPercentageRemaining(targetMoney, totalPrice)
}

// PercentageCalculationForPercentageRemaining yuzde hesaplama == geriye sadece kalan yuzde verir
func PercentageCalculationForPercentageRemaining(targetMoney, totalPrice float64) int {
	var temp1, percentage float64
	var percentageRemaining int
	remainingMoney := targetMoney - totalPrice
	temp1 = remainingMoney / targetMoney
	percentage = math.Round(temp1 * 100)
	percentageRemaining = 100 - int(percentage)
	return percentageRemaining
}
