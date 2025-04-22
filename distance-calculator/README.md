# Kratix Distance Calculator

Простой калькулятор расстояния между двумя точками на плоскости.

## Установка

```bash
go mod tidy
```

## Запуск

```bash
go run cmd/distance/main.go
```

## Использование

Введите координаты двух точек в формате:
```
x1 y1 x2 y2
```

Например:
```
1 2 4 6
```

## Структура проекта

```
kratix/
├── cmd/
│   └── distance/
│       └── main.go
├── internal/
│   └── calculator/
│       └── distance.go
└── go.mod
```

## Ссылки

- [GitHub Repository](https://github.com/saneaizog/kratix) 