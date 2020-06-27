package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := 4

	if a == 4 {
		fmt.Println("a는 4")
	} else if a == 3 {
		fmt.Println("a는 3이 아니다.")
	} else {
		fmt.Println("a는 3도 아니고 4도 아니다.")
	}
	fmt.Println("a의 값은 ", a)

	fmt.Println("숫자를 입력해주세요.")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.", n1, n2)
}
