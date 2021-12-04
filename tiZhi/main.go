package main

import (
	"fmt"
)

//体脂计算器
//单个人
func bodyFatCalculator() {
	for {
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

		var bmi float64 = weight / (tall * tall)

		var sexWeight int

		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}

		var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100

		if sex == "男" {
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
		} else {
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

		var weightControl string
		fmt.Println("请输入y/n")
		fmt.Scanf("%s", &weightControl)
		if weightControl != "y" {
			break
		}
	}
}

//一次输入三个人
func bodyFatCalculators() {
	for {
		var name [3]string
		fmt.Println("请输入姓名:")
		fmt.Scanf("%s %s %s", &name[0], &name[1], &name[2])
		var sex [3]string
		fmt.Println("请输入性别:")
		fmt.Scanf("%s %s %s", &sex[0], &sex[1], &sex[2])

		var weight [3]float64
		fmt.Println("请输入体重(kg):")
		fmt.Scanf("%f,%f,%f", &weight[0], &weight[1], &weight[2])

		var tall [3]float64
		fmt.Println("请输入身高(m):")
		fmt.Scanf("%f,%f,%f", &tall[0], &tall[1], &tall[2])

		var age [3]int
		fmt.Println("请输入年龄:")
		fmt.Scanf("%d,%d,%d", &age[0], &age[1], &age[2])

		//var bmi = [3]float64{weight[0] / (tall[0] * tall[0]), weight[1] / (tall[1] * tall[1]), weight[2] / (tall[2] * tall[2])}

		//var sexWeight [3]int
		//
		//for n, v := range sex {
		//	if v == "男" {
		//		sexWeight[n] = 1
		//	} else {
		//		sexWeight[n] = 0
		//	}
		//}

		//var fatRate = [...]float64{
		//	(1.2*bmi[0] + 0.23*float64(age[0]) - 5.4 - 10.8*float64(sexWeight)) / 100,
		//	(1.2*bmi[1] + 0.23*float64(age[1]) - 5.4 - 10.8*float64(sexWeight)) / 100,
		//	(1.2*bmi[2] + 0.23*float64(age[2]) - 5.4 - 10.8*float64(sexWeight)) / 100,
		//}
		var h float64
		for n, _ := range name {
			x := sex[n]
			y := age[n]
			var sexWeight int
			bmi := weight[n] / (tall[n] * tall[n])
			if x == "男" {
				sexWeight = 1
			} else {
				sexWeight = 0
			}

			fatRate := 1.2*bmi + 0.23*float64(y) - 5.4 - 10.8*float64(sexWeight)/100

			z := fatRate
			fmt.Print(name[n], "的体脂率为", z, "，结果是：")
			h = h + z
			if x == "男" {
				if 18 <= y && y <= 39 {
					if z <= 0.1 {
						fmt.Println("偏瘦")
					} else if z > 0.1 && z <= 0.16 {
						fmt.Println("标准")
					} else if z > 0.16 && z <= 0.21 {
						fmt.Println("偏胖")
					} else if z > 0.21 && z <= 0.26 {
						fmt.Println("肥胖")
					} else {
						fmt.Println("非常肥胖")
					}
				} else if y >= 40 && y <= 59 {
					if z <= 0.11 {
						fmt.Println("偏瘦")
					} else if z > 0.11 && z <= 0.17 {
						fmt.Println("标准")
					} else if z > 0.17 && z <= 0.22 {
						fmt.Println("偏胖")
					} else if z > 0.22 && z <= 0.27 {
						fmt.Println("肥胖")
					} else {
						fmt.Println("非常肥胖")
					}
				} else if y >= 60 {
					if z <= 0.13 {
						fmt.Println("偏瘦")
					} else if z > 0.13 && z <= 0.19 {
						fmt.Println("标准")
					} else if z > 0.19 && z <= 0.24 {
						fmt.Println("偏胖")
					} else if z > 0.24 && z <= 0.29 {
						fmt.Println("肥胖")
					} else {
						fmt.Println("非常肥胖")
					}
				} else {
					fmt.Println("未成年")
				}
			} else {
				if 18 <= y && y <= 39 {
					if z <= 0.2 {
						fmt.Println("偏瘦")
					} else if z > 0.2 && z <= 0.27 {
						fmt.Println("标准")
					} else if z > 0.27 && z <= 0.34 {
						fmt.Println("偏胖")
					} else if z > 0.34 && z <= 0.39 {
						fmt.Println("肥胖")
					} else {
						fmt.Println("非常肥胖")
					}
				} else if y >= 40 && y <= 59 {
					if z <= 0.21 {
						fmt.Println("偏瘦")
					} else if z > 0.21 && z <= 0.28 {
						fmt.Println("标准")
					} else if z > 0.28 && z <= 0.35 {
						fmt.Println("偏胖")
					} else if z > 0.35 && z <= 0.40 {
						fmt.Println("肥胖")
					} else {
						fmt.Println("非常肥胖")
					}
				} else if y >= 60 {
					if z <= 0.22 {
						fmt.Println("偏瘦")
					} else if z > 0.22 && z <= 0.29 {
						fmt.Println("标准")
					} else if z > 0.29 && z <= 0.36 {
						fmt.Println("偏胖")
					} else if z > 0.36 && z <= 0.41 {
						fmt.Println("肥胖")
					} else {
						fmt.Println("非常肥胖")
					}
				} else {
					fmt.Println("小孩子调皮捣蛋!")
				}
			}
		}

		fmt.Println("平均体脂率为:",h/3)

		var weightControl string
		fmt.Println("请输入y/n")
		fmt.Scanf("%s", &weightControl)
		if weightControl != "y" {
			break
		}
	}
}

//判断两条直线是否平行
func parallel() {
	var x1, x2, y1, y2 float64
	var a1, a2, b1, b2 float64
	fmt.Println("依次输入直线两个坐标:")
	fmt.Scanf("(%f,%f),(%f,%f),(%f,%f),(%f,%f)", &x1, &y1, &x2, &y2, &a1, &b1, &a2, &b2)
	inclination1:=(y2-y1)/(x2-x1)
	inclination2:=(b2-b1)/(a2-a1)
	if inclination1==inclination2 {
		fmt.Println("两直线平行")
	} else {
		fmt.Println("两直线不平行")
	}
}

func main() {
	//bodyFatCalculator()
	//bodyFatCalculators()
	parallel()
}
