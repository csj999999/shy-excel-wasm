package main

import "testing"
import "syscall/js"

func TestRegInteropFunc(t *testing.T) {
	js.Global().Set("shyexcel", map[string]interface{}{})
	regFuncs()
}
