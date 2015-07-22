package main

import "github.com/yuin/gopher-lua"

func square(L *lua.LState) int {
	i := L.ToInt(1)          // retrieve first function argument and convert to integer
	ln := lua.LNumber(i * i) // make calculation and cast to LNumber
	L.Push(ln)               // Push it to the stack
	return 1                 // Notify that we pushed one value to the stack
}

func main() {
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("square", L.NewFunction(square))
	if err := L.DoString(`print("4 * 4 = " .. square(4))`); err != nil {
		panic(err)
	}
}
