package gpuinfo

var AvailabilityMapping = map[uint16]string{
	1:  "Other",
	2:  "Unknown",
	3:  "Running/Full Power",
	4:  "Warning",
	5:  "In Test",
	6:  "Not Applicable",
	7:  "Power Off",
	8:  "Offline",
	9:  "Off Duty",
	10: "Degraded",
	11: "Not Installed",
	12: "Install Error",
	13: "Power Save - Unknown",
	14: "Power Save - Low Power Mode",
	15: "Power Save - Standby",
	16: "Power Cycle",
	17: "Power Save - Warning",
	18: "Paused",
	19: "Not Ready",
	20: "Not Configured",
	21: "Quiesced",
}

var VideoArchitectureMapping = map[uint16]string{
	1:  "Other",
	2:  "Unknown",
	3:  "VRAM",
	4:  "DRAM",
	5:  "SRAM",
	6:  "WRAM",
	7:  "EDO RAM",
	8:  "Burst Synchronous DRAM",
	9:  "Pipelined Burst SRAM",
	10: "CDRAM",
	11: "3DRAM",
	12: "SDRAM",
	13: "SGRAM",
}