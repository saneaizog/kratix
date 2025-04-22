package main

import (
	"fmt"
	"log"

	"github.com/saneaizog/kratix/tree/main/distance-calculator/distance-calculator/internal/calculator"
)

func main() {
	var a, b, c, d float64
	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil {
		log.Fatal("Ошибка ввода:", err)
	}

	result := calculator.CalculateDistance(a, b, c, d)
	fmt.Println(result)
}
