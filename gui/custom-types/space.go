package custom_types

import (
	"fmt"
	"math"
)

type Space struct {
	ByteAmount uint64
	ConvAmount float64
	ConvUnit   SpaceUnit
}

type SpaceUnit string

const (
	B  SpaceUnit = "B"
	KB           = "KB"
	MB           = "MB"
	GB           = "GB"
	TB           = "TB"
)

var unitToExponent = map[SpaceUnit]int{
	B:  0,
	KB: 3,
	MB: 6,
	GB: 9,
	TB: 12,
}

func (s Space) String() string {
	return fmt.Sprintf("%v %s", s.ConvAmount, s.ConvUnit)
}

func ConvertUnitFromByte(curr Space, newUnit SpaceUnit) Space {
	exp, found := unitToExponent[newUnit]

	if !found {
		return Space{
			ByteAmount: curr.ByteAmount,
			ConvAmount: float64(curr.ByteAmount),
			ConvUnit:   B,
		}
	}

	newUnitSpace := float64(curr.ByteAmount) / math.Pow(10, float64(exp))

	return Space{
		ByteAmount: curr.ByteAmount,
		ConvAmount: newUnitSpace,
		ConvUnit:   newUnit,
	}
}
