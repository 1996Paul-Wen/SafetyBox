package callstack

import (
	"fmt"
	"testing"
)

func TestGetfuncname(t *testing.T) {

	funcname := GetDirectCallerName()
	fmt.Println(funcname) // TestGetfuncname
}

func TestGetLevelfuncname(t *testing.T) {
	Level1Path()
}

func Level1Path() {
	Level2Path()
}

func Level2Path() {
	fmt.Println(GetCallerNameBySkip(3), GetCallerNameBySkip(2), GetCallerNameBySkip(1), GetCallerNameBySkip(0))
	// Level1Path Level2Path GetCallerNameBySkip Callers
}
