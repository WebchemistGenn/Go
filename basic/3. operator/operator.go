package main

import "fmt"

func main() {
	a := 4
	b := 2
	// 이항 연산자

	// 단항 연산자

	// binary operator
	// "||" | "&&" | rel | add | mul

	// rel operator
	// "==" | "!=" | "<" | "<=" | ">" | ">="
	fmt.Printf("a == b = %v\n", a == b)
	fmt.Printf("a != b = %v\n", a != b)
	fmt.Printf("a < b = %v\n", a < b)
	fmt.Printf("a <= b = %v\n", a <= b)
	fmt.Printf("a > b = %v\n", a > b)
	fmt.Printf("a >= b = %v\n", a >= b)

	// add operator
	// "+" | "-" | "|" | "^"
	fmt.Printf("a + b = %v\n", a+b)
	fmt.Printf("a - b = %v\n", a-b)
	fmt.Printf("a | b = %v\n", a|b)
	fmt.Printf("a ^ b = %v\n", a^b)

	// mul operator
	// "*" | "/" | "%" | "<<" | ">>" | "&" | "&^"
	fmt.Printf("a * b = %v\n", a*b)
	fmt.Printf("a / b = %v\n", a/b)
	fmt.Println("a % b = ", a%b)
	fmt.Printf("a << 1 = %v\n", a<<1)
	fmt.Printf("a >> 1 = %v\n", a>>1)

	// unary operator
	// "+" | "-" | "!" | "^" | "*" | "&" | "<-"

	// 산술연산자 - ( +, -, *, / )
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)

	// 비트연산자 - ( &, |, ^, ^ )
	fmt.Println("a / b = ", a/b)

	// 논리연산자 - <, <=, >, >=, ==, !=, &&, ||, !

	// 그외

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("a|b = %v\n", a|b)
	fmt.Printf("a^b = %v\n", a^b)

	fmt.Printf("a<<1 = %v\n", a<<1)
	fmt.Printf("a>>1 = %v\n", a>>1)
}
