package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Введите строку: ")
	var str1 string
	var ch string
	var str2 string
	var result string
	kovichki := '"'
	fmt.Scanln(&str1, &ch, &str2)
	str := str1 + ch + str2
	test(str, "+")
	test(str, "-")
	test(str, "*")
	test(str, "/")
	if is == 0 || ps != 1 {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(1)
	}
	ij := strings.Split(str, "")
	for _, v := range ij {
		switch v {
		case "+":
			ch = "+"
		case "-":
			ch = "-"
		case "*":
			ch = "*"
		case "/":
			ch = "/"
		}
	}
	sign := ij[is]
	var act int                     // какого типа строка
	number, _ := strconv.Atoi(str2) //число
	//------------------------------
	if strings.Index(str1, string(kovichki)) == 0 && strings.LastIndex(str1, string(kovichki)) == len(str1)-1 && strings.Index(str1, string(kovichki)) != strings.LastIndex(str1, string(kovichki)) {
		str1 = strings.ReplaceAll(str1, string(kovichki), "")
		if strings.Index(str1, string(kovichki)) != -1 {
			fmt.Println(fmt.Errorf("invalid input"))
			os.Exit(2)
		}
	} else {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(3)
	}
	if number > 1 {
		act = 1
	} else if strings.Index(str2, string(kovichki)) == 0 && strings.LastIndex(str2, string(kovichki)) == len(str2)-1 && strings.Index(str2, string(kovichki)) != strings.LastIndex(str2, string(kovichki)) {
		str2 = strings.ReplaceAll(str2, string(kovichki), "")
		act = 2
		if strings.Index(str2, string(kovichki)) != -1 {
			fmt.Println(fmt.Errorf("invalid input"))
			os.Exit(4)
		}
	} else {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(5)
	}
	// ----------------------
	if len(str1) > 10 {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(6)
	}
	if len(str2) > 10 {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(7)
	}
	if number > 10 {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(8)
	}
	//-----------------
	if sign == "+" && act == 1 {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(9)
	}
	if sign == "-" && act == 1 {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(10)
	}
	if sign == "*" && act == 2 {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(11)
	}
	if sign == "/" && act == 2 {
		fmt.Println(fmt.Errorf("invalid input"))
		os.Exit(12)
	}
	//------------
	if sign == "+" {
		result = str1 + str2
	}
	if sign == "-" {
		result = strings.ReplaceAll(str1, str2, "")
	}
	if sign == "*" {
		for i := 0; i < number; i++ {
			result += str1 // ****
		}
	}
	if sign == "/" {
		for i := 0; i < (len(str1) / number); i++ {
			k := strings.Split(str1, "")
			result += k[i]
		}
	}
	if len(result) > 40 {
		result1 := strings.Split(result, "")
		result = ""
		for i := 0; i < len(result1); i++ {
			if i == 40 {
				result += "..."
				break
			} else {
				result += result1[i]
			}
		}
	}
	fmt.Printf("\"%s\"\n", result)
}

var ps int
var is int

func test(st string, ch string) {
	if strings.Index(st, ch) == strings.LastIndex(st, ch) && strings.Index(st, ch) != -1 {
		is = strings.Index(st, ch)
		ps++
	}
}
