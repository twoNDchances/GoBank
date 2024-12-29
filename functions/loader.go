package functions

import (
	"fmt"
	"os"
	"strconv"
)

func importBalance(fileName string, balance *float64) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(customErrorText("File not found, load default balance"))
		*balance = 1000
		return
	}
	*balance, err = strconv.ParseFloat(fmt.Sprint(data), 64)
	if err != nil {
		fmt.Println(customErrorText("Can't parse balance, load default balance"))
		*balance = 1000
		return
	}
}

func exportBalance(fileName string, balance *float64)  {
	if os.WriteFile(fileName, []byte(fmt.Sprint(*balance)), 0644) != nil {
		fmt.Println(customErrorText("Can't write to file, data in file not changed"))
	}
}