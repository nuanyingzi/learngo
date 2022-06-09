package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type: %T value: %#v\n", i1, i1)
	}

	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type: %T value: %#v\n", s2, s2)

	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("can't convert str to bool")
	} else {
		fmt.Printf("b type: %T value: %#v\n", b, b)
	}

	c, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println("can't convert str to float64")
	} else {
		fmt.Printf("c type: %T value: %#v\n", c, c)
	}

	e, err := strconv.ParseInt("100", 10, 64)
	if err != nil {
		fmt.Println("can't convert str to int64 十进制")
	} else {
		fmt.Printf("e type: %T value: %#v\n", e, e)
	}

	ss1 := strconv.FormatBool(true)
	ss2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	ss3 := strconv.FormatInt(-2, 16)
	ss4 := strconv.FormatUint(2, 16)
	fmt.Println(ss1, ss2, ss3, ss4)
}
