package exercises

import (
	"fmt"
	"strconv"
)

func ConvertToInt(number string) (int, string) {
	intValue, err := strconv.Atoi(number)

	if err != nil {
		return 0, fmt.Sprintf("Error Inesperado: %v", err)
	}

	message := "Es Menor a 100"
	if intValue > 100 {
		message = "Es Mayor a 100"
	}

	return intValue, message
}
