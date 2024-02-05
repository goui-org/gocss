package gocss

import "syscall/js"

var global = js.Global()
var document = global.Get("document")
var documentElement = document.Get("documentElement")
var documentElementStyle = documentElement.Get("style")
var sheet js.Value
var cssRules js.Value

func init() {
	var head = document.Get("head")
	var style = document.Call("createElement", "style")
	head.Call("appendChild", style)
	sheet = style.Get("sheet")
	cssRules = sheet.Get("cssRules")
}

func insertRule(rule string) {
	sheet.Call("insertRule", rule, cssRules.Length())
}
