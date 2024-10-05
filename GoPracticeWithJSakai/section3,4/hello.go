package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
	runeKey = 'a'
)

const z = 20 * 10

func main() {
	// fmt.Println("Hello, World!")
	// var x int
	// x = 10
	// fmt.Println(x)

	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// x = x + 10
	// y = "world"

	// fmt.Println(x)
	// fmt.Println(y)

	// 	yagisawahodaka@LAB-N2964R ch1 % go run hello.go
	// # command-line-arguments
	// ./hello.go:26:2: cannot assign to x (neither addressable nor a map index expression)
	// ./hello.go:27:2: cannot assign to y (neither addressable nor a map index expression)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	t := 10
	fmt.Println(t)
	fmt.Println(1+1, 2+2)
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1", 10-1)
	fmt.Println("10 / 2", 10/2)
	fmt.Println("10 / 3", 10/3)
	fmt.Println("10.0 / 3", 10.0/3)
	fmt.Println("10 / 3.0", 10/3.0)
	fmt.Println("10 % 2", 10%2)
	fmt.Println("10 % 3", 10%3)

	r := 0
	fmt.Println(r)
	r++
	fmt.Println(r)
	r--
	fmt.Println(r)

	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000
	stringsExample()

	boo()

	cast()

	array()

	slice()

	mapEx()

	byteEx()

	add(1, 2)

	cal(100, 2)

	f := func(num int) {
		fmt.Println("inner func", num)
	}
	f(3)

	// 並列化で使える記述方法なので覚えておくが吉
	func(num int) {
		fmt.Println("inner func", num)
	}(30)

	foo(0)

	foo(10, 20)

	foo(30, 40, 50)

	// 可変長引数のスライスを渡す
	s := []int{1, 2, 3}
	foo(s...)

	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}

	// 上記の記述方法を簡略化
	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	/*
			// これはエラー
		fmt.Println(result2)
	*/

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("continueやbreak を使った場合")
	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("continue")
			continue
		}

		if i == 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	rangeEx()

	switchEx()

	deferEx()

	// errHndle()

	save()
}

func stringsExample() {
	fmt.Println("Hello World")
	fmt.Println("Hello" + "World")
	fmt.Println("Hello World"[0])
	fmt.Println(string("Hello World"[0:5]))

	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(s)

	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(`test
	                        test
		test
test`)
	fmt.Println(`"`)
	fmt.Println("\"")
}

func boo() {
	t, f := true, false
	fmt.Println(t, f)
	fmt.Printf("%t %v %t\n", t, t, t)
	fmt.Printf("%t %v %t\n", f, f, f)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
}

func cast() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	// z = int(s) キャストできない
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T %v %d\n", i, i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}

func array() {
	var a [2]int
	a[0] = 100
	a[1] = 200

	fmt.Println(a)
	/*
		    var b [2]int = [2]int{100, 200}
			b = b.append(b, 300)
			配列はappendできない
	*/

	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}

func slice() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)

	n = append(n, 7, 8, 9)
	fmt.Println(n)

	board := [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)

	// make関数
	m := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m), m)

	m = append(m, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m), m)

	m = append(m, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m), m)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d valuw=%v\n", len(a), cap(a), a)

	b := make([]int, 0)
	var c []int
	fmt.Printf("lew=%d cap=%d value=%v\n", len(b), cap(b), b) // 0のスライスをメモリに確保している
	fmt.Printf("lew=%d cap=%d value=%v\n", len(c), cap(c), c) // nilのスライスなのでメモリを確保していない

	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	// d = make([]int, 0, 5)
	// for i := 0; i < 5; i++ {
	// 	d = append(d, i)
	// 	fmt.Println(d)
	// }
	// fmt.Println(d)
}

func mapEx() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)

	m["banana"] = 300
	fmt.Println(m)

	m["new"] = 500
	fmt.Println(m)

	v3 := m["apple"]
	v, ok := m["apple"]
	fmt.Println(v, ok, v3)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)
}

func byteEx() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}

func add(x, y int) int {
	return x + y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

// 可変長引数
func foo(params ...int) {
	fmt.Println(len(params), params)

	for _, param := range params {
		fmt.Println(param)
	}
}

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	}
	return "no"
}

func rangeEx() {
	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// rangeoを使った場合
	for i, v := range l {
		fmt.Println(i, v)
	}

	// インデックスを使わない場合
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}
}

func getOs() string {
	return "mac"
}

func switchEx() {
	os := "windows"
	switch os {
	case "mac":
		fmt.Println("mac!!")
	case "windows":
		fmt.Println("windows!!")
	default:
		fmt.Println("default!!")
	}

	switch os2 := getOs(); os2 {
	case "mac":
		fmt.Println("mac!!")
	case "windows":
		fmt.Println("windows!!")
	default:
		fmt.Println("default!!")
	}

	fmt.Println(os)
	// 出力されない
	// fmt.Println(os2)

	t := time.Now()
	fmt.Printf("%T %v\n", t, t)

	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	case t.Hour() < 20:
		fmt.Println("evening")
	}
}

func deferEx() {
	defer fmt.Println("world")
	fmt.Println("hello")

	// stacking defers
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("done")
}

func errHndle() {
	file, err := os.Open("hello.go")
	if err != nil {
		log.Fatalln("Error")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)

	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	if err := os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}
}

func thirdPartyConnectDB() {
	/*
		基本的にpanicはエラー処理の最終手段として使う
		err　!= nil でエラー処理を行う
	*/
	panic("Unable to connet database")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}
