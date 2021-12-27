package utils

func GetFreeUsedPercents(total uint64, free uint64, used uint64) (freePercent float64, usedPercent float64) {
	totalF := float64(total)
	freeF := float64(free)
	usedF := float64(used)

	if totalF == 0 {
		return 0, 0
	}

	freePercent = freeF / totalF
	usedPercent = usedF / totalF

	return freePercent, usedPercent
}
