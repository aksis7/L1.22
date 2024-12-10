package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация двух больших чисел
	a := new(big.Int)
	b := new(big.Int)

	// Пример чисел больше 2^20
	a.SetString("1048577", 10) // 2^20 + 1
	b.SetString("2097154", 10) // 2^21

	// Вывод чисел
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Println("a + b =", sum)

	// Вычитание
	diff := new(big.Int).Sub(a, b)
	fmt.Println("a - b =", diff)

	// Умножение
	product := new(big.Int).Mul(a, b)
	fmt.Println("a * b =", product)

	// Деление (с округлением вниз)
	quotient := new(big.Int).Div(a, b)
	fmt.Println("a / b =", quotient)

	// Остаток от деления
	remainder := new(big.Int).Mod(a, b)
	fmt.Println("a % b =", remainder)
}
