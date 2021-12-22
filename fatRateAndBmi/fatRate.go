package fatRateAndBmi


import "fmt"

func CalcFatRate(bmi float64,age int, name string,sex string) float64 {

	var sexWeight int
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	var fatRate = 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)/100
	fmt.Print(name, "的体脂率为", fatRate, "，结果是：")
	return fatRate
}