package csvutil

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

// Weapon represents all attributes from the CSV
type Weapon struct {
	Name    string
	Type    string
	Phy     float64
	Mag     float64
	Fir     float64
	Lit     float64
	Hol     float64
	Cri     float64
	Sta     float64
	Str     float64
	Dex     float64
	Int     float64
	Fai     float64
	Arc     float64
	Any     float64
	PhyDF   float64
	MagDF   float64
	FirDF   float64
	LitDF   float64
	HolDF   float64
	Bst     float64
	Rst     float64
	Wgt     float64
	Upgrade string
}

// LoadWeapons reads and parses weapon data from a CSV file
func LoadWeapons(filename string) ([]Weapon, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var weapons []Weapon

	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}

		weapon := Weapon{
			Name:    record[0],
			Type:    record[1],
			Upgrade: record[22],
		}

		for j := 2; j < len(record)-1; j++ {
			parsed := strings.Replace(record[j], "(", "", 1)
			parsed = strings.Replace(parsed, ")", "", 1)
			parsed = strings.Replace(parsed, "-", "0", 1)
			parsed = strings.Replace(parsed, "E", "0", 1)
			parsed = strings.Replace(parsed, "D", "1", 1)
			parsed = strings.Replace(parsed, "C", "2", 1)
			parsed = strings.Replace(parsed, "B", "3", 1)
			parsed = strings.Replace(parsed, "A", "4", 1)
			parsed = strings.Replace(parsed, "S", "5", 1)
			val, err := strconv.ParseFloat(parsed, 64)
			if err != nil {
				println("it wasn't a float: ", record[j])
				val = 0 // Default on parse error
			}
			switch j {
			case 2:
				weapon.Phy = val
			case 3:
				weapon.Mag = val
			case 4:
				weapon.Fir = val
			case 5:
				weapon.Lit = val
			case 6:
				weapon.Hol = val
			case 7:
				weapon.Cri = val
			case 8:
				weapon.Sta = val
			case 9:
				weapon.Str = val
			case 10:
				weapon.Dex = val
			case 11:
				weapon.Int = val
			case 12:
				weapon.Fai = val
			case 13:
				weapon.Arc = val
			case 14:
				weapon.Any = val
			case 15:
				weapon.PhyDF = val
			case 16:
				weapon.MagDF = val
			case 17:
				weapon.FirDF = val
			case 18:
				weapon.LitDF = val
			case 19:
				weapon.HolDF = val
			case 20:
				weapon.Bst = val
			case 21:
				weapon.Rst = val
			case 22:
				weapon.Wgt = val
			}
		}

		weapons = append(weapons, weapon)
	}

	return weapons, nil
}
