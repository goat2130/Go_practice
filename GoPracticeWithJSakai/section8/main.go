package time

import (
	"GoPracticeWithJSakai/section8/mylib"
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(x))

	mylib.Say()
	p := mylib.Person{Name: "Milke", Age: 23}
	fmt.Println(p)

	fmt.Println(mylib.Public)
	// fmt.Println(mylib.private)
}
