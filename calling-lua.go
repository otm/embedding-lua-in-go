package main

import (
	"fmt"

	"github.com/yuin/gopher-lua"
)

// luaCode is the Lua code we want to call from Go
var luaCode = `
function concat(a, b)
  return a .. " + " .. b
end
`

func main() {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoString(luaCode); err != nil {
		panic(err)
	}

	// Call the Lua function concat
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("concat"),
		NRet:    1,
		Protect: true,
	}, lua.LString("Go"), lua.LString("Lua")); err != nil {
		panic(err)
	}

	// Get the returned value from the stack and cast it to a lua.LString
	if str, ok := L.Get(-1).(lua.LString); ok {
		fmt.Println(str)
	}

	// Pop the value from the stack
	L.Pop(1)
}
