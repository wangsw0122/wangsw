package calc

import (
	"github.com/armstrongli/go-bmi"
)

func CalcBmi(weight, tall float64) (bmi float64,err error) {

	bmi, err = gobmi.BMI(weight, tall)

	return
}
