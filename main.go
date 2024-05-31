package main

import "fmt"

type ttt struct {
	str string
	num int
}

func main() {
	var test []ttt
	var a, b ttt

	a.str = "aaa"
	a.num = 1

	b.str = "bbb"
	b.num = 2

	test = append(test, a)
	test = append(test, b)

	sub(&test)
	fmt.Println(test)
}

func sub(test *[]ttt) {
	for i := 0; i < len(*test); i++ {
		fmt.Println((*test)[i].str)
		fmt.Println((*test)[i].num)
		if (*test)[i].str == "aaa" {
			(*test)[i].str = "ccc"
		}
	}
}
