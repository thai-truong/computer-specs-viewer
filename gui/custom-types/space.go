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

var (
	unitToExponent = map[SpaceUnit]int{
		B:  0,
		KB: 3,
		MB: 6,
		GB: 9,
		TB: 12,
	}

	defaultUnit = MB
)

func (s Space) String() string {
	return fmt.Sprintf("%.2f %s", s.ConvAmount, s.ConvUnit)
}

func (s Space) ConvertToUnitFromByte(newUnit SpaceUnit) Space {
	exp, found := unitToExponent[newUnit]

	if !found {
		return Space{
			ByteAmount: s.ByteAmount,
			ConvAmount: float64(s.ByteAmount),
			ConvUnit:   B,
		}
	}

	newUnitSpace := float64(s.ByteAmount) / math.Pow(10, float64(exp))

	return Space{
		ByteAmount: s.ByteAmount,
		ConvAmount: newUnitSpace,
		ConvUnit:   newUnit,
	}
}

func (s Space) ToKb() Space {
	return s.ConvertToUnitFromByte(KB)
}

func (s Space) ToMb() Space {
	return s.ConvertToUnitFromByte(MB)
}

func (s Space) ToGb() Space {
	return s.ConvertToUnitFromByte(GB)
}

func (s Space) ToTb() Space {
	return s.ConvertToUnitFromByte(TB)
}

func NumToCustomSpaceType(spaceAmt uint64) Space {
	newSpace := Space{
		ByteAmount: spaceAmt,
		ConvAmount: float64(spaceAmt),
		ConvUnit:   B,
	}

	return newSpace.ConvertToUnitFromByte(MB)
}
