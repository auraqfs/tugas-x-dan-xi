package main

import (
	"fmt"
	"regexp"
	"strings"
)

func match(patern string, input string) {
	matched, _ := regexp.Match(patern, []byte(input))

	reg, _ := regexp.Compile("[^A-Za-z]+")

	e := reverse
	result := reg.ReplaceAllString(input, "")
	if matched {
		fmt.Println(strings.Title(strings.ToLower(e(result))))
	} else {
		fmt.Println("Error: Harus Terdapat string")
	}
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	var str, kalimat, strnum string
	fmt.Println("Input : \n")
	fmt.Scanln(&str)
	pattern := "[:word:]"
	fmt.Scanln(&kalimat)
	pattern1 := "[:word:]"
	fmt.Scanln(&strnum)
	pattern2 := "[:word:]"

	fmt.Println("\nOutput : \n")

	match(pattern, str)
	match(pattern1, kalimat)
	match(pattern2, strnum)

}
