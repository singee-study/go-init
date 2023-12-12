package main

import (
	_ "go-init/b"
	_ "go-init/a"
	_ "go-init/c/c1"
	_ "go-init/c/c2"
	_ "go-init/d/d1"
	_ "go-init/d/d2"
)

func init() {
	println("main.init")
}

func main() {
	println("main")
}