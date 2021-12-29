package calc

import (
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcFatRate(bmi float64, age int, name string, sex string) float64 {
	return gobmi.CalcFatRate(bmi, age, name, sex)
}
