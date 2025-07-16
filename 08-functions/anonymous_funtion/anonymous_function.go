package anonymousfuntion

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	triple := createTransformer(3)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transformation func(int) int) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transformation(value))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
