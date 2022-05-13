package plugin

import (
	"github.com/vela-security/pivot/pkg"
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
	"runtime"
	"syscall"
	"unsafe"
)

var (
	xEnv  assert.Environment
	_API_ = "WithEnv"
)

func newLuaPlugin(L *lua.LState) int {
	path := L.CheckString(1)
	dll, err := syscall.LoadLibrary(path)
	if err != nil {
		L.RaiseError("%v", err)
		return 0
	}

	inject, err := syscall.GetProcAddress(dll, _API_)
	if err != nil {
		L.RaiseError("%v", err)
		return 0
	}

	_, _, e := syscall.Syscall(inject, 1, uintptr(unsafe.Pointer(&xEnv)), 0, 0)
	if int(e) == 0 {
		L.RaiseError("%s", e.Error())
	}

	return 0
}

func Constructor(env *pkg.Environment) {
	xEnv = env
	env.Infof("plugin running in %s", runtime.GOOS)
	env.Set("load", lua.NewFunction(newLuaPlugin))
}
