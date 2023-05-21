package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // ถ้าแปลงเป็น float ไม่ได้มีตัวแปร err รองรับ
	if err != nil {
		msg, _ := fmt.Scanf("%v number only", promt)
		panic(msg) // print
	}
	return value
}
func getoperator() string {
	fmt.Print("operator(+ - * /) : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(val1, val2 float64) float64 {
	return val1 + val2
}
func del(val1, val2 float64) float64 {
	return val1 - val2
}
func mul(val1, val2 float64) float64 {
	return val1 * val2
}
func dev(val1, val2 float64) float64 {
	return val1 / val2
}
func main() {
	var num int
	fmt.Println("num :")
	fmt.Scanf("%d", &num)

	var cal []string
	for i := 0; i < num; i++ {
		var text string
		var op string
		fmt.Println("number" + strconv.Itoa(i+1) + " :")
		fmt.Scanf("%s", &text)
		cal = append(cal, text)
		if i < num {
			fmt.Println("op :")
			fmt.Scanf("%s", &op)
			cal = append(cal, op)
		}
	}
	fmt.Println(cal)

	// val1 := getInput("value 1 = ")
	// val2 := getInput("value 2 = ")
	// var result float64
	// switch operator := getoperator(); operator {
	// case "+":
	// 	result = add(val1, val2)
	// case "-":
	// 	result = del(val1, val2)
	// case "*":
	// 	result = mul(val1, val2)
	// case "/":
	// 	result = dev(val1, val2)
	// default:
	// 	panic("wrong operator")
	// }

	// fmt.Printf("result : %v", result)
}
