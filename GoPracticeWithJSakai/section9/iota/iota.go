package main

const (
	c1 = iota
	c2
	c3
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main() {
	println(c1, c2, c3)
	println(KB, MB, GB)
}
