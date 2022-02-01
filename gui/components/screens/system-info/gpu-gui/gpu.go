package gpugui

import (
	custom_types "computer-specs-viewer/gui/custom-types"
	"time"
)

type GpuInformationGUI struct {
	Index                int
	Name                 string
	Version              string
	LastModified         time.Time
	Resolution           string
	NumColors            uint64
	CurrRefreshRate      uint32
	RefreshRateRange     string
	BitsPerPixel         uint32
	Status               string
	Availability         string
	MemorySize           custom_types.Space
	AdapterCompatibility string
	AdapterDACType       string
	IsMonochrome         bool
	Manufacturer         string
	PresentOnSystem      bool
}
