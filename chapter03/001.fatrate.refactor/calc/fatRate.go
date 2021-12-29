package calc

import (
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcFatRate(bmi float64, age int, name string, sex string) (fatRate float64,err error) {
	return gobmi.CalcFatRate(bmi, age, sex)
}
