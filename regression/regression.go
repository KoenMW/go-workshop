package regression

import (
	"fmt"
	"go-workshop-1/csvutil"
	"log"

	"github.com/sajari/regression"
)

var weaponDataPoints regression.DataPoints

func Init(csvPath string) {

	weapons, err := csvutil.LoadWeapons(csvPath)
	if err != nil {
		log.Fatalf("Failed to load weapons: %v", err)
	}

	for _, w := range weapons {
		weaponDataPoints = append(weaponDataPoints, regression.DataPoint(w.PhyDF, []float64{
			w.Phy, w.Mag, w.Fir, w.Lit, w.Hol, w.Cri, w.Sta, w.Str, w.Dex, w.Int, w.Fai, w.Arc, w.Any, w.Bst, w.Rst, w.Wgt, w.FirDF, w.HolDF, w.LitDF, w.MagDF,
		}))
	}
}

func Reg() {

	if len(weaponDataPoints) == 0 {
		fmt.Println("weapon data not loaded please call regression.Init() first")
		return
	}

	// print(dataPoints)

	r := new(regression.Regression)
	r.SetObserved("PhyDF")
	r.SetVar(0, "Bst")
	// r.SetVar(1, "Percent with incomes below $5000")
	// r.SetVar(2, "Percent unemployed")
	r.Train(weaponDataPoints...)
	r.Run()

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)
}
