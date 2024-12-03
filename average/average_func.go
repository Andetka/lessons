package average

import (
	"fmt"
)

func Average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
func main() {
	fmt.Println(Average(100, 50))
	fmt.Println(Average(90.7, 89.7, 98.5, 92.3))
}
