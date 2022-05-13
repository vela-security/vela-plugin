package plugin

import (
	"github.com/vela-security/pivot/pkg"
	"github.com/vela-security/vela-public/lua"
	"runtime"
)

var xEnv *pkg.Environment

func newLuaPlugin(L *lua.LState) int {
	xEnv.Errorf("plugin not support darwin")
	return 0
}

func Constructor(env *pkg.Environment) {
	xEnv = env
	env.Infof("plugin running in %s", runtime.GOOS)
	env.Set("load", lua.NewFunction(newLuaPlugin))
}
