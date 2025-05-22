package main

import (
	"fmt"
	//"github.com/KoenMW/go-workshop/regression"
	"go-workshop/classification"
)

func main() {
	fmt.Println("Hello, Go 1.24.3!")
	classification.Init(5, "data/elden_ring_weapon.csv")
	//regression.Reg()
	prediction := classification.Predict(240, 62, 0, 0, 0, 47, 31, 31, 31, 2, 3, 0, 0, 3, 0, 3)
	fmt.Println("Predicted weapon type:", prediction)

}
