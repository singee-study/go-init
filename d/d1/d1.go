package d1

import "time"

func init() {
	println("d1.init")
	time.Sleep(2*time.Second)
	println("d1.init done")
}