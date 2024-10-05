package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func structEx() {
	v := Vertex{X: 1, Y: 2}
	v2 := Vertex{X: 1}
	fmt.Println(v)
	fmt.Println(v2)
	fmt.Println(v.X, v.Y, v2.X, v2.Y)

	v3 := Vertex{10, 20, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)
}

func one(x *int) {
	*x = 1
}

func newAndMake() {
	// メモリを確保しているので、アドレスが表示される
	var p *int = new(int)
	fmt.Println(p)
	fmt.Printf("値は%vになります\n", *p)
	*p++
	fmt.Printf("インクリメントも対応し値は%vになります\n", *p)

	// メモリを確保していないので、nilが表示される
	var p2 *int
	fmt.Println(p2)
	// メモリを確保せずこのポインターはnilなので、インクリメントはできない
	// *p2++
	// fmt.Println(p2)

	// makeはスライス、マップ、チャネルを作成する際に使用する
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	// ポインター型とストラクト型の場合はnewを使用し、メモリを確保する。
	var p3 *int = new(int)
	fmt.Printf("%T\n", p3)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	(*v).X = 1000
}

func compareStructAndPointer() {
	v := Vertex{X: 1, Y: 2}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{X: 1, Y: 2}
	changeVertex2(v2)
	fmt.Println(*v2)
}

func main() {
	var x int = 100
	fmt.Println(x)
	fmt.Println(&x)

	one(&x)
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*&x)

	var n int = 100

	fmt.Println(n)

	fmt.Println(&n)

	var p *int = &n

	fmt.Println(p)
	fmt.Println(*p)

	newAndMake()

	structEx()

	compareStructAndPointer()
}
