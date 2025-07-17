package main

import "fmt"

func LCM(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]

	for counter := 1; counter < len(numbers); counter++ {
		gcd := GCD(result, numbers[counter])
		result = (result * numbers[counter]) / gcd
	}
	return result
}

func GCD(numberOne int, numberTwo int) int {
	if numberTwo == 0 {
		return numberOne
	} else {
		return GCD(numberTwo, numberOne%numberTwo)
	}
}

func main() {
	lcm := LCM([]int{14, 12, 7, 8})
	fmt.Println(lcm)

}
