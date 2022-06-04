package fizzBuzz

import (
	"fmt"
	"strconv"
)

func fizzBuzz() {
	fmt.Println("hello")

	m := make(map[int]string)
	m[3] = "fizz"
	m[5] = "buzz"
	m[10] = "bazz"
	general(m, 100)
}

func general(m map[int]string, rng int) {
	msg := ""

	for i := 1; i < rng; i++ {
		for j, v := range m {
			if i%j == 0 {
				msg += v
			}
		}
		msg = strconv.Itoa(i) + ": " + msg

		fmt.Println(msg)
		msg = ""
	}
}

func naive() {
	for i := 1; i < 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		} else {
			fmt.Println(i)
		}
	}
}
