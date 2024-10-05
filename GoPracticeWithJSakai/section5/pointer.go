package main

import "fmt"

type Vertex struct {
	x, y int
	S    string
}

type Vertex2 struct {
	x, y int
}

// 値レシーバー
func (v Vertex2) Area() int {
	return v.x * v.y
}

// ポインターレシーバー
func (v *Vertex2) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

type Vertex2_3D struct {
	Vertex2
	z int
}

// 値レシーバー
func (v Vertex2_3D) Area3D() int {
	return v.x * v.y * v.z
}

// ポインターレシーバー
func (v *Vertex2_3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex2_3D {
	return &Vertex2_3D{Vertex2{x, y}, z}
}

func Area(v Vertex2) int {
	return v.x * v.y
}

func structEx() {
	v := Vertex{x: 1, y: 2}
	v2 := Vertex{x: 1}
	fmt.Println(v)
	fmt.Println(v2)
	fmt.Println(v.x, v.y, v2.x, v2.y)

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
	v.x = 1000
}

func changeVertex2(v *Vertex) {
	(*v).x = 1000
}

func compareStructAndPointer() {
	v := Vertex{x: 1, y: 2}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{x: 1, y: 2}
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

	// v8 := Vertex2{3, 4}
	v8 := New(3, 4, 5)
	v8.Scale(10)
	fmt.Println(v8.Area())
	fmt.Println(v8.Area3D())
}
