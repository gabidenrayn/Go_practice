package main

import (
	"fmt"

	"github.com/gabidenrayn/Go_practice/pkg/powerbill"

	"github.com/fatih/color"
)

func main() {
	prev := 1200.0
	curr := 1350.0
	tariff := 0.25

	kwh, err := powerbill.Consumption(prev, curr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cost, err := powerbill.EnergyCost(kwh, tariff)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = powerbill.ApplyDiscount(&cost, 10)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	report, err := powerbill.FormatEnergyReport("Maksoni Zhumatai", kwh, cost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ВОТ ЗДЕСЬ заменили fmt.Printf на цветной вывод
	color.Green(report)
	fmt.Printf(report)
}
