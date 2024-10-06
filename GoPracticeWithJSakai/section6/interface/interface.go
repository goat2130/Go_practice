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

func do(i interface{}) {
	/*
		// タイプアサーション
		ii := i.(int)
		ii *= 2
		fmt.Println(ii)
	*/

	// switch type文
	switch v := i.(type) {
	case int:
		ii := v * 2
		fmt.Println(ii)
	case string:
		ss := v + "!"
		fmt.Println(ss)
	case bool:
		if v == true {
			fmt.Printf("true is %v\n", v)
		} else {
			fmt.Printf("false is %v\n", v)
		}
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

type Person2 struct {
	Name string
	Age  int
}

func (p Person2) string() string {
	return fmt.Sprintf("My name is %v\n", p.Name)
}

// Custom error
type userNotFound struct {
	Username string
}

func (e *userNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &userNotFound{Username: "mike"}
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

	do(10)
	do("Mike")
	do(true)
	do(false)
	do(2.5)

	john := Person2{"John", 22}
	fmt.Println(john)
	fmt.Println(john.string())

	// Custom error
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
