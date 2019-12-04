package main

import (
	"github.com/davecgh/go-spew/spew"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("demo.lua"); err != nil {
		panic(err)
	}
	for index := 0; index < 10; index++ {

		if err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("request"), // 获取double这个函数的引用
			NRet:    2,                      // 指定返回值数量
			Protect: false,                  // 如果出现异常，是panic还是返回err
		}); err != nil {
			spew.Dump(err)
		}
		spew.Dump(L.Get(1))
		spew.Dump(L.Get(2))
		L.Pop(2)
		L.Remove(0)
	}

}
