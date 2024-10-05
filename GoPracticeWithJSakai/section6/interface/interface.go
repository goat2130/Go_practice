package main

import "fmt"

// 指定したメソッドを持つインターフェースを定義
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type dog struct {
	Name string
}

// func (p Person) Say() string {
// 	return fmt.Sprintf("Hello %v\n", p.Name)
// }

// ポインター型を引数に取るメソッドを定義
func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	message := fmt.Sprintf("Hello %v\n", p.Name)
	fmt.Print(message)
	return p.Name
}

func driveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Stop")
	}
}

func main() {
	// var mike Human = Person{"Mike"}

	// ポインター型を引数に取るメソッドを定義しているため、&をつける
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	// var dog dog = dog{"dog"}

	// fmt.Println(mike.Say())

	driveCar(mike)
	driveCar(x)
	// driveCar(dog) is error(missing say method)
}
