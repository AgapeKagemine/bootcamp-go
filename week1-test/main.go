package main

import (
	"fmt"
	"math/big"
)

func score(score int64) (result string) {
	result = "Anda mendapatkan nilai "
	if score < 60 {
		result += "E"
	} else if score < 70 {
		result += "D"
	} else if score < 80 {
		result += "C"
	} else if score < 90 {
		result += "B"
	} else {
		result = "Selamat! " + result + "A"
	}
	return result
}

func calculate(diskon bool, price, qty float64) (total float64) {
	total = qty * price
	if diskon {
		total -= total * 0.1
	}
	return total
}

func VowelCounter(text string) (vocal int64) {
	vocal = 0
	for _, char := range text {
		if char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' || char == 'A' || char == 'I' || char == 'U' || char == 'E' || char == 'O' {
			vocal += 1
		}
	}
	return vocal
}

func factorial(number *big.Int) (result *big.Int) {
	result = new(big.Int)
	if number.Cmp(big.NewInt(0)) == -1 {
		result.SetInt64(1)
	}
	if number.Cmp(big.NewInt(0)) == 0 {
		result.SetInt64(1)
	} else {
		result.Set(number)
		result.Mul(result, factorial(number.Sub(number, big.NewInt(1))))
	}
	return
}

func OddEven(number string) (result string) {
	num := 0
	if number == "satu" {
		num = 1
	}
	if number == "dua" {
		num = 2
	}
	if number == "tiga" {
		num = 3
	}
	if number == "empat" {
		num = 4
	}
	if number == "lima" {
		num = 5
	}
	if number == "enam" {
		num = 6
	}
	if number == "tujuh" {
		num = 7
	}
	if number == "delapan" {
		num = 8
	}
	if number == "sembilan" {
		num = 9
	}
	if number == "sepuluh" {
		num = 10
	}

	if num%2 == 0 {
		result = "genap"
	} else {
		result = "ganjil"
	}

	return result
}

func main() {
	// text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	fmt.Println(OddEven("satu"))
	fmt.Println(factorial(big.NewInt(30)))
}
