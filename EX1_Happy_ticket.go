package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	var first_part int = number / 1000
	var second_part int = number % 1000

	var first_part_1, first_part_2, first_part_3, first_part_sum int
	var second_part_1, second_part_2, second_part_3, second_part_sum int

	first_part_1 = first_part % 10
	first_part_2 = (first_part / 10) % 10
	first_part_3 = (first_part / 100) % 10
	first_part_sum = first_part_1 + first_part_2 + first_part_3

	second_part_1 = second_part % 10
	second_part_2 = (second_part / 10) % 10
	second_part_3 = (second_part / 100) % 10
	second_part_sum = second_part_1 + second_part_2 + second_part_3

	if first_part_sum == second_part_sum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
