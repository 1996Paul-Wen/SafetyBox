package callstack

import (
	"runtime"
	"strings"
)

// GetCallerNameBySkip 获取调用函数名字 , 可以配置skip 层级
func GetCallerNameBySkip(skip int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(skip, pc)
	f := runtime.FuncForPC(pc[0])
	spPath := strings.Split(f.Name(), ".")
	return spPath[len(spPath)-1]
}

// GetDirectCallerName 获取函数的直接调用方名字
func GetDirectCallerName() string {
	return GetCallerNameBySkip(3)
}
