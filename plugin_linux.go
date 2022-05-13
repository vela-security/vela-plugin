package plugin

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
	"plugin"
	"runtime"
)

var (
	xEnv assert.Environment
)

func newLuaPlugin(L *lua.LState) int {
	path := L.CheckString(1)

	p, err := plugin.Open(path)
	if err != nil {
		L.RaiseError("%v", err)
		return 0
	}

	sym, err := p.Lookup("WithEnv")
	if err != nil {
		L.RaiseError("%v", err)
		return 0
	}
	sym.(func(assert.Environment))(xEnv)
	return 0
}

func Constructor(env assert.Environment) {
	xEnv = env
	xEnv.Infof("plugin running in %s", runtime.GOOS)
	env.Set("load", lua.NewFunction(newLuaPlugin))
}
