package functions

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const fileName string = "balance.txt"

func Core() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Connect with us:", randomdata.PhoneNumber())
	var choice int
	var balance float64
	importBalance(fileName, &balance)
	for {
		presentMenu()
		inputInteger("your choice", &choice)
		switch choice {
		case 1:
			fmt.Println("Balance:", balance)
		case 2:
			depositBalance(&balance)
			fmt.Println("Balance:", balance)
			exportBalance(fileName, &balance)
		case 3:
			withdrawBalance(&balance)
			fmt.Println("Balance:", balance)
			exportBalance(fileName, &balance)
		default:
			fmt.Println("Goodbye & Thank you!")
			return
		}
	}
}