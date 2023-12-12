package c1

import "time"

func init() {
	println("c1.init")
	time.Sleep(2*time.Second)
	println("c1.init done")
}