package main

import (
	"fmt"
	"log"
	"os"

	"github.com/saneaizog/kratix/distance-calculator/internal/calculator"
)

func main() {
	formula := os.Getenv("FORMULA")
	fmt.Printf("Используемая формула: %s\n", formula)
	var x1, y1, x2, y2 float64
	fmt.Println("Введите координаты двух точек (x1 y1 x2 y2):")
	_, err := fmt.Scan(&x1, &y1, &x2, &y2)
	if err != nil {
		log.Fatal("Ошибка ввода:", err)
	}

	result := calculator.CalculateDistance(x1, y1, x2, y2)
	fmt.Printf("Расстояние между точками: %.2f\n", result)
}
