package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/saneaizog/kratix/distance-calculator/internal/calculator"
)

func main() {
	// Подключение к БД
	db, err := connectDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка ping к БД:", err)
	}
	log.Println("Успешное подключение к БД")

	// Ваша существующая логика
	formula := os.Getenv("FORMULA")
	fmt.Printf("Используемая формула: %s\n", formula)

	var x1, y1, x2, y2 float64
	fmt.Println("Введите координаты двух точек (x1 y1 x2 y2):")
	_, err = fmt.Scan(&x1, &y1, &x2, &y2)
	if err != nil {
		log.Fatal("Ошибка ввода:", err)
	}

	result := calculator.CalculateDistance(x1, y1, x2, y2)
	fmt.Printf("Расстояние между точками: %.2f\n", result)

	// Пример сохранения в БД
	log.Printf("Попытка сохранения данных: x1=%.2f, y1=%.2f, x2=%.2f, y2=%.2f, distance=%.2f", x1, y1, x2, y2, result)
	_, err = db.Exec("INSERT INTO distances (x1, y1, x2, y2, distance) VALUES ($1, $2, $3, $4, $5)",
		x1, y1, x2, y2, result)
	if err != nil {
		log.Printf("Ошибка сохранения в БД: %v", err)
	} else {
		log.Println("Данные успешно сохранены в БД")
	}
}

func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	log.Printf("Подключение к БД: %s", connStr)
	return sql.Open("postgres", connStr)
}
