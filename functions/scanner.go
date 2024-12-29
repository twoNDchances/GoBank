package functions

import "fmt"

func inputFloat(text string, target *float64) {
	fmt.Printf("Enter %s: ", text)
	fmt.Scan(target)
}

func inputInteger(text string, target *int) {
	fmt.Printf("Enter %s: ", text)
	fmt.Scan(target)
}
