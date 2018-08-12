package p34

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Field struct {
		H int `json:"h"`
		W int `json:"w"`
	} `json:"field"`
	Lives []Pos `json:"lives"`
}

type Pos struct {
	X int
	Y int
}

func (o Configuration) blankGrid() Grid {
	rows := make(Grid, o.Field.H+2)

	for i, _ := range rows {
		rows[i] = make([]bool, o.Field.W+2)
	}

	return rows
}

func (o Configuration) toGrid() Grid {
	rows := o.blankGrid()

	for _, life := range o.Lives {
		rows[life.Y+1][life.X+1] = true
	}

	return rows
}

func toConfigurationFromJSON(path string) Configuration {
	bytes, _ := ioutil.ReadFile(path)

	var c Configuration
	json.Unmarshal(bytes, &c)

	return c
}

func display(grid Grid) string {
	str := ""

	for y, cols := range grid {
		for x, col := range cols {
			switch {
			case y == 0 || x == 0 || y == len(grid)-1 || x == len(cols)-1:
				str += "+ "
			case col:
				str += "■ "
			default:
				str += "□ "
			}
		}
		str += "\n"
	}

	return str
}
