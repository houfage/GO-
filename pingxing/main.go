package main

import "fmt"

func main() {
	hengs := [2][2]float64{}
	shus := [2][2]float64{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("第" + string(i+1) + "条线的第" + string(j+1) + "个点的横坐标：")
			fmt.Scanln(&hengs[i][j])
			fmt.Print("第" + string(i+1) + "条线的第" + string(j+1) + "个点的纵坐标：")
			fmt.Scanln(&shus[i][j])

		}

	}

	if (shus[0][1]-shus[0][0])*(hengs[1][1]-hengs[1][0]) == (shus[1][1]-shus[1][0])*(hengs[0][1]-hengs[0][0]) {
		fmt.Println("两线平行")
	} else {
		fmt.Println("两线不平行")
	}

}
