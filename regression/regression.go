package regression

import (
	"fmt"
	"go-workshop-1/csvutil"
	"reflect"

	"github.com/sajari/regression"
)

type RegOptions struct {
	YField  string
	XFields []string
}

func CreateDataPoints(weapons []csvutil.Weapon, yField string, xFields []string) (datapoints regression.DataPoints, err error) {
	var weaponDataPoints regression.DataPoints
	for _, weapon := range weapons {
		values := reflect.ValueOf(weapon)

		yValue := values.FieldByName(yField)

		if !yValue.IsValid() || yValue.Kind() != reflect.Float64 {
			return nil, fmt.Errorf("invalid Y field: %s", yField)
		}

		xRow := []float64{}

		for _, xField := range xFields {
			xValue := values.FieldByName(xField)

			if !xValue.IsValid() || xValue.Kind() != reflect.Float64 {
				println("oh no")
				return nil, fmt.Errorf("invalid X field: %s", xField)
			}

			xRow = append(xRow, xValue.Float())
		}

		weaponDataPoints = append(weaponDataPoints, regression.DataPoint(yValue.Float(), xRow))
	}

	return weaponDataPoints, nil
}

func Reg(weapons regression.DataPoints, args ...interface{}) {
	yField := ""

	xFields := []string{}

	if len(args) >= 1 {
		if y, ok := args[0].(string); ok {
			yField = y
		}
	}
	if len(args) >= 2 {
		if x, ok := args[1].([]string); ok {
			xFields = x
		}
	}

	r := new(regression.Regression)
	r.SetObserved(yField)

	for i, v := range xFields {
		r.SetVar(i, v)
	}

	// r.SetVar(1, "Percent with incomes below $5000")
	// r.SetVar(2, "Percent unemployed")
	r.Train(weapons...)
	r.Run()

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)
}
