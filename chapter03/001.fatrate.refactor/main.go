package main

import (
	"fmt"
	"runtime/debug"
	calc2 "wangsw.go/chapter03/001.fatrate.refactor/calc"
)

func main() {
	for {

		mainFatRateBody()

		if cont := whetherContinue(); !cont {
			break
		}
	}
	//平均bmi在哪儿调用
	//calculateAvgOnSlice()
}
func recoverMainBody() {
	if re := recover(); re != nil {
		fmt.Printf("warning: catch critical error: %v\n", re)
		debug.PrintStack()
	}
}
func mainFatRateBody() {
	defer recoverMainBody()
	name, sex, weight, tall, age := getMaterialsFromInput()

	fatRate, err := calcFatRate(weight, tall, age, name, sex)
	if err != nil {
		fmt.Println("计算fatRate错误", err)
		return
	}
	if fatRate <= 0 {
		panic("fat rate is not allowed to be negative")
	}

	//定义函数类型的变量，函数的参数必须一致
	var checkHealthinessFunc func(age int, fatRate float64)
	if sex == "男" {
		checkHealthinessFunc = getHealthinessSuggestionsForMale
	} else {
		checkHealthinessFunc = getHealthinessSuggestionsForFeMale
	}

	//getHealthinessSuggestions(sex, age, fatRate)
	//getHealthinessSuggestions(age,fatRate,getHealthinessSuggestionsForMale)
	getHealthinessSuggestions(age, fatRate, checkHealthinessFunc)
}

func getHealthinessSuggestions(age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	getSuggestion(age, fatRate)
}
func generateCheckHealthinessForMale() func(age int, fatRate float64) {
	//定制方法
	return func(age int, fatRate float64) {
	}
}

func getHealthinessSuggestionsForFeMale(age int, fatRate float64) {
	if 18 <= age && age <= 39 {
		if fatRate <= 0.2 {
			fmt.Println("偏瘦")
		} else if fatRate > 0.2 && fatRate <= 0.27 {
			fmt.Println("标准")
		} else if fatRate > 0.27 && fatRate <= 0.34 {
			fmt.Println("偏胖")
		} else if fatRate > 0.34 && fatRate <= 0.39 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("非常肥胖")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.21 {
			fmt.Println("偏瘦")
		} else if fatRate > 0.21 && fatRate <= 0.28 {
			fmt.Println("标准")
		} else if fatRate > 0.28 && fatRate <= 0.35 {
			fmt.Println("偏胖")
		} else if fatRate > 0.35 && fatRate <= 0.40 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("非常肥胖")
		}
	} else if age >= 60 {
		if fatRate <= 0.22 {
			fmt.Println("偏瘦")
		} else if fatRate > 0.22 && fatRate <= 0.29 {
			fmt.Println("标准")
		} else if fatRate > 0.29 && fatRate <= 0.36 {
			fmt.Println("偏胖")
		} else if fatRate > 0.36 && fatRate <= 0.41 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("非常肥胖")
		}
	} else {
		fmt.Println("小孩子调皮捣蛋!")
	}
}

func getHealthinessSuggestionsForMale(age int, fatRate float64) {
	if 18 <= age && age <= 39 {
		if fatRate <= 0.1 {
			fmt.Println("偏瘦")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("标准")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("偏胖")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("非常肥胖")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.11 {
			fmt.Println("偏瘦")
		} else if fatRate > 0.11 && fatRate <= 0.17 {
			fmt.Println("标准")
		} else if fatRate > 0.17 && fatRate <= 0.22 {
			fmt.Println("偏胖")
		} else if fatRate > 0.22 && fatRate <= 0.27 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("非常肥胖")
		}
	} else if age >= 60 {
		if fatRate <= 0.13 {
			fmt.Println("偏瘦")
		} else if fatRate > 0.13 && fatRate <= 0.19 {
			fmt.Println("标准")
		} else if fatRate > 0.19 && fatRate <= 0.24 {
			fmt.Println("偏胖")
		} else if fatRate > 0.24 && fatRate <= 0.29 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("非常肥胖")
		}
	} else {
		fmt.Println("未成年")
	}
}

func calcFatRate(weight float64, tall float64, age int, name string, sex string) (fatRate float64, err error) {
	bmi, err := calc2.CalcBmi(weight, tall)
	if err != nil {
		return 0, err
	}
	fatRate = calc2.CalcFatRate(bmi, age, name, sex)
	fmt.Println(name, "的体脂率为", fatRate)
	return
}

func getMaterialsFromInput() (string, string, float64, float64, int) {
	var name string
	fmt.Println("请输入姓名:")
	fmt.Scanf("%s", &name)

	var sex string
	fmt.Println("请输入性别:")
	fmt.Scanf("%s", &sex)

	var weight float64
	fmt.Println("请输入体重(kg):")
	fmt.Scanf("%f", &weight)

	var tall float64
	fmt.Println("请输入身高(m):")
	fmt.Scanf("%f", &tall)

	var age int
	fmt.Println("请输入年龄:")
	fmt.Scanf("%d", &age)

	return name, sex, weight, tall, age
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Println("请输入y/n")
	fmt.Scanf("%s", &whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}

func calculateAvgOnSlice(bmis ...float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	avgBmi = total / float64(len(bmis))
	return avgBmi
}
