package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
}

type MyType struct {
	A string
	FooBar string
}

func main() {
	user := User{"test"}
	b, _ := json.Marshal(user)
	fmt.Printf("%s\n", string(b))
	var mt MyType
	json.Unmarshal([]byte(`{"A":"aaa", "FooBar":"baz"}`), &mt)
	fmt.Printf("%#v\n", mt)
}