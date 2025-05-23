package main

import (
	"fmt"
	"go-workshop-1/csvutil"
	"go-workshop-1/regression"
)

func main() {
	fmt.Println("Hello, Go 1.24.3!")

	weapons, err := csvutil.LoadWeapons("./data/elden_ring_weapon.csv")

	if err != nil {
		println("some error occured: ", err.Error())
		return
	}
	datapoints, err := regression.CreateDataPoints(weapons, "PhyDF", []string{"Bst"})

	if err != nil {
		println("some error occured: ", err.Error())
		return
	}

	regression.Reg(datapoints, "PhyDF", []string{"Bst"})
}
