package util

import (
	"math/rand"
	"time"

	lua "github.com/yuin/gopher-lua"
)

//生成随机数
func RandInt(L *lua.LState)  {
	gomin := L.ToInt(1)
	gomax := L.ToInt(2)
	rand.Seed(time.Now().UnixNano())
	if gomin >= gomax || gomin == 0 || gomax == 0 {
			L.Push(lua.LNumber(rand.Intn(gomax-gomin) + gomin))
		return
	}
	L.Push(lua.LNumber(rand.Intn(gomax-gomin) + gomin))
	return
}
