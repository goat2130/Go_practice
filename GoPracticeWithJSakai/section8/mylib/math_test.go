package mylib

import "testing"

/*
func Average(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total / len(x)
}
この関数をテストする
*/

func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5, 6, 7})
	if v != 3 {
		t.Error("Expected 3, but got", v)
	}
}

/*
yagisawahodaka@LAB-N2964R section8 % go test ./...
?       GoPracticeWithJSakai/section8   [no test files]
ok      GoPracticeWithJSakai/section8/mylib     0.289s


yagisawahodaka@LAB-N2964R section8 % go test -v ./...
?       GoPracticeWithJSakai/section8   [no test files]
=== RUN   TestAverage
--- PASS: TestAverage (0.00s)
PASS
ok      GoPracticeWithJSakai/section8/mylib     0.201s


// []int{1, 2, 3, 4, 5, 6, 7}の平均は4であるため、テストが失敗する
yagisawahodaka@LAB-N2964R section8 % go test -v ./...
?       GoPracticeWithJSakai/section8   [no test files]
=== RUN   TestAverage
    math_test.go:19: Expected 3, but got 4
--- FAIL: TestAverage (0.00s)
FAIL
FAIL    GoPracticeWithJSakai/section8/mylib     0.197s
FAIL
*/
