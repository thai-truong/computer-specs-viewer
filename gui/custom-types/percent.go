package custom_types

import "fmt"

type Percent struct {
	DecimalValue float64
}

func (p Percent) String() string {
	return fmt.Sprintf("%.2f%%", p.DecimalValue)
}

func CreatePercent(percentValue float64) Percent {
	return Percent{DecimalValue: percentValue * 100}
}
