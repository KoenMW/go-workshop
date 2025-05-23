package main

import (
	"fmt"
	"go-workshop-1/regression"
)

func main() {
	fmt.Println("Hello, Go 1.24.3!")

	regression.Init("./data/elden_ring_weapon.csv")
	regression.Reg()
}
