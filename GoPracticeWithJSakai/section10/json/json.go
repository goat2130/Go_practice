package main

import (
	"encoding/json"
	"fmt"
)

type T struct{}

type Person struct {
	Name      string   `json:"name"`                 // このように("-")指定することで、JSON形式のバイト配列に変換する際に無視される
	Age       int      `json:"age, string"`          // このように指定することで、JSON形式のバイト配列に変換する際に指定した名前に変換される
	Nicknames []string `json:"nicknames, omitempty"` // このように指定することで、ゼロ値の場合に無視される
	t         *T       `json:"T", omitempty`
}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

func main() {
	b := []byte(`{"name": "Gopher", "age": 10, "nicknames": ["Gopher", "Gopher-san", "Gopher-kun"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// 構造体の値をJSON形式のバイト配列に変換 (Marshal)`json:"nameXXXXX"`のように指定ができる
	v, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(v))
	//  10 [Gopher Gopher-san Gopher-kun]
	//{"nameXXXXX":"","age":10,"nicknames":["Gopher","Gopher-san","Gopher-kun"]}
}
