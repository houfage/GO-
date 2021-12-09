package main

import "fmt"

func main() {
	var totalfatrate float64
	names := [3]string{}
	weights := [3]float64{}
	talls := [3]float64{}
	ages := [3]int{}
	bmis := [3]float64{}
	fatrates := [3]float64{}

	for i := 0; i < 3; i++ {

		fmt.Print("姓名：")
		fmt.Scanln(&names[i])

		fmt.Print("体重：")
		fmt.Scanln(&weights[i])

		fmt.Print("身高：")
		fmt.Scanln(&talls[i])

		fmt.Print("年龄：")
		fmt.Scanln(&ages[i])

		bmis[i] = weights[i] / (talls[i] * talls[i])

		var sexweight int
		var sex string = "男"
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)

		if sex == "男" {
			sexweight = 1
		} else {
			sexweight = 0
		}

		var fatrate float64 = (1.2*bmis[i] + 0.23*float64(ages[i]) - 5.4 - 10.8*float64(sexweight)) / 100
		fatrates[i] = fatrate
		fmt.Println("体脂率是：", fatrate)

		if sex == "男" {
			if ages[i] >= 18 && ages[i] <= 39 {
				if fatrates[i] <= 0.1 {
					fmt.Println("偏瘦。要多吃多锻炼，增强体质。")
				} else if fatrates[i] > 0.1 && fatrates[i] <= 0.16 {
					fmt.Println("标准。很好，要保持。")
				} else if fatrates[i] > 0.16 && fatrates[i] <= 0.21 {
					fmt.Println("偏胖。吃完饭多散步，消化消化。")
				} else if fatrates[i] > 0.21 && fatrates[i] <= 0.26 {
					fmt.Println("肥胖。少吃点，多运动。")
				} else {
					fmt.Println("非常肥胖。健身游泳跑步。")
				}

			} else if ages[i] >= 40 && ages[i] <= 59 {
				if fatrates[i] <= 0.1 {
					fmt.Println("偏瘦。要多吃多锻炼，增强体质。")
				} else if fatrates[i] > 0.1 && fatrates[i] <= 0.16 {
					fmt.Println("标准。很好，要保持。")
				} else if fatrates[i] > 0.16 && fatrates[i] <= 0.21 {
					fmt.Println("偏胖。吃完饭多散步，消化消化。")
				} else if fatrates[i] > 0.21 && fatrates[i] <= 0.26 {
					fmt.Println("肥胖。少吃点，多运动。")
				} else {
					fmt.Println("非常肥胖。健身游泳跑步。")
				}

			} else if ages[i] >= 60 {
				if fatrates[i] <= 0.1 {
					fmt.Println("偏瘦。要多吃多锻炼，增强体质。")
				} else if fatrates[i] > 0.1 && fatrates[i] <= 0.16 {
					fmt.Println("标准。很好，要保持。")
				} else if fatrates[i] > 0.16 && fatrates[i] <= 0.21 {
					fmt.Println("偏胖。吃完饭多散步，消化消化。")
				} else if fatrates[i] > 0.21 && fatrates[i] <= 0.26 {
					fmt.Println("肥胖。少吃点，多运动。")
				} else {
					fmt.Println("非常肥胖。健身游泳跑步。")
				}

			} else {
				fmt.Println("无法评判未成年人。")
			}

		} else {
			if ages[i] >= 18 && ages[i] <= 39 {
				if fatrates[i] <= 0.1 {
					fmt.Println("偏瘦。要多吃多锻炼，增强体质。")
				} else if fatrates[i] > 0.1 && fatrates[i] <= 0.16 {
					fmt.Println("标准。很好，要保持。")
				} else if fatrates[i] > 0.16 && fatrates[i] <= 0.21 {
					fmt.Println("偏胖。吃完饭多散步，消化消化。")
				} else if fatrates[i] > 0.21 && fatrates[i] <= 0.26 {
					fmt.Println("肥胖。少吃点，多运动。")
				} else {
					fmt.Println("非常肥胖。健身游泳跑步。")
				}

			} else if ages[i] >= 40 && ages[i] <= 59 {
				if fatrates[i] <= 0.1 {
					fmt.Println("偏瘦。要多吃多锻炼，增强体质。")
				} else if fatrates[i] > 0.1 && fatrates[i] <= 0.16 {
					fmt.Println("标准。很好，要保持。")
				} else if fatrates[i] > 0.16 && fatrates[i] <= 0.21 {
					fmt.Println("偏胖。吃完饭多散步，消化消化。")
				} else if fatrates[i] > 0.21 && fatrates[i] <= 0.26 {
					fmt.Println("肥胖。少吃点，多运动。")
				} else {
					fmt.Println("非常肥胖。健身游泳跑步。")
				}

			} else if ages[i] >= 60 {
				if fatrates[i] <= 0.1 {
					fmt.Println("偏瘦。要多吃多锻炼，增强体质。")
				} else if fatrates[i] > 0.1 && fatrates[i] <= 0.16 {
					fmt.Println("标准。很好，要保持。")
				} else if fatrates[i] > 0.16 && fatrates[i] <= 0.21 {
					fmt.Println("偏胖。吃完饭多散步，消化消化。")
				} else if fatrates[i] > 0.21 && fatrates[i] <= 0.26 {
					fmt.Println("肥胖。少吃点，多运动。")
				} else {
					fmt.Println("非常肥胖。健身游泳跑步。")
				}

			} else {
				fmt.Println("无法评判未成年人。")
			}

		}

	}

	for i := 0; i < 3; i++ {
		totalfatrate += fatrates[i]
		fmt.Println(names[i], weights[i], talls[i], ages[i], bmis[i], fatrates[i])
	}
	fmt.Println("平均体脂率是：", totalfatrate/3)
	fmt.Println("总人数：", len(names))

}
