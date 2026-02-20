// Package powerbill предоставляет функции для расчета
// потребления и стоимости электроэнергии.
package powerbill

import (
	"errors"
	"fmt"
)

// Consumption вычисляет потребление по показаниям счетчика.
func Consumption(prev, curr float64) (float64, error) {
	if prev < 0 || curr < 0 {
		return 0, fmt.Errorf("meter readings cannot be negative")
	}
	if curr < prev {
		return 0, fmt.Errorf("current reading cannot be less than previous")
	}
	return curr - prev, nil
}

// EnergyCost рассчитывает стоимость электроэнергии.
func EnergyCost(kwh, tariff float64) (float64, error) {
	if kwh < 0 || tariff < 0 {
		return 0, fmt.Errorf("kwh and tariff must be non-negative")
	}
	return kwh * tariff, nil
}

// ApplyDiscount применяет скидку к стоимости (работа с указателем).
func ApplyDiscount(cost *float64, percent float64) error {
	if cost == nil {
		return errors.New("cost pointer is nil")
	}
	if percent < 0 || percent > 100 {
		return errors.New("discount must be between 0 and 100")
	}
	*cost -= *cost * percent / 100
	return nil
}

// FormatEnergyReport формирует форматированный отчет.
func FormatEnergyReport(owner string, kwh, cost float64) (string, error) {
	if owner == "" {
		return "", errors.New("owner cannot be empty")
	}

	report := fmt.Sprintf(
		"Owner: %s\nConsumption: %.2f kWh\nTotal cost: %.2f\n",
		owner, kwh, cost,
	)

	return report, nil
}