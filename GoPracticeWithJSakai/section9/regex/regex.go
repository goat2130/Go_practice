package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	ms2 := r.MatchString("anna crace")
	fmt.Println(ms, ms2)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	ms3 := r2.FindString("/view/test")
	fmt.Println(ms3)

	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

	fss2 := r2.FindStringSubmatch("/save/test")
	fmt.Println(fss2, fss2[0], fss2[1], fss2[2])

	fss3 := r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss3, fss3[0], fss3[1], fss3[2])
}
