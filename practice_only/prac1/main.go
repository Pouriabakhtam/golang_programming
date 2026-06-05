package main

import (
	"fmt"
)

func main() {
	printvalue(1)
	printvalue("Hello")
	printvalue(3.14)
	fmt.Println("== Using type assertion with ok ==")
	printvalue2(1)
	printvalue2("Hello")
	printvalue2(3.14)
	fmt.Println("== Using add function with type assertion ==")
	result := add2(1, 2)
	result += 3
	result2 := add2("Hello, ", "world!")
	result2 += " How are you?"

	fmt.Println("Result of adding integers: ", result)
	fmt.Println("Result of adding strings: ", result2)

	fmt.Println("== Using substract function with type parameter ==")
	result3 := substract(5.9, 3)
	fmt.Println(result3)
}
func printvalue(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Print("Integer value: ", value.(int), "\n")
	case string:
		fmt.Print("String value: ", value.(string), "\n")
	case float64:
		fmt.Print("Float value: ", value.(float64), "\n")
	default:
		fmt.Print("Unknown type")
	}
}

func printvalue2(value interface{}) {
	intvalue, ok := value.(int)
	if ok {
		fmt.Println("Interger value: ", intvalue)
	}
	stringvalue, ok := value.(string)
	if ok {
		fmt.Println("String Value :", stringvalue)
	}
	floatvalue, ok := value.(float64)
	if ok {
		fmt.Println("Float Value :", floatvalue)
	}
}

func add(a, b interface{}) interface{} {
	aInt, aOK := a.(int)
	bInt, bOk := b.(int)
	if aOK && bOk {
		return aInt + bInt
	}
	aString, aOk := a.(string)
	bString, bOK := b.(string)
	if aOk && bOK {
		return aString + bString
	}
	return nil
}

func add2[T any](a, b T) T {
	aInt, aOK := a.(int)
	bInt, bOk := b.(int)
	if aOK && bOk {
		return aInt + bInt
	}
	aString, aOk := a.(string)
	bString, bOK := b.(string)
	if aOk && bOK {
		return aString + bString
	}
	var zero T
	return zero
}

func substract[T int | float64 | string](a, b T) T {
	return a + b
}
