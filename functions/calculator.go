package functions

import "fmt"

func depositBalance(balance *float64) {
	var amount float64
	for {
		inputFloat("Amount", &amount)
		validator, err := validateNumber(&amount)
		if !validator {
			fmt.Println(err)
			continue
		} else {
			break
		}
	}
	*balance += amount
}

func withdrawBalance(balance *float64)  {
	var amount float64
	for {
		inputFloat("Amount", &amount)
		validator, err := limitNumber(&amount, balance)
		if !validator {
			fmt.Println(err)
			if *balance == 0 {
				break
			}
			continue
		} else {
			break
		}
	}
	*balance -= amount
}