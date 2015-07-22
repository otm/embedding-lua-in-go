package main

import "github.com/yuin/gopher-lua"

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`function sayHello() print("Hello Again") end`); err != nil {
		panic(err)
	}

	if err := L.DoString(`sayHello()`); err != nil {
		panic(err)
	}
}
