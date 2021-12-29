package gobmi

import "fmt"

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64, err error) {

	var sexWeight int
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	if bmi == 0 {
		err = fmt.Errorf("bmi cannot be 0")
		return
	}
	if age < 0 || age > 150 {
		err = fmt.Errorf("age cannot be negative or 大于 150")
		return
	}
	if sex != "男" && sex != "女" {
		err = fmt.Errorf("性别不是男或者女")
		return
	}
	return (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100, err
}
