package basic

import (
	"errors"
	"fmt"
	"math"
)

func basic() {
	// 메세지 출력
	fmt.Println("Hello Webchemist")

	// 변수선언 ( 다음과 같이 값에 의해 자동선언이 되게 할 수 있다. )
	value001 := "Welcome to Webchemist" // var value003 string = "Welcome to Webchemist"
	value002 := 10                      // var value001 int = 10
	value003 := 20                      // var value002 int = 20
	array001 := []int{5, 4, 3, 2, 1}
	array001 = append(array001, 13)
	array002 := make(map[string]int)

	array002["test001"] = 1
	array002["test002"] = 2

	// 출력
	fmt.Println(value001)
	fmt.Println(value002)
	fmt.Println(value003)
	fmt.Println(array001)
	fmt.Println(array002["test001"])

	// IF 문
	if value002 > 6 {
		fmt.Println("6보다 큽니다.")
	} else if value002 < 6 {
		fmt.Println("6보다 작습니다.")
	} else {
		fmt.Println("6입니다.")
	}

	// FOR 문
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 배열 반복
	array003 := []string{"a", "b", "c"}
	for index, value := range array003 {
		fmt.Println("index", index, "value", value)
	}

	// Map 반복
	array004 := make(map[string]string)
	array004["a"] = "test001"
	array004["b"] = "test002"
	for key, value := range array004 {
		fmt.Println("key", key, "value", value)
	}

	sum001 := 5
	sum002 := 7
	sumResult := sum(sum001, sum002)
	fmt.Println(sumResult)

	sqrtReulst, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrtReulst)
	}
}

// sum 함수
func sum(x int, y int) int {
	return x + y
}

// 함수 에러처리
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
